package proto

import (
	"math"

	"github.com/pkg/errors"
)

// Neat trick taken from strconv
const intSize = 32 << (^uint(0) >> 63)

type Reader struct {
	data []byte
}

func NewReader(data []byte) (r *Reader) {
	return &Reader{
		data: data,
	}
}

func (r *Reader) Len() (n int) {
	return len(r.data)
}

func (r *Reader) ReadKey() (id int, t WireType, err error) {
	temp, err := r.readInt()
	if err != nil {
		return 0, 0, err
	}

	id = temp >> 3
	t = WireType(temp & 0x7)

	return id, t, nil
}

func (r *Reader) ReadInt(t WireType) (x int, err error) {
	if t != WIRETYPE_VARINT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readInt()
}

func (r *Reader) readInt() (x int, err error) {
	if len(r.data) < 1 {
		return 0, errors.New("unexpected EOF")
	}

	x = int(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return x, nil
	}

	if len(r.data) < 2 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return x, nil
	}

	if len(r.data) < 3 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return x, nil
	}

	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return x, nil
	}

	if len(r.data) < 5 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return x, nil
	}

	if intSize == 32 {
		return 0, errors.New("integer overflow")
	}

	if len(r.data) < 6 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return x, nil
	}

	if len(r.data) < 7 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return x, nil
	}

	if len(r.data) < 8 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return x, nil
	}

	if len(r.data) < 9 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return x, nil
	}

	if len(r.data) < 10 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return x, nil
	}

	return 0, errors.New("integer overflow")
}

func (r *Reader) ReadUInt(t WireType) (x uint, err error) {
	if t != WIRETYPE_VARINT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readUInt()
}

func (r *Reader) readUInt() (x uint, err error) {
	if len(r.data) < 1 {
		return 0, errors.New("unexpected EOF")
	}

	x = uint(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return x, nil
	}

	if len(r.data) < 2 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return x, nil
	}

	if len(r.data) < 3 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return x, nil
	}

	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return x, nil
	}

	if len(r.data) < 5 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return x, nil
	}

	if intSize == 32 {
		return 0, errors.New("integer overflow")
	}

	if len(r.data) < 6 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return x, nil
	}

	if len(r.data) < 7 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return x, nil
	}

	if len(r.data) < 8 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return x, nil
	}

	if len(r.data) < 9 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return x, nil
	}

	if len(r.data) < 10 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return x, nil
	}

	return 0, errors.New("integer overflow")
}

func (r *Reader) ReadFloat64(t WireType) (x float64, err error) {
	if t != WIRETYPE_64BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readFloat64()
}

func (r *Reader) readFloat64() (x float64, err error) {
	temp, err := r.readFixed64()
	if err != nil {
		return 0, err
	}

	x = math.Float64frombits(temp)
	return x, nil
}

func (r *Reader) ReadFloat32(t WireType) (x float32, err error) {
	if t != WIRETYPE_32BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readFloat32()
}

func (r *Reader) readFloat32() (x float32, err error) {
	temp, err := r.readFixed32()
	if err != nil {
		return 0, err
	}

	x = math.Float32frombits(temp)
	return x, nil
}

func (r *Reader) ReadInt64(t WireType) (x int64, err error) {
	if t != WIRETYPE_VARINT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readInt64()
}

func (r *Reader) readInt64() (x int64, err error) {
	if len(r.data) < 1 {
		return 0, errors.New("unexpected EOF")
	}

	x = int64(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return x, nil
	}

	if len(r.data) < 2 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return x, nil
	}

	if len(r.data) < 3 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return x, nil
	}

	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return x, nil
	}

	if len(r.data) < 5 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return x, nil
	}

	if len(r.data) < 6 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return x, nil
	}

	if len(r.data) < 7 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return x, nil
	}

	if len(r.data) < 8 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return x, nil
	}

	if len(r.data) < 9 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return x, nil
	}

	if len(r.data) < 10 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int64(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return x, nil
	}

	return 0, errors.New("integer overflow")
}

func (r *Reader) ReadUInt64(t WireType) (x uint64, err error) {
	if t != WIRETYPE_VARINT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readUInt64()
}

func (r *Reader) readUInt64() (x uint64, err error) {
	if len(r.data) < 1 {
		return 0, errors.New("unexpected EOF")
	}

	x = uint64(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return x, nil
	}

	if len(r.data) < 2 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return x, nil
	}

	if len(r.data) < 3 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return x, nil
	}

	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return x, nil
	}

	if len(r.data) < 5 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return x, nil
	}

	if len(r.data) < 6 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return x, nil
	}

	if len(r.data) < 7 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return x, nil
	}

	if len(r.data) < 8 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return x, nil
	}

	if len(r.data) < 9 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return x, nil
	}

	if len(r.data) < 10 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint64(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return x, nil
	}

	return 0, errors.New("integer overflow")
}

func (r *Reader) ReadInt32(t WireType) (x int32, err error) {
	if t != WIRETYPE_VARINT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readInt32()
}

func (r *Reader) readInt32() (x int32, err error) {
	if len(r.data) < 1 {
		return 0, errors.New("unexpected EOF")
	}

	x = int32(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return x, nil
	}

	if len(r.data) < 2 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int32(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return x, nil
	}

	if len(r.data) < 3 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int32(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return x, nil
	}

	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int32(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return x, nil
	}

	if len(r.data) < 5 {
		return 0, errors.New("unexpected EOF")
	}

	x |= int32(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return x, nil
	}

	r.data = r.data[5:]

	return 0, errors.New("integer overflow")
}

func (r *Reader) ReadUInt32(t WireType) (x uint32, err error) {
	if t != WIRETYPE_VARINT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readUInt32()
}

func (r *Reader) readUInt32() (x uint32, err error) {
	if len(r.data) < 1 {
		return 0, errors.New("unexpected EOF")
	}

	x = uint32(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return x, nil
	}

	if len(r.data) < 2 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint32(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return x, nil
	}

	if len(r.data) < 3 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint32(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return x, nil
	}

	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint32(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return x, nil
	}

	if len(r.data) < 5 {
		return 0, errors.New("unexpected EOF")
	}

	x |= uint32(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return x, nil
	}

	r.data = r.data[5:]

	return 0, errors.New("integer overflow")
}

// TODO optimize
func (r *Reader) ReadEnum(t WireType, x *Enum) (err error) {
	temp, err := r.readInt()
	if err != nil {
		return err
	}

	*x = Enum(temp)
	return nil
}

func (r *Reader) ReadFixed64(t WireType) (x uint64, err error) {
	if t != WIRETYPE_64BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readFixed64()
}

func (r *Reader) readFixed64() (x uint64, err error) {
	if len(r.data) < 8 {
		return 0, errors.New("unexpected EOF")
	}

	x = uint64(r.data[0])
	x |= uint64(r.data[1] << 8)
	x |= uint64(r.data[2] << 16)
	x |= uint64(r.data[3] << 24)
	x |= uint64(r.data[4] << 32)
	x |= uint64(r.data[5] << 40)
	x |= uint64(r.data[6] << 48)
	x |= uint64(r.data[7] << 56)

	r.data = r.data[8:]

	return x, nil
}

func (r *Reader) ReadFixed32(t WireType) (x uint32, err error) {
	if t != WIRETYPE_32BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readFixed32()
}

func (r *Reader) readFixed32() (x uint32, err error) {
	if len(r.data) < 4 {
		return 0, errors.New("unexpected EOF")
	}

	x = uint32(r.data[0])
	x |= uint32(r.data[1] << 8)
	x |= uint32(r.data[2] << 16)
	x |= uint32(r.data[3] << 24)

	r.data = r.data[4:]

	return x, nil
}

func (r *Reader) ReadBool(t WireType) (x bool, err error) {
	if t != WIRETYPE_VARINT {
		return false, errors.New("unexpected wire type")
	}

	return r.readBool()
}

func (r *Reader) readBool() (x bool, err error) {
	if len(r.data) < 1 {
		return false, errors.New("unexpected EOF")
	}

	if r.data[0] > 1 {
		return false, errors.New("unexpected boolean value")
	}

	x = (r.data[0] == 1)

	return x, nil
}

func (r *Reader) ReadMessage(t WireType, x Unmarshaller) (err error) {
	if t != WIRETYPE_LENGTH {
		return errors.New("unexpected wire type")
	}

	size, err := r.readInt()
	if err != nil {
		return err
	}

	if len(r.data) < size {
		return errors.New("unexpected EOF")
	}

	err = x.Unmarshal(r.data[:size])
	r.data = r.data[size:]

	if err != nil {
		return err
	}

	return nil
}

func (r *Reader) ReadString(t WireType) (x string, err error) {
	if t != WIRETYPE_LENGTH {
		return "", errors.New("unexpected wire type")
	}

	return r.readString()
}

func (r *Reader) readString() (x string, err error) {
	temp, err := r.readBytes()
	if err != nil {
		return "", err
	}

	x = string(temp)

	return x, nil
}

func (r *Reader) ReadBytes(t WireType) (x []byte, err error) {
	if t != WIRETYPE_LENGTH {
		return nil, errors.New("unexpected wire type")
	}

	return r.readBytes()
}

func (r *Reader) readBytes() (x []byte, err error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	if len(r.data) < size {
		return nil, errors.New("unexpected EOF")
	}

	x = r.data[:size]
	r.data = r.data[size:]

	return x, nil
}

func (r *Reader) ReadSFixed32(t WireType) (x int32, err error) {
	if t != WIRETYPE_32BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readSFixed32()
}

func (r *Reader) readSFixed32() (x int32, err error) {
	// TODO Optimize
	temp, err := r.readFixed32()
	x = int32(temp)
	return 0, err
}

func (r *Reader) ReadSFixed64(t WireType) (x int64, err error) {
	if t != WIRETYPE_64BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readSFixed64()
}

func (r *Reader) readSFixed64() (x int64, err error) {
	// TODO Optimize
	temp, err := r.readFixed64()
	x = int64(temp)
	return 0, err
}

func (r *Reader) ReadSInt32(t WireType) (x int32, err error) {
	if t != WIRETYPE_32BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readSInt32()
}

func (r *Reader) readSInt32() (x int32, err error) {
	x, err = r.readInt32()
	if err != nil {
		return 0, err
	}

	// TODO This seems wrong?
	x = (x >> 1) ^ (((x & 0x1) << 31) >> 31)
	return x, nil
}

func (r *Reader) ReadSInt64(t WireType) (x int64, err error) {
	if t != WIRETYPE_64BIT {
		return 0, errors.New("unexpected wire type")
	}

	return r.readSInt64()
}

func (r *Reader) readSInt64() (x int64, err error) {
	x, err = r.readInt64()
	if err != nil {
		return 0, err
	}

	// TODO This seems wrong?
	x = (x >> 1) ^ (((x & 0x1) << 63) >> 63)
	return x, nil
}
