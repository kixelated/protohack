package generator

import (
	"path"
	"sort"
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
	g.writeMethods(file)
}

func goPackageName(file *desc.File) (name string) {
	name = file.Proto.Options.GetGoPackage()
	if name != "" {
		return name
	}

	name = file.Proto.GetPackage()
	if name != "" {
		parts := strings.Split(name, ".")
		name = parts[len(parts)-1]
		return name
	}

	panic("missing package name")
}

func (g *Generator) writeImports(file *desc.File) {
	g.w.Line(`import "github.com/kixelated/protohack/proto"`)

	// TODO
	g.w.Line(`var _ = proto.WIRETYPE_VARINT`)
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
	g.w.Line(`type ` + goEnumType(enum) + ` int`)
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
		name = "float64"
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		name = "float32"
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		name = "int64"
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		name = "uint64"
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		name = "int32"
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		name = "uint32"
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		name = "uint64"
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		name = "uint32"
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		name = "bool"
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		name = "string"
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		message := field.MessageType()
		name = "*" + goMessageType(message)
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		message := field.MessageType()
		name = "*" + goMessageType(message)
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		name = "[]byte"
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		enum := field.EnumType()
		name = goEnumType(enum)
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		name = "int32"
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		name = "int64"
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		name = "int32"
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		name = "int64"
	default:
		panic("unknown field type")
	}

	if field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED {
		name = "[]" + name
	}

	return name
}

func (g *Generator) writeMethods(file *desc.File) {
	for _, message := range file.Messages {
		g.writeMessageMethods(message)
	}
}

func (g *Generator) writeMessageMethods(message *desc.Message) {
	g.writeMessageMarshal(message)
	g.writeMessageMarshalTo(message)
	g.writeMessageMarshalSize(message)
	g.writeMessageUnmarshal(message)

	for _, nested := range message.Messages {
		g.writeMessageMethods(nested)
	}
}

func (g *Generator) writeMessageMarshal(message *desc.Message) {
	g.w.Line(`func (m ` + goMessageType(message) + `) Marshal() (data []byte, err error) {`).In()

	g.w.Line(`data = make([]byte, m.MarshalSize())`)
	g.w.Line(`_ = m.MarshalTo(data)`)
	g.w.Line(`return data, nil`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalTo(message *desc.Message) {
	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalTo(data []byte) (n int) {`).In()

	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Proto.GetNumber() < fields[j].Proto.GetNumber()
	})

	for _, field := range fields {
		varName := `m.` + goFieldName(field)
		id := strconv.FormatInt(int64(field.Proto.GetNumber()), 10)

		repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
		// TODO
		// packed := field.Proto.GetOptions().GetPacked()

		if repeated {
			g.w.Line(`for _, x := range ` + varName + ` {`).In()
			varName = `x`
		}

		switch field.Proto.GetType() {
		case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
			g.w.Line(`n += proto.WriteFieldDouble(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
			g.w.Line(`n += proto.WriteFieldFloat(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_INT64:
			g.w.Line(`n += proto.WriteFieldInt64(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_UINT64:
			g.w.Line(`n += proto.WriteFieldUInt64(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_INT32:
			g.w.Line(`n += proto.WriteFieldInt32(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_UINT32:
			g.w.Line(`n += proto.WriteFieldUInt32(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
			g.w.Line(`n += proto.WriteFieldFixed64(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
			g.w.Line(`n += proto.WriteFieldFixed32(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_BOOL:
			g.w.Line(`n += proto.WriteFieldBool(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_STRING:
			g.w.Line(`n += proto.WriteFieldString(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_GROUP:
			g.w.Line(`if ` + varName + ` != nil {`).In()
			g.w.Line(`n += proto.WriteFieldGroup(data[n:], ` + id + `, ` + varName + `)`)
			g.w.Out().Line(`}`)
		case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
			g.w.Line(`if ` + varName + ` != nil {`).In()
			g.w.Line(`n += proto.WriteFieldMessage(data[n:], ` + id + `, ` + varName + `)`)
			g.w.Out().Line(`}`)
		case proto_desc.FieldDescriptorProto_TYPE_BYTES:
			g.w.Line(`n += proto.WriteFieldBytes(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_ENUM:
			g.w.Line(`n += proto.WriteFieldEnum(data[n:], ` + id + `, int(` + varName + `))`)
		case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
			g.w.Line(`n += proto.WriteFieldSFixed32(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
			g.w.Line(`n += proto.WriteFieldSFixed64(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_SINT32:
			g.w.Line(`n += proto.WriteFieldSInt32(data[n:], ` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_SINT64:
			g.w.Line(`n += proto.WriteFieldSInt64(data[n:], ` + id + `, ` + varName + `)`)
		}

		if repeated {
			g.w.Out().Line(`}`)
		}
	}

	g.w.Line(`return n`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalSize(message *desc.Message) {
	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalSize() (n int) {`).In()

	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Proto.GetNumber() < fields[j].Proto.GetNumber()
	})

	for _, field := range fields {
		varName := `m.` + goFieldName(field)
		id := strconv.FormatInt(int64(field.Proto.GetNumber()), 10)

		repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
		// TODO
		// packed := field.Proto.GetOptions().GetPacked()

		if repeated {
			g.w.Line(`for _, x := range ` + varName + ` {`).In()
			varName = `x`
		}

		switch field.Proto.GetType() {
		case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
			g.w.Line(`n += proto.SizeFieldDouble(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
			g.w.Line(`n += proto.SizeFieldFloat(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_INT64:
			g.w.Line(`n += proto.SizeFieldInt64(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_UINT64:
			g.w.Line(`n += proto.SizeFieldUInt64(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_INT32:
			g.w.Line(`n += proto.SizeFieldInt32(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_UINT32:
			g.w.Line(`n += proto.SizeFieldUInt32(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
			g.w.Line(`n += proto.SizeFieldFixed64(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
			g.w.Line(`n += proto.SizeFieldFixed32(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_BOOL:
			g.w.Line(`n += proto.SizeFieldBool(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_STRING:
			g.w.Line(`n += proto.SizeFieldString(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_GROUP:
			g.w.Line(`if ` + varName + ` != nil {`).In()
			g.w.Line(`n += proto.SizeFieldGroup(` + id + `, ` + varName + `)`)
			g.w.Out().Line(`}`)
		case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
			g.w.Line(`if ` + varName + ` != nil {`).In()
			g.w.Line(`n += proto.SizeFieldMessage(` + id + `, ` + varName + `)`)
			g.w.Out().Line(`}`)
		case proto_desc.FieldDescriptorProto_TYPE_BYTES:
			g.w.Line(`n += proto.SizeFieldBytes(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_ENUM:
			g.w.Line(`n += proto.SizeFieldEnum(` + id + `, int(` + varName + `))`)
		case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
			g.w.Line(`n += proto.SizeFieldSFixed32(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
			g.w.Line(`n += proto.SizeFieldSFixed64(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_SINT32:
			g.w.Line(`n += proto.SizeFieldSInt32(` + id + `, ` + varName + `)`)
		case proto_desc.FieldDescriptorProto_TYPE_SINT64:
			g.w.Line(`n += proto.SizeFieldSInt64(` + id + `, ` + varName + `)`)
		}

		if repeated {
			g.w.Out().Line(`}`)
		}
	}

	g.w.Line(`return n`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageUnmarshal(message *desc.Message) {

}
