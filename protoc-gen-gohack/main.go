package main

import (
	"io/ioutil"
	"log"
	"os"

	proto "github.com/golang/protobuf/proto"
	proto_plugin "github.com/golang/protobuf/protoc-gen-go/plugin"

	plugin "github.com/kixelated/protohack/plugin"
	plugin_base "github.com/kixelated/protohack/plugin/base"
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

	request := new(proto_plugin.CodeGeneratorRequest)
	err = proto.Unmarshal(data, request)
	if err != nil {
		return err
	}

	plugins := loadPlugins()

	response := new(proto_plugin.CodeGeneratorResponse)

	for i, inputName := range request.FileToGenerate {
		inputProto := request.ProtoFile[i]

		for _, generate := range plugins {
			pluginRequest := &plugin.Request{
				Name: inputName,
				File: inputProto,
			}

			pluginResponse, err := generate(pluginRequest)
			if err != nil {
				return err
			}

			responseFile := new(proto_plugin.CodeGeneratorResponse_File)

			if pluginResponse.Name != "" {
				responseFile.Name = &pluginResponse.Name
			}

			if pluginResponse.InsertionPoint != "" {
				responseFile.InsertionPoint = &pluginResponse.InsertionPoint
			}

			if pluginResponse.Content != "" {
				responseFile.Content = &pluginResponse.Content
			}

			response.File = append(response.File, responseFile)
		}
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

func loadPlugins() (plugins []plugin.Plugin) {
	plugins = append(plugins, plugin_base.Generate)
	return plugins
}
