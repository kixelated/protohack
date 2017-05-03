package descriptor
type FileDescriptorSet struct {
	File *FileDescriptorProto
}
type FileDescriptorProto struct {
	Name string
	Package string
	Dependency string
	PublicDependency int32
	WeakDependency int32
	MessageType *DescriptorProto
	EnumType *EnumDescriptorProto
	Service *ServiceDescriptorProto
	Extension *FieldDescriptorProto
	Options *FileOptions
	SourceCodeInfo *SourceCodeInfo
	Syntax string
}
type DescriptorProto struct {
	Name string
	Field *FieldDescriptorProto
	Extension *FieldDescriptorProto
	NestedType *DescriptorProto
	EnumType *EnumDescriptorProto
	ExtensionRange *DescriptorProto_ExtensionRange
	OneofDecl *OneofDescriptorProto
	Options *MessageOptions
	ReservedRange *DescriptorProto_ReservedRange
	ReservedName string
}
type DescriptorProto_ExtensionRange struct {
	Start int32
	End int32
}
type DescriptorProto_ReservedRange struct {
	Start int32
	End int32
}
type FieldDescriptorProto struct {
	Name string
	Number int32
	Label *FieldDescriptorProto_Label
	Type *FieldDescriptorProto_Type
	TypeName string
	Extendee string
	DefaultValue string
	OneofIndex int32
	JsonName string
	Options *FieldOptions
}
type FieldDescriptorProto_Type int32
const (
	FieldDescriptorProto_TYPE_DOUBLE FieldDescriptorProto_Type = 1
	FieldDescriptorProto_TYPE_FLOAT FieldDescriptorProto_Type = 2
	FieldDescriptorProto_TYPE_INT64 FieldDescriptorProto_Type = 3
	FieldDescriptorProto_TYPE_UINT64 FieldDescriptorProto_Type = 4
	FieldDescriptorProto_TYPE_INT32 FieldDescriptorProto_Type = 5
	FieldDescriptorProto_TYPE_FIXED64 FieldDescriptorProto_Type = 6
	FieldDescriptorProto_TYPE_FIXED32 FieldDescriptorProto_Type = 7
	FieldDescriptorProto_TYPE_BOOL FieldDescriptorProto_Type = 8
	FieldDescriptorProto_TYPE_STRING FieldDescriptorProto_Type = 9
	FieldDescriptorProto_TYPE_GROUP FieldDescriptorProto_Type = 10
	FieldDescriptorProto_TYPE_MESSAGE FieldDescriptorProto_Type = 11
	FieldDescriptorProto_TYPE_BYTES FieldDescriptorProto_Type = 12
	FieldDescriptorProto_TYPE_UINT32 FieldDescriptorProto_Type = 13
	FieldDescriptorProto_TYPE_ENUM FieldDescriptorProto_Type = 14
	FieldDescriptorProto_TYPE_SFIXED32 FieldDescriptorProto_Type = 15
	FieldDescriptorProto_TYPE_SFIXED64 FieldDescriptorProto_Type = 16
	FieldDescriptorProto_TYPE_SINT32 FieldDescriptorProto_Type = 17
	FieldDescriptorProto_TYPE_SINT64 FieldDescriptorProto_Type = 18
)
type FieldDescriptorProto_Label int32
const (
	FieldDescriptorProto_LABEL_OPTIONAL FieldDescriptorProto_Label = 1
	FieldDescriptorProto_LABEL_REQUIRED FieldDescriptorProto_Label = 2
	FieldDescriptorProto_LABEL_REPEATED FieldDescriptorProto_Label = 3
)
type OneofDescriptorProto struct {
	Name string
	Options *OneofOptions
}
type EnumDescriptorProto struct {
	Name string
	Value *EnumValueDescriptorProto
	Options *EnumOptions
}
type EnumValueDescriptorProto struct {
	Name string
	Number int32
	Options *EnumValueOptions
}
type ServiceDescriptorProto struct {
	Name string
	Method *MethodDescriptorProto
	Options *ServiceOptions
}
type MethodDescriptorProto struct {
	Name string
	InputType string
	OutputType string
	Options *MethodOptions
	ClientStreaming bool
	ServerStreaming bool
}
type FileOptions struct {
	JavaPackage string
	JavaOuterClassname string
	JavaMultipleFiles bool
	JavaGenerateEqualsAndHash bool
	JavaStringCheckUtf8 bool
	OptimizeFor *FileOptions_OptimizeMode
	GoPackage string
	CcGenericServices bool
	JavaGenericServices bool
	PyGenericServices bool
	Deprecated bool
	CcEnableArenas bool
	ObjcClassPrefix string
	CsharpNamespace string
	SwiftPrefix string
	UninterpretedOption *UninterpretedOption
}
type FileOptions_OptimizeMode int32
const (
	FileOptions_SPEED FileOptions_OptimizeMode = 1
	FileOptions_CODE_SIZE FileOptions_OptimizeMode = 2
	FileOptions_LITE_RUNTIME FileOptions_OptimizeMode = 3
)
type MessageOptions struct {
	MessageSetWireFormat bool
	NoStandardDescriptorAccessor bool
	Deprecated bool
	MapEntry bool
	UninterpretedOption *UninterpretedOption
}
type FieldOptions struct {
	Ctype *FieldOptions_CType
	Packed bool
	Jstype *FieldOptions_JSType
	Lazy bool
	Deprecated bool
	Weak bool
	UninterpretedOption *UninterpretedOption
}
type FieldOptions_CType int32
const (
	FieldOptions_STRING FieldOptions_CType = 0
	FieldOptions_CORD FieldOptions_CType = 1
	FieldOptions_STRING_PIECE FieldOptions_CType = 2
)
type FieldOptions_JSType int32
const (
	FieldOptions_JS_NORMAL FieldOptions_JSType = 0
	FieldOptions_JS_STRING FieldOptions_JSType = 1
	FieldOptions_JS_NUMBER FieldOptions_JSType = 2
)
type OneofOptions struct {
	UninterpretedOption *UninterpretedOption
}
type EnumOptions struct {
	AllowAlias bool
	Deprecated bool
	UninterpretedOption *UninterpretedOption
}
type EnumValueOptions struct {
	Deprecated bool
	UninterpretedOption *UninterpretedOption
}
type ServiceOptions struct {
	Deprecated bool
	UninterpretedOption *UninterpretedOption
}
type MethodOptions struct {
	Deprecated bool
	IdempotencyLevel *MethodOptions_IdempotencyLevel
	UninterpretedOption *UninterpretedOption
}
type MethodOptions_IdempotencyLevel int32
const (
	MethodOptions_IDEMPOTENCY_UNKNOWN MethodOptions_IdempotencyLevel = 0
	MethodOptions_NO_SIDE_EFFECTS MethodOptions_IdempotencyLevel = 1
	MethodOptions_IDEMPOTENT MethodOptions_IdempotencyLevel = 2
)
type UninterpretedOption struct {
	Name *UninterpretedOption_NamePart
	IdentifierValue string
	PositiveIntValue uint64
	NegativeIntValue int64
	DoubleValue float64
	StringValue []byte
	AggregateValue string
}
type UninterpretedOption_NamePart struct {
	NamePart string
	IsExtension bool
}
type SourceCodeInfo struct {
	Location *SourceCodeInfo_Location
}
type SourceCodeInfo_Location struct {
	Path int32
	Span int32
	LeadingComments string
	TrailingComments string
	LeadingDetachedComments string
}
type GeneratedCodeInfo struct {
	Annotation *GeneratedCodeInfo_Annotation
}
type GeneratedCodeInfo_Annotation struct {
	Path int32
	SourceFile string
	Begin int32
	End int32
}
