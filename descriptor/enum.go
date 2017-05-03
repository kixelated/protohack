package descriptor

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Enum struct {
	Proto *desc.EnumDescriptorProto

	// Only one of the two will be set.
	ParentFile    *File
	ParentMessage *Message
}

func NewEnum(proto *desc.EnumDescriptorProto, parent interface{}) (enum *Enum) {
	enum = new(Enum)
	enum.Proto = proto

	switch realParent := parent.(type) {
	case *File:
		enum.ParentFile = realParent
	case *Message:
		enum.ParentMessage = realParent
	default:
		panic("unknown parent type")
	}

	return enum
}

func (enum *Enum) Parents() (set *FileSet, file *File, messages []*Message) {
	if enum.ParentFile != nil {
		return enum.ParentFile.Parent, enum.ParentFile, nil
	}

	set, file, messages = enum.ParentMessage.Parents()
	messages = append(messages, enum.ParentMessage)

	return set, file, messages
}

func (enum *Enum) FullName() (name string) {
	if enum.ParentFile != nil {
		return enum.ParentFile.FullName() + "." + enum.Proto.GetName()
	} else {
		return enum.ParentMessage.FullName() + "." + enum.Proto.GetName()
	}
}
