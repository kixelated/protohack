package proto

import (
	"github.com/pkg/errors"
)

func (r *Reader) ReadIntRepeated(t WireType, x []int) ([]int, error) {
	if t == WIRETYPE_VARINT {
		return r.readIntRepeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readIntPacked(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readIntRepeated(x []int) ([]int, error) {
	temp, err := r.readInt()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readIntPacked(x []int) ([]int, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readInt()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadUIntRepeated(t WireType, x []uint) ([]uint, error) {
	if t == WIRETYPE_VARINT {
		return r.readUIntRepeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readUIntPacked(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readUIntRepeated(x []uint) ([]uint, error) {
	temp, err := r.readUInt()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readUIntPacked(x []uint) ([]uint, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readUInt()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadFloat64Repeated(t WireType, x []float64) ([]float64, error) {
	if t == WIRETYPE_64BIT {
		return r.readFloat64Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readFloat64Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readFloat64Repeated(x []float64) ([]float64, error) {
	temp, err := r.readFloat64()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readFloat64Packed(x []float64) ([]float64, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readFloat64()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadFloat32Repeated(t WireType, x []float32) ([]float32, error) {
	if t == WIRETYPE_32BIT {
		return r.readFloat32Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readFloat32Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readFloat32Repeated(x []float32) ([]float32, error) {
	temp, err := r.readFloat32()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readFloat32Packed(x []float32) ([]float32, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readFloat32()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadInt64Repeated(t WireType, x []int64) ([]int64, error) {
	if t == WIRETYPE_VARINT {
		return r.readInt64Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readInt64Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readInt64Repeated(x []int64) ([]int64, error) {
	temp, err := r.readInt64()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readInt64Packed(x []int64) ([]int64, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readInt64()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadUInt64Repeated(t WireType, x []uint64) ([]uint64, error) {
	if t == WIRETYPE_VARINT {
		return r.readUInt64Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readUInt64Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readUInt64Repeated(x []uint64) ([]uint64, error) {
	temp, err := r.readUInt64()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readUInt64Packed(x []uint64) ([]uint64, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readUInt64()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadInt32Repeated(t WireType, x []int32) ([]int32, error) {
	if t == WIRETYPE_VARINT {
		return r.readInt32Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readInt32Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readInt32Repeated(x []int32) ([]int32, error) {
	temp, err := r.readInt32()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readInt32Packed(x []int32) ([]int32, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readInt32()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadUInt32Repeated(t WireType, x []uint32) ([]uint32, error) {
	if t == WIRETYPE_VARINT {
		return r.readUInt32Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readUInt32Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readUInt32Repeated(x []uint32) ([]uint32, error) {
	temp, err := r.readUInt32()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readUInt32Packed(x []uint32) ([]uint32, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readUInt32()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadFixed64Repeated(t WireType, x []uint64) ([]uint64, error) {
	if t == WIRETYPE_64BIT {
		return r.readFixed64Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readFixed64Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readFixed64Repeated(x []uint64) ([]uint64, error) {
	temp, err := r.readFixed64()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

// TODO optimize
func (r *Reader) readFixed64Packed(x []uint64) ([]uint64, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readFixed64()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadFixed32Repeated(t WireType, x []uint32) ([]uint32, error) {
	if t == WIRETYPE_32BIT {
		return r.readFixed32Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readFixed32Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readFixed32Repeated(x []uint32) ([]uint32, error) {
	temp, err := r.readFixed32()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

// TODO optimize
func (r *Reader) readFixed32Packed(x []uint32) ([]uint32, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readFixed32()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadBoolRepeated(t WireType, x []bool) ([]bool, error) {
	if t == WIRETYPE_VARINT {
		return r.readBoolRepeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readBoolPacked(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readBoolRepeated(x []bool) ([]bool, error) {
	temp, err := r.readBool()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

// TODO optimize
func (r *Reader) readBoolPacked(x []bool) ([]bool, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readBool()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadStringRepeated(t WireType, x []string) ([]string, error) {
	temp, err := r.readString()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) ReadBytesRepeated(t WireType, x [][]byte) ([][]byte, error) {
	temp, err := r.readBytes()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) ReadSFixed32Repeated(t WireType, x []int32) ([]int32, error) {
	if t == WIRETYPE_32BIT {
		return r.readSFixed32Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readSFixed32Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readSFixed32Repeated(x []int32) ([]int32, error) {
	temp, err := r.readSFixed32()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

// TODO optimize
func (r *Reader) readSFixed32Packed(x []int32) ([]int32, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readSFixed32()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadSFixed64Repeated(t WireType, x []int64) ([]int64, error) {
	if t == WIRETYPE_64BIT {
		return r.readSFixed64Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readSFixed64Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readSFixed64Repeated(x []int64) ([]int64, error) {
	temp, err := r.readSFixed64()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readSFixed64Packed(x []int64) ([]int64, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readSFixed64()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadSInt32Repeated(t WireType, x []int32) ([]int32, error) {
	if t == WIRETYPE_VARINT {
		return r.readSInt32Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readSInt32Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readSInt32Repeated(x []int32) ([]int32, error) {
	temp, err := r.readSInt32()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readSInt32Packed(x []int32) ([]int32, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readSInt32()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}

func (r *Reader) ReadSInt64Repeated(t WireType, x []int64) ([]int64, error) {
	if t == WIRETYPE_VARINT {
		return r.readSInt64Repeated(x)
	} else if t == WIRETYPE_LENGTH {
		return r.readSInt64Packed(x)
	} else {
		return nil, errors.New("unexpected wire type")
	}
}

func (r *Reader) readSInt64Repeated(x []int64) ([]int64, error) {
	temp, err := r.readSInt64()
	if err != nil {
		return nil, err
	}

	x = append(x, temp)
	return x, nil
}

func (r *Reader) readSInt64Packed(x []int64) ([]int64, error) {
	size, err := r.readInt()
	if err != nil {
		return nil, err
	}

	startLen := r.Len()

	for startLen-r.Len() < size {
		temp, err := r.readSInt64()
		if err != nil {
			return nil, err
		}

		x = append(x, temp)
	}

	return x, nil
}
