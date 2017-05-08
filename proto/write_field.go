package proto

import (
	"math"
)

func WriteDoubleField(data []byte, id int, x float64) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_64BIT)
	n += WriteFixed64(data[n:], math.Float64bits(x))

	return n
}

func WriteFloatField(data []byte, id int, x float32) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_32BIT)
	n += WriteFixed32(data[n:], math.Float32bits(x))

	return n
}

func WriteInt64Field(data []byte, id int, x int64) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt64(data[n:], uint64(x))

	return n
}

func WriteUInt64Field(data []byte, id int, x uint64) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt64(data[n:], x)

	return n
}

func WriteInt32Field(data []byte, id int, x int32) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt32(data[n:], uint32(x))

	return n
}

func WriteUInt32Field(data []byte, id int, x uint32) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt32(data[n:], x)

	return n
}

func WriteFixed64Field(data []byte, id int, x uint64) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_64BIT)
	n += WriteFixed64(data[n:], x)

	return n
}

func WriteFixed32Field(data []byte, id int, x uint32) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_32BIT)
	n += WriteFixed32(data[n:], x)

	return n
}

func WriteBoolField(data []byte, id int, x bool) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteFixed8(data[n:], 1)

	return n
}

func WriteStringField(data []byte, id int, x string) (n int) {
	return WriteBytesField(data, id, []byte(x))
}

func WriteGroupField(data []byte, id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()

	n += WriteKey(data[n:], id, WIRETYPE_GROUP_START)
	n += WriteVarInt(data[n:], uint(size))

	temp, err := x.MarshalTo(data[n:])
	if err != nil {
		panic(err.Error()) // TODO
	}

	n += temp
	n += WriteKey(data[n:], id, WIRETYPE_GROUP_END)

	return n
}

func WriteMessageField(data []byte, id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()

	n += WriteKey(data[n:], id, WIRETYPE_LENGTH)
	n += WriteVarInt(data[n:], uint(size))

	temp, err := x.MarshalTo(data[n:])
	if err != nil {
		panic(err.Error()) // TODO
	}

	n += temp

	return n
}

func WriteBytesField(data []byte, id int, x []byte) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_LENGTH)
	n += WriteVarInt(data[n:], uint(len(x)))
	n += copy(data[n:], x)

	return n
}

func WriteEnumField(data []byte, id int, x int) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt(data[n:], uint(x))

	return n
}

func WriteSFixed32Field(data []byte, id int, x int32) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_32BIT)
	n += WriteFixed32(data[n:], uint32(x))

	return n
}

func WriteSFixed64Field(data []byte, id int, x int64) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_64BIT)
	n += WriteFixed64(data[n:], uint64(x))

	return n
}

func WriteSInt32Field(data []byte, id int, x int32) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteZigZag32(data[n:], uint32(x))

	return n
}

func WriteSInt64Field(data []byte, id int, x int64) (n int) {
	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteZigZag64(data[n:], uint64(x))

	return n
}
