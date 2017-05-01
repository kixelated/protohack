package base

import (
	"path"
	"strings"

	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/kixelated/protohack/codegen"
	"github.com/kixelated/protohack/plugin"
)

type gen struct {
	w codegen.Writer
}

func Generate(request *plugin.Request) (response *plugin.Response, err error) {
	g := new(gen)
	g.writeContent(request.File)

	response = &plugin.Response{
		Name:    outputName(request.Name),
		Content: g.w.String(),
	}

	return response, nil
}

func outputName(inputName string) (name string) {
	name = inputName

	ext := path.Ext(name)
	if ext == ".proto" {
		name = name[:len(name)-len(ext)]
	}

	name = name + ".pb.go"
	return name
}

func (g *gen) writeContent(file *desc.FileDescriptorProto) {
	g.writePackage(file)
	g.writeImports(file)
	g.writeEnums(file)
	g.writeTypes(file)
}

func packageName(file *desc.FileDescriptorProto) (name string) {
	name = file.Options.GetGoPackage()
	if name != "" {
		return name
	}

	parts := strings.Split(*file.Package, ".")
	name = parts[len(parts)-1]

	return name
}

func (g *gen) writePackage(file *desc.FileDescriptorProto) {
	g.w.Line(`package ` + packageName(file))
}

func (g *gen) writeImports(file *desc.FileDescriptorProto) {
	g.w.Comment(`@@protoc_insertion_point(import)`)
}

func (g *gen) writeEnums(file *desc.FileDescriptorProto) {
	g.w.Comment(`@@protoc_insertion_point(enum)`)
}

func (g *gen) writeTypes(file *desc.FileDescriptorProto) {
	for _, messageType := range file.MessageType {
		g.writeMessageType(messageType)
	}

	g.w.Comment(`@@protoc_insertion_point(type)`)
}

func (g *gen) writeMessageType(message *desc.DescriptorProto) {
	g.w.Line(`type ` + messageGoType(message) + ` struct {`)
	g.w.In()

	for _, field := range message.Field {
		g.w.Line(fieldGoName(field) + ` ` + fieldGoType(field))
	}

	g.w.Comment(`@@protoc_insertion_point(type_` + messageGoType(message) + `)`)

	g.w.Out()
	g.w.Line(`}`)
}

func messageGoType(message *desc.DescriptorProto) (name string) {
	if message.Name == nil {
		panic("message has no name")
	}

	name = *message.Name
	return name
}

func fieldGoName(field *desc.FieldDescriptorProto) (name string) {
	if field.Name == nil {
		panic("field has no name")
	}

	parts := strings.Split(*field.Name, "_")

	nextParts := make([]string, len(parts))
	for i, part := range parts {
		nextParts[i] = strings.Title(part)
	}

	name = strings.Join(nextParts, "")
	return name
}

func fieldGoType(field *desc.FieldDescriptorProto) (name string) {
	switch *field.Type {
	case desc.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64"
	case desc.FieldDescriptorProto_TYPE_FLOAT:
		return "float32"
	case desc.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case desc.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case desc.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case desc.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case desc.FieldDescriptorProto_TYPE_FIXED64:
		return "uint64"
	case desc.FieldDescriptorProto_TYPE_FIXED32:
		return "uint32"
	case desc.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case desc.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are unsupported")
	case desc.FieldDescriptorProto_TYPE_MESSAGE:
		return *field.TypeName
	case desc.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte"
	case desc.FieldDescriptorProto_TYPE_SFIXED32:
		return "int32"
	case desc.FieldDescriptorProto_TYPE_SFIXED64:
		return "int64"
	case desc.FieldDescriptorProto_TYPE_SINT32:
		return "int32"
	case desc.FieldDescriptorProto_TYPE_SINT64:
		return "int64"
	default:
		panic("unknown field type")
	}
}
