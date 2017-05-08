package proto

import (
	"math"
)

func WriteFieldDouble(data []byte, id int, x float64) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_64BIT)
	n += WriteFixed64(data[n:], math.Float64bits(x))

	return n
}

func WriteFieldFloat(data []byte, id int, x float32) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_32BIT)
	n += WriteFixed32(data[n:], math.Float32bits(x))

	return n
}

func WriteFieldInt64(data []byte, id int, x int64) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt64(data[n:], uint64(x))

	return n
}

func WriteFieldUInt64(data []byte, id int, x uint64) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt64(data[n:], x)

	return n
}

func WriteFieldInt32(data []byte, id int, x int32) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt32(data[n:], uint32(x))

	return n
}

func WriteFieldUInt32(data []byte, id int, x uint32) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt32(data[n:], x)

	return n
}

func WriteFieldFixed64(data []byte, id int, x uint64) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_64BIT)
	n += WriteFixed64(data[n:], x)

	return n
}

func WriteFieldFixed32(data []byte, id int, x uint32) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_32BIT)
	n += WriteFixed32(data[n:], x)

	return n
}

func WriteFieldBool(data []byte, id int, x bool) (n int) {
	if !x {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteFixed8(data[n:], 1)

	return n
}

func WriteFieldString(data []byte, id int, x string) (n int) {
	return WriteFieldBytes(data, id, []byte(x))
}

func WriteFieldGroup(data []byte, id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()
	if size == 0 {
		return 0
	}

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

func WriteFieldMessage(data []byte, id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()
	if size == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_LENGTH)
	n += WriteVarInt(data[n:], uint(size))

	temp, err := x.MarshalTo(data[n:])
	if err != nil {
		panic(err.Error()) // TODO
	}

	n += temp

	return n
}

func WriteFieldBytes(data []byte, id int, x []byte) (n int) {
	if len(x) == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_LENGTH)
	n += WriteVarInt(data[n:], uint(len(x)))
	n += copy(data[n:], x)

	return n
}

func WriteFieldEnum(data []byte, id int, x int) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteVarInt(data[n:], uint(x))

	return n
}

func WriteFieldSFixed32(data []byte, id int, x int32) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_32BIT)
	n += WriteFixed32(data[n:], uint32(x))

	return n
}

func WriteFieldSFixed64(data []byte, id int, x int64) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_64BIT)
	n += WriteFixed64(data[n:], uint64(x))

	return n
}

func WriteFieldSInt32(data []byte, id int, x int32) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteZigZag32(data[n:], uint32(x))

	return n
}

func WriteFieldSInt64(data []byte, id int, x int64) (n int) {
	if x == 0 {
		return 0
	}

	n += WriteKey(data[n:], id, WIRETYPE_VARINT)
	n += WriteZigZag64(data[n:], uint64(x))

	return n
}
