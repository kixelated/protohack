package generator

import (
	"errors"
	"path"
	"strings"

	desc "github.com/kixelated/protohack/descriptor"

	"github.com/kixelated/protohack/codegen"
)

type Generator struct {
	files []*desc.FileDescriptorProto
}

func New(files []*desc.FileDescriptorProto) (g *Generator) {
	return &Generator{files: files}
}

func (g *Generator) File(inputName string) (name string, content string, err error) {
	proto, err := g.lookupFile(inputName)
	if err != nil {
		return "", "", err
	}

	name, err = g.goFileName(proto)
	if err != nil {
		return "", "", err
	}

	w := new(codegen.Writer)
	err = g.writeFile(w, proto)
	if err != nil {
		return "", "", err
	}

	content = w.String()

	return name, content, nil
}

func (g *Generator) goFileName(file *desc.FileDescriptorProto) (name string, err error) {
	name = file.GetName()

	ext := path.Ext(name)
	if ext != ".proto" {
		err = errors.New("file does not end with .proto")
		return "", err
	}

	name = name[:len(name)-len(ext)]
	name = name + ".pb.go"
	return name, nil
}

func (g *Generator) writeFile(w *codegen.Writer, file *desc.FileDescriptorProto) (err error) {
	w.Comment(`PROTO name: ` + file.GetName())
	w.Comment(`PROTO package: ` + file.GetPackage())

	w.Line(`package ` + g.goPackageName(file))

	g.writeImports(w, file)
	g.writeEnums(w, file)

	err = g.writeTypes(w, file)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) goPackageName(file *desc.FileDescriptorProto) (name string) {
	name = file.Options.GetGoPackage()
	if name != "" {
		return name
	}

	parts := strings.Split(file.GetPackage(), ".")
	name = parts[len(parts)-1]

	return name
}

func (g *Generator) writeImports(w *codegen.Writer, file *desc.FileDescriptorProto) {
	for _, dependency := range file.GetDependency() {
		w.Comment(`PROTO dependency: ` + dependency)
	}

	for _, publicDependency := range file.GetPublicDependency() {
		w.Comment(`PROTO public dependency: ` + file.GetDependency()[publicDependency])
	}

	w.Comment(`@@protoc_insertion_point(import)`)
}

func (g *Generator) writeEnums(w *codegen.Writer, file *desc.FileDescriptorProto) {
	w.Comment(`@@protoc_insertion_point(enum)`)
}

func (g *Generator) writeTypes(w *codegen.Writer, file *desc.FileDescriptorProto) (err error) {
	for _, messageType := range file.MessageType {
		err = g.writeMessageType(w, messageType)
		if err != nil {
			return err
		}
	}

	w.Comment(`@@protoc_insertion_point(type)`)
	return nil
}

func (g *Generator) writeMessageType(w *codegen.Writer, message *desc.DescriptorProto) (err error) {
	w.Comment(`PROTO name: ` + message.GetName())

	w.Line(`type ` + g.goMessageType(message) + ` struct {`)
	w.In()

	for _, field := range message.Field {
		err = g.writeFieldType(w, field)
		if err != nil {
			return err
		}
	}

	w.Comment(`@@protoc_insertion_point(type_` + g.goMessageType(message) + `)`)

	w.Out()
	w.Line(`}`)

	return nil
}

func (g *Generator) writeFieldType(w *codegen.Writer, field *desc.FieldDescriptorProto) (err error) {
	w.Comment(`PROTO name: ` + field.GetName())
	w.Comment(`PROTO label: ` + field.GetLabel().String())
	w.Comment(`PROTO type: ` + field.GetType().String())

	if field.TypeName != nil {
		w.Comment(`PROTO type_name: ` + field.GetTypeName())
	}

	fieldName := g.goFieldName(field)
	fieldType, err := g.goFieldType(field)
	if err != nil {
		return err
	}

	w.Line(fieldName + ` ` + fieldType)
	return nil
}

func (g *Generator) goMessageType(message *desc.DescriptorProto) (name string) {
	return message.GetName()
}

func (g *Generator) goFieldName(field *desc.FieldDescriptorProto) (name string) {
	parts := strings.Split(field.GetName(), "_")

	nextParts := make([]string, len(parts))
	for i, part := range parts {
		nextParts[i] = strings.Title(part)
	}

	name = strings.Join(nextParts, "")
	return name
}

func (g *Generator) goFieldType(field *desc.FieldDescriptorProto) (name string, err error) {
	switch field.GetType() {
	case desc.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64", nil
	case desc.FieldDescriptorProto_TYPE_FLOAT:
		return "float32", nil
	case desc.FieldDescriptorProto_TYPE_INT64:
		return "int64", nil
	case desc.FieldDescriptorProto_TYPE_UINT64:
		return "uint64", nil
	case desc.FieldDescriptorProto_TYPE_INT32:
		return "int32", nil
	case desc.FieldDescriptorProto_TYPE_UINT32:
		return "uint32", nil
	case desc.FieldDescriptorProto_TYPE_FIXED64:
		return "uint64", nil
	case desc.FieldDescriptorProto_TYPE_FIXED32:
		return "uint32", nil
	case desc.FieldDescriptorProto_TYPE_BOOL:
		return "bool", nil
	case desc.FieldDescriptorProto_TYPE_STRING:
		return "string", nil
	case desc.FieldDescriptorProto_TYPE_GROUP:
		err = errors.New("groups are unsupported")
		return "", err
	case desc.FieldDescriptorProto_TYPE_MESSAGE:
		message := g.lookupMessage(field.GetTypeName())
		if message == nil {
			panic("could not find field type: " + field.GetTypeName())
		}

		name = "*" + g.goMessageType(message)
		return name, nil
	case desc.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte", nil
	case desc.FieldDescriptorProto_TYPE_ENUM:
		// TODO
		return field.GetTypeName(), nil
	case desc.FieldDescriptorProto_TYPE_SFIXED32:
		return "int32", nil
	case desc.FieldDescriptorProto_TYPE_SFIXED64:
		return "int64", nil
	case desc.FieldDescriptorProto_TYPE_SINT32:
		return "int32", nil
	case desc.FieldDescriptorProto_TYPE_SINT64:
		return "int64", nil
	default:
		err = errors.New("unknown field type")
		return "", err
	}
}
