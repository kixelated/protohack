package proto

type Sizer struct {
	n     int
	stack []int
}

func (s *Sizer) Grow(size int) {
	if cap(s.stack) > size {
		return
	}

	temp := make([]int, len(s.stack), size)
	copy(temp, s.stack)
	s.stack = temp
}

func (s *Sizer) Size() int {
	return s.n
}

func (s *Sizer) Push(n int) {
	s.stack = append(s.stack, n)
}

func (s *Sizer) Shift() (n int) {
	n = s.stack[0]
	s.stack = s.stack[1:]
	return n
}
