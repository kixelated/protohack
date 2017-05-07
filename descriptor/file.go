package descriptor

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type File struct {
	Proto  *desc.FileDescriptorProto
	Parent *FileSet

	Messages []*Message
	Enums    []*Enum
	//Services   []*Service
	//Extensions []*Extension
}

func NewFile(proto *desc.FileDescriptorProto, parent *FileSet) (file *File) {
	file = new(File)
	file.Proto = proto
	file.Parent = parent

	for _, messageProto := range proto.GetMessageType() {
		message := NewMessage(messageProto, file)
		file.Messages = append(file.Messages, message)
	}

	for _, enumProto := range proto.GetEnumType() {
		enum := NewEnum(enumProto, file)
		file.Enums = append(file.Enums, enum)
	}

	return file
}

func (file *File) FindMessage(name string) (message *Message) {
	for _, message := range file.Messages {
		if message.FullName() == name {
			return message
		}

		message = message.FindMessage(name)
		if message != nil {
			return message
		}
	}

	return nil
}

func (file *File) FindEnum(name string) (enum *Enum) {
	for _, enum := range file.Enums {
		if enum.FullName() == name {
			return enum
		}
	}

	for _, message := range file.Messages {
		enum := message.FindEnum(name)
		if enum != nil {
			return enum
		}
	}

	return nil
}

func (file *File) FullName() (name string) {
	if file.Proto.Package != nil {
		return "." + file.Proto.GetPackage()
	} else {
		return ""
	}

}
