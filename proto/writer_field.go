package proto

func (w *Writer) TagRaw(tagRaw []byte) {
	w.Bytes(tagRaw)
}

func (w *Writer) FixedInt32Field(tagRaw []byte, x int32) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int32(x)
}

func (w *Writer) FixedUInt32Field(tagRaw []byte, x uint32) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.UInt32(x)
}

func (w *Writer) FixedInt64Field(tagRaw []byte, x int64) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int64(x)
}

func (w *Writer) FixedUInt64Field(tagRaw []byte, x uint64) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.UInt64(x)
}

func (w *Writer) IntField(tagRaw []byte, x int) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int(x)
}

func (w *Writer) UIntField(tagRaw []byte, x uint) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.UInt(x)
}

func (w *Writer) Int32Field(tagRaw []byte, x int32) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int32(x)
}

func (w *Writer) SInt32Field(tagRaw []byte, x int32) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.SInt32(x)
}

func (w *Writer) UInt32Field(tagRaw []byte, x uint32) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.UInt32(x)
}

func (w *Writer) Int64Field(tagRaw []byte, x int64) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int64(x)
}

func (w *Writer) SInt64Field(tagRaw []byte, x int64) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.SInt64(x)
}

func (w *Writer) UInt64Field(tagRaw []byte, x uint64) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.UInt64(x)
}

func (w *Writer) Float64Field(tagRaw []byte, x float64) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Float64(x)
}

func (w *Writer) Float32Field(tagRaw []byte, x float32) {
	if x == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Float32(x)
}

func (w *Writer) BoolField(tagRaw []byte, x bool) {
	if !x {
		return
	}

	w.TagRaw(tagRaw)
	w.Int(1)
}

func (w *Writer) StringField(tagRaw []byte, x string) {
	size := len(x)
	if size == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int(size)
	w.String(x)
}

func (w *Writer) MessageField(tagRaw []byte, x MarshallerTo) {
	size := w.s.Shift()
	if size == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int(size)
	w.Message(x)
}

func (w *Writer) BytesField(tagRaw []byte, x []byte) {
	size := len(x)
	if size == 0 {
		return
	}

	w.TagRaw(tagRaw)
	w.Int(size)
	w.Bytes(x)
}
