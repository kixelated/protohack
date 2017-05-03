package descriptor

import (
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type EnumValue struct {
	Proto  *desc.EnumValueDescriptorProto
	Parent *Enum
}

func NewEnumValue(proto *desc.EnumValueDescriptorProto, parent *Enum) (ev *EnumValue) {
	ev = new(EnumValue)
	ev.Proto = proto
	ev.Parent = parent

	return ev
}

func (ev *EnumValue) FullName() (name string) {
	return ev.Parent.FullName() + "." + ev.Proto.GetName()
}
