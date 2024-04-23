install-gimports:
	go install golang.org/x/tools/cmd/goimports@latest

install-protoc:
	go install github.com/golang/protobuf/protoc-gen-go@v1.5.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
