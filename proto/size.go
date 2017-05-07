package proto

const SizeFixed8 = 1
const SizeFixed32 = 4
const SizeFixed64 = 8

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
