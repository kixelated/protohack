package proto

import (
	"math"
)

func WriteFixed8(data []byte, x uint8) (n int) {
	data[0] = byte(x)
	return 1
}

func WriteFixed32(data []byte, x uint32) (n int) {
	data[0] = byte(x)
	data[1] = byte(x >> 8)
	data[2] = byte(x >> 16)
	data[3] = byte(x >> 24)
	return 4
}

func WriteFixed64(data []byte, x uint64) (n int) {
	data[0] = byte(x)
	data[1] = byte(x >> 8)
	data[2] = byte(x >> 16)
	data[3] = byte(x >> 24)
	data[4] = byte(x >> 32)
	data[5] = byte(x >> 40)
	data[6] = byte(x >> 48)
	data[7] = byte(x >> 56)

	return 8
}

func WriteVarInt(data []byte, x uint) (n int) {
	for x >= 1<<7 {
		data[n] = byte(x&0x7f | 0x80)

		x >>= 7
		n += 1
	}

	data[n] = byte(x)
	return n + 1
}

func WriteVarInt32(data []byte, x uint32) (n int) {
	for x >= 1<<7 {
		data[n] = byte(x&0x7f | 0x80)

		x >>= 7
		n += 1
	}

	data[n] = byte(x)
	return n + 1
}

func WriteVarInt64(data []byte, x uint64) (n int) {
	for x >= 1<<7 {
		data[n] = byte(x&0x7f | 0x80)

		x >>= 7
		n += 1
	}

	data[n] = byte(x)
	return n + 1
}

func WriteZigZag32(data []byte, x uint32) (n int) {
	// TODO Optimize
	return WriteVarInt32(data, (x<<1)^(x>>31))
}

func WriteZigZag64(data []byte, x uint64) (n int) {
	// TODO Optimize
	return WriteVarInt64(data, (x<<1)^(x>>63))
}

func WriteKey(data []byte, id int, t WireType) (n int) {
	return WriteVarInt64(data, uint64(id<<3)|uint64(t))
}

func WriteDouble(data []byte, x float64) (n int) {
	return WriteFixed64(data[n:], math.Float64bits(x))
}

func WriteFloat(data []byte, x float32) (n int) {
	return WriteFixed32(data[n:], math.Float32bits(x))
}

func WriteInt64(data []byte, x int64) (n int) {
	return WriteVarInt64(data[n:], uint64(x))
}

func WriteUInt64(data []byte, x uint64) (n int) {
	return WriteVarInt64(data[n:], x)
}

func WriteInt32(data []byte, x int32) (n int) {
	return WriteVarInt32(data[n:], uint32(x))
}

func WriteUInt32(data []byte, x uint32) (n int) {
	return WriteVarInt32(data[n:], x)
}

func WriteBool(data []byte, x bool) (n int) {
	return WriteByte(data[n:], 1)
}

func WriteString(data []byte, x string) (n int) {
	return WriteBytes(data, []byte(x))
}

func WriteStringLength(data []byte, x string) (n int) {
	return WriteBytesLength(data, []byte(x))
}

func WriteMessage(data []byte, x MarshallerTo) (n int) {
	n, err := x.MarshalTo(data[n:])
	if err != nil {
		panic(err.Error()) // TODO
	}

	return n
}

func WriteMessageLength(data []byte, x MarshallerTo) (n int) {
	size := x.MarshalSize()
	n += WriteVarInt(data[n:], uint(size))

	temp, err := x.MarshalTo(data[n:])
	if err != nil {
		panic(err.Error()) // TODO
	}

	n += temp
	return n
}

func WriteByte(data []byte, x byte) (n int) {
	data[0] = x
	return 1
}

func WriteBytes(data []byte, x []byte) (n int) {
	return copy(data[n:], x)
}

func WriteBytesLength(data []byte, x []byte) (n int) {
	size := len(x)

	n += WriteVarInt(data[n:], uint(size))
	n += copy(data[n:], x)

	return n
}

func WriteEnum(data []byte, x int) (n int) {
	return WriteVarInt(data[n:], uint(x))
}

func WriteSFixed32(data []byte, x int32) (n int) {
	return WriteFixed32(data[n:], uint32(x))
}

func WriteSFixed64(data []byte, x int64) (n int) {
	return WriteFixed64(data[n:], uint64(x))
}

func WriteSInt32(data []byte, x int32) (n int) {
	return WriteZigZag32(data[n:], uint32(x))
}

func WriteSInt64(data []byte, x int64) (n int) {
	return WriteZigZag64(data[n:], uint64(x))
}
