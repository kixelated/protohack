// PROTO name: descriptor.proto
// PROTO package: google.protobuf
package descriptor
// @@protoc_insertion_point(import)
// @@protoc_insertion_point(enum)
// PROTO name: FileDescriptorSet
type FileDescriptorSet struct {
	// PROTO name: file
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.FileDescriptorProto
	File *FileDescriptorProto
	// @@protoc_insertion_point(type_FileDescriptorSet)
}
// PROTO name: FileDescriptorProto
type FileDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: package
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Package string
	// PROTO name: dependency
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_STRING
	Dependency string
	// PROTO name: public_dependency
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_INT32
	PublicDependency int32
	// PROTO name: weak_dependency
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_INT32
	WeakDependency int32
	// PROTO name: message_type
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.DescriptorProto
	MessageType *DescriptorProto
	// PROTO name: enum_type
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.EnumDescriptorProto
	EnumType *EnumDescriptorProto
	// PROTO name: service
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.ServiceDescriptorProto
	Service *ServiceDescriptorProto
	// PROTO name: extension
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.FieldDescriptorProto
	Extension *FieldDescriptorProto
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.FileOptions
	Options *FileOptions
	// PROTO name: source_code_info
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.SourceCodeInfo
	SourceCodeInfo *SourceCodeInfo
	// PROTO name: syntax
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Syntax string
	// @@protoc_insertion_point(type_FileDescriptorProto)
}
// PROTO name: DescriptorProto
type DescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: field
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.FieldDescriptorProto
	Field *FieldDescriptorProto
	// PROTO name: extension
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.FieldDescriptorProto
	Extension *FieldDescriptorProto
	// PROTO name: nested_type
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.DescriptorProto
	NestedType *DescriptorProto
	// PROTO name: enum_type
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.EnumDescriptorProto
	EnumType *EnumDescriptorProto
	// PROTO name: extension_range
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.DescriptorProto.ExtensionRange
	ExtensionRange *ExtensionRange
	// PROTO name: oneof_decl
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.OneofDescriptorProto
	OneofDecl *OneofDescriptorProto
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.MessageOptions
	Options *MessageOptions
	// PROTO name: reserved_range
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.DescriptorProto.ReservedRange
	ReservedRange *ReservedRange
	// PROTO name: reserved_name
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_STRING
	ReservedName string
	// @@protoc_insertion_point(type_DescriptorProto)
}
// PROTO name: FieldDescriptorProto
type FieldDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: number
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_INT32
	Number int32
	// PROTO name: label
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_ENUM
	// PROTO type_name: .google.protobuf.FieldDescriptorProto.Label
	Label .google.protobuf.FieldDescriptorProto.Label
	// PROTO name: type
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_ENUM
	// PROTO type_name: .google.protobuf.FieldDescriptorProto.Type
	Type .google.protobuf.FieldDescriptorProto.Type
	// PROTO name: type_name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	TypeName string
	// PROTO name: extendee
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Extendee string
	// PROTO name: default_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	DefaultValue string
	// PROTO name: oneof_index
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_INT32
	OneofIndex int32
	// PROTO name: json_name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	JsonName string
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.FieldOptions
	Options *FieldOptions
	// @@protoc_insertion_point(type_FieldDescriptorProto)
}
// PROTO name: OneofDescriptorProto
type OneofDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.OneofOptions
	Options *OneofOptions
	// @@protoc_insertion_point(type_OneofDescriptorProto)
}
// PROTO name: EnumDescriptorProto
type EnumDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: value
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.EnumValueDescriptorProto
	Value *EnumValueDescriptorProto
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.EnumOptions
	Options *EnumOptions
	// @@protoc_insertion_point(type_EnumDescriptorProto)
}
// PROTO name: EnumValueDescriptorProto
type EnumValueDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: number
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_INT32
	Number int32
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.EnumValueOptions
	Options *EnumValueOptions
	// @@protoc_insertion_point(type_EnumValueDescriptorProto)
}
// PROTO name: ServiceDescriptorProto
type ServiceDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: method
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.MethodDescriptorProto
	Method *MethodDescriptorProto
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.ServiceOptions
	Options *ServiceOptions
	// @@protoc_insertion_point(type_ServiceDescriptorProto)
}
// PROTO name: MethodDescriptorProto
type MethodDescriptorProto struct {
	// PROTO name: name
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	Name string
	// PROTO name: input_type
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	InputType string
	// PROTO name: output_type
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	OutputType string
	// PROTO name: options
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.MethodOptions
	Options *MethodOptions
	// PROTO name: client_streaming
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	ClientStreaming bool
	// PROTO name: server_streaming
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	ServerStreaming bool
	// @@protoc_insertion_point(type_MethodDescriptorProto)
}
// PROTO name: FileOptions
type FileOptions struct {
	// PROTO name: java_package
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	JavaPackage string
	// PROTO name: java_outer_classname
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	JavaOuterClassname string
	// PROTO name: java_multiple_files
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	JavaMultipleFiles bool
	// PROTO name: java_generate_equals_and_hash
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	JavaGenerateEqualsAndHash bool
	// PROTO name: java_string_check_utf8
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	JavaStringCheckUtf8 bool
	// PROTO name: optimize_for
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_ENUM
	// PROTO type_name: .google.protobuf.FileOptions.OptimizeMode
	OptimizeFor .google.protobuf.FileOptions.OptimizeMode
	// PROTO name: go_package
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	GoPackage string
	// PROTO name: cc_generic_services
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	CcGenericServices bool
	// PROTO name: java_generic_services
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	JavaGenericServices bool
	// PROTO name: py_generic_services
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	PyGenericServices bool
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: cc_enable_arenas
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	CcEnableArenas bool
	// PROTO name: objc_class_prefix
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	ObjcClassPrefix string
	// PROTO name: csharp_namespace
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	CsharpNamespace string
	// PROTO name: swift_prefix
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	SwiftPrefix string
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_FileOptions)
}
// PROTO name: MessageOptions
type MessageOptions struct {
	// PROTO name: message_set_wire_format
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	MessageSetWireFormat bool
	// PROTO name: no_standard_descriptor_accessor
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	NoStandardDescriptorAccessor bool
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: map_entry
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	MapEntry bool
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_MessageOptions)
}
// PROTO name: FieldOptions
type FieldOptions struct {
	// PROTO name: ctype
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_ENUM
	// PROTO type_name: .google.protobuf.FieldOptions.CType
	Ctype .google.protobuf.FieldOptions.CType
	// PROTO name: packed
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Packed bool
	// PROTO name: jstype
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_ENUM
	// PROTO type_name: .google.protobuf.FieldOptions.JSType
	Jstype .google.protobuf.FieldOptions.JSType
	// PROTO name: lazy
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Lazy bool
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: weak
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Weak bool
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_FieldOptions)
}
// PROTO name: OneofOptions
type OneofOptions struct {
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_OneofOptions)
}
// PROTO name: EnumOptions
type EnumOptions struct {
	// PROTO name: allow_alias
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	AllowAlias bool
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_EnumOptions)
}
// PROTO name: EnumValueOptions
type EnumValueOptions struct {
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_EnumValueOptions)
}
// PROTO name: ServiceOptions
type ServiceOptions struct {
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_ServiceOptions)
}
// PROTO name: MethodOptions
type MethodOptions struct {
	// PROTO name: deprecated
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BOOL
	Deprecated bool
	// PROTO name: idempotency_level
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_ENUM
	// PROTO type_name: .google.protobuf.MethodOptions.IdempotencyLevel
	IdempotencyLevel .google.protobuf.MethodOptions.IdempotencyLevel
	// PROTO name: uninterpreted_option
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption
	UninterpretedOption *UninterpretedOption
	// @@protoc_insertion_point(type_MethodOptions)
}
// PROTO name: UninterpretedOption
type UninterpretedOption struct {
	// PROTO name: name
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.UninterpretedOption.NamePart
	Name *NamePart
	// PROTO name: identifier_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	IdentifierValue string
	// PROTO name: positive_int_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_UINT64
	PositiveIntValue uint64
	// PROTO name: negative_int_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_INT64
	NegativeIntValue int64
	// PROTO name: double_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_DOUBLE
	DoubleValue float64
	// PROTO name: string_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_BYTES
	StringValue []byte
	// PROTO name: aggregate_value
	// PROTO label: LABEL_OPTIONAL
	// PROTO type: TYPE_STRING
	AggregateValue string
	// @@protoc_insertion_point(type_UninterpretedOption)
}
// PROTO name: SourceCodeInfo
type SourceCodeInfo struct {
	// PROTO name: location
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.SourceCodeInfo.Location
	Location *Location
	// @@protoc_insertion_point(type_SourceCodeInfo)
}
// PROTO name: GeneratedCodeInfo
type GeneratedCodeInfo struct {
	// PROTO name: annotation
	// PROTO label: LABEL_REPEATED
	// PROTO type: TYPE_MESSAGE
	// PROTO type_name: .google.protobuf.GeneratedCodeInfo.Annotation
	Annotation *Annotation
	// @@protoc_insertion_point(type_GeneratedCodeInfo)
}
// @@protoc_insertion_point(type)
