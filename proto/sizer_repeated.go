package proto

func (s *Sizer) FixedInt32Packed(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x) * 4
	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedUInt32Packed(tag int, x []uint32) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x) * 4
	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedInt64Packed(tag int, x []int64) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x) * 8
	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedUInt64Packed(tag int, x []uint64) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x) * 8
	s.Int(size)
	s.n += size
}

func (s *Sizer) IntPacked(tag int, x []int) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.Int(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) UIntPacked(tag int, x []uint) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.UInt(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) Int32Packed(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.Int32(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) Int32Repeated(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		s.Tag(tag)
		s.Int32(y)
	}
}

func (s *Sizer) SInt32Packed(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.SInt32(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) UInt32Packed(tag int, x []uint32) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.UInt32(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) Int64Packed(tag int, x []int64) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.Int64(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) SInt64Packed(tag int, x []int64) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.SInt64(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) UInt64Packed(tag int, x []uint64) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	start := s.Size()
	for _, y := range x {
		s.UInt64(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) Float32Packed(tag int, x []float32) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x) * 4

	s.Int(size)
	s.n += size
}

func (s *Sizer) Float64Packed(tag int, x []float64) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x) * 8

	s.Int(size)
	s.n += size
}

func (s *Sizer) BoolPacked(tag int, x []bool) {
	if len(x) == 0 {
		return
	}

	s.Tag(tag)

	size := len(x)

	s.Int(size)
	s.n += size
}

func (s *Sizer) StringRepeated(tag int, x []string) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		// Don't perform a nil check.
		s.Tag(tag)

		size := len(y)

		s.Int(size)
		s.n += size
	}
}

func (s *Sizer) MessageRepeated(tag int, x []MarshallerSize) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		// Don't perform a nil check.
		start := s.Size()
		y.MarshalSize(s)

		size := s.Size() - start
		s.Push(size)

		s.Tag(tag)
		s.Int(size)
	}
}

func (s *Sizer) BytesRepeated(tag int, x [][]byte) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		// Don't perform a nil check.
		s.Tag(tag)

		size := len(y)

		s.Int(size)
		s.n += size
	}
}
