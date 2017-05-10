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
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m FileDescriptorSet) MarshalSize() (n int) {
	for _, x := range m.File {
		n += 1 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *FileDescriptorSet) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			temp := new(FileDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.File = append(m.File, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteStringLength(data[n:], x)
	}
	for _, x := range m.MessageType {
		n += copy(data[n:], []byte{0x22})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.EnumType {
		n += copy(data[n:], []byte{0x2a})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.Service {
		n += copy(data[n:], []byte{0x32})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.Extension {
		n += copy(data[n:], []byte{0x3a})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += copy(data[n:], []byte{0x50})
		n += proto.WriteInt32(data[n:], x)
	}
	for _, x := range m.WeakDependency {
		n += copy(data[n:], []byte{0x58})
		n += proto.WriteInt32(data[n:], x)
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
		n += 1 + proto.SizeStringLength(x)
	}
	for _, x := range m.MessageType {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.EnumType {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.Service {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.Extension {
		n += 1 + proto.SizeMessageLength(x)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	if m.SourceCodeInfo != nil {
		n += 1 + proto.SizeMessageLength(m.SourceCodeInfo)
	}
	for _, x := range m.PublicDependency {
		n += 1 + proto.SizeInt32(x)
	}
	for _, x := range m.WeakDependency {
		n += 1 + proto.SizeInt32(x)
	}
	if m.Syntax != "" {
		n += 1 + proto.SizeStringLength(m.Syntax)
	}
	return n
}
func (m *FileDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			m.Package, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 3:
			temp, err := r.ReadString(t)
			if err != nil {
				return err
			}
			m.Dependency = append(m.Dependency, temp)
		case 4:
			temp := new(DescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.MessageType = append(m.MessageType, temp)
		case 5:
			temp := new(EnumDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.EnumType = append(m.EnumType, temp)
		case 6:
			temp := new(ServiceDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Service = append(m.Service, temp)
		case 7:
			temp := new(FieldDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Extension = append(m.Extension, temp)
		case 8:
			m.Options = new(FileOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		case 9:
			m.SourceCodeInfo = new(SourceCodeInfo)
			err = r.ReadToMessage(t, m.SourceCodeInfo)
			if err != nil {
				return err
			}
		case 10:
			temp, err := r.ReadInt32(t)
			if err != nil {
				return err
			}
			m.PublicDependency = append(m.PublicDependency, temp)
		case 11:
			temp, err := r.ReadInt32(t)
			if err != nil {
				return err
			}
			m.WeakDependency = append(m.WeakDependency, temp)
		case 12:
			m.Syntax, err = r.ReadString(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.NestedType {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.EnumType {
		n += copy(data[n:], []byte{0x22})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.ExtensionRange {
		n += copy(data[n:], []byte{0x2a})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.Extension {
		n += copy(data[n:], []byte{0x32})
		n += proto.WriteMessageLength(data[n:], x)
	}
	if m.Options != nil {
		n += copy(data[n:], []byte{0x3a})
		n += proto.WriteMessageLength(data[n:], m.Options)
	}
	for _, x := range m.OneofDecl {
		n += copy(data[n:], []byte{0x42})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.ReservedRange {
		n += copy(data[n:], []byte{0x4a})
		n += proto.WriteMessageLength(data[n:], x)
	}
	for _, x := range m.ReservedName {
		n += copy(data[n:], []byte{0x52})
		n += proto.WriteStringLength(data[n:], x)
	}
	return n, nil
}
func (m DescriptorProto) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeStringLength(m.Name)
	}
	for _, x := range m.Field {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.NestedType {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.EnumType {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.ExtensionRange {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.Extension {
		n += 1 + proto.SizeMessageLength(x)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	for _, x := range m.OneofDecl {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.ReservedRange {
		n += 1 + proto.SizeMessageLength(x)
	}
	for _, x := range m.ReservedName {
		n += 1 + proto.SizeStringLength(x)
	}
	return n
}
func (m *DescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			temp := new(FieldDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Field = append(m.Field, temp)
		case 3:
			temp := new(DescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.NestedType = append(m.NestedType, temp)
		case 4:
			temp := new(EnumDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.EnumType = append(m.EnumType, temp)
		case 5:
			temp := new(DescriptorProto_ExtensionRange)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.ExtensionRange = append(m.ExtensionRange, temp)
		case 6:
			temp := new(FieldDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Extension = append(m.Extension, temp)
		case 7:
			m.Options = new(MessageOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		case 8:
			temp := new(OneofDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.OneofDecl = append(m.OneofDecl, temp)
		case 9:
			temp := new(DescriptorProto_ReservedRange)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.ReservedRange = append(m.ReservedRange, temp)
		case 10:
			temp, err := r.ReadString(t)
			if err != nil {
				return err
			}
			m.ReservedName = append(m.ReservedName, temp)
		}
	}
	return nil
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
func (m *DescriptorProto_ExtensionRange) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Start, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		case 2:
			m.End, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
func (m *DescriptorProto_ReservedRange) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Start, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		case 2:
			m.End, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
func (m *FieldDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			m.Extendee, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 3:
			m.Number, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		case 4:
			temp, err := r.ReadEnum(t)
			if err != nil {
				return err
			}
			m.Label = FieldDescriptorProto_Label(temp)
		case 5:
			temp, err := r.ReadEnum(t)
			if err != nil {
				return err
			}
			m.Type = FieldDescriptorProto_Type(temp)
		case 6:
			m.TypeName, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 7:
			m.DefaultValue, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 8:
			m.Options = new(FieldOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		case 9:
			m.OneofIndex, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		case 10:
			m.JsonName, err = r.ReadString(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
func (m *OneofDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			m.Options = new(OneofOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 1 + proto.SizeMessageLength(x)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	return n
}
func (m *EnumDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			temp := new(EnumValueDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Value = append(m.Value, temp)
		case 3:
			m.Options = new(EnumOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
func (m *EnumValueDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			m.Number, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		case 3:
			m.Options = new(EnumValueOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 1 + proto.SizeMessageLength(x)
	}
	if m.Options != nil {
		n += 1 + proto.SizeMessageLength(m.Options)
	}
	return n
}
func (m *ServiceDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			temp := new(MethodDescriptorProto)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Method = append(m.Method, temp)
		case 3:
			m.Options = new(ServiceOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
func (m *MethodDescriptorProto) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			m.InputType, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 3:
			m.OutputType, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 4:
			m.Options = new(MethodOptions)
			err = r.ReadToMessage(t, m.Options)
			if err != nil {
				return err
			}
		case 5:
			m.ClientStreaming, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 6:
			m.ServerStreaming, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *FileOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.JavaPackage, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 8:
			m.JavaOuterClassname, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 9:
			temp, err := r.ReadEnum(t)
			if err != nil {
				return err
			}
			m.OptimizeFor = FileOptions_OptimizeMode(temp)
		case 10:
			m.JavaMultipleFiles, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 11:
			m.GoPackage, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 16:
			m.CcGenericServices, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 17:
			m.JavaGenericServices, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 18:
			m.PyGenericServices, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 20:
			m.JavaGenerateEqualsAndHash, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 23:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 27:
			m.JavaStringCheckUtf8, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 31:
			m.CcEnableArenas, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 36:
			m.ObjcClassPrefix, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 37:
			m.CsharpNamespace, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 39:
			m.SwiftPrefix, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *MessageOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.MessageSetWireFormat, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 2:
			m.NoStandardDescriptorAccessor, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 3:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 7:
			m.MapEntry, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *FieldOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			temp, err := r.ReadEnum(t)
			if err != nil {
				return err
			}
			m.Ctype = FieldOptions_CType(temp)
		case 2:
			m.Packed, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 3:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 5:
			m.Lazy, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 6:
			temp, err := r.ReadEnum(t)
			if err != nil {
				return err
			}
			m.Jstype = FieldOptions_JSType(temp)
		case 10:
			m.Weak, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
}
func (m OneofOptions) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m OneofOptions) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.UninterpretedOption {
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m OneofOptions) MarshalSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *OneofOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *EnumOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 2:
			m.AllowAlias, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 3:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m EnumValueOptions) MarshalSize() (n int) {
	if m.Deprecated != false {
		n += 1 + proto.SizeBool(m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *EnumValueOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m ServiceOptions) MarshalSize() (n int) {
	if m.Deprecated != false {
		n += 2 + proto.SizeBool(m.Deprecated)
	}
	for _, x := range m.UninterpretedOption {
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *ServiceOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 33:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
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
		n += copy(data[n:], []byte{0xba, 0x3e})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 2 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *MethodOptions) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 33:
			m.Deprecated, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		case 34:
			temp, err := r.ReadEnum(t)
			if err != nil {
				return err
			}
			m.IdempotencyLevel = MethodOptions_IdempotencyLevel(temp)
		case 999:
			temp := new(UninterpretedOption)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.UninterpretedOption = append(m.UninterpretedOption, temp)
		}
	}
	return nil
}
func (m UninterpretedOption) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m UninterpretedOption) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Name {
		n += copy(data[n:], []byte{0x12})
		n += proto.WriteMessageLength(data[n:], x)
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
		n += 1 + proto.SizeMessageLength(x)
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
func (m *UninterpretedOption) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 2:
			temp := new(UninterpretedOption_NamePart)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Name = append(m.Name, temp)
		case 3:
			m.IdentifierValue, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 4:
			m.PositiveIntValue, err = r.ReadUInt64(t)
			if err != nil {
				return err
			}
		case 5:
			m.NegativeIntValue, err = r.ReadInt64(t)
			if err != nil {
				return err
			}
		case 6:
			m.DoubleValue, err = r.ReadDouble(t)
			if err != nil {
				return err
			}
		case 7:
			m.StringValue, err = r.ReadBytes(t)
			if err != nil {
				return err
			}
		case 8:
			m.AggregateValue, err = r.ReadString(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		return n, errors.New("missing required field: NamePart")
	}
	if m.IsExtension != false {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteBool(data[n:], m.IsExtension)
	} else {
		return n, errors.New("missing required field: IsExtension")
	}
	return n, nil
}
func (m UninterpretedOption_NamePart) MarshalSize() (n int) {
	n += 1 + proto.SizeStringLength(m.NamePart)
	n += 1 + proto.SizeBool(m.IsExtension)
	return n
}
func (m *UninterpretedOption_NamePart) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.NamePart, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 2:
			m.IsExtension, err = r.ReadBool(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (m SourceCodeInfo) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m SourceCodeInfo) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Location {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m SourceCodeInfo) MarshalSize() (n int) {
	for _, x := range m.Location {
		n += 1 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *SourceCodeInfo) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			temp := new(SourceCodeInfo_Location)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Location = append(m.Location, temp)
		}
	}
	return nil
}
func (m SourceCodeInfo_Location) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m SourceCodeInfo_Location) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Path {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteInt32(data[n:], x)
	}
	for _, x := range m.Span {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteInt32(data[n:], x)
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
		n += copy(data[n:], []byte{0x32})
		n += proto.WriteStringLength(data[n:], x)
	}
	return n, nil
}
func (m SourceCodeInfo_Location) MarshalSize() (n int) {
	for _, x := range m.Path {
		n += 1 + proto.SizeInt32(x)
	}
	for _, x := range m.Span {
		n += 1 + proto.SizeInt32(x)
	}
	if m.LeadingComments != "" {
		n += 1 + proto.SizeStringLength(m.LeadingComments)
	}
	if m.TrailingComments != "" {
		n += 1 + proto.SizeStringLength(m.TrailingComments)
	}
	for _, x := range m.LeadingDetachedComments {
		n += 1 + proto.SizeStringLength(x)
	}
	return n
}
func (m *SourceCodeInfo_Location) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			temp, err := r.ReadInt32(t)
			if err != nil {
				return err
			}
			m.Path = append(m.Path, temp)
		case 2:
			temp, err := r.ReadInt32(t)
			if err != nil {
				return err
			}
			m.Span = append(m.Span, temp)
		case 3:
			m.LeadingComments, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 4:
			m.TrailingComments, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 6:
			temp, err := r.ReadString(t)
			if err != nil {
				return err
			}
			m.LeadingDetachedComments = append(m.LeadingDetachedComments, temp)
		}
	}
	return nil
}
func (m GeneratedCodeInfo) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m GeneratedCodeInfo) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Annotation {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m GeneratedCodeInfo) MarshalSize() (n int) {
	for _, x := range m.Annotation {
		n += 1 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *GeneratedCodeInfo) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			temp := new(GeneratedCodeInfo_Annotation)
			err = r.ReadToMessage(t, temp)
			if err != nil {
				return err
			}
			m.Annotation = append(m.Annotation, temp)
		}
	}
	return nil
}
func (m GeneratedCodeInfo_Annotation) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m GeneratedCodeInfo_Annotation) MarshalTo(data []byte) (n int, err error) {
	for _, x := range m.Path {
		n += copy(data[n:], []byte{0x8})
		n += proto.WriteInt32(data[n:], x)
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
		n += 1 + proto.SizeInt32(x)
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
func (m *GeneratedCodeInfo_Annotation) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			temp, err := r.ReadInt32(t)
			if err != nil {
				return err
			}
			m.Path = append(m.Path, temp)
		case 2:
			m.SourceFile, err = r.ReadString(t)
			if err != nil {
				return err
			}
		case 3:
			m.Begin, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		case 4:
			m.End, err = r.ReadInt32(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
