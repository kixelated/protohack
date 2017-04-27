package generator

import (
	"errors"
	"path"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func Generate(file *descriptor.FileDescriptorProto) (output string, err error) {
	w := newWriter()

	if file.Options.GoPackage == nil {
		err = errors.New("missing go package name")
		return "", err
	}

	w.Package(*file.Options.GoPackage)

	output, err = w.Output()
	if err != nil {
		return "", err
	}

	return output, nil
}

func FileName(name string) string {
	ext := path.Ext(name)

	if ext == ".proto" {
		name = name[:len(name)-len(ext)]
	}

	return name + ".pb.go"
}
