package proto

import (
	"math"
)

func EncodeUInt32(x uint32) (data []byte) {
	return []byte{
		byte(x),
		byte(x >> 8),
		byte(x >> 16),
		byte(x >> 24),
	}
}

func EncodeUInt64(x uint64) (data []byte) {
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

func EncodeVarUInt(x uint) (data []byte) {
	for x >= 1<<7 {
		data = append(data, byte(x&0x7f|0x80))
		x >>= 7
	}

	return append(data, byte(x))
}

func EncodeVarUInt32(x uint32) (data []byte) {
	for x >= 1<<7 {
		data = append(data, byte(x&0x7f|0x80))
		x >>= 7
	}

	return append(data, byte(x))
}

func EncodeVarUInt64(x uint64) (data []byte) {
	for x >= 1<<7 {
		data = append(data, byte(x&0x7f|0x80))
		x >>= 7
	}

	return append(data, byte(x))
}

func EncodeVarSInt32(x int32) (data []byte) {
	// TODO Optimize
	return EncodeVarUInt32(uint32(x<<1) ^ uint32(x>>31))
}

func EncodeVarSInt64(x int64) (data []byte) {
	// TODO Optimize
	return EncodeVarUInt64(uint64(x<<1) ^ uint64(x>>63))
}

func EncodeTag(id int, t WireType) (data []byte) {
	return EncodeVarUInt64(uint64(id<<3) | uint64(t))
}

func EncodeFloat64(x float64) (data []byte) {
	return EncodeUInt64(math.Float64bits(x))
}

func EncodeFloat32(x float32) (data []byte) {
	return EncodeUInt32(math.Float32bits(x))
}

func EncodeBool(x bool) (data []byte) {
	if x {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func EncodeString(x string) (data []byte) {
	return EncodeBytes([]byte(x))
}

func EncodeBytes(x []byte) (data []byte) {
	return x
}

func EncodeSInt32(x int32) (data []byte) {
	return EncodeUInt32(uint32(x))
}

func EncodeSInt64(x int64) (data []byte) {
	return EncodeUInt64(uint64(x))
}
