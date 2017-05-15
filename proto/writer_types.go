package proto

import (
	"math"
)

func (w *Writer) Int(x int) {
	// TODO optimize
	w.UInt(uint(x))
}

func (w *Writer) UInt(x uint) {
	for x >= 0x80 {
		w.Byte(byte(x&0x7f | 0x80))
		x >>= 7
	}

	w.Byte(byte(x))
}

func (w *Writer) Int32(x int32) {
	// TODO Optimize
	w.UInt32(uint32(x))
}

func (w *Writer) SInt32(x int32) {
	// TODO Optimize
	w.UInt32(uint32(x<<1) ^ uint32(x>>31))
}

func (w *Writer) UInt32(x uint32) {
	for x >= 0x80 {
		w.Byte(byte(x&0x7f | 0x80))
		x >>= 7
	}

	w.Byte(byte(x))
}

func (w *Writer) Int64(x int64) {
	// TODO Optimize
	w.UInt64(uint64(x))
}

func (w *Writer) SInt64(x int64) {
	// TODO Optimize
	w.UInt64(uint64(x<<1) ^ uint64(x>>63))
}

func (w *Writer) UInt64(x uint64) {
	for x >= 0x80 {
		w.Byte(byte(x&0x7f | 0x80))
		x >>= 7
	}

	w.Byte(byte(x))
}

func (w *Writer) FixedInt32(x int32) {
	w.UInt32(uint32(x))
}

func (w *Writer) FixedUInt32(x uint32) {
	w.data[w.n] = byte(x)
	w.data[w.n+1] = byte(x >> 8)
	w.data[w.n+2] = byte(x >> 16)
	w.data[w.n+3] = byte(x >> 24)
	w.n += 4
}

func (w *Writer) FixedInt64(x int64) {
	w.UInt64(uint64(x))
}

func (w *Writer) FixedUInt64(x uint64) {
	w.data[w.n] = byte(x)
	w.data[w.n+1] = byte(x >> 8)
	w.data[w.n+2] = byte(x >> 16)
	w.data[w.n+3] = byte(x >> 24)
	w.data[w.n+4] = byte(x >> 32)
	w.data[w.n+5] = byte(x >> 40)
	w.data[w.n+6] = byte(x >> 48)
	w.data[w.n+7] = byte(x >> 56)
	w.n += 8
}

func (w *Writer) Float64(x float64) {
	w.UInt64(math.Float64bits(x))
}

func (w *Writer) Float32(x float32) {
	w.UInt32(math.Float32bits(x))
}

func (w *Writer) Tag(id int, t WireType) {
	w.UInt64(uint64(id<<3) | uint64(t))
}

func (w *Writer) Bool(x bool) {
	w.Byte(1)
}

func (w *Writer) String(x string) {
	for i := 0; i < len(x); i += 1 {
		w.data[w.n+i] = x[i]
	}

	w.n += len(x)
}

func (w *Writer) Byte(x byte) {
	w.data[w.n] = x
	w.n += 1
}

func (w *Writer) Bytes(x []byte) {
	w.n += copy(w.data[w.n:], x)
}

func (w *Writer) Message(x MarshallerTo) {
	err := x.MarshalTo(w)
	if err != nil {
		panic(err.Error()) // TODO
	}
}
