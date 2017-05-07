package example
import "github.com/kixelated/protohack/proto"
var _ = proto.WIRETYPE_VARINT
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
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m Person) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Name)
	n += proto.WriteFieldInt32(data[n:], 2, m.Id)
	n += proto.WriteFieldString(data[n:], 3, m.Email)
	for _, x := range m.Phone {
		if x != nil {
			n += proto.WriteFieldMessage(data[n:], 4, x)
		}
	}
	return n
}
func (m Person) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Name)
	n += proto.SizeFieldInt32(2, m.Id)
	n += proto.SizeFieldString(3, m.Email)
	for _, x := range m.Phone {
		if x != nil {
			n += proto.SizeFieldMessage(4, x)
		}
	}
	return n
}
func (m Person_PhoneNumber) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_ = m.MarshalTo(data)
	return data, nil
}
func (m Person_PhoneNumber) MarshalTo(data []byte) (n int) {
	n += proto.WriteFieldString(data[n:], 1, m.Number)
	n += proto.WriteFieldEnum(data[n:], 2, int(m.Type))
	return n
}
func (m Person_PhoneNumber) MarshalSize() (n int) {
	n += proto.SizeFieldString(1, m.Number)
	n += proto.SizeFieldEnum(2, int(m.Type))
	return n
}
