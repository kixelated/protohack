package proto

func (s *Sizer) Int(x int) {
	s.UInt(uint(x))
}

func (s *Sizer) UInt(x uint) {
	for x > 0x7f {
		s.n += 1
		x >>= 7
	}

	s.n += 1
}

func (s *Sizer) Int32(x int32) {
	s.UInt32(uint32(x))
}

func (s *Sizer) SInt32(x int32) {
	// TODO optimize
	s.UInt32(uint32(x<<1) ^ uint32(x>>31))
}

func (s *Sizer) UInt32(x uint32) {
	for x > 0x7f {
		s.n += 1
		x >>= 7
	}

	s.n += 1
}

func (s *Sizer) Int64(x int64) {
	s.UInt64(uint64(x))
}

func (s *Sizer) SInt64(x int64) {
	// TODO optimize
	s.UInt64(uint64(x<<1) ^ uint64(x>>63))
}

func (s *Sizer) UInt64(x uint64) {
	for x > 0x7f {
		s.n += 1
		x >>= 7
	}

	s.n += 1
}

func (s *Sizer) Bool(x bool) {
	s.n += 1
}

func (s *Sizer) Tag(id int) {
	s.UInt64(uint64(id << 3))
}

func (s *Sizer) Float32(x float32) {
	s.n += 4
}

func (s *Sizer) Float64(x float64) {
	s.n += 8
}

func (s *Sizer) FixedInt32(x int32) {
	s.n += 4
}

func (s *Sizer) FixedUInt32(x uint32) {
	s.n += 4
}

func (s *Sizer) FixedInt64(x int64) {
	s.n += 8
}

func (s *Sizer) FixedUInt64(x uint64) {
	s.n += 8
}

func (s *Sizer) String(x string) {
	s.n += len(x)
}

func (s *Sizer) Message(x MarshallerSize) {
	x.MarshalSize(s)
}

func (s *Sizer) Bytes(x []byte) {
	s.n += len(x)
}
