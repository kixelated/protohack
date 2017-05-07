package proto

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
