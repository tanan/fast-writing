### index api
```
curl -XPOST /fast-writing/_doc/ -d '
{
"title": "test",
"category": "category",
"contents_id": "1234567890asdfghjk",
"username": "anan",
"quiz": "私は誰ですか?",
"answer": "Who am I?"
}
'
```

### search

```
$ curl -XGET -H "Content-Type: application/json" localhost:9200/fast-writing/_search/ -d '
{
"query": {
"match": {
"username": "anan"
}
}
}
'
```

```
$ cat save-test.json | grpcurl -plaintext -d @ localhost:10002 fastwriting.SearchService.SaveSearchContents
```