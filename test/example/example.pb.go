package example
import "github.com/kixelated/protohack/proto"
var _ = proto.WIRETYPE_VARINT
type Test struct {
	Label string
	Type int32
	Reps []int64
	Optionalgroup *Test_OptionalGroup
	Number int32
	Name string
}
type Test_OptionalGroup struct {
	RequiredField string
}
type FOO int
const (
	X FOO = 17
)
func (m Test) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(1, m.Label)
	w.WriteInt32(2, m.Type)
	for _, x := range m.Reps {
		w.WriteInt64(3, x)
	}
	err = w.WriteGroup(4, m.Optionalgroup)
	if err != nil {
		return nil, err
	}
	w.WriteInt32(6, m.Number)
	w.WriteString(7, m.Name)
	return w.Bytes(), nil
}
func (m Test_OptionalGroup) Marshal() (data []byte, err error) {
	var w proto.Writer
	w.WriteString(5, m.RequiredField)
	return w.Bytes(), nil
}
