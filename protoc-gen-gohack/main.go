package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/kixelated/protohack/generator"
	"github.com/kixelated/protohack/protoc-gen-gohack/plugin"
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

	request := new(plugin.CodeGeneratorRequest)
	err = proto.Unmarshal(data, request)
	if err != nil {
		return err
	}

	// Create the generator with all of the file descriptors.
	generator := generator.New(request.GetProtoFile())

	response := new(plugin.CodeGeneratorResponse)

	for _, name := range request.FileToGenerate {
		name, content, err := generator.File(name)
		if err != nil {
			return err
		}

		responseFile := &plugin.CodeGeneratorResponse_File{
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
