package descriptor

type FileDescriptorSet struct {
	File *FileDescriptorProto
}
type FileDescriptorProto struct {
	Name             string
	Package          string
	Dependency       string
	PublicDependency int32
	WeakDependency   int32
	MessageType      *DescriptorProto
	EnumType         *EnumDescriptorProto
	Service          *ServiceDescriptorProto
	Extension        *FieldDescriptorProto
	Options          *FileOptions
	SourceCodeInfo   *SourceCodeInfo
	Syntax           string
}
type DescriptorProto struct {
	Name           string
	Field          *FieldDescriptorProto
	Extension      *FieldDescriptorProto
	NestedType     *DescriptorProto
	EnumType       *EnumDescriptorProto
	ExtensionRange *DescriptorProto_ExtensionRange
	OneofDecl      *OneofDescriptorProto
	Options        *MessageOptions
	ReservedRange  *DescriptorProto_ReservedRange
	ReservedName   string
}
type DescriptorProto_ExtensionRange struct {
	Start int32
	End   int32
}
type DescriptorProto_ReservedRange struct {
	Start int32
	End   int32
}
type FieldDescriptorProto struct {
	Name         string
	Number       int32
	Label        *FieldDescriptorProto_Label
	Type         *FieldDescriptorProto_Type
	TypeName     string
	Extendee     string
	DefaultValue string
	OneofIndex   int32
	JsonName     string
	Options      *FieldOptions
}
type FieldDescriptorProto_Type int32
type FieldDescriptorProto_Label int32
type OneofDescriptorProto struct {
	Name    string
	Options *OneofOptions
}
type EnumDescriptorProto struct {
	Name    string
	Value   *EnumValueDescriptorProto
	Options *EnumOptions
}
type EnumValueDescriptorProto struct {
	Name    string
	Number  int32
	Options *EnumValueOptions
}
type ServiceDescriptorProto struct {
	Name    string
	Method  *MethodDescriptorProto
	Options *ServiceOptions
}
type MethodDescriptorProto struct {
	Name            string
	InputType       string
	OutputType      string
	Options         *MethodOptions
	ClientStreaming bool
	ServerStreaming bool
}
type FileOptions struct {
	JavaPackage               string
	JavaOuterClassname        string
	JavaMultipleFiles         bool
	JavaGenerateEqualsAndHash bool
	JavaStringCheckUtf8       bool
	OptimizeFor               *FileOptions_OptimizeMode
	GoPackage                 string
	CcGenericServices         bool
	JavaGenericServices       bool
	PyGenericServices         bool
	Deprecated                bool
	CcEnableArenas            bool
	ObjcClassPrefix           string
	CsharpNamespace           string
	SwiftPrefix               string
	UninterpretedOption       *UninterpretedOption
}
type FileOptions_OptimizeMode int32
type MessageOptions struct {
	MessageSetWireFormat         bool
	NoStandardDescriptorAccessor bool
	Deprecated                   bool
	MapEntry                     bool
	UninterpretedOption          *UninterpretedOption
}
type FieldOptions struct {
	Ctype               *FieldOptions_CType
	Packed              bool
	Jstype              *FieldOptions_JSType
	Lazy                bool
	Deprecated          bool
	Weak                bool
	UninterpretedOption *UninterpretedOption
}
type FieldOptions_CType int32
type FieldOptions_JSType int32
type OneofOptions struct {
	UninterpretedOption *UninterpretedOption
}
type EnumOptions struct {
	AllowAlias          bool
	Deprecated          bool
	UninterpretedOption *UninterpretedOption
}
type EnumValueOptions struct {
	Deprecated          bool
	UninterpretedOption *UninterpretedOption
}
type ServiceOptions struct {
	Deprecated          bool
	UninterpretedOption *UninterpretedOption
}
type MethodOptions struct {
	Deprecated          bool
	IdempotencyLevel    *MethodOptions_IdempotencyLevel
	UninterpretedOption *UninterpretedOption
}
type MethodOptions_IdempotencyLevel int32
type UninterpretedOption struct {
	Name             *UninterpretedOption_NamePart
	IdentifierValue  string
	PositiveIntValue uint64
	NegativeIntValue int64
	DoubleValue      float64
	StringValue      []byte
	AggregateValue   string
}
type UninterpretedOption_NamePart struct {
	NamePart    string
	IsExtension bool
}
type SourceCodeInfo struct {
	Location *SourceCodeInfo_Location
}
type SourceCodeInfo_Location struct {
	Path                    int32
	Span                    int32
	LeadingComments         string
	TrailingComments        string
	LeadingDetachedComments string
}
type GeneratedCodeInfo struct {
	Annotation *GeneratedCodeInfo_Annotation
}
type GeneratedCodeInfo_Annotation struct {
	Path       int32
	SourceFile string
	Begin      int32
	End        int32
}
