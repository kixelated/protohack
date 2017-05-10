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
		n += proto.WriteStringLength(data[n:], m.Name)
	} else {
		return n, errors.New("missing required field: Name")
	}
	if m.Id != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteInt32(data[n:], m.Id)
	} else {
		return n, errors.New("missing required field: Id")
	}
	if m.Email != "" {
		n += copy(data[n:], []byte{0x1a})
		n += proto.WriteStringLength(data[n:], m.Email)
	}
	for _, x := range m.Phone {
		n += copy(data[n:], []byte{0x22})
		n += proto.WriteMessageLength(data[n:], x)
	}
	return n, nil
}
func (m Person) MarshalSize() (n int) {
	n += 1 + proto.SizeStringLength(m.Name)
	n += 1 + proto.SizeInt32(m.Id)
	if m.Email != "" {
		n += 1 + proto.SizeStringLength(m.Email)
	}
	for _, x := range m.Phone {
		n += 1 + proto.SizeMessageLength(x)
	}
	return n
}
func (m *Person) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, _, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Name, err = r.ReadString()
			if err != nil {
				return err
			}
		case 2:
			m.Id, err = r.ReadInt32()
			if err != nil {
				return err
			}
		case 3:
			m.Email, err = r.ReadString()
			if err != nil {
				return err
			}
		case 4:
			temp := new(Person_PhoneNumber)
			err = r.ReadToMessage(temp)
			if err != nil {
				return err
			}
			m.Phone = append(m.Phone, temp)
		}
	}
	return nil
}
func (m Person_PhoneNumber) Marshal() (data []byte, err error) {
	data = make([]byte, m.MarshalSize())
	_, err = m.MarshalTo(data)
	return data, err
}
func (m Person_PhoneNumber) MarshalTo(data []byte) (n int, err error) {
	if m.Number != "" {
		n += copy(data[n:], []byte{0xa})
		n += proto.WriteStringLength(data[n:], m.Number)
	} else {
		return n, errors.New("missing required field: Number")
	}
	if int(m.Type) != 0 {
		n += copy(data[n:], []byte{0x10})
		n += proto.WriteEnum(data[n:], int(m.Type))
	}
	return n, nil
}
func (m Person_PhoneNumber) MarshalSize() (n int) {
	n += 1 + proto.SizeStringLength(m.Number)
	if int(m.Type) != 0 {
		n += 1 + proto.SizeEnum(int(m.Type))
	}
	return n
}
func (m *Person_PhoneNumber) Unmarshal(data []byte) (err error) {
	r := proto.NewReader(data)
	for r.Len() > 0 {
		id, _, err := r.ReadKey()
		if err != nil {
			return err
		}
		switch id {
		case 1:
			m.Number, err = r.ReadString()
			if err != nil {
				return err
			}
		case 2:
			temp, err := r.ReadEnum()
			if err != nil {
				return err
			}
			m.Type = Person_PhoneType(temp)
		}
	}
	return nil
}
