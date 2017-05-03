package generator

import (
	"path"
	"strconv"
	"strings"

	proto_desc "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/kixelated/protohack/codegen"
	desc "github.com/kixelated/protohack/descriptor"
)

type Generator struct {
	files *desc.FileSet
	w     codegen.Writer
}

func New(files *desc.FileSet) (g *Generator) {
	return &Generator{files: files}
}

func (g *Generator) File(inputName string) (name string, content string) {
	file := g.files.FindFile(inputName)
	if file == nil {
		panic("failed to find file: " + inputName)
	}

	g.writeFile(file)

	name = goFileName(file)
	content = g.w.String()

	g.w.Reset()

	return name, content
}

func goFileName(file *desc.File) (name string) {
	name = file.Proto.GetName()

	ext := path.Ext(name)
	if ext != ".proto" {
		panic("file does not end with .proto: " + name)
	}

	name = name[:len(name)-len(ext)]
	name = name + ".pb.go"

	return name
}

func (g *Generator) writeFile(file *desc.File) {
	g.w.Line(`package ` + goPackageName(file))

	g.writeImports(file)
	g.writeTypes(file)
}

func goPackageName(file *desc.File) (name string) {
	name = file.Proto.Options.GetGoPackage()
	if name != "" {
		return name
	}

	parts := strings.Split(file.Proto.GetPackage(), ".")
	name = parts[len(parts)-1]

	return name
}

func (g *Generator) writeImports(file *desc.File) {
}

func (g *Generator) writeTypes(file *desc.File) {
	for _, message := range file.Messages {
		g.writeMessageType(message)
	}

	for _, enum := range file.Enums {
		g.writeEnumType(enum)
		g.writeEnumValues(enum)
	}
}

func (g *Generator) writeMessageType(message *desc.Message) {
	g.w.Line(`type ` + goMessageType(message) + ` struct {`)
	g.w.In()

	for _, field := range message.Fields {
		g.writeFieldType(field)
	}

	g.w.Out()
	g.w.Line(`}`)

	for _, enum := range message.Enums {
		g.writeEnumType(enum)
		g.writeEnumValues(enum)
	}

	for _, message := range message.Messages {
		g.writeMessageType(message)
	}
}

func (g *Generator) writeEnumType(enum *desc.Enum) {
	g.w.Line(`type ` + goEnumType(enum) + ` int32`)
}

func (g *Generator) writeEnumValues(enum *desc.Enum) {
	g.w.Line(`const (`)
	g.w.In()

	enumType := goEnumType(enum)
	for _, enumValue := range enum.Values {
		enumValueType := goEnumValueType(enumValue)
		enumValueNumber := strconv.FormatInt(int64(enumValue.Proto.GetNumber()), 10)

		g.w.Line(enumValueType + ` ` + enumType + ` = ` + enumValueNumber)
	}

	g.w.Out()
	g.w.Line(`)`)
}

func (g *Generator) writeFieldType(field *desc.Field) {
	fieldName := goFieldName(field)
	fieldType := goFieldType(field)

	g.w.Line(fieldName + ` ` + fieldType)
}

func goMessageType(message *desc.Message) (name string) {
	_, _, messages := message.Parents()

	var parts []string
	for _, parent := range messages {
		parts = append(parts, parent.Proto.GetName())
	}

	parts = append(parts, message.Proto.GetName())

	name = strings.Join(parts, "_")
	return name
}

func goEnumType(enum *desc.Enum) (name string) {
	_, _, messages := enum.Parents()

	var parts []string
	for _, parent := range messages {
		parts = append(parts, parent.Proto.GetName())
	}

	parts = append(parts, enum.Proto.GetName())

	name = strings.Join(parts, "_")
	return name
}

func goEnumValueType(enumValue *desc.EnumValue) (name string) {
	name = enumValue.Proto.GetName()

	enum := enumValue.Parent
	if enum.ParentMessage != nil {
		name = goMessageType(enum.ParentMessage) + "_" + name
	}

	return name
}

func goFieldName(field *desc.Field) (name string) {
	parts := strings.Split(field.Proto.GetName(), "_")

	nextParts := make([]string, len(parts))
	for i, part := range parts {
		nextParts[i] = strings.Title(part)
	}

	name = strings.Join(nextParts, "")
	return name
}

func goFieldType(field *desc.Field) (name string) {
	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64"
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		return "float32"
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		return "uint64"
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		return "uint32"
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are unsupported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		message := field.MessageType()
		return "*" + goMessageType(message)
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte"
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		enum := field.EnumType()
		return "*" + goEnumType(enum)
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		return "int32"
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		return "int64"
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		return "int32"
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		return "int64"
	default:
		panic("unknown field type")
	}
}