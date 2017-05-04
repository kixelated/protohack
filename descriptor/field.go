package descriptor

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Field struct {
	Proto  *desc.FieldDescriptorProto
	Parent *Message
}

func NewField(proto *desc.FieldDescriptorProto, parent *Message) (message *Field) {
	message = new(Field)
	message.Proto = proto
	message.Parent = parent

	return message
}

func (field *Field) Parents() (set *FileSet, file *File, messages []*Message) {
	set, file, messages = field.Parent.Parents()
	messages = append(messages, field.Parent)

	return set, file, messages
}

func (field *Field) MessageType() (message *Message) {
	switch field.Proto.GetType() {
	case desc.FieldDescriptorProto_TYPE_MESSAGE:
	case desc.FieldDescriptorProto_TYPE_GROUP:
	default:
		panic("field is not a message type")
	}

	set, _, _ := field.Parents()

	name := field.Proto.GetTypeName()
	message = set.FindMessage(name)
	if message == nil {
		panic("failed to find message type: " + name)
	}

	return message
}

func (field *Field) EnumType() (enum *Enum) {
	if field.Proto.GetType() != desc.FieldDescriptorProto_TYPE_ENUM {
		panic("field is not an enum type")
	}

	set, _, _ := field.Parents()

	name := field.Proto.GetTypeName()
	enum = set.FindEnum(name)
	if enum == nil {
		panic("failed to find enum type: " + name)
	}

	return enum
}
