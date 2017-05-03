package descriptor

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Message struct {
	Proto *desc.DescriptorProto

	// Only one of the two will be set.
	ParentFile    *File
	ParentMessage *Message

	Fields   []*Field
	Messages []*Message
	Enums    []*Enum
}

func NewMessage(proto *desc.DescriptorProto, parent interface{}) (message *Message) {
	message = new(Message)
	message.Proto = proto

	switch realParent := parent.(type) {
	case *File:
		message.ParentFile = realParent
	case *Message:
		message.ParentMessage = realParent
	default:
		panic("unknown parent type")
	}

	for _, fieldProto := range proto.GetField() {
		field := NewField(fieldProto, message)
		message.Fields = append(message.Fields, field)
	}

	for _, messageProto := range proto.GetNestedType() {
		subMessage := NewMessage(messageProto, message)
		message.Messages = append(message.Messages, subMessage)
	}

	for _, enumProto := range proto.GetEnumType() {
		enum := NewEnum(enumProto, message)
		message.Enums = append(message.Enums, enum)
	}

	return message
}

func (m *Message) Parents() (set *FileSet, file *File, messages []*Message) {
	if m.ParentFile != nil {
		return m.ParentFile.Parent, m.ParentFile, nil
	}

	set, file, messages = m.ParentMessage.Parents()
	messages = append(messages, m.ParentMessage)

	return set, file, messages
}

func (m *Message) FindMessage(name string) (message *Message) {
	for _, subMessage := range m.Messages {
		if subMessage.FullName() == name {
			return subMessage
		}

		message = subMessage.FindMessage(name)
		if message != nil {
			return message
		}
	}

	return nil
}

func (m *Message) FindEnum(name string) (enum *Enum) {
	for _, enum := range m.Enums {
		if enum.FullName() == name {
			return enum
		}
	}

	for _, message := range m.Messages {
		enum = message.FindEnum(name)
		if enum != nil {
			return enum
		}
	}

	return nil
}

func (m *Message) FullName() (name string) {
	if m.ParentFile != nil {
		return m.ParentFile.FullName() + "." + m.Proto.GetName()
	} else {
		return m.ParentMessage.FullName() + "." + m.Proto.GetName()
	}
}
