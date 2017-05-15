package proto

type Writer struct {
	s *Sizer

	data []byte
	n    int
}

func (w *Writer) WithSizer(s *Sizer) {
	w.data = make([]byte, s.Size())
	w.s = s
}

func (w *Writer) Data() []byte {
	return w.data[:w.n]
}
