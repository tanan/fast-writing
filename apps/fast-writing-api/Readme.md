## grpc

- install

```
## grpc
% go get -u google.golang.org/grpc

## protobuf compiler
% brew install protobuf

## protoc-gen-go
% go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
% go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- compile
```
% protoc -I../proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ../proto/*.proto
```

- request
```
% grpc_cli call localhost:10001 fastwriting.UserService.GetUser 'id: "test"'
```