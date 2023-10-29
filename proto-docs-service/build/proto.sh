export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
protoc -I=. --go_out=./grpc grpc/docs.proto