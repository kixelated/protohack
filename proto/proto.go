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

type MarshallerTo interface {
	MarshalTo(data []byte) (n int, err error)
	MarshalSize() (n int)
}

type Unmarshaller interface {
	Unmarshal(data []byte) (err error)
}
