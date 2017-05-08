package proto

func SizeDoubleField(id int, x float64) (n int) {
	n += SizeKey(id)
	n += 8

	return n
}

func SizeFloatField(id int, x float32) (n int) {
	n += SizeKey(id)
	n += 4

	return n
}

func SizeInt64Field(id int, x int64) (n int) {
	n += SizeKey(id)
	n += SizeVarInt64(uint64(x))

	return n
}

func SizeUInt64Field(id int, x uint64) (n int) {
	n += SizeKey(id)
	n += SizeVarInt64(x)

	return n
}

func SizeInt32Field(id int, x int32) (n int) {
	n += SizeKey(id)
	n += SizeVarInt32(uint32(x))

	return n
}

func SizeUInt32Field(id int, x uint32) (n int) {
	n += SizeKey(id)
	n += SizeVarInt32(x)

	return n
}

func SizeFixed64Field(id int, x uint64) (n int) {
	n += SizeKey(id)
	n += 8

	return n
}

func SizeFixed32Field(id int, x uint32) (n int) {
	n += SizeKey(id)
	n += 4

	return n
}

func SizeBoolField(id int, x bool) (n int) {
	n += SizeKey(id)
	n += 1

	return n
}

func SizeStringField(id int, x string) (n int) {
	return SizeBytesField(id, []byte(x))
}

func SizeMessageField(id int, x MarshallerTo) (n int) {
	size := x.MarshalSize()

	n += SizeKey(id)
	n += SizeVarInt(uint(size))
	n += size

	return n
}

func SizeBytesField(id int, x []byte) (n int) {
	size := len(x)

	n += SizeKey(id)
	n += SizeVarInt(uint(size))
	n += size

	return n
}

func SizeEnumField(id int, x int) (n int) {
	n += SizeKey(id)
	n += SizeVarInt(uint(x))

	return n
}

func SizeSFixed32Field(id int, x int32) (n int) {
	n += SizeKey(id)
	n += 4

	return n
}

func SizeSFixed64Field(id int, x int64) (n int) {
	n += SizeKey(id)
	n += 8

	return n
}

func SizeSInt32Field(id int, x int32) (n int) {
	n += SizeKey(id)
	n += SizeZigZag32(uint32(x))

	return n
}

func SizeSInt64Field(id int, x int64) (n int) {
	n += SizeKey(id)
	n += SizeZigZag64(uint64(x))

	return n
}
