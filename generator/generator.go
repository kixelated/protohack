package generator

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	proto_desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/kixelated/protohack/proto"

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
	g.w.Line(`import "errors"`)

	// TODO
	g.w.Line(`var _ = proto.WIRETYPE_VARINT`)
	g.w.Line(`var _ = errors.New("")`)
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

func goElemType(field *desc.Field) (name string) {
	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		name = `float64`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		name = `float32`
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		name = `int64`
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		name = `uint64`
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		name = `int32`
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		name = `uint32`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		name = `uint64`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		name = `uint32`
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		name = `bool`
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		name = `string`
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are not supported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		message := field.MessageType()
		name = goMessageType(message)
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		name = `[]byte`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		enum := field.EnumType()
		name = goEnumType(enum)
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		name = `int32`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		name = `int64`
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		name = `int32`
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		name = `int64`
	default:
		panic("unknown element type")
	}

	return name
}

func goFieldType(field *desc.Field) (name string) {
	name = goElemType(field)

	if field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_MESSAGE {
		name = "*" + name
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
	g.w.Line(`_, err = m.MarshalTo(data)`)
	g.w.Line(`return data, err`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalTo(message *desc.Message) {
	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Proto.GetNumber() < fields[j].Proto.GetNumber()
	})

	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalTo(data []byte) (n int, err error) {`).In()

	for _, field := range fields {
		g.writeMessageMarshalToField(field)
	}

	g.w.Line(`return n, nil`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalToField(field *desc.Field) {
	repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
	required := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REQUIRED)
	// packed := field.Proto.GetOptions().GetPacked() // TODO implement

	var wireType proto.WireType
	var method string
	var zero string
	var cast string

	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		wireType = proto.WIRETYPE_64BIT
		method = `proto.WriteDouble`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		wireType = proto.WIRETYPE_32BIT
		method = `proto.WriteFloat`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteInt64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteUInt64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteInt32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteUInt32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		wireType = proto.WIRETYPE_64BIT
		method = `proto.WriteFixed64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		wireType = proto.WIRETYPE_32BIT
		method = `proto.WriteFixed32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteBool`
		zero = `false`
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		wireType = proto.WIRETYPE_LENGTH
		method = `proto.WriteStringLength`
		zero = `""`
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are not supported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		wireType = proto.WIRETYPE_LENGTH
		method = `proto.WriteMessageLength`
		zero = `nil`
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		wireType = proto.WIRETYPE_LENGTH
		method = `proto.WriteBytesLength`
		zero = `nil`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteEnum`
		cast = `int`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		wireType = proto.WIRETYPE_32BIT
		method = `proto.WriteSFixed32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		wireType = proto.WIRETYPE_64BIT
		method = `proto.WriteSFixed64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteSInt32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		wireType = proto.WIRETYPE_VARINT
		method = `proto.WriteSInt64`
		zero = `0`
	}

	name := `m.` + goFieldName(field)

	if repeated {
		g.w.Line(`for _, x := range ` + name + ` {`).In()
		name = `x`
	}

	if cast != "" {
		name = cast + `(` + name + `)`
	}

	if !repeated {
		g.w.Line(`if ` + name + ` != ` + zero + ` {`).In()
	}

	// TODO Make a better interface.
	keySize := proto.SizeKey(int(field.Proto.GetNumber()))
	keyData := make([]byte, keySize)
	proto.WriteKey(keyData, int(field.Proto.GetNumber()), wireType)

	g.w.Line(`n += copy(data[n:], ` + fmt.Sprintf("%#v", keyData) + `)`)

	g.w.Line(`n += ` + method + `(data[n:], ` + name + `)`)

	if !repeated && required {
		g.w.Out().Line(`} else {`).In()
		g.w.Line(`return n, errors.New("missing required field: ` + goFieldName(field) + `")`)
	}

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalSize(message *desc.Message) {
	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Proto.GetNumber() < fields[j].Proto.GetNumber()
	})

	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalSize() (n int) {`).In()

	for _, field := range fields {
		g.writeMessageMarshalSizeField(field)
	}

	g.w.Line(`return n`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalSizeField(field *desc.Field) {
	repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
	required := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REQUIRED)
	// TODO
	// packed := field.Proto.GetOptions().GetPacked()

	var method string
	var zero string
	var cast string

	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		method = `proto.SizeDouble`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		method = `proto.SizeFloat`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		method = `proto.SizeInt64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		method = `proto.SizeUInt64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		method = `proto.SizeInt32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		method = `proto.SizeUInt32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		method = `proto.SizeFixed64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		method = `proto.SizeFixed32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		method = `proto.SizeBool`
		zero = `false`
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		method = `proto.SizeStringLength`
		zero = `""`
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are not supported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		method = `proto.SizeMessageLength`
		zero = `nil`
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		method = `proto.SizeBytesLength`
		zero = `nil`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		method = `proto.SizeEnum`
		cast = `int`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		method = `proto.SizeSFixed32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		method = `proto.SizeSFixed64`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		method = `proto.SizeSInt32`
		zero = `0`
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		method = `proto.SizeSInt64`
		zero = `0`
	}

	name := `m.` + goFieldName(field)

	if repeated {
		g.w.Line(`for _, x := range ` + name + ` {`).In()
		name = `x`
	}

	if cast != "" {
		name = cast + `(` + name + `)`
	}

	if !required && !repeated {
		g.w.Line(`if ` + name + ` != ` + zero + ` {`).In()
	}

	size := proto.SizeKey(int(field.Proto.GetNumber()))
	sizeStr := strconv.FormatInt(int64(size), 10)

	g.w.Line(`n += ` + sizeStr + ` + ` + method + `(` + name + `)`)

	if !required && !repeated {
		g.w.Out().Line(`}`)
	}

	if repeated {
		g.w.Out().Line(`}`)
	}
}

func (g *Generator) writeMessageUnmarshal(message *desc.Message) {
	g.w.Line(`func (m *` + goMessageType(message) + `) Unmarshal(data []byte) (err error) {`).In()

	g.w.Line(`r := proto.NewReader(data)`)
	g.w.Line(`for r.Len() > 0 {`).In()

	g.w.Line(`id, _, err := r.ReadKey()`)
	g.w.Line(`if err != nil {`).In()
	g.w.Line(`return err`)
	g.w.Out().Line(`}`)

	g.w.Line(`switch id {`)

	for _, field := range message.Fields {
		g.writeMessageUnmarshalField(field)
	}

	g.w.Line(`}`)
	g.w.Out().Line(`}`)

	g.w.Line(`return nil`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageUnmarshalField(field *desc.Field) {
	repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
	message := (field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_MESSAGE)

	numberStr := strconv.Itoa(int(field.Proto.GetNumber()))
	g.w.Line(`case ` + numberStr + `:`).In()

	var method string
	var cast string

	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		method = `ReadDouble`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		method = `ReadFloat`
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		method = `ReadInt64`
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		method = `ReadUInt64`
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		method = `ReadInt32`
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		method = `ReadUInt32`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		method = `ReadFixed64`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		method = `ReadFixed32`
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		method = `ReadBool`
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		method = `ReadString`
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are not supported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		method = `ReadToMessage`
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		method = `ReadBytes`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		method = `ReadEnum`
		cast = goElemType(field)
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		method = `ReadSFixed32`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		method = `ReadSFixed64`
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		method = `ReadSInt32`
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		method = `ReadSInt64`
	}

	name := `m.` + goFieldName(field)
	tempName := name

	if repeated || cast != "" {
		tempName = `temp`
	}

	if message {
		if tempName != name {
			g.w.Line(tempName + ` := new(` + goElemType(field) + `)`)
			g.w.Line(`err = r.` + method + `(` + tempName + `)`)
		} else {
			g.w.Line(name + ` = new(` + goElemType(field) + `)`)
			g.w.Line(`err = r.` + method + `(` + name + `)`)
		}
	} else {
		if tempName != name {
			g.w.Line(tempName + `, err := r.` + method + `()`)
		} else {
			g.w.Line(name + `, err = r.` + method + `()`)
		}
	}

	g.w.Line(`if err != nil {`).In()
	g.w.Line(`return err`)
	g.w.Out().Line(`}`)

	if cast != "" {
		tempName = cast + `(` + tempName + `)`
	}

	if repeated {
		g.w.Line(name + ` = append(` + name + `, ` + tempName + `)`)
	} else if tempName != name {
		g.w.Line(name + ` = ` + tempName)
	}

	g.w.Out()
}
