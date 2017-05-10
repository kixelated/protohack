package example_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/kixelated/protohack/test/example"
	gold "github.com/kixelated/protohack/test/example/gold"
	gold_gogo "github.com/kixelated/protohack/test/example/gold_gogo"
	"github.com/kixelated/protohack/test/testutil"
)

func sample() (person *example.Person) {
	return &example.Person{
		Name:  "Velstadt, the Royal Aegis",
		Id:    1245,
		Email: "velstadt@undead.crypt",

		Phone: []*example.Person_PhoneNumber{
			&example.Person_PhoneNumber{
				Number: "(123) 456-7890",
				Type:   example.Person_WORK,
			},
			&example.Person_PhoneNumber{
				Number: "(726) 154-0983",
				Type:   example.Person_HOME,
			},
		},
	}
}

func TestExample(t *testing.T) {
	person := sample()

	gold_person := new(gold.Person)
	err := testutil.ConvertProto(gold_person, person)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := proto.Marshal(gold_person)
	if err != nil {
		t.Fatal(err)
	}

	data, err := person.Marshal()
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, expected) {
		t.Error("wrong output")
	}
}

func TestUnmarshal(t *testing.T) {
	person := sample()

	gold_person := new(gold.Person)
	err := testutil.ConvertProto(gold_person, person)
	if err != nil {
		t.Fatal(err)
	}

	data, err := proto.Marshal(gold_person)
	if err != nil {
		t.Fatal(err)
	}

	person2 := new(example.Person)
	err = person2.Unmarshal(data)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(person, person2) {
		t.Error("wrong value")
	}
}

func BenchmarkMarshal(b *testing.B) {
	person := sample()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := person.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalGold(b *testing.B) {
	person := sample()

	gold_person := new(gold.Person)
	err := testutil.ConvertProto(gold_person, person)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := proto.Marshal(gold_person)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalGoldGogo(b *testing.B) {
	person := sample()

	gold_person := new(gold_gogo.Person)
	err := testutil.ConvertProto(gold_person, person)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := gold_person.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	data, err := sample().Marshal()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var person example.Person

		err := person.Unmarshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalGold(b *testing.B) {
	data, err := sample().Marshal()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var person gold.Person

		err := proto.Unmarshal(data, &person)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalGoldGogo(b *testing.B) {
	data, err := sample().Marshal()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var person gold_gogo.Person

		err := person.Unmarshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
