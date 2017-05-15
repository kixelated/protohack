package example
import "github.com/kixelated/protohack/proto"
import "errors"
var _ = proto.WIRETYPE_VARINT
var _ = errors.New("")
type Person struct {
	Name string
	Id int32
	Email string
	Phone []*Person_PhoneNumber
}
type Person_PhoneType int
const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME Person_PhoneType = 1
	Person_WORK Person_PhoneType = 2
)
type Person_PhoneNumber struct {
	Number string
	Type Person_PhoneType
}
func (m Person) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m Person) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Name)
	w.Int32Field([]byte{0x10}, m.Id)
	w.StringField([]byte{0x1a}, m.Email)
	for _, x := range m.Phone{
		w.MessageField([]byte{0x22}, x)
	}
	return nil
}
func (m Person) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Name)
	s.Int32Field(1, m.Id)
	s.StringField(1, m.Email)
	for _, x := range m.Phone{
		s.MessageField(1, x)
	}
}
func (m Person) MarshalStackSize() (n int) {
	for _, x := range m.Phone {
		n += 1 + x.MarshalStackSize()
	}
	return n
}
func (m *Person) Unmarshal(data []byte) (err error) {
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
			m.Id, err = r.ReadInt32(t)
		case 3:
			m.Email, err = r.ReadString(t)
		case 4:
			x := new(Person_PhoneNumber)
			err = r.ReadMessage(t, x)
			if err == nil {
				m.Phone = append(m.Phone, x)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (m Person_PhoneNumber) Marshal() (data []byte, err error) {
	var s proto.Sizer
	s.Grow(m.MarshalStackSize())
	m.MarshalSize(&s)
	
	var w proto.Writer
	w.WithSizer(&s)
	err = m.MarshalTo(&w)
	
	return w.Data(), err
}
func (m Person_PhoneNumber) MarshalTo(w *proto.Writer) (err error) {
	w.StringField([]byte{0xa}, m.Number)
	w.IntField([]byte{0x10}, int(m.Type))
	return nil
}
func (m Person_PhoneNumber) MarshalSize(s *proto.Sizer) {
	s.StringField(1, m.Number)
	s.IntField(1, int(m.Type))
}
func (m Person_PhoneNumber) MarshalStackSize() (n int) {
	return n
}
func (m *Person_PhoneNumber) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, t, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Number, err = r.ReadString(t)
		case 2:
			x, err := r.ReadInt(t)
			if err != nil {
				return err
			}
			m.Type = Person_PhoneType(x)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
