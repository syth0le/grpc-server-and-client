gen-go-proto:
	protoc -I=api/proto --go_out=pkg/api api/proto/simple.proto
	// generate go code

gen-grpc-proto:
	protoc -I api/proto --go-grpc_out=require_unimplemented_servers=false:pkg api/proto/simple.proto
	// generate grpc code

call-evans:
	evans api/proto/simple.proto -p=50051

call-client:
	go run ./cmd/client/main.go Name 0