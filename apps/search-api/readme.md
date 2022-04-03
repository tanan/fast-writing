### request

```
## create
$ cat save-test.json | grpcurl -plaintext -d @ localhost:10002 fastwriting.SearchService.SaveSearchContents

## get
$ grpcurl -plaintext -d '{ "title": "greeting" }'  localhost:10002 fastwriting.SearchService.FindContentsIdListByTitle
```

### index api
```
$ curl -XPOST -H "Content-Type: application/json" localhost:9200/fast-writing/_doc/ -d '
{
  "title": "空港での会話",
  "category": "category",
  "contents_id": "123",
  "username": "anan",
  "question": "いつ関西空港に到着しますか？",
  "answer": "When do you arrive Kansai Airport?"
}
'
```

### search

```
$ curl -XGET -H "Content-Type: application/json" localhost:9200/fast-writing/_search/ -d '
{
  "query": {
    "multi_match": {
      "query": "空港",
      "type": "most_fields",
      "operator": "or",
      "fields": ["title", "question"]
    }
  }
}
'
```

### analyze
```
$ curl -XGET -H "Content-Type: application/json" "localhost:9200/fast-writing/_analyze" -d '
{
  "analyzer": "japanese_analyzer",
  "text": "いつ関西空港に到着しますか？"
}
'
```
