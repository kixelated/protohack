package proto

import (
	"errors"
	"math"
)

var ErrIntOverflow = errors.New("integer overflow")
var ErrUnexpectedEOF = errors.New("unexpected EOF")

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
	temp, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	*id = int(temp >> 3)
	*t = WireType(temp & 0x7)

	return nil
}

func (r *Reader) ReadVarInt64() (x uint64, err error) {
	err = r.ReadToVarInt64(&x)
	return x, err
}

func (r *Reader) ReadToVarInt64(x *uint64) (err error) {
	for shift := uint(0); ; shift += 7 {
		if len(r.data) < 1 {
			return ErrUnexpectedEOF
		}

		if shift >= 64 {
			return ErrIntOverflow
		}

		temp := r.data[0]
		r.data = r.data[1:]

		*x |= (uint64(temp) & 0x7F << shift)
		if temp < 0x80 {
			break
		}
	}

	return nil
}

func (r *Reader) ReadVarInt32() (x uint32, err error) {
	err = r.ReadToVarInt32(&x)
	return x, err
}

func (r *Reader) ReadToVarInt32(x *uint32) (err error) {
	for shift := uint(0); ; shift += 7 {
		if len(r.data) < 1 {
			return ErrUnexpectedEOF
		}

		if shift >= 32 {
			return ErrIntOverflow
		}

		temp := r.data[0]
		r.data = r.data[1:]

		*x |= (uint32(temp) & 0x7F << shift)
		if temp < 0x80 {
			break
		}
	}

	return nil
}

func (r *Reader) ReadDouble() (x float64, err error) {
	err = r.ReadToDouble(&x)
	return x, err
}

func (r *Reader) ReadToDouble(x *float64) (err error) {
	temp, err := r.ReadFixed64()
	if err != nil {
		return err
	}

	*x = math.Float64frombits(temp)
	return nil
}

func (r *Reader) ReadFloat() (x float32, err error) {
	err = r.ReadToFloat(&x)
	return x, err
}

func (r *Reader) ReadToFloat(x *float32) (err error) {
	temp, err := r.ReadFixed32()
	if err != nil {
		return err
	}

	*x = math.Float32frombits(temp)
	return nil
}

func (r *Reader) ReadInt64() (x int64, err error) {
	err = r.ReadToInt64(&x)
	return x, err
}

func (r *Reader) ReadToInt64(x *int64) (err error) {
	// TODO Optimize
	temp, err := r.ReadVarInt64()
	*x = int64(temp)
	return err
}

func (r *Reader) ReadUInt64() (x uint64, err error) {
	err = r.ReadToUInt64(&x)
	return x, err
}

func (r *Reader) ReadToUInt64(x *uint64) (err error) {
	return r.ReadToVarInt64(x)
}

func (r *Reader) ReadInt32() (x int32, err error) {
	err = r.ReadToInt32(&x)
	return x, err
}

func (r *Reader) ReadToInt32(x *int32) (err error) {
	// TODO Optimize
	temp, err := r.ReadVarInt32()
	*x = int32(temp)
	return err
}

func (r *Reader) ReadUInt32() (x uint32, err error) {
	err = r.ReadToUInt32(&x)
	return x, err
}

func (r *Reader) ReadToUInt32(x *uint32) (err error) {
	return r.ReadToVarInt32(x)
}

func (r *Reader) ReadFixed64() (x uint64, err error) {
	err = r.ReadToFixed64(&x)
	return x, err
}

func (r *Reader) ReadToFixed64(x *uint64) (err error) {
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

func (r *Reader) ReadFixed32() (x uint32, err error) {
	err = r.ReadToFixed32(&x)
	return x, err
}

func (r *Reader) ReadToFixed32(x *uint32) (err error) {
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

func (r *Reader) ReadBool() (x bool, err error) {
	err = r.ReadToBool(&x)
	return x, err
}

func (r *Reader) ReadToBool(x *bool) (err error) {
	if len(r.data) < 1 {
		return ErrUnexpectedEOF
	}

	if r.data[0] > 1 {
		return errors.New("unexpected boolean value")
	}

	*x = (r.data[0] == 1)
	return nil
}

func (r *Reader) ReadToMessage(x Unmarshaller) (err error) {
	size, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	if len(r.data) < int(size) {
		return ErrUnexpectedEOF
	}

	temp := r.data[:size]
	r.data = r.data[size:]

	return x.Unmarshal(temp)
}

func (r *Reader) ReadString() (x string, err error) {
	err = r.ReadToString(&x)
	return x, err
}

func (r *Reader) ReadToString(x *string) (err error) {
	temp, err := r.ReadBytes()
	if err != nil {
		return err
	}

	*x = string(temp)
	return nil
}

func (r *Reader) ReadBytes() (x []byte, err error) {
	err = r.ReadToBytes(&x)
	return x, err
}

func (r *Reader) ReadToBytes(x *[]byte) (err error) {
	size, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	if len(r.data) < int(size) {
		return ErrUnexpectedEOF
	}

	*x = r.data[:size]
	r.data = r.data[size:]

	return nil
}

func (r *Reader) ReadEnum() (x int, err error) {
	err = r.ReadToEnum(&x)
	return x, err
}

func (r *Reader) ReadToEnum(x *int) (err error) {
	temp, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	*x = int(temp)
	return nil
}

func (r *Reader) ReadSFixed32() (x int32, err error) {
	err = r.ReadToSFixed32(&x)
	return x, err
}

func (r *Reader) ReadToSFixed32(x *int32) (err error) {
	// TODO Optimize
	temp, err := r.ReadFixed32()
	*x = int32(temp)
	return err
}

func (r *Reader) ReadSFixed64() (x int64, err error) {
	err = r.ReadToSFixed64(&x)
	return x, err
}

func (r *Reader) ReadToSFixed64(x *int64) (err error) {
	// TODO Optimize
	temp, err := r.ReadFixed64()
	*x = int64(temp)
	return err
}

func (r *Reader) ReadSInt32() (x int32, err error) {
	panic("TODO")
}

func (r *Reader) ReadSInt64() (x int64, err error) {
	panic("TODO")
}
