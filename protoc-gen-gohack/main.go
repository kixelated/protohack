package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	proto_descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	proto_plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"github.com/kixelated/protohack/descriptor"
	"github.com/kixelated/protohack/generator"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	request := new(proto_plugin.CodeGeneratorRequest)
	err = proto.Unmarshal(data, request)
	if err != nil {
		return err
	}

	// For some reason, the request contains a slice of files instead of a set.
	// This conversion is not entirely neccisary but is more consistent.
	filesProto := &proto_descriptor.FileDescriptorSet{
		File: request.GetProtoFile(),
	}

	// Wrap the proto objects with helper fields and methods.
	files := descriptor.NewFileSet(filesProto)

	// Create the generator with all of the file descriptors.
	generator := generator.New(files)

	response := new(proto_plugin.CodeGeneratorResponse)

	for _, name := range request.FileToGenerate {
		name, content := generator.File(name)
		if err != nil {
			return err
		}

		responseFile := &proto_plugin.CodeGeneratorResponse_File{
			Name:    proto.String(name),
			Content: proto.String(content),
		}

		response.File = append(response.File, responseFile)
	}

	data, err = proto.Marshal(response)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		return err
	}

	return nil
}
