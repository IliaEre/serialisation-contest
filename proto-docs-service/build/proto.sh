# if it is your first downloading:
#brew install protobuf
#go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
#export PATH="$PATH:$(go env GOPATH)/bin"

# only generator
cd .. && cd grpc
mkdir "docs"
protoc --go_out=./docs --go_opt=paths=source_relative \
    --go-grpc_out=./docs --go-grpc_opt=paths=source_relative docs.proto
