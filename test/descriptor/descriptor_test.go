package descriptor_test

import (
	"bytes"
	"compress/gzip"
	//	"encoding/hex"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/kixelated/protohack/test/descriptor"
	descriptor_gold "github.com/kixelated/protohack/test/descriptor/gold"
)

func getData() (data []byte, err error) {
	data_gz := proto.FileDescriptor("descriptor.proto")

	reader, err := gzip.NewReader(bytes.NewReader(data_gz))
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	data, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestMarshal(t *testing.T) {
	data_gold, err := getData()
	if err != nil {
		t.Fatal(err)
	}

	desc_gold := new(descriptor_gold.FileDescriptorProto)
	err = proto.Unmarshal(data_gold, desc_gold)
	if err != nil {
		t.Fatal(err)
	}

	desc := new(descriptor.FileDescriptorProto)

	portProto(reflect.ValueOf(desc_gold), reflect.ValueOf(desc))

	data, err := desc.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, data_gold) {
		//t.Errorf("wrong output\noutput: %s\ngold: %s", hex.EncodeToString(data), hex.EncodeToString(data_gold))
	}
}

func portProto(v1 reflect.Value, v2 reflect.Value) {
	if !v1.IsValid() {
		return
	}

	switch v1.Kind() {
	case reflect.Ptr:
		if v2.Kind() == reflect.Ptr {
			if !v1.Elem().IsValid() {
				return
			}

			v2v := reflect.New(v2.Type().Elem())

			// I don't fully understand why this is required.
			if v2.CanSet() {
				v2.Set(v2v)
			} else {
				v2.Elem().Set(v2v.Elem())
			}

			portProto(v1.Elem(), v2.Elem())
		} else {
			portProto(v1.Elem(), v2)
		}
	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			name := v1.Type().Field(i).Name

			// Skip these two fields we don't add.
			if name == "XXX_InternalExtensions" || name == "XXX_unrecognized" {
				continue
			}

			portProto(v1.Field(i), v2.FieldByName(name))
		}
	case reflect.Slice:
		if v2.CanSet() {
			v2.Set(reflect.MakeSlice(v2.Type(), v1.Len(), v1.Cap()))
		} else {
			fmt.Println(v2)
			v2.Elem().Set(reflect.MakeSlice(v2.Elem().Type(), v1.Len(), v1.Cap()))
		}

		for i := 0; i < v1.Len(); i++ {
			portProto(v1.Index(i), v2.Index(i))
		}
	case reflect.Map:
		v2.Set(reflect.MakeMap(v2.Type()))

		for _, key := range v1.MapKeys() {
			v1Elem := v1.MapIndex(key)
			v2Elem := reflect.New(v1Elem.Type()).Elem()
			portProto(v1Elem, v2Elem)
			v2.SetMapIndex(key, v2Elem)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v2.SetInt(v1.Int())
	default:
		v2.Set(v1)
	}
}

func BenchmarkMarshal(b *testing.B) {
	data, err := getData()
	if err != nil {
		b.Fatal(err)
	}

	desc_gold := new(descriptor_gold.FileDescriptorProto)
	err = proto.Unmarshal(data, desc_gold)
	if err != nil {
		b.Fatal(err)
	}

	desc := new(descriptor.FileDescriptorProto)
	portProto(reflect.ValueOf(desc_gold), reflect.ValueOf(desc))

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := desc.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalGold(b *testing.B) {
	data, err := getData()
	if err != nil {
		b.Fatal(err)
	}

	desc_gold := new(descriptor_gold.FileDescriptorProto)
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
