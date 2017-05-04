package proto

func EncodeId(id int, t WireType) (data []byte) {
	return EncodeVarInt(uint((id << 3) | int(t)))
}

// TODO Investigate using int instead of uint everywhere.
func EncodeFixed32(x uint32) (data []byte) {
	return []byte{
		byte(x),
		byte(x >> 8),
		byte(x >> 16),
		byte(x >> 24),
	}
}

func EncodeFixed64(x uint64) (data []byte) {
	return []byte{
		byte(x),
		byte(x >> 8),
		byte(x >> 16),
		byte(x >> 24),
		byte(x >> 32),
		byte(x >> 40),
		byte(x >> 48),
		byte(x >> 56),
	}
}

func EncodeVarInt(x uint) (data []byte) {
	for x >= 1<<7 {
		data = append(data, byte(x&0x7f|0x80))
		x >>= 7
	}

	data = append(data, byte(x))
	return data
}

func EncodeVarInt32(x uint32) (data []byte) {
	for x >= 1<<7 {
		data = append(data, byte(x&0x7f|0x80))
		x >>= 7
	}

	data = append(data, byte(x))
	return data
}

func EncodeVarInt64(x uint64) (data []byte) {
	for x >= 1<<7 {
		data = append(data, byte(x&0x7f|0x80))
		x >>= 7
	}

	data = append(data, byte(x))
	return data
}

func EncodeZigZag32(x uint32) (data []byte) {
	return EncodeVarInt32(uint32((x << 1) ^ (x >> 31)))
}

func EncodeZigZag64(x uint64) (data []byte) {
	return EncodeVarInt64(uint64((x << 1) ^ (x >> 63)))
}
