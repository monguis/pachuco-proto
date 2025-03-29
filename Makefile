setup:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

compile:
	bash recompile-protos.sh
