package descriptor

import "github.com/kixelated/protohack/proto"

var _ = proto.WIRETYPE_VARINT

type FileDescriptorSet struct {
	File []*FileDescriptorProto
}
type FileDescriptorProto struct {
	Name             string
	Package          string
	Dependency       []string
	PublicDependency []int32
	WeakDependency   []int32
	MessageType      []*DescriptorProto
	EnumType         []*EnumDescriptorProto
	Service          []*ServiceDescriptorProto
	Extension        []*FieldDescriptorProto
	Options          *FileOptions
	SourceCodeInfo   *SourceCodeInfo
	Syntax           string
}
type DescriptorProto struct {
	Name           string
	Field          []*FieldDescriptorProto
	Extension      []*FieldDescriptorProto
	NestedType     []*DescriptorProto
	EnumType       []*EnumDescriptorProto
	ExtensionRange []*DescriptorProto_ExtensionRange
	OneofDecl      []*OneofDescriptorProto
	Options        *MessageOptions
	ReservedRange  []*DescriptorProto_ReservedRange
	ReservedName   []string
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
	Label        FieldDescriptorProto_Label
	Type         FieldDescriptorProto_Type
	TypeName     string
	Extendee     string
	DefaultValue string
	OneofIndex   int32
	JsonName     string
	Options      *FieldOptions
}
type FieldDescriptorProto_Type int

const (
	FieldDescriptorProto_TYPE_DOUBLE   FieldDescriptorProto_Type = 1
	FieldDescriptorProto_TYPE_FLOAT    FieldDescriptorProto_Type = 2
	FieldDescriptorProto_TYPE_INT64    FieldDescriptorProto_Type = 3
	FieldDescriptorProto_TYPE_UINT64   FieldDescriptorProto_Type = 4
	FieldDescriptorProto_TYPE_INT32    FieldDescriptorProto_Type = 5
	FieldDescriptorProto_TYPE_FIXED64  FieldDescriptorProto_Type = 6
	FieldDescriptorProto_TYPE_FIXED32  FieldDescriptorProto_Type = 7
	FieldDescriptorProto_TYPE_BOOL     FieldDescriptorProto_Type = 8
	FieldDescriptorProto_TYPE_STRING   FieldDescriptorProto_Type = 9
	FieldDescriptorProto_TYPE_GROUP    FieldDescriptorProto_Type = 10
	FieldDescriptorProto_TYPE_MESSAGE  FieldDescriptorProto_Type = 11
	FieldDescriptorProto_TYPE_BYTES    FieldDescriptorProto_Type = 12
	FieldDescriptorProto_TYPE_UINT32   FieldDescriptorProto_Type = 13
	FieldDescriptorProto_TYPE_ENUM     FieldDescriptorProto_Type = 14
	FieldDescriptorProto_TYPE_SFIXED32 FieldDescriptorProto_Type = 15
	FieldDescriptorProto_TYPE_SFIXED64 FieldDescriptorProto_Type = 16
	FieldDescriptorProto_TYPE_SINT32   FieldDescriptorProto_Type = 17
	FieldDescriptorProto_TYPE_SINT64   FieldDescriptorProto_Type = 18
)

type FieldDescriptorProto_Label int

const (
	FieldDescriptorProto_LABEL_OPTIONAL FieldDescriptorProto_Label = 1
	FieldDescriptorProto_LABEL_REQUIRED FieldDescriptorProto_Label = 2
	FieldDescriptorProto_LABEL_REPEATED FieldDescriptorProto_Label = 3
)

type OneofDescriptorProto struct {
	Name    string
	Options *OneofOptions
}
type EnumDescriptorProto struct {
	Name    string
	Value   []*EnumValueDescriptorProto
	Options *EnumOptions
}
type EnumValueDescriptorProto struct {
	Name    string
	Number  int32
	Options *EnumValueOptions
}
type ServiceDescriptorProto struct {
	Name    string
	Method  []*MethodDescriptorProto
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
	OptimizeFor               FileOptions_OptimizeMode
	GoPackage                 string
	CcGenericServices         bool
	JavaGenericServices       bool
	PyGenericServices         bool
	Deprecated                bool
	CcEnableArenas            bool
	ObjcClassPrefix           string
	CsharpNamespace           string
	SwiftPrefix               string
	UninterpretedOption       []*UninterpretedOption
}
type FileOptions_OptimizeMode int

const (
	FileOptions_SPEED        FileOptions_OptimizeMode = 1
	FileOptions_CODE_SIZE    FileOptions_OptimizeMode = 2
	FileOptions_LITE_RUNTIME FileOptions_OptimizeMode = 3
)

type MessageOptions struct {
	MessageSetWireFormat         bool
	NoStandardDescriptorAccessor bool
	Deprecated                   bool
	MapEntry                     bool
	UninterpretedOption          []*UninterpretedOption
}
type FieldOptions struct {
	Ctype               FieldOptions_CType
	Packed              bool
	Jstype              FieldOptions_JSType
	Lazy                bool
	Deprecated          bool
	Weak                bool
	UninterpretedOption []*UninterpretedOption
}
type FieldOptions_CType int

const (
	FieldOptions_STRING       FieldOptions_CType = 0
	FieldOptions_CORD         FieldOptions_CType = 1
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
	AllowAlias          bool
	Deprecated          bool
	UninterpretedOption []*UninterpretedOption
}
type EnumValueOptions struct {
	Deprecated          bool
	UninterpretedOption []*UninterpretedOption
}
type ServiceOptions struct {
	Deprecated          bool
	UninterpretedOption []*UninterpretedOption
}
type MethodOptions struct {
	Deprecated          bool
	IdempotencyLevel    MethodOptions_IdempotencyLevel
	UninterpretedOption []*UninterpretedOption
}
type MethodOptions_IdempotencyLevel int

const (
	MethodOptions_IDEMPOTENCY_UNKNOWN MethodOptions_IdempotencyLevel = 0
	MethodOptions_NO_SIDE_EFFECTS     MethodOptions_IdempotencyLevel = 1
	MethodOptions_IDEMPOTENT          MethodOptions_IdempotencyLevel = 2
)

type UninterpretedOption struct {
	Name             []*UninterpretedOption_NamePart
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
	Location []*SourceCodeInfo_Location
}
type SourceCodeInfo_Location struct {
	Path                    []int32
	Span                    []int32
	LeadingComments         string
	TrailingComments        string
	LeadingDetachedComments []string
}
type GeneratedCodeInfo struct {
	Annotation []*GeneratedCodeInfo_Annotation
}
type GeneratedCodeInfo_Annotation struct {
	Path       []int32
	SourceFile string
	Begin      int32
	End        int32
}

func (m FileDescriptorSet) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.File {
		err = w.WriteMessage(1, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m FileDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	w.WriteString(2, m.Package)
	for _, x := range m.Dependency {
		w.WriteString(3, x)
	}
	for _, x := range m.PublicDependency {
		w.WriteInt32(10, x)
	}
	for _, x := range m.WeakDependency {
		w.WriteInt32(11, x)
	}
	for _, x := range m.MessageType {
		err = w.WriteMessage(4, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.EnumType {
		err = w.WriteMessage(5, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.Service {
		err = w.WriteMessage(6, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.Extension {
		err = w.WriteMessage(7, x)
		if err != nil {
			return nil, err
		}
	}
	err = w.WriteMessage(8, m.Options)
	if err != nil {
		return nil, err
	}
	err = w.WriteMessage(9, m.SourceCodeInfo)
	if err != nil {
		return nil, err
	}
	w.WriteString(12, m.Syntax)
	return w.Bytes(), nil
}
func (m DescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	for _, x := range m.Field {
		err = w.WriteMessage(2, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.Extension {
		err = w.WriteMessage(6, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.NestedType {
		err = w.WriteMessage(3, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.EnumType {
		err = w.WriteMessage(4, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.ExtensionRange {
		err = w.WriteMessage(5, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.OneofDecl {
		err = w.WriteMessage(8, x)
		if err != nil {
			return nil, err
		}
	}
	err = w.WriteMessage(7, m.Options)
	if err != nil {
		return nil, err
	}
	for _, x := range m.ReservedRange {
		err = w.WriteMessage(9, x)
		if err != nil {
			return nil, err
		}
	}
	for _, x := range m.ReservedName {
		w.WriteString(10, x)
	}
	return w.Bytes(), nil
}
func (m DescriptorProto_ExtensionRange) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteInt32(1, m.Start)
	w.WriteInt32(2, m.End)
	return w.Bytes(), nil
}
func (m DescriptorProto_ReservedRange) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteInt32(1, m.Start)
	w.WriteInt32(2, m.End)
	return w.Bytes(), nil
}
func (m FieldDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	w.WriteInt32(3, m.Number)
	w.WriteEnum(4, int(m.Label))
	w.WriteEnum(5, int(m.Type))
	w.WriteString(6, m.TypeName)
	w.WriteString(2, m.Extendee)
	w.WriteString(7, m.DefaultValue)
	w.WriteInt32(9, m.OneofIndex)
	w.WriteString(10, m.JsonName)
	err = w.WriteMessage(8, m.Options)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
func (m OneofDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	err = w.WriteMessage(2, m.Options)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
func (m EnumDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	for _, x := range m.Value {
		err = w.WriteMessage(2, x)
		if err != nil {
			return nil, err
		}
	}
	err = w.WriteMessage(3, m.Options)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
func (m EnumValueDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	w.WriteInt32(2, m.Number)
	err = w.WriteMessage(3, m.Options)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
func (m ServiceDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	for _, x := range m.Method {
		err = w.WriteMessage(2, x)
		if err != nil {
			return nil, err
		}
	}
	err = w.WriteMessage(3, m.Options)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
func (m MethodDescriptorProto) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Name)
	w.WriteString(2, m.InputType)
	w.WriteString(3, m.OutputType)
	err = w.WriteMessage(4, m.Options)
	if err != nil {
		return nil, err
	}
	w.WriteBool(5, m.ClientStreaming)
	w.WriteBool(6, m.ServerStreaming)
	return w.Bytes(), nil
}
func (m FileOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.JavaPackage)
	w.WriteString(8, m.JavaOuterClassname)
	w.WriteBool(10, m.JavaMultipleFiles)
	w.WriteBool(20, m.JavaGenerateEqualsAndHash)
	w.WriteBool(27, m.JavaStringCheckUtf8)
	w.WriteEnum(9, int(m.OptimizeFor))
	w.WriteString(11, m.GoPackage)
	w.WriteBool(16, m.CcGenericServices)
	w.WriteBool(17, m.JavaGenericServices)
	w.WriteBool(18, m.PyGenericServices)
	w.WriteBool(23, m.Deprecated)
	w.WriteBool(31, m.CcEnableArenas)
	w.WriteString(36, m.ObjcClassPrefix)
	w.WriteString(37, m.CsharpNamespace)
	w.WriteString(39, m.SwiftPrefix)
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m MessageOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteBool(1, m.MessageSetWireFormat)
	w.WriteBool(2, m.NoStandardDescriptorAccessor)
	w.WriteBool(3, m.Deprecated)
	w.WriteBool(7, m.MapEntry)
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m FieldOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteEnum(1, int(m.Ctype))
	w.WriteBool(2, m.Packed)
	w.WriteEnum(6, int(m.Jstype))
	w.WriteBool(5, m.Lazy)
	w.WriteBool(3, m.Deprecated)
	w.WriteBool(10, m.Weak)
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m OneofOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m EnumOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteBool(2, m.AllowAlias)
	w.WriteBool(3, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m EnumValueOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteBool(1, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m ServiceOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteBool(33, m.Deprecated)
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m MethodOptions) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteBool(33, m.Deprecated)
	w.WriteEnum(34, int(m.IdempotencyLevel))
	for _, x := range m.UninterpretedOption {
		err = w.WriteMessage(999, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m UninterpretedOption) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.Name {
		err = w.WriteMessage(2, x)
		if err != nil {
			return nil, err
		}
	}
	w.WriteString(3, m.IdentifierValue)
	w.WriteUInt64(4, m.PositiveIntValue)
	w.WriteInt64(5, m.NegativeIntValue)
	w.WriteDouble(6, m.DoubleValue)
	w.WriteBytes(7, m.StringValue)
	w.WriteString(8, m.AggregateValue)
	return w.Bytes(), nil
}
func (m UninterpretedOption_NamePart) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.NamePart)
	w.WriteBool(2, m.IsExtension)
	return w.Bytes(), nil
}
func (m SourceCodeInfo) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.Location {
		err = w.WriteMessage(1, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m SourceCodeInfo_Location) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.Path {
		w.WriteInt32(1, x)
	}
	for _, x := range m.Span {
		w.WriteInt32(2, x)
	}
	w.WriteString(3, m.LeadingComments)
	w.WriteString(4, m.TrailingComments)
	for _, x := range m.LeadingDetachedComments {
		w.WriteString(6, x)
	}
	return w.Bytes(), nil
}
func (m GeneratedCodeInfo) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.Annotation {
		err = w.WriteMessage(1, x)
		if err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
func (m GeneratedCodeInfo_Annotation) Marshal() (data []byte, err error) {
	var w proto.Writer
	for _, x := range m.Path {
		w.WriteInt32(1, x)
	}
	w.WriteString(2, m.SourceFile)
	w.WriteInt32(3, m.Begin)
	w.WriteInt32(4, m.End)
	return w.Bytes(), nil
}
