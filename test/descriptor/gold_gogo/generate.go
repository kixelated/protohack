package descriptor

//go:generate go get github.com/gogo/protobuf/protoc-gen-gogofast
//go:generate protoc -I.. --gogofast_out=. ../descriptor_test.proto
