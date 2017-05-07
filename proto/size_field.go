package proto

func SizeKey(id int, t WireType) (n int) {
	// TODO Optimize
	return SizeVarInt64((uint64(id<<3) | uint64(t)))
}

func SizeFieldDouble(id int, x float64) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_64BIT)
	n += SizeFixed64

	return n
}

func SizeFieldFloat(id int, x float32) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_32BIT)
	n += SizeFixed32

	return n
}

func SizeFieldInt64(id int, x int64) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeVarInt64(uint64(x))

	return n
}

func SizeFieldUInt64(id int, x uint64) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeVarInt64(x)

	return n
}

func SizeFieldInt32(id int, x int32) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeVarInt32(uint32(x))

	return n
}

func SizeFieldUInt32(id int, x uint32) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeVarInt32(x)

	return n
}

func SizeFieldFixed64(id int, x uint64) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_64BIT)
	n += SizeFixed64

	return n
}

func SizeFieldFixed32(id int, x uint32) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_32BIT)
	n += SizeFixed32

	return n
}

func SizeFieldBool(id int, x bool) (n int) {
	if !x {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeFixed8

	return n
}

func SizeFieldString(id int, x string) (n int) {
	return SizeFieldBytes(id, []byte(x))
}

func SizeFieldGroup(id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()
	if size == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_GROUP_START)
	n += SizeVarInt(uint(size))
	n += size
	n += SizeKey(id, WIRETYPE_GROUP_END)

	return n
}

func SizeFieldMessage(id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()
	if size == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_LENGTH)
	n += SizeVarInt(uint(size))
	n += size

	return n
}

func SizeFieldBytes(id int, x []byte) (n int) {
	size := len(x)
	if size == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_LENGTH)
	n += SizeVarInt(uint(size))
	n += size

	return n
}

func SizeFieldEnum(id int, x int) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeVarInt(uint(x))

	return n
}

func SizeFieldSFixed32(id int, x int32) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_32BIT)
	n += SizeFixed32

	return n
}

func SizeFieldSFixed64(id int, x int64) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_64BIT)
	n += SizeFixed64

	return n
}

func SizeFieldSInt32(id int, x int32) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeZigZag32(uint32(x))

	return n
}

func SizeFieldSInt64(id int, x int64) (n int) {
	if x == 0 {
		return 0
	}

	n += SizeKey(id, WIRETYPE_VARINT)
	n += SizeZigZag64(uint64(x))

	return n
}
