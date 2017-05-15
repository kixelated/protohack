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
	g.writeMessageMarshalStackSize(message)
	g.writeMessageUnmarshal(message)

	for _, nested := range message.Messages {
		g.writeMessageMethods(nested)
	}
}

func (g *Generator) writeMessageMarshal(message *desc.Message) {
	g.w.Line(`func (m ` + goMessageType(message) + `) Marshal() (data []byte, err error) {`).In()

	g.w.Line(`var s proto.Sizer`)
	g.w.Line(`s.Grow(m.MarshalStackSize())`)
	g.w.Line(`m.MarshalSize(&s)`)
	g.w.Line()
	g.w.Line(`var w proto.Writer`)
	g.w.Line(`w.WithSizer(&s)`)
	g.w.Line(`err = m.MarshalTo(&w)`)
	g.w.Line()
	g.w.Line(`return w.Data(), err`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalTo(message *desc.Message) {
	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Proto.GetNumber() < fields[j].Proto.GetNumber()
	})

	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalTo(w *proto.Writer) (err error) {`).In()

	for _, field := range fields {
		g.writeMessageMarshalToField(field)
	}

	g.w.Line(`return nil`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalToField(field *desc.Field) {
	repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
	message := (field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_MESSAGE)
	packed := field.Proto.GetOptions().GetPacked()

	var wireType proto.WireType
	var method string
	var cast string

	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		wireType = proto.WIRETYPE_64BIT
		method = `Float64`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		wireType = proto.WIRETYPE_32BIT
		method = `Float32`
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		wireType = proto.WIRETYPE_VARINT
		method = `Int64`
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		wireType = proto.WIRETYPE_VARINT
		method = `UInt64`
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		wireType = proto.WIRETYPE_VARINT
		method = `Int32`
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		wireType = proto.WIRETYPE_VARINT
		method = `UInt32`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		wireType = proto.WIRETYPE_64BIT
		method = `FixedUInt64`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		wireType = proto.WIRETYPE_32BIT
		method = `FixedUInt32`
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		wireType = proto.WIRETYPE_VARINT
		method = `Bool`
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		wireType = proto.WIRETYPE_LENGTH
		method = `String`
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are not supported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		wireType = proto.WIRETYPE_LENGTH
		method = `Message`
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		wireType = proto.WIRETYPE_LENGTH
		method = `Bytes`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		wireType = proto.WIRETYPE_VARINT
		method = `Int`
		cast = `int`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		wireType = proto.WIRETYPE_32BIT
		method = `FixedInt32`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		wireType = proto.WIRETYPE_64BIT
		method = `FixedInt64`
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		wireType = proto.WIRETYPE_VARINT
		method = `SInt32`
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		wireType = proto.WIRETYPE_VARINT
		method = `SInt64`
	}

	name := `m.` + goFieldName(field)

	if packed {
		method = method + `Packed`
		wireType = proto.WIRETYPE_LENGTH
	} else if repeated {
		method = method + `Repeated`
	} else {
		method = method + `Field`
	}

	if cast != "" {
		name = cast + `(` + name + `)`
	}

	tag := proto.EncodeTag(int(field.Proto.GetNumber()), wireType)
	tagStr := fmt.Sprintf("%#v", tag)

	if repeated && message {
		g.w.Line(`for _, x := range ` + name + `{`).In()
		g.w.Line(`w.MessageField(` + tagStr + `, x)`)
		g.w.Out().Line(`}`)
	} else if message {
		g.w.Line(`if ` + name + ` != nil {`).In()
		g.w.Line(`w.MessageField(` + tagStr + `, ` + name + `)`)
		g.w.Out().Line(`}`)
	} else {
		g.w.Line(`w.` + method + `(` + tagStr + `, ` + name + `)`)
	}
}

func (g *Generator) writeMessageMarshalSize(message *desc.Message) {
	fields := message.Fields
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Proto.GetNumber() < fields[j].Proto.GetNumber()
	})

	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalSize(s *proto.Sizer) {`).In()

	for _, field := range fields {
		g.writeMessageMarshalSizeField(field)
	}

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageMarshalSizeField(field *desc.Field) {
	repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
	message := (field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_MESSAGE)
	packed := field.Proto.GetOptions().GetPacked()

	var method string
	var cast string

	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		method = `Float64`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		method = `Float32`
	case proto_desc.FieldDescriptorProto_TYPE_INT64:
		method = `Int64`
	case proto_desc.FieldDescriptorProto_TYPE_UINT64:
		method = `UInt64`
	case proto_desc.FieldDescriptorProto_TYPE_INT32:
		method = `Int32`
	case proto_desc.FieldDescriptorProto_TYPE_UINT32:
		method = `UInt32`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED64:
		method = `FixedUInt64`
	case proto_desc.FieldDescriptorProto_TYPE_FIXED32:
		method = `FixedUInt32`
	case proto_desc.FieldDescriptorProto_TYPE_BOOL:
		method = `Bool`
	case proto_desc.FieldDescriptorProto_TYPE_STRING:
		method = `String`
	case proto_desc.FieldDescriptorProto_TYPE_GROUP:
		panic("groups are not supported")
	case proto_desc.FieldDescriptorProto_TYPE_MESSAGE:
		method = `Message`
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		method = `Bytes`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		method = `Int`
		cast = `int`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED32:
		method = `FixedInt32`
	case proto_desc.FieldDescriptorProto_TYPE_SFIXED64:
		method = `FixedInt64`
	case proto_desc.FieldDescriptorProto_TYPE_SINT32:
		method = `Int32`
	case proto_desc.FieldDescriptorProto_TYPE_SINT64:
		method = `Int64`
	}

	name := `m.` + goFieldName(field)

	if cast != "" {
		name = cast + `(` + name + `)`
	}

	if packed {
		method = method + `Packed`
	} else if repeated {
		method = method + `Repeated`
	} else {
		method = method + `Field`
	}

	tag := proto.EncodeTag(int(field.Proto.GetNumber()), 0)
	tagSize := len(tag)
	tagSizeStr := fmt.Sprintf("%d", tagSize)

	if repeated && message {
		g.w.Line(`for _, x := range ` + name + `{`).In()
		g.w.Line(`s.MessageField(` + tagSizeStr + `, x)`)
		g.w.Out().Line(`}`)
	} else if message {
		g.w.Line(`if ` + name + ` != nil {`).In()
		g.w.Line(`s.MessageField(` + tagSizeStr + `, ` + name + `)`)
		g.w.Out().Line(`}`)
	} else {
		g.w.Line(`s.` + method + `(` + tagSizeStr + `, ` + name + `)`)
	}
}

func (g *Generator) writeMessageMarshalStackSize(message *desc.Message) {
	g.w.Line(`func (m ` + goMessageType(message) + `) MarshalStackSize() (n int) {`).In()

	var n int

	for _, field := range message.Fields {
		repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
		message := (field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_MESSAGE)
		packed := field.Proto.GetOptions().GetPacked()

		if repeated && packed {
			// TODO Ignore fixed types
			n += 1
		} else if repeated && message {
			name := `m.` + goFieldName(field)
			g.w.Line(`for _, x := range ` + name + ` {`).In()
			g.w.Line(`n += 1 + x.MarshalStackSize()`)
			g.w.Out().Line(`}`)
		} else if message {
			n += 1
		}
	}

	if n > 0 {
		g.w.Line(`n += ` + fmt.Sprintf("%d", n))
	}

	g.w.Line(`return n`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageUnmarshal(message *desc.Message) {
	g.w.Line(`func (m *` + goMessageType(message) + `) Unmarshal(data []byte) (err error) {`).In()

	g.w.Line(`r := proto.NewReader(data)`)
	g.w.Line(`for r.Len() > 0 {`).In()

	g.w.Line(`id, t, err := r.ReadKey()`)
	g.w.Line(`if err != nil {`).In()
	g.w.Line(`return err`)
	g.w.Out().Line(`}`)

	g.w.Line(`switch id {`)

	for _, field := range message.Fields {
		g.writeMessageUnmarshalField(field)
	}

	g.w.Line(`}`)

	g.w.Line(`if err != nil {`).In()
	g.w.Line(`return err`)
	g.w.Out().Line(`}`)

	g.w.Out().Line(`}`)

	g.w.Line(`return nil`)

	g.w.Out().Line(`}`)
}

func (g *Generator) writeMessageUnmarshalField(field *desc.Field) {
	repeated := (field.Proto.GetLabel() == proto_desc.FieldDescriptorProto_LABEL_REPEATED)
	message := (field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_MESSAGE)
	enum := (field.Proto.GetType() == proto_desc.FieldDescriptorProto_TYPE_ENUM)

	numberStr := strconv.Itoa(int(field.Proto.GetNumber()))
	g.w.Line(`case ` + numberStr + `:`).In()

	var method string

	switch field.Proto.GetType() {
	case proto_desc.FieldDescriptorProto_TYPE_DOUBLE:
		method = `ReadFloat64`
	case proto_desc.FieldDescriptorProto_TYPE_FLOAT:
		method = `ReadFloat32`
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
	case proto_desc.FieldDescriptorProto_TYPE_BYTES:
		method = `ReadBytes`
	case proto_desc.FieldDescriptorProto_TYPE_ENUM:
		method = `ReadInt`
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

	if message {
		if repeated {
			g.w.Line(`x := new(` + goElemType(field) + `)`)
			g.w.Line(`err = r.ReadMessage(t, x)`)
			g.w.Line(`if err == nil {`).In()
			g.w.Line(name + ` = append(` + name + `, x)`)
			g.w.Out().Line(`}`)
			// We check err later.
		} else {
			g.w.Line(name + ` = new(` + goElemType(field) + `)`)
			g.w.Line(`err = r.ReadMessage(t, ` + name + `)`)
		}
	} else if enum {
		if repeated {
			g.w.Line(`x, err := r.ReadIntRepeated(t, nil)`)
			g.w.Line(`if err != nil {`).In()
			g.w.Line(`return err`)
			g.w.Out().Line(`}`)

			g.w.Line(`for _, y := range x {`).In()
			g.w.Line(name + ` = append(` + name + `, ` + goElemType(field) + `(y))`)
			g.w.Out().Line(`}`)
		} else {
			g.w.Line(`x, err := r.ReadInt(t)`)
			g.w.Line(`if err != nil {`).In()
			g.w.Line(`return err`)
			g.w.Out().Line(`}`)

			g.w.Line(name + ` = ` + goElemType(field) + `(x)`)
		}
	} else {
		if repeated {
			g.w.Line(name + `, err = r.` + method + `Repeated(t, ` + name + `)`)
		} else {
			g.w.Line(name + `, err = r.` + method + `(t)`)
		}
	}

	g.w.Out()
}
