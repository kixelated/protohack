package descriptor
import "github.com/kixelated/protohack/proto"
var _ = proto.WIRETYPE_VARINT
type FileDescriptorSet struct {
	File []*FileDescriptorProto
}
type FileDescriptorProto struct {
	Name string
	Package string
	Dependency []string
	PublicDependency []int32
	WeakDependency []int32
	MessageType []*DescriptorProto
	EnumType []*EnumDescriptorProto
	Service []*ServiceDescriptorProto
	Extension []*FieldDescriptorProto
	Options *FileOptions
	SourceCodeInfo *SourceCodeInfo
	Syntax string
}
type DescriptorProto struct {
	Name string
	Field []*FieldDescriptorProto
	Extension []*FieldDescriptorProto
	NestedType []*DescriptorProto
	EnumType []*EnumDescriptorProto
	ExtensionRange []*DescriptorProto_ExtensionRange
	OneofDecl []*OneofDescriptorProto
	Options *MessageOptions
	ReservedRange []*DescriptorProto_ReservedRange
	ReservedName []string
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
	Label FieldDescriptorProto_Label
	Type FieldDescriptorProto_Type
	TypeName string
	Extendee string
	DefaultValue string
	OneofIndex int32
	JsonName string
	Options *FieldOptions
}
type FieldDescriptorProto_Type int
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
type FieldDescriptorProto_Label int
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
	Value []*EnumValueDescriptorProto
	Options *EnumOptions
}
type EnumValueDescriptorProto struct {
	Name string
	Number int32
	Options *EnumValueOptions
}
type ServiceDescriptorProto struct {
	Name string
	Method []*MethodDescriptorProto
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
	OptimizeFor FileOptions_OptimizeMode
	GoPackage string
	CcGenericServices bool
	JavaGenericServices bool
	PyGenericServices bool
	Deprecated bool
	CcEnableArenas bool
	ObjcClassPrefix string
	CsharpNamespace string
	SwiftPrefix string
	UninterpretedOption []*UninterpretedOption
}
type FileOptions_OptimizeMode int
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
	UninterpretedOption []*UninterpretedOption
}
type FieldOptions struct {
	Ctype FieldOptions_CType
	Packed bool
	Jstype FieldOptions_JSType
	Lazy bool
	Deprecated bool
	Weak bool
	UninterpretedOption []*UninterpretedOption
}
type FieldOptions_CType int
const (
	FieldOptions_STRING FieldOptions_CType = 0
	FieldOptions_CORD FieldOptions_CType = 1
	FieldOptions_STRING_PIECE FieldOptions_CType = 2
)
type FieldOptions_JSType int
const (
	FieldOptions_JS_NORMAL FieldOptions_JSType = 0
	FieldOptions_JS_STRING FieldOptions_JSType = 1
	FieldOptions_JS_NUMBER FieldOptions_JSType = 2
)
type OneofOptions struct {
	UninterpretedOption []*UninterpretedOption
}
type EnumOptions struct {
	AllowAlias bool
	Deprecated bool
	UninterpretedOption []*UninterpretedOption
}
type EnumValueOptions struct {
	Deprecated bool
	UninterpretedOption []*UninterpretedOption
}
type ServiceOptions struct {
	Deprecated bool
	UninterpretedOption []*UninterpretedOption
}
type MethodOptions struct {
	Deprecated bool
	IdempotencyLevel MethodOptions_IdempotencyLevel
	UninterpretedOption []*UninterpretedOption
}
type MethodOptions_IdempotencyLevel int
const (
	MethodOptions_IDEMPOTENCY_UNKNOWN MethodOptions_IdempotencyLevel = 0
	MethodOptions_NO_SIDE_EFFECTS MethodOptions_IdempotencyLevel = 1
	MethodOptions_IDEMPOTENT MethodOptions_IdempotencyLevel = 2
)
type UninterpretedOption struct {
	Name []*UninterpretedOption_NamePart
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
	Location []*SourceCodeInfo_Location
}
type SourceCodeInfo_Location struct {
	Path []int32
	Span []int32
	LeadingComments string
	TrailingComments string
	LeadingDetachedComments []string
}
type GeneratedCodeInfo struct {
	Annotation []*GeneratedCodeInfo_Annotation
}
type GeneratedCodeInfo_Annotation struct {
	Path []int32
	SourceFile string
	Begin int32
	End int32
}
func (m FileDescriptorSet) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m FileDescriptorSet) MarshalTo(data []byte) (n int) {
	for _, x := range m.File {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 1, x)
		}
	}
	return n
}
func (m FileDescriptorSet) MarshalSize() (n int) {
	for _, x := range m.File {
		if x != nil {
			n += proto.SizeFieldMessage(1, x)
		}
	}
	return n
}
func (m FileDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m FileDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	n += proto.WriteFieldString(data[n:], 2, m.Package)
	for _, x := range m.Dependency {
		n += proto.WriteFieldString(data[n:], 3, x)
	}
	for _, x := range m.MessageType {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 4, x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 5, x)
		}
	}
	for _, x := range m.Service {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 6, x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 7, x)
		}
	}
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 8, m.Options)
	}
	if m.SourceCodeInfo != nil {
		n += proto.WriteFieldMessage(data[n:], 9, m.SourceCodeInfo)
	}
	for _, x := range m.PublicDependency {
		n += proto.WriteFieldInt32(data[n:], 10, x)
	}
	for _, x := range m.WeakDependency {
		n += proto.WriteFieldInt32(data[n:], 11, x)
	}
	n += proto.WriteFieldString(data[n:], 12, m.Syntax)
	return n
}
func (m FileDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	n += proto.SizeFieldString(2, m.Package)
	for _, x := range m.Dependency {
		n += proto.SizeFieldString(3, x)
	}
	for _, x := range m.MessageType {
		if x != nil {
			n += proto.SizeFieldMessage(4, x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += proto.SizeFieldMessage(5, x)
		}
	}
	for _, x := range m.Service {
		if x != nil {
			n += proto.SizeFieldMessage(6, x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += proto.SizeFieldMessage(7, x)
		}
	}
	if m.Options != nil {
		n += proto.SizeFieldMessage(8, m.Options)
	}
	if m.SourceCodeInfo != nil {
		n += proto.SizeFieldMessage(9, m.SourceCodeInfo)
	}
	for _, x := range m.PublicDependency {
		n += proto.SizeFieldInt32(10, x)
	}
	for _, x := range m.WeakDependency {
		n += proto.SizeFieldInt32(11, x)
	}
	n += proto.SizeFieldString(12, m.Syntax)
	return n
}
func (m DescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m DescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	for _, x := range m.Field {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 2, x)
		}
	}
	for _, x := range m.NestedType {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 3, x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 4, x)
		}
	}
	for _, x := range m.ExtensionRange {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 5, x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 6, x)
		}
	}
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 7, m.Options)
	}
	for _, x := range m.OneofDecl {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 8, x)
		}
	}
	for _, x := range m.ReservedRange {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 9, x)
		}
	}
	for _, x := range m.ReservedName {
		n += proto.WriteFieldString(data[n:], 10, x)
	}
	return n
}
func (m DescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	for _, x := range m.Field {
		if x != nil {
			n += proto.SizeFieldMessage(2, x)
		}
	}
	for _, x := range m.NestedType {
		if x != nil {
			n += proto.SizeFieldMessage(3, x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += proto.SizeFieldMessage(4, x)
		}
	}
	for _, x := range m.ExtensionRange {
		if x != nil {
			n += proto.SizeFieldMessage(5, x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += proto.SizeFieldMessage(6, x)
		}
	}
	if m.Options != nil {
		n += proto.SizeFieldMessage(7, m.Options)
	}
	for _, x := range m.OneofDecl {
		if x != nil {
			n += proto.SizeFieldMessage(8, x)
		}
	}
	for _, x := range m.ReservedRange {
		if x != nil {
			n += proto.SizeFieldMessage(9, x)
		}
	}
	for _, x := range m.ReservedName {
		n += proto.SizeFieldString(10, x)
	}
	return n
}
func (m DescriptorProto_ExtensionRange) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m DescriptorProto_ExtensionRange) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldInt32(data[n:], 1, m.Start)
	n += proto.WriteFieldInt32(data[n:], 2, m.End)
	return n
}
func (m DescriptorProto_ExtensionRange) MarshalSize() (n int) {
	n += proto.SizeFieldInt32(1, m.Start)
	n += proto.SizeFieldInt32(2, m.End)
	return n
}
func (m DescriptorProto_ReservedRange) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m DescriptorProto_ReservedRange) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldInt32(data[n:], 1, m.Start)
	n += proto.WriteFieldInt32(data[n:], 2, m.End)
	return n
}
func (m DescriptorProto_ReservedRange) MarshalSize() (n int) {
	n += proto.SizeFieldInt32(1, m.Start)
	n += proto.SizeFieldInt32(2, m.End)
	return n
}
func (m FieldDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m FieldDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	n += proto.WriteFieldString(data[n:], 2, m.Extendee)
	n += proto.WriteFieldInt32(data[n:], 3, m.Number)
	n += proto.WriteFieldEnum(data[n:], 4, int(m.Label))
	n += proto.WriteFieldEnum(data[n:], 5, int(m.Type))
	n += proto.WriteFieldString(data[n:], 6, m.TypeName)
	n += proto.WriteFieldString(data[n:], 7, m.DefaultValue)
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 8, m.Options)
	}
	n += proto.WriteFieldInt32(data[n:], 9, m.OneofIndex)
	n += proto.WriteFieldString(data[n:], 10, m.JsonName)
	return n
}
func (m FieldDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	n += proto.SizeFieldString(2, m.Extendee)
	n += proto.SizeFieldInt32(3, m.Number)
	n += proto.SizeFieldEnum(4, int(m.Label))
	n += proto.SizeFieldEnum(5, int(m.Type))
	n += proto.SizeFieldString(6, m.TypeName)
	n += proto.SizeFieldString(7, m.DefaultValue)
	if m.Options != nil {
		n += proto.SizeFieldMessage(8, m.Options)
	}
	n += proto.SizeFieldInt32(9, m.OneofIndex)
	n += proto.SizeFieldString(10, m.JsonName)
	return n
}
func (m OneofDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m OneofDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 2, m.Options)
	}
	return n
}
func (m OneofDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	if m.Options != nil {
		n += proto.SizeFieldMessage(2, m.Options)
	}
	return n
}
func (m EnumDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m EnumDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	for _, x := range m.Value {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 2, x)
		}
	}
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 3, m.Options)
	}
	return n
}
func (m EnumDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	for _, x := range m.Value {
		if x != nil {
			n += proto.SizeFieldMessage(2, x)
		}
	}
	if m.Options != nil {
		n += proto.SizeFieldMessage(3, m.Options)
	}
	return n
}
func (m EnumValueDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m EnumValueDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	n += proto.WriteFieldInt32(data[n:], 2, m.Number)
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 3, m.Options)
	}
	return n
}
func (m EnumValueDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	n += proto.SizeFieldInt32(2, m.Number)
	if m.Options != nil {
		n += proto.SizeFieldMessage(3, m.Options)
	}
	return n
}
func (m ServiceDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m ServiceDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	for _, x := range m.Method {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 2, x)
		}
	}
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 3, m.Options)
	}
	return n
}
func (m ServiceDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	for _, x := range m.Method {
		if x != nil {
			n += proto.SizeFieldMessage(2, x)
		}
	}
	if m.Options != nil {
		n += proto.SizeFieldMessage(3, m.Options)
	}
	return n
}
func (m MethodDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m MethodDescriptorProto) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	n += proto.WriteFieldString(data[n:], 2, m.InputType)
	n += proto.WriteFieldString(data[n:], 3, m.OutputType)
	if m.Options != nil {
		n += proto.WriteFieldMessage(data[n:], 4, m.Options)
	}
	n += proto.WriteFieldBool(data[n:], 5, m.ClientStreaming)
	n += proto.WriteFieldBool(data[n:], 6, m.ServerStreaming)
	return n
}
func (m MethodDescriptorProto) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	n += proto.SizeFieldString(2, m.InputType)
	n += proto.SizeFieldString(3, m.OutputType)
	if m.Options != nil {
		n += proto.SizeFieldMessage(4, m.Options)
	}
	n += proto.SizeFieldBool(5, m.ClientStreaming)
	n += proto.SizeFieldBool(6, m.ServerStreaming)
	return n
}
func (m FileOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m FileOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.JavaPackage)
	n += proto.WriteFieldString(data[n:], 8, m.JavaOuterClassname)
	n += proto.WriteFieldEnum(data[n:], 9, int(m.OptimizeFor))
	n += proto.WriteFieldBool(data[n:], 10, m.JavaMultipleFiles)
	n += proto.WriteFieldString(data[n:], 11, m.GoPackage)
	n += proto.WriteFieldBool(data[n:], 16, m.CcGenericServices)
	n += proto.WriteFieldBool(data[n:], 17, m.JavaGenericServices)
	n += proto.WriteFieldBool(data[n:], 18, m.PyGenericServices)
	n += proto.WriteFieldBool(data[n:], 20, m.JavaGenerateEqualsAndHash)
	n += proto.WriteFieldBool(data[n:], 23, m.Deprecated)
	n += proto.WriteFieldBool(data[n:], 27, m.JavaStringCheckUtf8)
	n += proto.WriteFieldBool(data[n:], 31, m.CcEnableArenas)
	n += proto.WriteFieldString(data[n:], 36, m.ObjcClassPrefix)
	n += proto.WriteFieldString(data[n:], 37, m.CsharpNamespace)
	n += proto.WriteFieldString(data[n:], 39, m.SwiftPrefix)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m FileOptions) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.JavaPackage)
	n += proto.SizeFieldString(8, m.JavaOuterClassname)
	n += proto.SizeFieldEnum(9, int(m.OptimizeFor))
	n += proto.SizeFieldBool(10, m.JavaMultipleFiles)
	n += proto.SizeFieldString(11, m.GoPackage)
	n += proto.SizeFieldBool(16, m.CcGenericServices)
	n += proto.SizeFieldBool(17, m.JavaGenericServices)
	n += proto.SizeFieldBool(18, m.PyGenericServices)
	n += proto.SizeFieldBool(20, m.JavaGenerateEqualsAndHash)
	n += proto.SizeFieldBool(23, m.Deprecated)
	n += proto.SizeFieldBool(27, m.JavaStringCheckUtf8)
	n += proto.SizeFieldBool(31, m.CcEnableArenas)
	n += proto.SizeFieldString(36, m.ObjcClassPrefix)
	n += proto.SizeFieldString(37, m.CsharpNamespace)
	n += proto.SizeFieldString(39, m.SwiftPrefix)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m MessageOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m MessageOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldBool(data[n:], 1, m.MessageSetWireFormat)
	n += proto.WriteFieldBool(data[n:], 2, m.NoStandardDescriptorAccessor)
	n += proto.WriteFieldBool(data[n:], 3, m.Deprecated)
	n += proto.WriteFieldBool(data[n:], 7, m.MapEntry)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m MessageOptions) MarshalSize() (n int) {
	n += proto.SizeFieldBool(1, m.MessageSetWireFormat)
	n += proto.SizeFieldBool(2, m.NoStandardDescriptorAccessor)
	n += proto.SizeFieldBool(3, m.Deprecated)
	n += proto.SizeFieldBool(7, m.MapEntry)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m FieldOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m FieldOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldEnum(data[n:], 1, int(m.Ctype))
	n += proto.WriteFieldBool(data[n:], 2, m.Packed)
	n += proto.WriteFieldBool(data[n:], 3, m.Deprecated)
	n += proto.WriteFieldBool(data[n:], 5, m.Lazy)
	n += proto.WriteFieldEnum(data[n:], 6, int(m.Jstype))
	n += proto.WriteFieldBool(data[n:], 10, m.Weak)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m FieldOptions) MarshalSize() (n int) {
	n += proto.SizeFieldEnum(1, int(m.Ctype))
	n += proto.SizeFieldBool(2, m.Packed)
	n += proto.SizeFieldBool(3, m.Deprecated)
	n += proto.SizeFieldBool(5, m.Lazy)
	n += proto.SizeFieldEnum(6, int(m.Jstype))
	n += proto.SizeFieldBool(10, m.Weak)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m OneofOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m OneofOptions) MarshalTo(data []byte) (n int) {
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m OneofOptions) MarshalSize() (n int) {
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m EnumOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m EnumOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldBool(data[n:], 2, m.AllowAlias)
	n += proto.WriteFieldBool(data[n:], 3, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m EnumOptions) MarshalSize() (n int) {
	n += proto.SizeFieldBool(2, m.AllowAlias)
	n += proto.SizeFieldBool(3, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m EnumValueOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m EnumValueOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldBool(data[n:], 1, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m EnumValueOptions) MarshalSize() (n int) {
	n += proto.SizeFieldBool(1, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m ServiceOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m ServiceOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldBool(data[n:], 33, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m ServiceOptions) MarshalSize() (n int) {
	n += proto.SizeFieldBool(33, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m MethodOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m MethodOptions) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldBool(data[n:], 33, m.Deprecated)
	n += proto.WriteFieldEnum(data[n:], 34, int(m.IdempotencyLevel))
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 999, x)
		}
	}
	return n
}
func (m MethodOptions) MarshalSize() (n int) {
	n += proto.SizeFieldBool(33, m.Deprecated)
	n += proto.SizeFieldEnum(34, int(m.IdempotencyLevel))
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += proto.SizeFieldMessage(999, x)
		}
	}
	return n
}
func (m UninterpretedOption) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m UninterpretedOption) MarshalTo(data []byte) (n int) {
	for _, x := range m.Name {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 2, x)
		}
	}
	n += proto.WriteFieldString(data[n:], 3, m.IdentifierValue)
	n += proto.WriteFieldUInt64(data[n:], 4, m.PositiveIntValue)
	n += proto.WriteFieldInt64(data[n:], 5, m.NegativeIntValue)
	n += proto.WriteFieldDouble(data[n:], 6, m.DoubleValue)
	n += proto.WriteFieldBytes(data[n:], 7, m.StringValue)
	n += proto.WriteFieldString(data[n:], 8, m.AggregateValue)
	return n
}
func (m UninterpretedOption) MarshalSize() (n int) {
	for _, x := range m.Name {
		if x != nil {
			n += proto.SizeFieldMessage(2, x)
		}
	}
	n += proto.SizeFieldString(3, m.IdentifierValue)
	n += proto.SizeFieldUInt64(4, m.PositiveIntValue)
	n += proto.SizeFieldInt64(5, m.NegativeIntValue)
	n += proto.SizeFieldDouble(6, m.DoubleValue)
	n += proto.SizeFieldBytes(7, m.StringValue)
	n += proto.SizeFieldString(8, m.AggregateValue)
	return n
}
func (m UninterpretedOption_NamePart) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m UninterpretedOption_NamePart) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.NamePart)
	n += proto.WriteFieldBool(data[n:], 2, m.IsExtension)
	return n
}
func (m UninterpretedOption_NamePart) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.NamePart)
	n += proto.SizeFieldBool(2, m.IsExtension)
	return n
}
func (m SourceCodeInfo) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m SourceCodeInfo) MarshalTo(data []byte) (n int) {
	for _, x := range m.Location {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 1, x)
		}
	}
	return n
}
func (m SourceCodeInfo) MarshalSize() (n int) {
	for _, x := range m.Location {
		if x != nil {
			n += proto.SizeFieldMessage(1, x)
		}
	}
	return n
}
func (m SourceCodeInfo_Location) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m SourceCodeInfo_Location) MarshalTo(data []byte) (n int) {
	for _, x := range m.Path {
		n += proto.WriteFieldInt32(data[n:], 1, x)
	}
	for _, x := range m.Span {
		n += proto.WriteFieldInt32(data[n:], 2, x)
	}
	n += proto.WriteFieldString(data[n:], 3, m.LeadingComments)
	n += proto.WriteFieldString(data[n:], 4, m.TrailingComments)
	for _, x := range m.LeadingDetachedComments {
		n += proto.WriteFieldString(data[n:], 6, x)
	}
	return n
}
func (m SourceCodeInfo_Location) MarshalSize() (n int) {
	for _, x := range m.Path {
		n += proto.SizeFieldInt32(1, x)
	}
	for _, x := range m.Span {
		n += proto.SizeFieldInt32(2, x)
	}
	n += proto.SizeFieldString(3, m.LeadingComments)
	n += proto.SizeFieldString(4, m.TrailingComments)
	for _, x := range m.LeadingDetachedComments {
		n += proto.SizeFieldString(6, x)
	}
	return n
}
func (m GeneratedCodeInfo) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m GeneratedCodeInfo) MarshalTo(data []byte) (n int) {
	for _, x := range m.Annotation {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 1, x)
		}
	}
	return n
}
func (m GeneratedCodeInfo) MarshalSize() (n int) {
	for _, x := range m.Annotation {
		if x != nil {
			n += proto.SizeFieldMessage(1, x)
		}
	}
	return n
}
func (m GeneratedCodeInfo_Annotation) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m GeneratedCodeInfo_Annotation) MarshalTo(data []byte) (n int) {
	for _, x := range m.Path {
		n += proto.WriteFieldInt32(data[n:], 1, x)
	}
	n += proto.WriteFieldString(data[n:], 2, m.SourceFile)
	n += proto.WriteFieldInt32(data[n:], 3, m.Begin)
	n += proto.WriteFieldInt32(data[n:], 4, m.End)
	return n
}
func (m GeneratedCodeInfo_Annotation) MarshalSize() (n int) {
	for _, x := range m.Path {
		n += proto.SizeFieldInt32(1, x)
	}
	n += proto.SizeFieldString(2, m.SourceFile)
	n += proto.SizeFieldInt32(3, m.Begin)
	n += proto.SizeFieldInt32(4, m.End)
	return n
}
