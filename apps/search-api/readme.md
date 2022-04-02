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
    "match": {
      "title": "空港"
    }
  }
}
'
```

```
$ cat save-test.json | grpcurl -plaintext -d @ localhost:10002 fastwriting.SearchService.SaveSearchContents
```