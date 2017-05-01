package descriptor
// @@protoc_insertion_point(import)
// @@protoc_insertion_point(enum)
type FileDescriptorSet struct {
	File .google.protobuf.FileDescriptorProto
	// @@protoc_insertion_point(type_FileDescriptorSet)
}
type FileDescriptorProto struct {
	Name string
	Package string
	Dependency string
	PublicDependency int32
	WeakDependency int32
	MessageType .google.protobuf.DescriptorProto
	EnumType .google.protobuf.EnumDescriptorProto
	Service .google.protobuf.ServiceDescriptorProto
	Extension .google.protobuf.FieldDescriptorProto
	Options .google.protobuf.FileOptions
	SourceCodeInfo .google.protobuf.SourceCodeInfo
	Syntax string
	// @@protoc_insertion_point(type_FileDescriptorProto)
}
type DescriptorProto struct {
	Name string
	Field .google.protobuf.FieldDescriptorProto
	Extension .google.protobuf.FieldDescriptorProto
	NestedType .google.protobuf.DescriptorProto
	EnumType .google.protobuf.EnumDescriptorProto
	ExtensionRange .google.protobuf.DescriptorProto.ExtensionRange
	OneofDecl .google.protobuf.OneofDescriptorProto
	Options .google.protobuf.MessageOptions
	ReservedRange .google.protobuf.DescriptorProto.ReservedRange
	ReservedName string
	// @@protoc_insertion_point(type_DescriptorProto)
}
type FieldDescriptorProto struct {
	Name string
	Number int32
	Label .google.protobuf.FieldDescriptorProto.Label
	Type .google.protobuf.FieldDescriptorProto.Type
	TypeName string
	Extendee string
	DefaultValue string
	OneofIndex int32
	JsonName string
	Options .google.protobuf.FieldOptions
	// @@protoc_insertion_point(type_FieldDescriptorProto)
}
type OneofDescriptorProto struct {
	Name string
	Options .google.protobuf.OneofOptions
	// @@protoc_insertion_point(type_OneofDescriptorProto)
}
type EnumDescriptorProto struct {
	Name string
	Value .google.protobuf.EnumValueDescriptorProto
	Options .google.protobuf.EnumOptions
	// @@protoc_insertion_point(type_EnumDescriptorProto)
}
type EnumValueDescriptorProto struct {
	Name string
	Number int32
	Options .google.protobuf.EnumValueOptions
	// @@protoc_insertion_point(type_EnumValueDescriptorProto)
}
type ServiceDescriptorProto struct {
	Name string
	Method .google.protobuf.MethodDescriptorProto
	Options .google.protobuf.ServiceOptions
	// @@protoc_insertion_point(type_ServiceDescriptorProto)
}
type MethodDescriptorProto struct {
	Name string
	InputType string
	OutputType string
	Options .google.protobuf.MethodOptions
	ClientStreaming bool
	ServerStreaming bool
	// @@protoc_insertion_point(type_MethodDescriptorProto)
}
type FileOptions struct {
	JavaPackage string
	JavaOuterClassname string
	JavaMultipleFiles bool
	JavaGenerateEqualsAndHash bool
	JavaStringCheckUtf8 bool
	OptimizeFor .google.protobuf.FileOptions.OptimizeMode
	GoPackage string
	CcGenericServices bool
	JavaGenericServices bool
	PyGenericServices bool
	Deprecated bool
	CcEnableArenas bool
	ObjcClassPrefix string
	CsharpNamespace string
	SwiftPrefix string
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_FileOptions)
}
type MessageOptions struct {
	MessageSetWireFormat bool
	NoStandardDescriptorAccessor bool
	Deprecated bool
	MapEntry bool
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_MessageOptions)
}
type FieldOptions struct {
	Ctype .google.protobuf.FieldOptions.CType
	Packed bool
	Jstype .google.protobuf.FieldOptions.JSType
	Lazy bool
	Deprecated bool
	Weak bool
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_FieldOptions)
}
type OneofOptions struct {
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_OneofOptions)
}
type EnumOptions struct {
	AllowAlias bool
	Deprecated bool
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_EnumOptions)
}
type EnumValueOptions struct {
	Deprecated bool
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_EnumValueOptions)
}
type ServiceOptions struct {
	Deprecated bool
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_ServiceOptions)
}
type MethodOptions struct {
	Deprecated bool
	IdempotencyLevel .google.protobuf.MethodOptions.IdempotencyLevel
	UninterpretedOption .google.protobuf.UninterpretedOption
	// @@protoc_insertion_point(type_MethodOptions)
}
type UninterpretedOption struct {
	Name .google.protobuf.UninterpretedOption.NamePart
	IdentifierValue string
	PositiveIntValue uint64
	NegativeIntValue int64
	DoubleValue float64
	StringValue []byte
	AggregateValue string
	// @@protoc_insertion_point(type_UninterpretedOption)
}
type SourceCodeInfo struct {
	Location .google.protobuf.SourceCodeInfo.Location
	// @@protoc_insertion_point(type_SourceCodeInfo)
}
type GeneratedCodeInfo struct {
	Annotation .google.protobuf.GeneratedCodeInfo.Annotation
	// @@protoc_insertion_point(type_GeneratedCodeInfo)
}
// @@protoc_insertion_point(type)
