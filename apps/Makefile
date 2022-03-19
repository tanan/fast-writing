build:
	protoc -I./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/models/*.proto
	protoc -I./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/*.proto
	protoc -I./proto --js_out=import_style=commonjs,binary:front/src/pb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:front/src/pb ./proto/models/*.proto
	protoc -I./proto --js_out=import_style=commonjs,binary:front/src/pb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:front/src/pb ./proto/*.proto