package descriptor
import "github.com/kixelated/protohack/proto"
import "errors"
var _ = proto.WIRETYPE_VARINT
var _ = errors.New("")
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
	_, err = m.MarshalTo(data)
	return data, err
}
func (m FileDescriptorSet) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.File {
		if x != nil {
			n += copy(data[n:], []byte{0xa})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m FileDescriptorSet) MarshalSize() (n int) {
	for _, x := range m.File {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m FileDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m FileDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	if m.Package != "" {
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteStringLength(data[n:], m.Package)
	}
	for _, x := range m.Dependency {
		if x != "" {
			n += copy(data[n:], []byte{0x1a})
			n += proto.WriteStringLength(data[n:], x)
		}
	}
	for _, x := range m.MessageType {
		if x != nil {
			n += copy(data[n:], []byte{0x22})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += copy(data[n:], []byte{0x2a})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.Service {
		if x != nil {
			n += copy(data[n:], []byte{0x32})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += copy(data[n:], []byte{0x3a})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x42})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	if m.SourceCodeInfo != nil {
		n += copy(data[n:], []byte{0x4a})
		n += proto.WriteMessageLength(data[n:], m.SourceCodeInfo)
	}
	for _, x := range m.PublicDependency {
		if x != 0 {
			n += copy(data[n:], []byte{0x50})
			n += proto.WriteInt32(data[n:], x)
		}
	}
	for _, x := range m.WeakDependency {
		if x != 0 {
			n += copy(data[n:], []byte{0x58})
			n += proto.WriteInt32(data[n:], x)
		}
	}
	if m.Syntax != "" {
		n += copy(data[n:], []byte{0x62})
		n += proto.WriteStringLength(data[n:], m.Syntax)
	}
	return n, nil
}
func (m FileDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	if m.Package != "" {
		n += 1 + proto.SizeStringLength(m.Package)
	}
	for _, x := range m.Dependency {
		if x != "" {
			n += 1 + proto.SizeStringLength(x)
		}
	}
	for _, x := range m.MessageType {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.Service {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	if m.SourceCodeInfo != nil {
		n += 1 + proto.SizeMessageLength(m.SourceCodeInfo)
	}
	for _, x := range m.PublicDependency {
		if x != 0 {
			n += 1 + proto.SizeInt32(x)
		}
	}
	for _, x := range m.WeakDependency {
		if x != 0 {
			n += 1 + proto.SizeInt32(x)
		}
	}
	if m.Syntax != "" {
		n += 1 + proto.SizeStringLength(m.Syntax)
	}
	return n
}
func (m DescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m DescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	for _, x := range m.Field {
		if x != nil {
			n += copy(data[n:], []byte{0x12})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.NestedType {
		if x != nil {
			n += copy(data[n:], []byte{0x1a})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += copy(data[n:], []byte{0x22})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.ExtensionRange {
		if x != nil {
			n += copy(data[n:], []byte{0x2a})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += copy(data[n:], []byte{0x32})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x3a})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	for _, x := range m.OneofDecl {
		if x != nil {
			n += copy(data[n:], []byte{0x42})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.ReservedRange {
		if x != nil {
			n += copy(data[n:], []byte{0x4a})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	for _, x := range m.ReservedName {
		if x != "" {
			n += copy(data[n:], []byte{0x52})
			n += proto.WriteStringLength(data[n:], x)
		}
	}
	return n, nil
}
func (m DescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	for _, x := range m.Field {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.NestedType {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.EnumType {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.ExtensionRange {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.Extension {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	for _, x := range m.OneofDecl {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.ReservedRange {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	for _, x := range m.ReservedName {
		if x != "" {
			n += 1 + proto.SizeStringLength(x)
		}
	}
	return n
}
func (m DescriptorProto_ExtensionRange) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m DescriptorProto_ExtensionRange) MarshalTo(data []byte) (n int, err error) {
	if m.Start != 0 {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteInt32(data[n:], m.Start)
	}
	if m.End != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteInt32(data[n:], m.End)
	}
	return n, nil
}
func (m DescriptorProto_ExtensionRange) MarshalSize() (n int) {
	if m.Start != 0 {
		n += 1 + proto.SizeInt32(m.Start)
	}
	if m.End != 0 {
		n += 1 + proto.SizeInt32(m.End)
	}
	return n
}
func (m DescriptorProto_ReservedRange) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m DescriptorProto_ReservedRange) MarshalTo(data []byte) (n int, err error) {
	if m.Start != 0 {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteInt32(data[n:], m.Start)
	}
	if m.End != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteInt32(data[n:], m.End)
	}
	return n, nil
}
func (m DescriptorProto_ReservedRange) MarshalSize() (n int) {
	if m.Start != 0 {
		n += 1 + proto.SizeInt32(m.Start)
	}
	if m.End != 0 {
		n += 1 + proto.SizeInt32(m.End)
	}
	return n
}
func (m FieldDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m FieldDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	if m.Extendee != "" {
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteStringLength(data[n:], m.Extendee)
	}
	if m.Number != 0 {
		n += copy(data[n:], []byte{0x18})
		n += proto.WriteInt32(data[n:], m.Number)
	}
	if int(m.Label) != 0 {
		n += copy(data[n:], []byte{0x20})
		n += proto.WriteEnum(data[n:], int(m.Label))
	}
	if int(m.Type) != 0 {
		n += copy(data[n:], []byte{0x28})
		n += proto.WriteEnum(data[n:], int(m.Type))
	}
	if m.TypeName != "" {
		n += copy(data[n:], []byte{0x32})
		n += proto.WriteStringLength(data[n:], m.TypeName)
	}
	if m.DefaultValue != "" {
		n += copy(data[n:], []byte{0x3a})
		n += proto.WriteStringLength(data[n:], m.DefaultValue)
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x42})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	if m.OneofIndex != 0 {
		n += copy(data[n:], []byte{0x48})
		n += proto.WriteInt32(data[n:], m.OneofIndex)
	}
	if m.JsonName != "" {
		n += copy(data[n:], []byte{0x52})
		n += proto.WriteStringLength(data[n:], m.JsonName)
	}
	return n, nil
}
func (m FieldDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	if m.Extendee != "" {
		n += 1 + proto.SizeStringLength(m.Extendee)
	}
	if m.Number != 0 {
		n += 1 + proto.SizeInt32(m.Number)
	}
	if int(m.Label) != 0 {
		n += 1 + proto.SizeEnum(int(m.Label))
	}
	if int(m.Type) != 0 {
		n += 1 + proto.SizeEnum(int(m.Type))
	}
	if m.TypeName != "" {
		n += 1 + proto.SizeStringLength(m.TypeName)
	}
	if m.DefaultValue != "" {
		n += 1 + proto.SizeStringLength(m.DefaultValue)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	if m.OneofIndex != 0 {
		n += 1 + proto.SizeInt32(m.OneofIndex)
	}
	if m.JsonName != "" {
		n += 1 + proto.SizeStringLength(m.JsonName)
	}
	return n
}
func (m OneofDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m OneofDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	return n, nil
}
func (m OneofDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	return n
}
func (m EnumDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m EnumDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	for _, x := range m.Value {
		if x != nil {
			n += copy(data[n:], []byte{0x12})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	return n, nil
}
func (m EnumDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	for _, x := range m.Value {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	return n
}
func (m EnumValueDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m EnumValueDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	if m.Number != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteInt32(data[n:], m.Number)
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	return n, nil
}
func (m EnumValueDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	if m.Number != 0 {
		n += 1 + proto.SizeInt32(m.Number)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	return n
}
func (m ServiceDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m ServiceDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	for _, x := range m.Method {
		if x != nil {
			n += copy(data[n:], []byte{0x12})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	return n, nil
}
func (m ServiceDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	for _, x := range m.Method {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	return n
}
func (m MethodDescriptorProto) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m MethodDescriptorProto) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Name)
	}
	if m.InputType != "" {
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteStringLength(data[n:], m.InputType)
	}
	if m.OutputType != "" {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteStringLength(data[n:], m.OutputType)
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x22})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	if m.ClientStreaming != false {
		n += copy(data[n:], []byte{0x28})
		n += proto.WriteBool(data[n:], m.ClientStreaming)
	}
	if m.ServerStreaming != false {
		n += copy(data[n:], []byte{0x30})
		n += proto.WriteBool(data[n:], m.ServerStreaming)
	}
	return n, nil
}
func (m MethodDescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	if m.InputType != "" {
		n += 1 + proto.SizeStringLength(m.InputType)
	}
	if m.OutputType != "" {
		n += 1 + proto.SizeStringLength(m.OutputType)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	if m.ClientStreaming != false {
		n += 1 + proto.SizeBool(m.ClientStreaming)
	}
	if m.ServerStreaming != false {
		n += 1 + proto.SizeBool(m.ServerStreaming)
	}
	return n
}
func (m FileOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m FileOptions) MarshalTo(data []byte) (n int, err error) {
	if m.JavaPackage != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.JavaPackage)
	}
	if m.JavaOuterClassname != "" {
		n += copy(data[n:], []byte{0x42})
		n += proto.WriteStringLength(data[n:], m.JavaOuterClassname)
	}
	if int(m.OptimizeFor) != 0 {
		n += copy(data[n:], []byte{0x48})
		n += proto.WriteEnum(data[n:], int(m.OptimizeFor))
	}
	if m.JavaMultipleFiles != false {
		n += copy(data[n:], []byte{0x50})
		n += proto.WriteBool(data[n:], m.JavaMultipleFiles)
	}
	if m.GoPackage != "" {
		n += copy(data[n:], []byte{0x5a})
		n += proto.WriteStringLength(data[n:], m.GoPackage)
	}
	if m.CcGenericServices != false {
		n += copy(data[n:], []byte{0x80, 0x1})
		n += proto.WriteBool(data[n:], m.CcGenericServices)
	}
	if m.JavaGenericServices != false {
		n += copy(data[n:], []byte{0x88, 0x1})
		n += proto.WriteBool(data[n:], m.JavaGenericServices)
	}
	if m.PyGenericServices != false {
		n += copy(data[n:], []byte{0x90, 0x1})
		n += proto.WriteBool(data[n:], m.PyGenericServices)
	}
	if m.JavaGenerateEqualsAndHash != false {
		n += copy(data[n:], []byte{0xa0, 0x1})
		n += proto.WriteBool(data[n:], m.JavaGenerateEqualsAndHash)
	}
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0xb8, 0x1})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	if m.JavaStringCheckUtf8 != false {
		n += copy(data[n:], []byte{0xd8, 0x1})
		n += proto.WriteBool(data[n:], m.JavaStringCheckUtf8)
	}
	if m.CcEnableArenas != false {
		n += copy(data[n:], []byte{0xf8, 0x1})
		n += proto.WriteBool(data[n:], m.CcEnableArenas)
	}
	if m.ObjcClassPrefix != "" {
		n += copy(data[n:], []byte{0xa2, 0x2})
		n += proto.WriteStringLength(data[n:], m.ObjcClassPrefix)
	}
	if m.CsharpNamespace != "" {
		n += copy(data[n:], []byte{0xaa, 0x2})
		n += proto.WriteStringLength(data[n:], m.CsharpNamespace)
	}
	if m.SwiftPrefix != "" {
		n += copy(data[n:], []byte{0xba, 0x2})
		n += proto.WriteStringLength(data[n:], m.SwiftPrefix)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m FileOptions) MarshalSize() (n int) {
	if m.JavaPackage != "" {
		n += 1 + proto.SizeStringLength(m.JavaPackage)
	}
	if m.JavaOuterClassname != "" {
		n += 1 + proto.SizeStringLength(m.JavaOuterClassname)
	}
	if int(m.OptimizeFor) != 0 {
		n += 1 + proto.SizeEnum(int(m.OptimizeFor))
	}
	if m.JavaMultipleFiles != false {
		n += 1 + proto.SizeBool(m.JavaMultipleFiles)
	}
	if m.GoPackage != "" {
		n += 1 + proto.SizeStringLength(m.GoPackage)
	}
	if m.CcGenericServices != false {
		n += 2 + proto.SizeBool(m.CcGenericServices)
	}
	if m.JavaGenericServices != false {
		n += 2 + proto.SizeBool(m.JavaGenericServices)
	}
	if m.PyGenericServices != false {
		n += 2 + proto.SizeBool(m.PyGenericServices)
	}
	if m.JavaGenerateEqualsAndHash != false {
		n += 2 + proto.SizeBool(m.JavaGenerateEqualsAndHash)
	}
	if m.Deprecated != false {
		n += 2 + proto.SizeBool(m.Deprecated)
	}
	if m.JavaStringCheckUtf8 != false {
		n += 2 + proto.SizeBool(m.JavaStringCheckUtf8)
	}
	if m.CcEnableArenas != false {
		n += 2 + proto.SizeBool(m.CcEnableArenas)
	}
	if m.ObjcClassPrefix != "" {
		n += 2 + proto.SizeStringLength(m.ObjcClassPrefix)
	}
	if m.CsharpNamespace != "" {
		n += 2 + proto.SizeStringLength(m.CsharpNamespace)
	}
	if m.SwiftPrefix != "" {
		n += 2 + proto.SizeStringLength(m.SwiftPrefix)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m MessageOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m MessageOptions) MarshalTo(data []byte) (n int, err error) {
	if m.MessageSetWireFormat != false {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteBool(data[n:], m.MessageSetWireFormat)
	}
	if m.NoStandardDescriptorAccessor != false {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteBool(data[n:], m.NoStandardDescriptorAccessor)
	}
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0x18})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	if m.MapEntry != false {
		n += copy(data[n:], []byte{0x38})
		n += proto.WriteBool(data[n:], m.MapEntry)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m MessageOptions) MarshalSize() (n int) {
	if m.MessageSetWireFormat != false {
		n += 1 + proto.SizeBool(m.MessageSetWireFormat)
	}
	if m.NoStandardDescriptorAccessor != false {
		n += 1 + proto.SizeBool(m.NoStandardDescriptorAccessor)
	}
	if m.Deprecated != false {
		n += 1 + proto.SizeBool(m.Deprecated)
	}
	if m.MapEntry != false {
		n += 1 + proto.SizeBool(m.MapEntry)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m FieldOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m FieldOptions) MarshalTo(data []byte) (n int, err error) {
	if int(m.Ctype) != 0 {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteEnum(data[n:], int(m.Ctype))
	}
	if m.Packed != false {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteBool(data[n:], m.Packed)
	}
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0x18})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	if m.Lazy != false {
		n += copy(data[n:], []byte{0x28})
		n += proto.WriteBool(data[n:], m.Lazy)
	}
	if int(m.Jstype) != 0 {
		n += copy(data[n:], []byte{0x30})
		n += proto.WriteEnum(data[n:], int(m.Jstype))
	}
	if m.Weak != false {
		n += copy(data[n:], []byte{0x50})
		n += proto.WriteBool(data[n:], m.Weak)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m FieldOptions) MarshalSize() (n int) {
	if int(m.Ctype) != 0 {
		n += 1 + proto.SizeEnum(int(m.Ctype))
	}
	if m.Packed != false {
		n += 1 + proto.SizeBool(m.Packed)
	}
	if m.Deprecated != false {
		n += 1 + proto.SizeBool(m.Deprecated)
	}
	if m.Lazy != false {
		n += 1 + proto.SizeBool(m.Lazy)
	}
	if int(m.Jstype) != 0 {
		n += 1 + proto.SizeEnum(int(m.Jstype))
	}
	if m.Weak != false {
		n += 1 + proto.SizeBool(m.Weak)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m OneofOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m OneofOptions) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m OneofOptions) MarshalSize() (n int) {
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m EnumOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m EnumOptions) MarshalTo(data []byte) (n int, err error) {
	if m.AllowAlias != false {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteBool(data[n:], m.AllowAlias)
	}
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0x18})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m EnumOptions) MarshalSize() (n int) {
	if m.AllowAlias != false {
		n += 1 + proto.SizeBool(m.AllowAlias)
	}
	if m.Deprecated != false {
		n += 1 + proto.SizeBool(m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m EnumValueOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m EnumValueOptions) MarshalTo(data []byte) (n int, err error) {
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m EnumValueOptions) MarshalSize() (n int) {
	if m.Deprecated != false {
		n += 1 + proto.SizeBool(m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m ServiceOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m ServiceOptions) MarshalTo(data []byte) (n int, err error) {
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0x88, 0x2})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m ServiceOptions) MarshalSize() (n int) {
	if m.Deprecated != false {
		n += 2 + proto.SizeBool(m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m MethodOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m MethodOptions) MarshalTo(data []byte) (n int, err error) {
	if m.Deprecated != false {
		n += copy(data[n:], []byte{0x88, 0x2})
		n += proto.WriteBool(data[n:], m.Deprecated)
	}
	if int(m.IdempotencyLevel) != 0 {
		n += copy(data[n:], []byte{0x90, 0x2})
		n += proto.WriteEnum(data[n:], int(m.IdempotencyLevel))
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += copy(data[n:], []byte{0xba, 0x3e})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m MethodOptions) MarshalSize() (n int) {
	if m.Deprecated != false {
		n += 2 + proto.SizeBool(m.Deprecated)
	}
	if int(m.IdempotencyLevel) != 0 {
		n += 2 + proto.SizeEnum(int(m.IdempotencyLevel))
	}
	for _, x := range m.UninterpretedOption {
		if x != nil {
			n += 2 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m UninterpretedOption) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m UninterpretedOption) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Name {
		if x != nil {
			n += copy(data[n:], []byte{0x12})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	if m.IdentifierValue != "" {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteStringLength(data[n:], m.IdentifierValue)
	}
	if m.PositiveIntValue != 0 {
		n += copy(data[n:], []byte{0x20})
		n += proto.WriteUInt64(data[n:], m.PositiveIntValue)
	}
	if m.NegativeIntValue != 0 {
		n += copy(data[n:], []byte{0x28})
		n += proto.WriteInt64(data[n:], m.NegativeIntValue)
	}
	if m.DoubleValue != 0 {
		n += copy(data[n:], []byte{0x31})
		n += proto.WriteDouble(data[n:], m.DoubleValue)
	}
	if m.StringValue != nil {
		n += copy(data[n:], []byte{0x3a})
		n += proto.WriteBytesLength(data[n:], m.StringValue)
	}
	if m.AggregateValue != "" {
		n += copy(data[n:], []byte{0x42})
		n += proto.WriteStringLength(data[n:], m.AggregateValue)
	}
	return n, nil
}
func (m UninterpretedOption) MarshalSize() (n int) {
	for _, x := range m.Name {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	if m.IdentifierValue != "" {
		n += 1 + proto.SizeStringLength(m.IdentifierValue)
	}
	if m.PositiveIntValue != 0 {
		n += 1 + proto.SizeUInt64(m.PositiveIntValue)
	}
	if m.NegativeIntValue != 0 {
		n += 1 + proto.SizeInt64(m.NegativeIntValue)
	}
	if m.DoubleValue != 0 {
		n += 1 + proto.SizeDouble(m.DoubleValue)
	}
	if m.StringValue != nil {
		n += 1 + proto.SizeBytesLength(m.StringValue)
	}
	if m.AggregateValue != "" {
		n += 1 + proto.SizeStringLength(m.AggregateValue)
	}
	return n
}
func (m UninterpretedOption_NamePart) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m UninterpretedOption_NamePart) MarshalTo(data []byte) (n int, err error) {
	if m.NamePart != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.NamePart)
	} else {
		return 0, errors.New("missing required field: NamePart")
	}
	if m.IsExtension != false {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteBool(data[n:], m.IsExtension)
	} else {
		return 0, errors.New("missing required field: IsExtension")
	}
	return n, nil
}
func (m UninterpretedOption_NamePart) MarshalSize() (n int) {
	if m.NamePart != "" {
		n += 1 + proto.SizeStringLength(m.NamePart)
	}
	if m.IsExtension != false {
		n += 1 + proto.SizeBool(m.IsExtension)
	}
	return n
}
func (m SourceCodeInfo) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m SourceCodeInfo) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Location {
		if x != nil {
			n += copy(data[n:], []byte{0xa})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m SourceCodeInfo) MarshalSize() (n int) {
	for _, x := range m.Location {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m SourceCodeInfo_Location) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m SourceCodeInfo_Location) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Path {
		if x != 0 {
			n += copy(data[n:], []byte{0x8})
			n += proto.WriteInt32(data[n:], x)
		}
	}
	for _, x := range m.Span {
		if x != 0 {
			n += copy(data[n:], []byte{0x10})
			n += proto.WriteInt32(data[n:], x)
		}
	}
	if m.LeadingComments != "" {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteStringLength(data[n:], m.LeadingComments)
	}
	if m.TrailingComments != "" {
		n += copy(data[n:], []byte{0x22})
		n += proto.WriteStringLength(data[n:], m.TrailingComments)
	}
	for _, x := range m.LeadingDetachedComments {
		if x != "" {
			n += copy(data[n:], []byte{0x32})
			n += proto.WriteStringLength(data[n:], x)
		}
	}
	return n, nil
}
func (m SourceCodeInfo_Location) MarshalSize() (n int) {
	for _, x := range m.Path {
		if x != 0 {
			n += 1 + proto.SizeInt32(x)
		}
	}
	for _, x := range m.Span {
		if x != 0 {
			n += 1 + proto.SizeInt32(x)
		}
	}
	if m.LeadingComments != "" {
		n += 1 + proto.SizeStringLength(m.LeadingComments)
	}
	if m.TrailingComments != "" {
		n += 1 + proto.SizeStringLength(m.TrailingComments)
	}
	for _, x := range m.LeadingDetachedComments {
		if x != "" {
			n += 1 + proto.SizeStringLength(x)
		}
	}
	return n
}
func (m GeneratedCodeInfo) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m GeneratedCodeInfo) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Annotation {
		if x != nil {
			n += copy(data[n:], []byte{0xa})
			n += proto.WriteMessageLength(data[n:], x)
		}
	}
	return n, nil
}
func (m GeneratedCodeInfo) MarshalSize() (n int) {
	for _, x := range m.Annotation {
		if x != nil {
			n += 1 + proto.SizeMessageLength(x)
		}
	}
	return n
}
func (m GeneratedCodeInfo_Annotation) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m GeneratedCodeInfo_Annotation) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Path {
		if x != 0 {
			n += copy(data[n:], []byte{0x8})
			n += proto.WriteInt32(data[n:], x)
		}
	}
	if m.SourceFile != "" {
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteStringLength(data[n:], m.SourceFile)
	}
	if m.Begin != 0 {
		n += copy(data[n:], []byte{0x18})
		n += proto.WriteInt32(data[n:], m.Begin)
	}
	if m.End != 0 {
		n += copy(data[n:], []byte{0x20})
		n += proto.WriteInt32(data[n:], m.End)
	}
	return n, nil
}
func (m GeneratedCodeInfo_Annotation) MarshalSize() (n int) {
	for _, x := range m.Path {
		if x != 0 {
			n += 1 + proto.SizeInt32(x)
		}
	}
	if m.SourceFile != "" {
		n += 1 + proto.SizeStringLength(m.SourceFile)
	}
	if m.Begin != 0 {
		n += 1 + proto.SizeInt32(m.Begin)
	}
	if m.End != 0 {
		n += 1 + proto.SizeInt32(m.End)
	}
	return n
}
