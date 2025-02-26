package m1binary

import (
	"fmt"
	"reflect"
)

func SizeOf(v any) int {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return sizeOf(rv)
	}

	n := 0
	for i := range rv.NumField() {
		n += sizeOf(rv.Field(i))
	}

	return n
}

func sizeOf(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Bool:
		return 1

	case reflect.Int8:
		return 1

	case reflect.Int16:
		return 2

	case reflect.Int32:
		return 4

	case reflect.Int64:
		return 8

	case reflect.Uint8:
		return 1

	case reflect.Uint16:
		return 2

	case reflect.Uint32:
		return 4

	case reflect.Uint64:
		return 8

	case reflect.Float32:
		return 4

	case reflect.Float64:
		return 8

	case reflect.Array, reflect.Slice:
		return v.Len() * sizeOf(reflect.Zero(v.Type().Elem()))

	case reflect.Interface:
		return sizeOf(v.Elem())

	case reflect.String:
		return v.Len()

	case reflect.Struct:
		return SizeOf(v.Interface())

	default:
		panic(fmt.Errorf("sizeof not implemented for type %v", v.Type()))
	}
}
