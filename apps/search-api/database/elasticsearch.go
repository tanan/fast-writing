package database

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"search-api/domain"
	"strconv"
	"time"
	"unicode"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

type SearchHandler struct {
	*elasticsearch8.Client
}

const INDEX = "fast-writing"

func NewElasticsearch() *SearchHandler {
	cfg := elasticsearch8.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return &SearchHandler{
		es,
	}
}

func (handler *SearchHandler) Version() {
	res, err := handler.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}

func (handler *SearchHandler) getQuizFieldByLanguage(isJapanese bool) string {
	if isJapanese {
		return "question"
	}
	return "answer"
}

func (handler *SearchHandler) generateQuery(keyword string, limit int32, offset int32, isJapanese bool) map[string]interface{} {
	return map[string]interface{}{
		"size": limit,
		"from": offset,
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":    keyword,
				"type":     "most_fields",
				"operator": "or",
				"fields":   []string{"title", handler.getQuizFieldByLanguage(isJapanese)},
			},
		},
	}
}

func (handler *SearchHandler) isJapanese(keyword string) bool {
	for _, v := range []rune(keyword) {
		if !unicode.In(v, unicode.Hiragana, unicode.Katakana, unicode.Han) {
			return false
		}
	}
	return true
}

func (handler *SearchHandler) FindContents(keyword string, limit int32, offset int32) ([]domain.ContentsScore, error) {
	query := handler.generateQuery(keyword, limit, offset, handler.isJapanese(keyword))
	data, err := json.Marshal(&query)
	if err != nil {
		log.Fatalf("Cannot encode: %d", err)
	}
	req := esapi.SearchRequest{
		Index: []string{INDEX},
		Body:  bytes.NewReader(data),
	}
	res, err := req.Do(context.Background(), handler.Client)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(fmt.Sprintf("[%s] Error indexing document", res.Status()))
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return handler.toContentsScoreList(r["hits"].(map[string]interface{})["hits"].([]interface{})), nil
}

func (handler *SearchHandler) toContentsScoreList(hits []interface{}) []domain.ContentsScore {
	var contentsScores []domain.ContentsScore
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		contentsScores = append(contentsScores, domain.ContentsScore{
			Id:          domain.ContentsId(int64(source.(map[string]interface{})["contents_id"].(float64))),
			Score:       float32(hit.(map[string]interface{})["_score"].(float64)),
			LastUpdated: time.Now(),
		})
	}
	return contentsScores
}

func (handler *SearchHandler) toContentsList(hits []interface{}) []domain.Contents {
	var contentsList []domain.Contents
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		contentsList = append(contentsList, domain.Contents{
			Id:       domain.ContentsId(int64(source.(map[string]interface{})["contents_id"].(float64))),
			Title:    source.(map[string]interface{})["title"].(string),
			Category: source.(map[string]interface{})["category"].(string),
			Username: source.(map[string]interface{})["username"].(string),
			Question: source.(map[string]interface{})["question"].(string),
			Answer:   source.(map[string]interface{})["answer"].(string),
		})
	}
	return contentsList
}

func (handler *SearchHandler) SaveContents(contents domain.Contents) (string, error) {
	data, err := json.Marshal(&contents)
	if err != nil {
		log.Fatalf("Cannot encode: %d", err)
	}
	req := esapi.IndexRequest{
		Index:      INDEX,
		DocumentID: strconv.Itoa(int(contents.Id)),
		Body:       bytes.NewReader(data),
	}
	res, err := req.Do(context.Background(), handler.Client)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	if res.IsError() {
		return "", errors.New(fmt.Sprintf("[%s] Error indexing document", res.Status()))
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing the response body: %s", err))
	}
	return r["_id"].(string), nil
}
