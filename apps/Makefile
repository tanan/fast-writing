build:
	protoc -I./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/models/*.proto
	protoc -I./proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/*.proto
	protoc -I./proto --js_out=import_style=commonjs,binary:front/src/pb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:front/src/pb ./proto/models/*.proto
	protoc -I./proto --js_out=import_style=commonjs,binary:front/src/pb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:front/src/pb ./proto/*.proto
	cp -pr fast-writing/pkg/pb/* ../pkg/pb/
	rm -rf fast-writing google.golang.org
descriptor:
	protoc -I./proto --go_out=. --experimental_allow_proto3_optional --include_imports --include_source_info --descriptor_set_out=api_descriptor.pb --go-grpc_out=require_unimplemented_servers=false:. ./proto/**/*.proto
	cp -rp api_descriptor.pb ../environments/endpoints/
	rm -rf fast-writing api_descriptor.pb google.golang.org
