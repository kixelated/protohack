package descriptor

//go:generate go install github.com/kixelated/protohack/protoc-gen-gohack
//go:generate protoc --go_out=gold --gohack_out=. descriptor.proto
