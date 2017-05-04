package proto

import (
	"bytes"
	"math"
)

type Marshaller interface {
	Marshal() (data []byte, err error)
}

type WireType int

const (
	WIRETYPE_VARINT      WireType = 0
	WIRETYPE_64BIT       WireType = 1
	WIRETYPE_LENGTH      WireType = 2
	WIRETYPE_GROUP_START WireType = 3
	WIRETYPE_GROUP_END   WireType = 4
	WIRETYPE_32BIT       WireType = 5
)

type Writer struct {
	b bytes.Buffer
}

func NewWriter() (w *Writer) {
	return new(Writer)
}

func (w *Writer) WriteId(id int, t WireType) {
	w.Write(EncodeId(id, t))
}

func (w *Writer) WriteDouble(id int, x float64) {
	w.WriteId(id, WIRETYPE_64BIT)
	w.Write(EncodeFixed64(math.Float64bits(x)))
}

func (w *Writer) WriteFloat(id int, x float32) {
	w.WriteId(id, WIRETYPE_32BIT)
	w.Write(EncodeFixed32(math.Float32bits(x)))
}

func (w *Writer) WriteInt64(id int, x int64) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeVarInt64(uint64(x)))
}

func (w *Writer) WriteUInt64(id int, x uint64) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeVarInt64(x))
}

func (w *Writer) WriteInt32(id int, x int32) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeVarInt32(uint32(x)))
}

func (w *Writer) WriteUInt32(id int, x uint32) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeVarInt32(x))
}

func (w *Writer) WriteFixed64(id int, x uint64) {
	w.WriteId(id, WIRETYPE_64BIT)
	w.Write(EncodeFixed64(x))
}

func (w *Writer) WriteFixed32(id int, x uint32) {
	w.WriteId(id, WIRETYPE_32BIT)
	w.Write(EncodeFixed32(x))
}

func (w *Writer) WriteBool(id int, x bool) {
	w.WriteId(id, WIRETYPE_VARINT)
	if x {
		w.WriteByte(2)
	} else {
		w.WriteByte(0)
	}
}

func (w *Writer) WriteString(id int, x string) {
	w.WriteBytes(id, []byte(x))
}

func (w *Writer) WriteGroup(id int, x Marshaller) (err error) {
	w.WriteId(id, WIRETYPE_GROUP_START)

	data, err := x.Marshal()
	if err != nil {
		return err
	}

	w.Write(EncodeVarInt(uint(len(data))))
	w.Write(data)

	w.WriteId(id, WIRETYPE_GROUP_END)
	return nil
}

func (w *Writer) WriteMessage(id int, x Marshaller) (err error) {
	w.WriteId(id, WIRETYPE_LENGTH)

	data, err := x.Marshal()
	if err != nil {
		return err
	}

	w.Write(EncodeVarInt(uint(len(data))))
	w.Write(data)

	return nil
}

func (w *Writer) WriteBytes(id int, x []byte) {
	w.WriteId(id, WIRETYPE_LENGTH)
	w.Write(EncodeVarInt(uint(len(x))))
	w.Write(x)
}

func (w *Writer) WriteEnum(id int, x int) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeVarInt(uint(x)))
}

func (w *Writer) WriteSFixed32(id int, x int32) {
	w.WriteId(id, WIRETYPE_32BIT)
	w.Write(EncodeFixed32(uint32(x)))
}

func (w *Writer) WriteSFixed64(id int, x int64) {
	w.WriteId(id, WIRETYPE_64BIT)
	w.Write(EncodeFixed64(uint64(x)))
}

func (w *Writer) WriteSInt32(id int, x int32) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeZigZag32(uint32(x)))
}

func (w *Writer) WriteSInt64(id int, x int64) {
	w.WriteId(id, WIRETYPE_VARINT)
	w.Write(EncodeZigZag64(uint64(x)))
}

func (w *Writer) Write(x []byte) {
	_, _ = w.b.Write(x)
}

func (w *Writer) WriteByte(x byte) {
	_ = w.b.WriteByte(x)
}

func (w *Writer) Bytes() []byte {
	return w.b.Bytes()
}
