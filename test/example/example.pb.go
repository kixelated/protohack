package example
type Test struct {
	Label string
	Type int32
	Reps []int64
	Optionalgroup *Test_OptionalGroup
	Number int32
	Name string
}
type Test_OptionalGroup struct {
	RequiredField string
}
type FOO int32
const (
	X FOO = 17
)
