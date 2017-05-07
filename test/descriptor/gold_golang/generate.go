package descriptor

//go:generate go get github.com/golang/protobuf/protoc-gen-go
//go:generate protoc -I.. --go_out=. ../descriptor_test.proto
