package descriptor_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/kixelated/protohack/test/descriptor"
	"github.com/kixelated/protohack/test/descriptor/bindata"
	gold_gogo "github.com/kixelated/protohack/test/descriptor/gold_gogo"
	gold_golang "github.com/kixelated/protohack/test/descriptor/gold_golang"
	"github.com/kixelated/protohack/test/testutil"
)

func TestMarshal(t *testing.T) {
	data_gold, err := bindata.Asset("descriptor_test.pb")
	if err != nil {
		t.Fatal(err)
	}

	desc_gold := new(gold_golang.FileDescriptorProto)
	err = proto.Unmarshal(data_gold, desc_gold)
	if err != nil {
		t.Fatal(err)
	}

	desc := new(descriptor.FileDescriptorProto)

	err = testutil.ConvertProto(desc, desc_gold)
	if err != nil {
		t.Fatal(err)
	}

	data, err := desc.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, data_gold) {
		t.Errorf("wrong output")

		fmt.Printf("output: %s\n", hex.EncodeToString(data))
		fmt.Printf("expected: %s\n", hex.EncodeToString(data_gold))
	}
}

func BenchmarkMarshal(b *testing.B) {
	data, err := bindata.Asset("descriptor_test.pb")
	if err != nil {
		b.Fatal(err)
	}

	desc_gold := new(gold_golang.FileDescriptorProto)
	err = proto.Unmarshal(data, desc_gold)
	if err != nil {
		b.Fatal(err)
	}

	desc := new(descriptor.FileDescriptorProto)
	err = testutil.ConvertProto(desc, desc_gold)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := desc.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalGoldGolang(b *testing.B) {
	data, err := bindata.Asset("descriptor_test.pb")
	if err != nil {
		b.Fatal(err)
	}

	desc_gold := new(gold_golang.FileDescriptorProto)
	err = proto.Unmarshal(data, desc_gold)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := proto.Marshal(desc_gold)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalGoldGogo(b *testing.B) {
	data, err := bindata.Asset("descriptor_test.pb")
	if err != nil {
		b.Fatal(err)
	}

	desc_gold := new(gold_gogo.FileDescriptorProto)
	err = proto.Unmarshal(data, desc_gold)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := proto.Marshal(desc_gold)
		if err != nil {
			b.Fatal(err)
		}
	}
}
