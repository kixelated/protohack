package proto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kixelated/protohack/proto"
)

func TestUInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    uint32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    270,
			Expected: []byte{0x8e, 0x2},
		},
		{
			Input:    86942,
			Expected: []byte{0x9e, 0xa7, 0x05},
		},
	}

	for i, test := range tests {
		var s proto.Sizer
		s.UInt32(test.Input)

		assert.Equal(len(test.Expected), s.Size(), "%d: wrong size", i)

		var w proto.Writer
		w.WithSizer(&s)
		w.UInt32(test.Input)

		output := w.Data()
		assert.Equal(test.Expected, output, "%d: wrong output", i)
	}
}

func TestInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xf},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff, 0xf},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    -127,
			Expected: []byte{0x81, 0xff, 0xff, 0xff, 0x0f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -128,
			Expected: []byte{0x80, 0xff, 0xff, 0xff, 0x0f},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    -129,
			Expected: []byte{0xff, 0xfe, 0xff, 0xff, 0x0f},
		},
	}

	for i, test := range tests {
		var s proto.Sizer
		s.Int32(test.Input)

		assert.Equal(len(test.Expected), s.Size(), "%d: wrong size", i)

		var w proto.Writer
		w.WithSizer(&s)
		w.Int32(test.Input)

		output := w.Data()
		assert.Equal(test.Expected, output, "%d: wrong output", i)
	}
}

func TestSInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    -1,
			Expected: []byte{1},
		},
		{
			Input:    1,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{3},
		},
		{
			Input:    2,
			Expected: []byte{4},
		},
		{
			Input:    63,
			Expected: []byte{0x7e},
		},
		{
			Input:    -64,
			Expected: []byte{0x7f},
		},
		{
			Input:    64,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -65,
			Expected: []byte{0x81, 0x1},
		},
	}

	for i, test := range tests {
		var s proto.Sizer
		s.SInt32(test.Input)

		assert.Equal(len(test.Expected), s.Size(), "%d: wrong size", i)

		var w proto.Writer
		w.WithSizer(&s)
		w.SInt32(test.Input)

		output := w.Data()
		assert.Equal(test.Expected, output, "%d: wrong output", i)
	}
}

func TestUInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    uint64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    270,
			Expected: []byte{0x8e, 0x2},
		},
		{
			Input:    86942,
			Expected: []byte{0x9e, 0xa7, 0x05},
		},
	}

	for i, test := range tests {
		var s proto.Sizer
		s.UInt64(test.Input)

		assert.Equal(len(test.Expected), s.Size(), "%d: wrong size", i)

		var w proto.Writer
		w.WithSizer(&s)
		w.UInt64(test.Input)

		output := w.Data()
		assert.Equal(test.Expected, output, "%d: wrong output", i)
	}
}

func TestInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    -127,
			Expected: []byte{0x81, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -128,
			Expected: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    -129,
			Expected: []byte{0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
	}

	for i, test := range tests {
		var s proto.Sizer
		s.Int64(test.Input)

		assert.Equal(len(test.Expected), s.Size(), "%d: wrong size", i)

		var w proto.Writer
		w.WithSizer(&s)
		w.Int64(test.Input)

		output := w.Data()
		assert.Equal(test.Expected, output, "%d: wrong output", i)
	}
}

func TestSInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    -1,
			Expected: []byte{1},
		},
		{
			Input:    1,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{3},
		},
		{
			Input:    2,
			Expected: []byte{4},
		},
		{
			Input:    63,
			Expected: []byte{0x7e},
		},
		{
			Input:    -64,
			Expected: []byte{0x7f},
		},
		{
			Input:    64,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -65,
			Expected: []byte{0x81, 0x1},
		},
	}

	for i, test := range tests {
		var s proto.Sizer
		s.SInt64(test.Input)

		assert.Equal(len(test.Expected), s.Size(), "%d: wrong size", i)

		var w proto.Writer
		w.WithSizer(&s)
		w.SInt64(test.Input)

		output := w.Data()
		assert.Equal(test.Expected, output, "%d: wrong output", i)
	}
}
