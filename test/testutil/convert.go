package testutil

import (
	"errors"
	"reflect"
)

// Convert from one proto/struct to a similar one.
// Loops over each field/index and tries to cast them to the correct value.
// Also handles mismatching pointers.
func ConvertProto(dst interface{}, src interface{}) (err error) {
	return convertProto(reflect.ValueOf(dst), reflect.ValueOf(src))
}

func convertProto(dst reflect.Value, src reflect.Value) (err error) {
	if (dst.Kind() == reflect.Ptr) && (src.Kind() != reflect.Ptr) {
		// Make sure src is a pointer.
		return convertProto(dst, src.Addr())
	}

	if (dst.Kind() != reflect.Ptr) && (src.Kind() == reflect.Ptr) {
		// Make sure dst is a pointer.
		return convertProto(dst.Addr(), src)
	}

	if (dst.Kind() != reflect.Ptr) && (src.Kind() != reflect.Ptr) {
		// Both are values, convert to pointers.
		// I don't think this is possible but it's easy to handle.
		return convertProto(dst.Addr(), src.Addr())
	}

	if !src.Elem().IsValid() {
		// Skip zero values.
		return nil
	}

	if !dst.Elem().CanSet() {
		// Allocate new data if the pointer is nil.
		dst.Set(reflect.New(dst.Type().Elem()))
	}

	if dst.Elem().Kind() == reflect.Ptr || src.Elem().Kind() == reflect.Ptr {
		// Both values are pointers to pointers, alright then.
		return convertProto(dst.Elem(), src.Elem())
	}

	switch src.Elem().Kind() {
	case reflect.Struct:
		for i := 0; i < src.Type().Elem().NumField(); i++ {
			name := src.Type().Elem().Field(i).Name

			// Skip these two fields we don't add.
			if name == "XXX_InternalExtensions" || name == "XXX_unrecognized" {
				continue
			}

			_, ok := dst.Type().Elem().FieldByName(name)
			if !ok {
				err = errors.New("unable to find field: " + name)
				return err
			}

			err = convertProto(dst.Elem().FieldByName(name), src.Elem().Field(i))
			if err != nil {
				return err
			}
		}
	case reflect.Slice:
		dst.Elem().Set(reflect.MakeSlice(dst.Elem().Type(), src.Elem().Len(), src.Elem().Cap()))

		for i := 0; i < src.Elem().Len(); i++ {
			err = convertProto(dst.Elem().Index(i), src.Elem().Index(i))
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		dst.Elem().Set(reflect.MakeMap(dst.Elem().Type().Elem()))

		for _, key := range src.Elem().MapKeys() {
			err = convertProto(dst.Elem().MapIndex(key), src.Elem().MapIndex(key))
			if err != nil {
				return err
			}
		}
	default:
		if !src.Type().Elem().ConvertibleTo(dst.Type().Elem()) {
			err = errors.New("unable to convert: " + src.Type().Elem().Name() + " to " + dst.Type().Elem().Name())
			return err
		}

		dst.Elem().Set(src.Elem().Convert(dst.Elem().Type()))
	}

	return nil
}
