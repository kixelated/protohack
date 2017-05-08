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
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m Person) MarshalTo(data []byte) (n int, err error) {
	if m.Name != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteString(data[n:], m.Name)
	} else {
		return 0, errors.New("missing required field: Name")
	}
	if m.Id != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteInt32(data[n:], m.Id)
	} else {
		return 0, errors.New("missing required field: Id")
	}
	if m.Email != "" {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteString(data[n:], m.Email)
	}
	for _, x := range m.Phone {
		if x != nil {
			n += copy(data[n:], []byte{0x22})
			n += proto.WriteMessage(data[n:], x)
		}
	}
	return n, nil
}
func (m Person) MarshalSize() (n int) {
	if m.Name != "" {
		n += 1 + proto.SizeString(m.Name)
	}
	if m.Id != 0 {
		n += 1 + proto.SizeInt32(m.Id)
	}
	if m.Email != "" {
		n += 1 + proto.SizeString(m.Email)
	}
	for _, x := range m.Phone {
		if x != nil {
			n += 1 + proto.SizeMessage(x)
		}
	}
	return n
}
func (m Person_PhoneNumber) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m Person_PhoneNumber) MarshalTo(data []byte) (n int, err error) {
	if m.Number != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteString(data[n:], m.Number)
	} else {
		return 0, errors.New("missing required field: Number")
	}
	if int(m.Type) != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteEnum(data[n:], int(m.Type))
	}
	return n, nil
}
func (m Person_PhoneNumber) MarshalSize() (n int) {
	if m.Number != "" {
		n += 1 + proto.SizeString(m.Number)
	}
	if int(m.Type) != 0 {
		n += 1 + proto.SizeEnum(int(m.Type))
	}
	return n
}
