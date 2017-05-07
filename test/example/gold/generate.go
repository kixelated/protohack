package example

//go:generate go get github.com/golang/protobuf/protoc-gen-go
//go:generate protoc -I.. --go_out=. ../example.proto
