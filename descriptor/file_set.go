package descriptor

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type FileSet struct {
	Proto *desc.FileDescriptorSet

	Files []*File
}

func NewFileSet(proto *desc.FileDescriptorSet) (fs *FileSet) {
	fs = new(FileSet)
	fs.Proto = proto

	for _, fileProto := range proto.GetFile() {
		file := NewFile(fileProto, fs)
		fs.Files = append(fs.Files, file)
	}

	return fs
}

func (fs *FileSet) FindFile(name string) (file *File) {
	for _, file := range fs.Files {
		if file.Proto.GetName() == name {
			return file
		}
	}

	return nil
}

func (fs *FileSet) FindMessage(name string) (message *Message) {
	for _, file := range fs.Files {
		message = file.FindMessage(name)
		if message != nil {
			return message
		}
	}

	return nil
}

func (fs *FileSet) FindEnum(name string) (enum *Enum) {
	for _, file := range fs.Files {
		enum = file.FindEnum(name)
		if enum != nil {
			return enum
		}
	}

	return nil
}
