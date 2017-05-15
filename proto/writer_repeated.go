package proto

func (w *Writer) IntPacked(tagRaw []byte, x []int) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.Int(y)
	}
}

func (w *Writer) UIntPacked(tagRaw []byte, x []uint) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.UInt(y)
	}
}

func (w *Writer) Int32Packed(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.Int32(y)
	}
}

func (w *Writer) Int32Repeated(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		w.TagRaw(tagRaw)
		w.Int32(y)
	}
}

func (w *Writer) UInt32Packed(tagRaw []byte, x []uint32) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.UInt32(y)
	}
}

func (w *Writer) SInt32Packed(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.SInt32(y)
	}
}

func (w *Writer) Int64Packed(tagRaw []byte, x []int64) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.Int64(y)
	}
}

func (w *Writer) UInt64Packed(tagRaw []byte, x []uint64) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.UInt64(y)
	}
}

func (w *Writer) SInt64Packed(tagRaw []byte, x []int64) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := w.s.Shift()
	w.Int(size)

	for _, y := range x {
		w.SInt64(y)
	}
}

func (w *Writer) FixedInt32Packed(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 4
	w.Int(size)

	for _, y := range x {
		w.Int32(y)
	}
}

func (w *Writer) FixedUInt32Packed(tagRaw []byte, x []uint32) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 4
	w.Int(size)

	for _, y := range x {
		w.UInt32(y)
	}
}

func (w *Writer) FixedInt64Packed(tagRaw []byte, x []int64) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 8
	w.Int(size)

	for _, y := range x {
		w.Int64(y)
	}
}

func (w *Writer) FixedUInt64Packed(tagRaw []byte, x []uint64) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 8
	w.Int(size)

	for _, y := range x {
		w.UInt64(y)
	}
}

func (w *Writer) Float32Packed(tagRaw []byte, x []float32) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 4
	w.Int(size)

	for _, y := range x {
		w.Float32(y)
	}
}

func (w *Writer) Float64Packed(tagRaw []byte, x []float64) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 8
	w.Int(size)

	for _, y := range x {
		w.Float64(y)
	}
}

func (w *Writer) BoolPacked(tagRaw []byte, x []bool) {
	if len(x) == 0 {
		return
	}

	w.TagRaw(tagRaw)

	size := len(x) * 1
	w.Int(size)

	for _, y := range x {
		w.Bool(y)
	}
}

func (w *Writer) BytesRepeated(tagRaw []byte, x [][]byte) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		w.TagRaw(tagRaw)

		size := len(y)
		w.Int(size)
		w.Bytes(y)
	}
}

func (w *Writer) StringRepeated(tagRaw []byte, x []string) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		w.TagRaw(tagRaw)

		size := len(y)
		w.Int(size)
		w.String(y)
	}
}

func (w *Writer) MessageRepeated(tagRaw []byte, x []MarshallerTo) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		w.TagRaw(tagRaw)

		size := w.s.Shift()
		w.Int(size)
		y.MarshalTo(w)
	}
}
