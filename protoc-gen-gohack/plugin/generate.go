package plugin

//go:generate go get github.com/golang/protobuf/protoc-gen-go
//go:generate protoc -I. --go_out=Mgoogle/protobuf/descriptor.proto=github.com/kixelated/protohack/descriptor:. ./plugin.proto
