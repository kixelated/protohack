package generator

import (
	"errors"

	desc "github.com/kixelated/protohack/descriptor"
)

func (g *Generator) lookupFile(name string) (file *desc.FileDescriptorProto, err error) {
	for _, file := range g.files {
		if file.GetName() == name {
			return file, nil
		}
	}

	err = errors.New("could not lookup file: " + name)
	return nil, err
}

func (g *Generator) lookupMessage(name string) (message *desc.DescriptorProto) {
	for _, file := range g.files {
		prefix := "." + file.GetPackage()

		message = g.lookupMessageNested(name, prefix, file.GetMessageType())
		if message != nil {
			return message
		}
	}

	return nil
}

func (g *Generator) lookupMessageNested(name string, prefix string, messages []*desc.DescriptorProto) (message *desc.DescriptorProto) {
	for _, message := range messages {
		messageName := prefix + "." + message.GetName()
		if messageName == name {
			return message
		}

		message = g.lookupMessageNested(name, messageName, message.GetNestedType())
		if message != nil {
			return message
		}
	}

	return nil
}
