package proto

func SizeKey(id int) (n int) {
	// TODO Optimize
	return SizeVarInt64(uint64(id << 3))
}

func SizeVarInt(x uint) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}

	return n
}

func SizeVarInt32(x uint32) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}

	return n
}

func SizeVarInt64(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}

	return n
}

func SizeZigZag32(x uint32) (n int) {
	// TODO Optimize
	return SizeVarInt32((x << 1) ^ (x >> 31))
}

func SizeZigZag64(x uint64) (n int) {
	// TODO Optimize
	return SizeVarInt64((x << 1) ^ (x >> 63))
}

func SizeDouble(x float64) (n int) {
	return 8
}

func SizeFloat(x float32) (n int) {
	return 4
}

func SizeInt64(x int64) (n int) {
	return SizeVarInt64(uint64(x))
}

func SizeUInt64(x uint64) (n int) {
	return SizeVarInt64(x)
}

func SizeInt32(x int32) (n int) {
	return SizeVarInt32(uint32(x))
}

func SizeUInt32(x uint32) (n int) {
	return SizeVarInt32(x)
}

func SizeFixed64(x uint64) (n int) {
	return 8
}

func SizeFixed32(x uint32) (n int) {
	return 4
}

func SizeBool(x bool) (n int) {
	return 1
}

func SizeString(x string) (n int) {
	return SizeBytes([]byte(x))
}

func SizeGroup(x MarshallerTo) (n int) {
	size := x.MarshalSize()
	return SizeVarInt(uint(size)) + size
}

func SizeMessage(x MarshallerTo) (n int) {
	size := x.MarshalSize()
	return SizeVarInt(uint(size)) + size
}

func SizeBytes(x []byte) (n int) {
	size := len(x)
	return SizeVarInt(uint(size)) + size
}

func SizeEnum(x int) (n int) {
	return SizeVarInt32(uint32(x))
}

func SizeSFixed32(x int32) (n int) {
	return 4
}

func SizeSFixed64(x int64) (n int) {
	return 8
}

func SizeSInt32(x int32) (n int) {
	return SizeVarInt32(uint32(x))
}

func SizeSInt64(x int64) (n int) {
	return SizeVarInt64(uint64(x))
}
