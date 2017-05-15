package proto

type WireType int

const (
	WIRETYPE_VARINT      WireType = 0
	WIRETYPE_64BIT       WireType = 1
	WIRETYPE_LENGTH      WireType = 2
	WIRETYPE_GROUP_START WireType = 3
	WIRETYPE_GROUP_END   WireType = 4
	WIRETYPE_32BIT       WireType = 5
)

type Marshaller interface {
	Marshal() (data []byte, err error)
}

type Unmarshaller interface {
	Unmarshal(data []byte) (err error)
}

type MarshallerTo interface {
	MarshalTo(w *Writer) (err error)
}

type MarshallerSize interface {
	MarshalSize(s *Sizer)
}
