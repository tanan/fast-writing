clean:
	rm -rf ./proto ./pb

build:
	cp -rp ../proto .
	protoc -I./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/models/*.proto
	protoc -I./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/*.proto
