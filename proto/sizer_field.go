package proto

func (s *Sizer) TagSize(size int) {
	s.n += size
}

func (s *Sizer) FixedInt32Field(tagSize int, x int32) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.FixedInt32(x)
}

func (s *Sizer) FixedUInt32Field(tagSize int, x uint32) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.FixedUInt32(x)
}

func (s *Sizer) FixedInt64Field(tagSize int, x int64) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.FixedInt64(x)
}

func (s *Sizer) FixedUInt64Field(tagSize int, x uint64) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.FixedUInt64(x)
}

func (s *Sizer) IntField(tagSize int, x int) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Int(x)
}

func (s *Sizer) UIntField(tagSize int, x uint) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.UInt(x)
}

func (s *Sizer) Int32Field(tagSize int, x int32) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Int32(x)
}

func (s *Sizer) SInt32Field(tagSize int, x int32) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.SInt32(x)
}

func (s *Sizer) UInt32Field(tagSize int, x uint32) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.UInt32(x)
}

func (s *Sizer) Int64Field(tagSize int, x int64) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Int64(x)
}

func (s *Sizer) SInt64Field(tagSize int, x int64) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.SInt64(x)
}

func (s *Sizer) UInt64Field(tagSize int, x uint64) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.UInt64(x)
}

func (s *Sizer) Float64Field(tagSize int, x float64) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Float64(x)
}

func (s *Sizer) Float32Field(tagSize int, x float32) {
	if x == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Float32(x)
}

func (s *Sizer) BoolField(tagSize int, x bool) {
	if !x {
		return
	}

	s.TagSize(tagSize)
	s.Int(1)
}

func (s *Sizer) StringField(tagSize int, x string) {
	size := len(x)
	if size == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Int(size)
	s.n += size
}

func (s *Sizer) MessageField(tagSize int, x MarshallerSize) {
	start := s.Size()
	x.MarshalSize(s)

	size := s.Size() - start
	s.Push(size)

	if size > 0 {
		s.TagSize(tagSize)
		s.Int(size)
	}
}

func (s *Sizer) BytesField(tagSize int, x []byte) {
	size := len(x)
	if size == 0 {
		return
	}

	s.TagSize(tagSize)
	s.Int(size)
	s.n += size
}
