package database

import (
	"encoding/json"
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func (h *SQLHandler) FindContentsList(limit int32, offset int32) ([]*domain.Contents, error) {
	var m []model.Contents
	db := h.Conn.Where("id > ? and scope = ?", offset, "public").Order("id").Limit(int(limit)).Find(&m)
	if db.Error != nil {
		return []*domain.Contents{}, errors.New("cannot find contents list: " + db.Error.Error())
	}
	return h.toContentsList(m)
}

func (h *SQLHandler) FindContentsListByTags(tags string, limit int32, offset int32) ([]*domain.Contents, error) {
	var m []model.Contents
	// TODO 検索するタグは１つのみ設定される想定。ESを稼働し始めたら複数タグに対応する
	db := h.Conn.Where("id > ? and scope = ? and tags like ?", offset, "public", "%"+tags+"%").Order("id").Limit(int(limit)).Find(&m)
	if db.Error != nil {
		return []*domain.Contents{}, errors.New("cannot find contents list: " + db.Error.Error())
	}
	return h.toContentsList(m)
}

func (h *SQLHandler) FindContentsListByUserId(userId domain.UserId, limit int32, offset int32) ([]*domain.Contents, error) {
	var m []model.Contents
	db := h.Conn.Where("id > ? and user_id = ?", offset, userId).Order("id").Limit(int(limit)).Find(&m)
	if db.Error != nil {
		return []*domain.Contents{}, errors.New("cannot find contents list by user: " + db.Error.Error())
	}
	return h.toContentsList(m)
}

func (h *SQLHandler) FindContentsById(id domain.ContentsId) (domain.Contents, error) {
	var m model.Contents
	db := h.Conn.Where("id = ?", id).First(&m)
	if db.Error != nil {
		return domain.Contents{}, errors.New("cannot find contents by id: " + db.Error.Error())
	}

	var user model.User
	userDb := h.Conn.Where("id = ?", m.UserId).First(&user)
	if userDb.Error != nil {
		return domain.Contents{}, errors.New("cannot find user by id: " + db.Error.Error())
	}

	quizJson, err := m.Quiz.MarshalJSON()
	if err != nil {
		return domain.Contents{}, errors.New("failed to marshal json: " + err.Error())
	}
	return *h.toContents(m, user.Name, quizJson), nil
}

func (h *SQLHandler) CreateContents(contents domain.Contents, userId domain.UserId) (domain.Contents, error) {
	j, _ := json.Marshal(&contents.QuizList)
	m := model.Contents{
		Id:          int64(contents.ContentsId),
		UserId:      string(userId),
		Title:       contents.Title,
		Description: contents.Description,
		Scope:       contents.Scope,
		Tags:        contents.Tags,
		LastUpdated: time.Now(),
		Quiz:        j,
	}
	if db := h.Conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"user_id": m.UserId, "title": m.Title, "description": m.Description, "scope": m.Scope, "tags": m.Tags, "last_updated": time.Now(), "quiz": j}),
	}).Create(&m); db.Error != nil {
		return domain.Contents{}, errors.New("cannot create contents: " + db.Error.Error())
	}

	contents.ContentsId = domain.ContentsId(m.Id)
	contents.LastUpdated = m.LastUpdated
	return contents, nil
}

func (h *SQLHandler) DeleteContents(userId domain.UserId, contentsId domain.ContentsId) (int64, error) {
	db := h.Conn.Where("user_id = ?", userId).Delete(&model.Contents{}, contentsId)
	if db.Error != nil {
		return 0, errors.New("cannot delete contents: " + db.Error.Error())
	}
	return 1, nil
}

func (h *SQLHandler) toContentsList(m []model.Contents) ([]*domain.Contents, error) {
	var contentsList []*domain.Contents
	userList := map[string]string{}
	for _, v := range m {
		if userName, ok := h.contains(v.UserId, userList); ok {
			contentsList = append(contentsList, h.toContents(v, userName, nil))
			continue
		}
		var user model.User
		db := h.Conn.Where("id = ?", v.UserId).First(&user)
		if db.Error != nil {
			return []*domain.Contents{}, errors.New("cannot find user by id: " + db.Error.Error())
		}
		userList[user.Id] = user.Name
		contentsList = append(contentsList, h.toContents(v, user.Name, nil))
	}
	return contentsList, nil
}

func (h *SQLHandler) contains(userId string, userList map[string]string) (string, bool) {
	for k, v := range userList {
		if userId == k {
			return v, true
		}
	}
	return "", false
}

func (h *SQLHandler) toContents(contents model.Contents, userName string, quizJson []byte) *domain.Contents {
	return &domain.Contents{
		ContentsId:  domain.ContentsId(contents.Id),
		Creator:     userName,
		Title:       contents.Title,
		Description: contents.Description,
		Scope:       contents.Scope,
		Tags:        contents.Tags,
		QuizList:    h.toQuizList(quizJson),
		LastUpdated: contents.LastUpdated,
	}
}

func (h *SQLHandler) toQuizList(l []byte) []domain.Quiz {
	var quizList []domain.Quiz
	if l == nil {
		return nil
	}
	err := json.Unmarshal(l, &quizList)
	if err != nil {
		log.Printf("unmarshal error: %v", err.Error())
		log.Printf("target json: %v", string(l))
		return nil
	}
	return quizList
}
