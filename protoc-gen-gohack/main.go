package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	"github.com/kixelated/protohack/generator"
)

func main() {
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

	request := &plugin.CodeGeneratorRequest{}
	err = proto.Unmarshal(data, request)
	if err != nil {
		return err
	}

	response := &plugin.CodeGeneratorResponse{}
	for _, requestFile := range request.ProtoFile {
		content, err := generator.Generate(requestFile)
		if err != nil {
			response.Error = proto.String(err.Error())
			break
		}

		name := generator.FileName(*requestFile.Name)

		fileResponse := &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(name),
			Content: proto.String(content),
		}

		response.File = append(response.File, fileResponse)
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
