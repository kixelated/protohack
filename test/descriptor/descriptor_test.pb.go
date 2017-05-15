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
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m FileDescriptorSet) MarshalTo(w *proto.Writer) (err error) {
	for _, x := range m.File{
		w.MessageField([]byte{0xa}, x)
	}
	return nil
}
func (m FileDescriptorSet) MarshalSize(s *proto.Sizer) {
	for _, x := range m.File{
		s.MessageField(1, x)
	}
}
func (m FileDescriptorSet) MarshalStackSize() (n int) {
	for _, x := range m.File {
		n += 1 + x.MarshalStackSize()
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
			x := new(FileDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.File = append(m.File, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m FileDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m FileDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	w.StringField([]byte{0x12}, m.Package)
	w.StringRepeated([]byte{0x1a}, m.Dependency)
	for _, x := range m.MessageType{
		w.MessageField([]byte{0x22}, x)
	}
	for _, x := range m.EnumType{
		w.MessageField([]byte{0x2a}, x)
	}
	for _, x := range m.Service{
		w.MessageField([]byte{0x32}, x)
	}
	for _, x := range m.Extension{
		w.MessageField([]byte{0x3a}, x)
	}
	if m.Options != nil {
		w.MessageField([]byte{0x42}, m.Options)
	}
	if m.SourceCodeInfo != nil {
		w.MessageField([]byte{0x4a}, m.SourceCodeInfo)
	}
	w.Int32Repeated([]byte{0x50}, m.PublicDependency)
	w.Int32Repeated([]byte{0x58}, m.WeakDependency)
	w.StringField([]byte{0x62}, m.Syntax)
	return nil
}
func (m FileDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	s.StringField(1, m.Package)
	s.StringRepeated(1, m.Dependency)
	for _, x := range m.MessageType{
		s.MessageField(1, x)
	}
	for _, x := range m.EnumType{
		s.MessageField(1, x)
	}
	for _, x := range m.Service{
		s.MessageField(1, x)
	}
	for _, x := range m.Extension{
		s.MessageField(1, x)
	}
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
	if m.SourceCodeInfo != nil {
		s.MessageField(1, m.SourceCodeInfo)
	}
	s.Int32Repeated(1, m.PublicDependency)
	s.Int32Repeated(1, m.WeakDependency)
	s.StringField(1, m.Syntax)
}
func (m FileDescriptorProto) MarshalStackSize() (n int) {
	for _, x := range m.MessageType {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.EnumType {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.Service {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.Extension {
		n += 1 + x.MarshalStackSize()
	}
	n += 2
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
		case 2:
			m.Package, err = r.ReadString(t)
		case 3:
			m.Dependency, err = r.ReadStringRepeated(t, m.Dependency)
		case 4:
			x := new(DescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.MessageType = append(m.MessageType, x)
			}
		case 5:
			x := new(EnumDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.EnumType = append(m.EnumType, x)
			}
		case 6:
			x := new(ServiceDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Service = append(m.Service, x)
			}
		case 7:
			x := new(FieldDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Extension = append(m.Extension, x)
			}
		case 8:
			m.Options = new(FileOptions)
			err = r.ReadMessage(t, m.Options)
		case 9:
			m.SourceCodeInfo = new(SourceCodeInfo)
			err = r.ReadMessage(t, m.SourceCodeInfo)
		case 10:
			m.PublicDependency, err = r.ReadInt32Repeated(t, m.PublicDependency)
		case 11:
			m.WeakDependency, err = r.ReadInt32Repeated(t, m.WeakDependency)
		case 12:
			m.Syntax, err = r.ReadString(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m DescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m DescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	for _, x := range m.Field{
		w.MessageField([]byte{0x12}, x)
	}
	for _, x := range m.NestedType{
		w.MessageField([]byte{0x1a}, x)
	}
	for _, x := range m.EnumType{
		w.MessageField([]byte{0x22}, x)
	}
	for _, x := range m.ExtensionRange{
		w.MessageField([]byte{0x2a}, x)
	}
	for _, x := range m.Extension{
		w.MessageField([]byte{0x32}, x)
	}
	if m.Options != nil {
		w.MessageField([]byte{0x3a}, m.Options)
	}
	for _, x := range m.OneofDecl{
		w.MessageField([]byte{0x42}, x)
	}
	for _, x := range m.ReservedRange{
		w.MessageField([]byte{0x4a}, x)
	}
	w.StringRepeated([]byte{0x52}, m.ReservedName)
	return nil
}
func (m DescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	for _, x := range m.Field{
		s.MessageField(1, x)
	}
	for _, x := range m.NestedType{
		s.MessageField(1, x)
	}
	for _, x := range m.EnumType{
		s.MessageField(1, x)
	}
	for _, x := range m.ExtensionRange{
		s.MessageField(1, x)
	}
	for _, x := range m.Extension{
		s.MessageField(1, x)
	}
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
	for _, x := range m.OneofDecl{
		s.MessageField(1, x)
	}
	for _, x := range m.ReservedRange{
		s.MessageField(1, x)
	}
	s.StringRepeated(1, m.ReservedName)
}
func (m DescriptorProto) MarshalStackSize() (n int) {
	for _, x := range m.Field {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.NestedType {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.EnumType {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.ExtensionRange {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.Extension {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.OneofDecl {
		n += 1 + x.MarshalStackSize()
	}
	for _, x := range m.ReservedRange {
		n += 1 + x.MarshalStackSize()
	}
	n += 1
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
		case 2:
			x := new(FieldDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Field = append(m.Field, x)
			}
		case 3:
			x := new(DescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.NestedType = append(m.NestedType, x)
			}
		case 4:
			x := new(EnumDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.EnumType = append(m.EnumType, x)
			}
		case 5:
			x := new(DescriptorProto_ExtensionRange)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.ExtensionRange = append(m.ExtensionRange, x)
			}
		case 6:
			x := new(FieldDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Extension = append(m.Extension, x)
			}
		case 7:
			m.Options = new(MessageOptions)
			err = r.ReadMessage(t, m.Options)
		case 8:
			x := new(OneofDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.OneofDecl = append(m.OneofDecl, x)
			}
		case 9:
			x := new(DescriptorProto_ReservedRange)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.ReservedRange = append(m.ReservedRange, x)
			}
		case 10:
			m.ReservedName, err = r.ReadStringRepeated(t, m.ReservedName)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m DescriptorProto_ExtensionRange) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m DescriptorProto_ExtensionRange) MarshalTo(w *proto.Writer) (err error) {
	w.Int32Field([]byte{0x8}, m.Start)
	w.Int32Field([]byte{0x10}, m.End)
	return nil
}
func (m DescriptorProto_ExtensionRange) MarshalSize(s *proto.Sizer) {
	s.Int32Field(1, m.Start)
	s.Int32Field(1, m.End)
}
func (m DescriptorProto_ExtensionRange) MarshalStackSize() (n int) {
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
		case 2:
			m.End, err = r.ReadInt32(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m DescriptorProto_ReservedRange) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m DescriptorProto_ReservedRange) MarshalTo(w *proto.Writer) (err error) {
	w.Int32Field([]byte{0x8}, m.Start)
	w.Int32Field([]byte{0x10}, m.End)
	return nil
}
func (m DescriptorProto_ReservedRange) MarshalSize(s *proto.Sizer) {
	s.Int32Field(1, m.Start)
	s.Int32Field(1, m.End)
}
func (m DescriptorProto_ReservedRange) MarshalStackSize() (n int) {
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
		case 2:
			m.End, err = r.ReadInt32(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m FieldDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m FieldDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	w.StringField([]byte{0x12}, m.Extendee)
	w.Int32Field([]byte{0x18}, m.Number)
	w.IntField([]byte{0x20}, int(m.Label))
	w.IntField([]byte{0x28}, int(m.Type))
	w.StringField([]byte{0x32}, m.TypeName)
	w.StringField([]byte{0x3a}, m.DefaultValue)
	if m.Options != nil {
		w.MessageField([]byte{0x42}, m.Options)
	}
	w.Int32Field([]byte{0x48}, m.OneofIndex)
	w.StringField([]byte{0x52}, m.JsonName)
	return nil
}
func (m FieldDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	s.StringField(1, m.Extendee)
	s.Int32Field(1, m.Number)
	s.IntField(1, int(m.Label))
	s.IntField(1, int(m.Type))
	s.StringField(1, m.TypeName)
	s.StringField(1, m.DefaultValue)
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
	s.Int32Field(1, m.OneofIndex)
	s.StringField(1, m.JsonName)
}
func (m FieldDescriptorProto) MarshalStackSize() (n int) {
	n += 1
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
		case 2:
			m.Extendee, err = r.ReadString(t)
		case 3:
			m.Number, err = r.ReadInt32(t)
		case 4:
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.Label = FieldDescriptorProto_Label(x)
		case 5:
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.Type = FieldDescriptorProto_Type(x)
		case 6:
			m.TypeName, err = r.ReadString(t)
		case 7:
			m.DefaultValue, err = r.ReadString(t)
		case 8:
			m.Options = new(FieldOptions)
			err = r.ReadMessage(t, m.Options)
		case 9:
			m.OneofIndex, err = r.ReadInt32(t)
		case 10:
			m.JsonName, err = r.ReadString(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m OneofDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m OneofDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	if m.Options != nil {
		w.MessageField([]byte{0x12}, m.Options)
	}
	return nil
}
func (m OneofDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
}
func (m OneofDescriptorProto) MarshalStackSize() (n int) {
	n += 1
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
		case 2:
			m.Options = new(OneofOptions)
			err = r.ReadMessage(t, m.Options)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m EnumDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m EnumDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	for _, x := range m.Value{
		w.MessageField([]byte{0x12}, x)
	}
	if m.Options != nil {
		w.MessageField([]byte{0x1a}, m.Options)
	}
	return nil
}
func (m EnumDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	for _, x := range m.Value{
		s.MessageField(1, x)
	}
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
}
func (m EnumDescriptorProto) MarshalStackSize() (n int) {
	for _, x := range m.Value {
		n += 1 + x.MarshalStackSize()
	}
	n += 1
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
		case 2:
			x := new(EnumValueDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Value = append(m.Value, x)
			}
		case 3:
			m.Options = new(EnumOptions)
			err = r.ReadMessage(t, m.Options)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m EnumValueDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m EnumValueDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	w.Int32Field([]byte{0x10}, m.Number)
	if m.Options != nil {
		w.MessageField([]byte{0x1a}, m.Options)
	}
	return nil
}
func (m EnumValueDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	s.Int32Field(1, m.Number)
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
}
func (m EnumValueDescriptorProto) MarshalStackSize() (n int) {
	n += 1
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
		case 2:
			m.Number, err = r.ReadInt32(t)
		case 3:
			m.Options = new(EnumValueOptions)
			err = r.ReadMessage(t, m.Options)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m ServiceDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m ServiceDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	for _, x := range m.Method{
		w.MessageField([]byte{0x12}, x)
	}
	if m.Options != nil {
		w.MessageField([]byte{0x1a}, m.Options)
	}
	return nil
}
func (m ServiceDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	for _, x := range m.Method{
		s.MessageField(1, x)
	}
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
}
func (m ServiceDescriptorProto) MarshalStackSize() (n int) {
	for _, x := range m.Method {
		n += 1 + x.MarshalStackSize()
	}
	n += 1
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
		case 2:
			x := new(MethodDescriptorProto)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Method = append(m.Method, x)
			}
		case 3:
			m.Options = new(ServiceOptions)
			err = r.ReadMessage(t, m.Options)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m MethodDescriptorProto) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m MethodDescriptorProto) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	w.StringField([]byte{0x12}, m.InputType)
	w.StringField([]byte{0x1a}, m.OutputType)
	if m.Options != nil {
		w.MessageField([]byte{0x22}, m.Options)
	}
	w.BoolField([]byte{0x28}, m.ClientStreaming)
	w.BoolField([]byte{0x30}, m.ServerStreaming)
	return nil
}
func (m MethodDescriptorProto) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	s.StringField(1, m.InputType)
	s.StringField(1, m.OutputType)
	if m.Options != nil {
		s.MessageField(1, m.Options)
	}
	s.BoolField(1, m.ClientStreaming)
	s.BoolField(1, m.ServerStreaming)
}
func (m MethodDescriptorProto) MarshalStackSize() (n int) {
	n += 1
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
		case 2:
			m.InputType, err = r.ReadString(t)
		case 3:
			m.OutputType, err = r.ReadString(t)
		case 4:
			m.Options = new(MethodOptions)
			err = r.ReadMessage(t, m.Options)
		case 5:
			m.ClientStreaming, err = r.ReadBool(t)
		case 6:
			m.ServerStreaming, err = r.ReadBool(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m FileOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m FileOptions) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.JavaPackage)
	w.StringField([]byte{0x42}, m.JavaOuterClassname)
	w.IntField([]byte{0x48}, int(m.OptimizeFor))
	w.BoolField([]byte{0x50}, m.JavaMultipleFiles)
	w.StringField([]byte{0x5a}, m.GoPackage)
	w.BoolField([]byte{0x80, 0x1}, m.CcGenericServices)
	w.BoolField([]byte{0x88, 0x1}, m.JavaGenericServices)
	w.BoolField([]byte{0x90, 0x1}, m.PyGenericServices)
	w.BoolField([]byte{0xa0, 0x1}, m.JavaGenerateEqualsAndHash)
	w.BoolField([]byte{0xb8, 0x1}, m.Deprecated)
	w.BoolField([]byte{0xd8, 0x1}, m.JavaStringCheckUtf8)
	w.BoolField([]byte{0xf8, 0x1}, m.CcEnableArenas)
	w.StringField([]byte{0xa2, 0x2}, m.ObjcClassPrefix)
	w.StringField([]byte{0xaa, 0x2}, m.CsharpNamespace)
	w.StringField([]byte{0xba, 0x2}, m.SwiftPrefix)
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m FileOptions) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.JavaPackage)
	s.StringField(1, m.JavaOuterClassname)
	s.IntField(1, int(m.OptimizeFor))
	s.BoolField(1, m.JavaMultipleFiles)
	s.StringField(1, m.GoPackage)
	s.BoolField(2, m.CcGenericServices)
	s.BoolField(2, m.JavaGenericServices)
	s.BoolField(2, m.PyGenericServices)
	s.BoolField(2, m.JavaGenerateEqualsAndHash)
	s.BoolField(2, m.Deprecated)
	s.BoolField(2, m.JavaStringCheckUtf8)
	s.BoolField(2, m.CcEnableArenas)
	s.StringField(2, m.ObjcClassPrefix)
	s.StringField(2, m.CsharpNamespace)
	s.StringField(2, m.SwiftPrefix)
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m FileOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
		case 8:
			m.JavaOuterClassname, err = r.ReadString(t)
		case 9:
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.OptimizeFor = FileOptions_OptimizeMode(x)
		case 10:
			m.JavaMultipleFiles, err = r.ReadBool(t)
		case 11:
			m.GoPackage, err = r.ReadString(t)
		case 16:
			m.CcGenericServices, err = r.ReadBool(t)
		case 17:
			m.JavaGenericServices, err = r.ReadBool(t)
		case 18:
			m.PyGenericServices, err = r.ReadBool(t)
		case 20:
			m.JavaGenerateEqualsAndHash, err = r.ReadBool(t)
		case 23:
			m.Deprecated, err = r.ReadBool(t)
		case 27:
			m.JavaStringCheckUtf8, err = r.ReadBool(t)
		case 31:
			m.CcEnableArenas, err = r.ReadBool(t)
		case 36:
			m.ObjcClassPrefix, err = r.ReadString(t)
		case 37:
			m.CsharpNamespace, err = r.ReadString(t)
		case 39:
			m.SwiftPrefix, err = r.ReadString(t)
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m MessageOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m MessageOptions) MarshalTo(w *proto.Writer) (err error) {
	w.BoolField([]byte{0x8}, m.MessageSetWireFormat)
	w.BoolField([]byte{0x10}, m.NoStandardDescriptorAccessor)
	w.BoolField([]byte{0x18}, m.Deprecated)
	w.BoolField([]byte{0x38}, m.MapEntry)
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m MessageOptions) MarshalSize(s *proto.Sizer) {
	s.BoolField(1, m.MessageSetWireFormat)
	s.BoolField(1, m.NoStandardDescriptorAccessor)
	s.BoolField(1, m.Deprecated)
	s.BoolField(1, m.MapEntry)
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m MessageOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
		case 2:
			m.NoStandardDescriptorAccessor, err = r.ReadBool(t)
		case 3:
			m.Deprecated, err = r.ReadBool(t)
		case 7:
			m.MapEntry, err = r.ReadBool(t)
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m FieldOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m FieldOptions) MarshalTo(w *proto.Writer) (err error) {
	w.IntField([]byte{0x8}, int(m.Ctype))
	w.BoolField([]byte{0x10}, m.Packed)
	w.BoolField([]byte{0x18}, m.Deprecated)
	w.BoolField([]byte{0x28}, m.Lazy)
	w.IntField([]byte{0x30}, int(m.Jstype))
	w.BoolField([]byte{0x50}, m.Weak)
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m FieldOptions) MarshalSize(s *proto.Sizer) {
	s.IntField(1, int(m.Ctype))
	s.BoolField(1, m.Packed)
	s.BoolField(1, m.Deprecated)
	s.BoolField(1, m.Lazy)
	s.IntField(1, int(m.Jstype))
	s.BoolField(1, m.Weak)
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m FieldOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.Ctype = FieldOptions_CType(x)
		case 2:
			m.Packed, err = r.ReadBool(t)
		case 3:
			m.Deprecated, err = r.ReadBool(t)
		case 5:
			m.Lazy, err = r.ReadBool(t)
		case 6:
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.Jstype = FieldOptions_JSType(x)
		case 10:
			m.Weak, err = r.ReadBool(t)
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m OneofOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m OneofOptions) MarshalTo(w *proto.Writer) (err error) {
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m OneofOptions) MarshalSize(s *proto.Sizer) {
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m OneofOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m EnumOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m EnumOptions) MarshalTo(w *proto.Writer) (err error) {
	w.BoolField([]byte{0x10}, m.AllowAlias)
	w.BoolField([]byte{0x18}, m.Deprecated)
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m EnumOptions) MarshalSize(s *proto.Sizer) {
	s.BoolField(1, m.AllowAlias)
	s.BoolField(1, m.Deprecated)
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m EnumOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
		case 3:
			m.Deprecated, err = r.ReadBool(t)
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m EnumValueOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m EnumValueOptions) MarshalTo(w *proto.Writer) (err error) {
	w.BoolField([]byte{0x8}, m.Deprecated)
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m EnumValueOptions) MarshalSize(s *proto.Sizer) {
	s.BoolField(1, m.Deprecated)
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m EnumValueOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m ServiceOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m ServiceOptions) MarshalTo(w *proto.Writer) (err error) {
	w.BoolField([]byte{0x88, 0x2}, m.Deprecated)
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m ServiceOptions) MarshalSize(s *proto.Sizer) {
	s.BoolField(2, m.Deprecated)
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m ServiceOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m MethodOptions) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m MethodOptions) MarshalTo(w *proto.Writer) (err error) {
	w.BoolField([]byte{0x88, 0x2}, m.Deprecated)
	w.IntField([]byte{0x90, 0x2}, int(m.IdempotencyLevel))
	for _, x := range m.UninterpretedOption{
		w.MessageField([]byte{0xba, 0x3e}, x)
	}
	return nil
}
func (m MethodOptions) MarshalSize(s *proto.Sizer) {
	s.BoolField(2, m.Deprecated)
	s.IntField(2, int(m.IdempotencyLevel))
	for _, x := range m.UninterpretedOption{
		s.MessageField(2, x)
	}
}
func (m MethodOptions) MarshalStackSize() (n int) {
	for _, x := range m.UninterpretedOption {
		n += 1 + x.MarshalStackSize()
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
		case 34:
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.IdempotencyLevel = MethodOptions_IdempotencyLevel(x)
		case 999:
			x := new(UninterpretedOption)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.UninterpretedOption = append(m.UninterpretedOption, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m UninterpretedOption) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m UninterpretedOption) MarshalTo(w *proto.Writer) (err error) {
	for _, x := range m.Name{
		w.MessageField([]byte{0x12}, x)
	}
	w.StringField([]byte{0x1a}, m.IdentifierValue)
	w.UInt64Field([]byte{0x20}, m.PositiveIntValue)
	w.Int64Field([]byte{0x28}, m.NegativeIntValue)
	w.Float64Field([]byte{0x31}, m.DoubleValue)
	w.BytesField([]byte{0x3a}, m.StringValue)
	w.StringField([]byte{0x42}, m.AggregateValue)
	return nil
}
func (m UninterpretedOption) MarshalSize(s *proto.Sizer) {
	for _, x := range m.Name{
		s.MessageField(1, x)
	}
	s.StringField(1, m.IdentifierValue)
	s.UInt64Field(1, m.PositiveIntValue)
	s.Int64Field(1, m.NegativeIntValue)
	s.Float64Field(1, m.DoubleValue)
	s.BytesField(1, m.StringValue)
	s.StringField(1, m.AggregateValue)
}
func (m UninterpretedOption) MarshalStackSize() (n int) {
	for _, x := range m.Name {
		n += 1 + x.MarshalStackSize()
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
			x := new(UninterpretedOption_NamePart)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Name = append(m.Name, x)
			}
		case 3:
			m.IdentifierValue, err = r.ReadString(t)
		case 4:
			m.PositiveIntValue, err = r.ReadUInt64(t)
		case 5:
			m.NegativeIntValue, err = r.ReadInt64(t)
		case 6:
			m.DoubleValue, err = r.ReadFloat64(t)
		case 7:
			m.StringValue, err = r.ReadBytes(t)
		case 8:
			m.AggregateValue, err = r.ReadString(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m UninterpretedOption_NamePart) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m UninterpretedOption_NamePart) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.NamePart)
	w.BoolField([]byte{0x10}, m.IsExtension)
	return nil
}
func (m UninterpretedOption_NamePart) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.NamePart)
	s.BoolField(1, m.IsExtension)
}
func (m UninterpretedOption_NamePart) MarshalStackSize() (n int) {
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
		case 2:
			m.IsExtension, err = r.ReadBool(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m SourceCodeInfo) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m SourceCodeInfo) MarshalTo(w *proto.Writer) (err error) {
	for _, x := range m.Location{
		w.MessageField([]byte{0xa}, x)
	}
	return nil
}
func (m SourceCodeInfo) MarshalSize(s *proto.Sizer) {
	for _, x := range m.Location{
		s.MessageField(1, x)
	}
}
func (m SourceCodeInfo) MarshalStackSize() (n int) {
	for _, x := range m.Location {
		n += 1 + x.MarshalStackSize()
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
			x := new(SourceCodeInfo_Location)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Location = append(m.Location, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m SourceCodeInfo_Location) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m SourceCodeInfo_Location) MarshalTo(w *proto.Writer) (err error) {
	w.Int32Packed([]byte{0xa}, m.Path)
	w.Int32Packed([]byte{0x12}, m.Span)
	w.StringField([]byte{0x1a}, m.LeadingComments)
	w.StringField([]byte{0x22}, m.TrailingComments)
	w.StringRepeated([]byte{0x32}, m.LeadingDetachedComments)
	return nil
}
func (m SourceCodeInfo_Location) MarshalSize(s *proto.Sizer) {
	s.Int32Packed(1, m.Path)
	s.Int32Packed(1, m.Span)
	s.StringField(1, m.LeadingComments)
	s.StringField(1, m.TrailingComments)
	s.StringRepeated(1, m.LeadingDetachedComments)
}
func (m SourceCodeInfo_Location) MarshalStackSize() (n int) {
	n += 2
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
			m.Path, err = r.ReadInt32Repeated(t, m.Path)
		case 2:
			m.Span, err = r.ReadInt32Repeated(t, m.Span)
		case 3:
			m.LeadingComments, err = r.ReadString(t)
		case 4:
			m.TrailingComments, err = r.ReadString(t)
		case 6:
			m.LeadingDetachedComments, err = r.ReadStringRepeated(t, m.LeadingDetachedComments)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m GeneratedCodeInfo) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m GeneratedCodeInfo) MarshalTo(w *proto.Writer) (err error) {
	for _, x := range m.Annotation{
		w.MessageField([]byte{0xa}, x)
	}
	return nil
}
func (m GeneratedCodeInfo) MarshalSize(s *proto.Sizer) {
	for _, x := range m.Annotation{
		s.MessageField(1, x)
	}
}
func (m GeneratedCodeInfo) MarshalStackSize() (n int) {
	for _, x := range m.Annotation {
		n += 1 + x.MarshalStackSize()
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
			x := new(GeneratedCodeInfo_Annotation)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Annotation = append(m.Annotation, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m GeneratedCodeInfo_Annotation) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m GeneratedCodeInfo_Annotation) MarshalTo(w *proto.Writer) (err error) {
	w.Int32Packed([]byte{0xa}, m.Path)
	w.StringField([]byte{0x12}, m.SourceFile)
	w.Int32Field([]byte{0x18}, m.Begin)
	w.Int32Field([]byte{0x20}, m.End)
	return nil
}
func (m GeneratedCodeInfo_Annotation) MarshalSize(s *proto.Sizer) {
	s.Int32Packed(1, m.Path)
	s.StringField(1, m.SourceFile)
	s.Int32Field(1, m.Begin)
	s.Int32Field(1, m.End)
}
func (m GeneratedCodeInfo_Annotation) MarshalStackSize() (n int) {
	n += 1
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
			m.Path, err = r.ReadInt32Repeated(t, m.Path)
		case 2:
			m.SourceFile, err = r.ReadString(t)
		case 3:
			m.Begin, err = r.ReadInt32(t)
		case 4:
			m.End, err = r.ReadInt32(t)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
