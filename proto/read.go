package proto

import (
	"errors"
	"math"
)

var ErrIntOverflow = errors.New("integer overflow")
var ErrUnexpectedEOF = errors.New("unexpected EOF")
var ErrWireType = errors.New("wrong wire type")

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
	err = r.ReadToKey(&id, &t)
	return id, t, err
}

func (r *Reader) ReadToKey(id *int, t *WireType) (err error) {
	temp, err := r.ReadInt(WIRETYPE_VARINT)
	if err != nil {
		return err
	}

	*id = temp >> 3
	*t = WireType(temp & 0x7)

	return nil
}

func (r *Reader) ReadInt(t WireType) (x int, err error) {
	err = r.ReadToInt(t, &x)
	return x, err
}

func (r *Reader) ReadToInt(t WireType, x *int) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	*x = int(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return nil
	}

	if len(r.data) < 2 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return nil
	}

	if len(r.data) < 3 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return nil
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return nil
	}

	if len(r.data) < 5 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return nil
	}

	if intSize == 32 {
		return ErrIntOverflow
	}

	if len(r.data) < 6 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return nil
	}

	if len(r.data) < 7 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return nil
	}

	if len(r.data) < 8 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return nil
	}

	if len(r.data) < 9 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return nil
	}

	if len(r.data) < 10 {
		return ErrUnexpectedEOF
	}

	*x |= int(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return nil
	}

	return ErrIntOverflow
}

func (r *Reader) ReadUInt(t WireType) (x uint, err error) {
	err = r.ReadToUInt(t, &x)
	return x, err
}

func (r *Reader) ReadToUInt(t WireType, x *uint) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	*x = uint(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return nil
	}

	if len(r.data) < 2 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return nil
	}

	if len(r.data) < 3 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return nil
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return nil
	}

	if len(r.data) < 5 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return nil
	}

	if intSize == 32 {
		return ErrIntOverflow
	}

	if len(r.data) < 6 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return nil
	}

	if len(r.data) < 7 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return nil
	}

	if len(r.data) < 8 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return nil
	}

	if len(r.data) < 9 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return nil
	}

	if len(r.data) < 10 {
		return ErrUnexpectedEOF
	}

	*x |= uint(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return nil
	}

	return ErrIntOverflow
}

func (r *Reader) ReadDouble(t WireType) (x float64, err error) {
	err = r.ReadToDouble(t, &x)
	return x, err
}

func (r *Reader) ReadToDouble(t WireType, x *float64) (err error) {
	temp, err := r.ReadFixed64(t)
	if err != nil {
		return err
	}

	*x = math.Float64frombits(temp)

	return nil
}

func (r *Reader) ReadFloat(t WireType) (x float32, err error) {
	err = r.ReadToFloat(t, &x)
	return x, err
}

func (r *Reader) ReadToFloat(t WireType, x *float32) (err error) {
	temp, err := r.ReadFixed32(t)
	if err != nil {
		return err
	}

	*x = math.Float32frombits(temp)

	return nil
}

func (r *Reader) ReadInt64(t WireType) (x int64, err error) {
	err = r.ReadToInt64(t, &x)
	return x, err
}

func (r *Reader) ReadToInt64(t WireType, x *int64) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	*x = int64(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return nil
	}

	if len(r.data) < 2 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return nil
	}

	if len(r.data) < 3 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return nil
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return nil
	}

	if len(r.data) < 5 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return nil
	}

	if len(r.data) < 6 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return nil
	}

	if len(r.data) < 7 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return nil
	}

	if len(r.data) < 8 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return nil
	}

	if len(r.data) < 9 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return nil
	}

	if len(r.data) < 10 {
		return ErrUnexpectedEOF
	}

	*x |= int64(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return nil
	}

	return ErrIntOverflow
}

func (r *Reader) ReadUInt64(t WireType) (x uint64, err error) {
	err = r.ReadToUInt64(t, &x)
	return x, err
}

func (r *Reader) ReadToUInt64(t WireType, x *uint64) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	*x = uint64(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return nil
	}

	if len(r.data) < 2 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return nil
	}

	if len(r.data) < 3 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return nil
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return nil
	}

	if len(r.data) < 5 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return nil
	}

	if len(r.data) < 6 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[5]&0x7F) << 35
	if (r.data[5] & 0x80) == 0 {
		r.data = r.data[6:]
		return nil
	}

	if len(r.data) < 7 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[6]&0x7F) << 42
	if (r.data[6] & 0x80) == 0 {
		r.data = r.data[7:]
		return nil
	}

	if len(r.data) < 8 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[7]&0x7F) << 49
	if (r.data[7] & 0x80) == 0 {
		r.data = r.data[8:]
		return nil
	}

	if len(r.data) < 9 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[8]&0x7F) << 56
	if (r.data[8] & 0x80) == 0 {
		r.data = r.data[9:]
		return nil
	}

	if len(r.data) < 10 {
		return ErrUnexpectedEOF
	}

	*x |= uint64(r.data[9]&0x7F) << 63
	if (r.data[9] & 0x80) == 0 {
		r.data = r.data[10:]
		return nil
	}

	return ErrIntOverflow
}

func (r *Reader) ReadInt32(t WireType) (x int32, err error) {
	err = r.ReadToInt32(t, &x)
	return x, err
}

func (r *Reader) ReadToInt32(t WireType, x *int32) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	*x = int32(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return nil
	}

	if len(r.data) < 2 {
		return ErrUnexpectedEOF
	}

	*x |= int32(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return nil
	}

	if len(r.data) < 3 {
		return ErrUnexpectedEOF
	}

	*x |= int32(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return nil
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x |= int32(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return nil
	}

	if len(r.data) < 5 {
		return ErrUnexpectedEOF
	}

	*x |= int32(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return nil
	}

	r.data = r.data[5:]

	return ErrIntOverflow
}

func (r *Reader) ReadUInt32(t WireType) (x uint32, err error) {
	err = r.ReadToUInt32(t, &x)
	return x, err
}

func (r *Reader) ReadToUInt32(t WireType, x *uint32) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	*x = uint32(r.data[0] & 0x7F)
	if (r.data[0] & 0x80) == 0 {
		r.data = r.data[1:]
		return nil
	}

	if len(r.data) < 2 {
		return ErrUnexpectedEOF
	}

	*x |= uint32(r.data[1]&0x7F) << 7
	if (r.data[1] & 0x80) == 0 {
		r.data = r.data[2:]
		return nil
	}

	if len(r.data) < 3 {
		return ErrUnexpectedEOF
	}

	*x |= uint32(r.data[2]&0x7F) << 14
	if (r.data[2] & 0x80) == 0 {
		r.data = r.data[3:]
		return nil
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x |= uint32(r.data[3]&0x7F) << 21
	if (r.data[3] & 0x80) == 0 {
		r.data = r.data[4:]
		return nil
	}

	if len(r.data) < 5 {
		return ErrUnexpectedEOF
	}

	*x |= uint32(r.data[4]&0x7F) << 28
	if (r.data[4] & 0x80) == 0 {
		r.data = r.data[5:]
		return nil
	}

	r.data = r.data[5:]

	return ErrIntOverflow
}

func (r *Reader) ReadFixed64(t WireType) (x uint64, err error) {
	err = r.ReadToFixed64(t, &x)
	return x, err
}

func (r *Reader) ReadToFixed64(t WireType, x *uint64) (err error) {
	if t != WIRETYPE_64BIT {
		return ErrWireType
	}

	if len(r.data) < 8 {
		return ErrUnexpectedEOF
	}

	*x = uint64(r.data[0])
	*x |= uint64(r.data[1] << 8)
	*x |= uint64(r.data[2] << 16)
	*x |= uint64(r.data[3] << 24)
	*x |= uint64(r.data[4] << 32)
	*x |= uint64(r.data[5] << 40)
	*x |= uint64(r.data[6] << 48)
	*x |= uint64(r.data[7] << 56)

	r.data = r.data[8:]

	return nil
}

func (r *Reader) ReadFixed32(t WireType) (x uint32, err error) {
	err = r.ReadToFixed32(t, &x)
	return x, err
}

func (r *Reader) ReadToFixed32(t WireType, x *uint32) (err error) {
	if t != WIRETYPE_32BIT {
		return ErrWireType
	}

	if len(r.data) < 4 {
		return ErrUnexpectedEOF
	}

	*x = uint32(r.data[0])
	*x |= uint32(r.data[1] << 8)
	*x |= uint32(r.data[2] << 16)
	*x |= uint32(r.data[3] << 24)

	r.data = r.data[4:]

	return nil
}

func (r *Reader) ReadBool(t WireType) (x bool, err error) {
	err = r.ReadToBool(t, &x)
	return x, err
}

func (r *Reader) ReadToBool(t WireType, x *bool) (err error) {
	if t != WIRETYPE_VARINT {
		return ErrWireType
	}

	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	if r.data[0] > 1 {
		return errors.New("unexpected boolean value")
	}

	*x = (r.data[0] == 1)

	return nil
}

func (r *Reader) ReadToMessage(t WireType, x Unmarshaller) (err error) {
	if t != WIRETYPE_LENGTH {
		return ErrWireType
	}

	size, err := r.ReadInt(WIRETYPE_VARINT)
	if err != nil {
		return err
	}

	if len(r.data) < size {
		return ErrUnexpectedEOF
	}

	temp := r.data[:size]
	r.data = r.data[size:]

	err = x.Unmarshal(temp)
	if err != nil {
		return err
	}

	return nil
}

func (r *Reader) ReadString(t WireType) (x string, err error) {
	err = r.ReadToString(t, &x)
	return x, err
}

func (r *Reader) ReadToString(t WireType, x *string) (err error) {
	temp, err := r.ReadBytes(t)
	if err != nil {
		return err
	}

	*x = string(temp)

	return nil
}

func (r *Reader) ReadBytes(t WireType) (x []byte, err error) {
	err = r.ReadToBytes(t, &x)
	return x, err
}

func (r *Reader) ReadToBytes(t WireType, x *[]byte) (err error) {
	if t != WIRETYPE_LENGTH {
		return ErrWireType
	}

	size, err := r.ReadInt(WIRETYPE_VARINT)
	if err != nil {
		return err
	}

	if len(r.data) < size {
		return ErrUnexpectedEOF
	}

	*x = r.data[:size]
	r.data = r.data[size:]

	return nil
}

func (r *Reader) ReadEnum(t WireType) (x Enum, err error) {
	err = r.ReadToEnum(t, &x)
	return x, err
}

func (r *Reader) ReadToEnum(t WireType, x *Enum) (err error) {
	return r.ReadToInt(t, (*int)(x))
}

func (r *Reader) ReadSFixed32(t WireType) (x int32, err error) {
	err = r.ReadToSFixed32(t, &x)
	return x, err
}

func (r *Reader) ReadToSFixed32(t WireType, x *int32) (err error) {
	// TODO Optimize
	temp, err := r.ReadFixed32(t)
	*x = int32(temp)
	return err
}

func (r *Reader) ReadSFixed64(t WireType) (x int64, err error) {
	err = r.ReadToSFixed64(t, &x)
	return x, err
}

func (r *Reader) ReadToSFixed64(t WireType, x *int64) (err error) {
	// TODO Optimize
	temp, err := r.ReadFixed64(t)
	*x = int64(temp)
	return err
}

func (r *Reader) ReadSInt32(t WireType) (x int32, err error) {
	err = r.ReadToSInt32(t, &x)
	return x, err
}

func (r *Reader) ReadToSInt32(t WireType, x *int32) (err error) {
	err = r.ReadToInt32(t, x)
	if err != nil {
		return err
	}

	// TODO This seems wrong?
	*x = (*x >> 1) ^ (((*x & 0x1) << 31) >> 31)
	return nil
}

func (r *Reader) ReadSInt64(t WireType) (x int64, err error) {
	err = r.ReadToSInt64(t, &x)
	return x, err
}

func (r *Reader) ReadToSInt64(t WireType, x *int64) (err error) {
	err = r.ReadToInt64(t, x)
	if err != nil {
		return err
	}

	// TODO This seems wrong?
	*x = (*x >> 1) ^ (((*x & 0x1) << 63) >> 63)
	return nil
}
