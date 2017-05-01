package plugin

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Plugin func(request *Request) (response *Response, err error)

type Request struct {
	Name string
	File *desc.FileDescriptorProto
}

type Response struct {
	Name           string
	InsertionPoint string
	Content        string
}
