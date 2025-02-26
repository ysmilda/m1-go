package m1binary

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"

	"github.com/ysmilda/m1-go/internals/ptr"
)

func Encode(v any) ([]byte, error) {
	return encode(v)
}

func encode(v any) ([]byte, error) {
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return encodeField(rv, Tag{})
	}

	result := make([]byte, 0, 2048)

	for i := range rv.NumField() {
		tag := parseTag(rt.Field(i).Tag.Get(structTag))

		if tag.LengthRef != nil {
			ref := rv.FieldByName(*tag.LengthRef)
			if !ref.IsValid() {
				panic("given lengthRef doesn't exist")
			}

			switch ref.Kind() {
			default:
				panic("lengthRef references non integer field")

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				tag.Length = ptr.For(int(ref.Uint()))

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				tag.Length = ptr.For(int(ref.Int()))
				if *tag.Length < 0 {
					return nil, fmt.Errorf("value pointed to by lengthRef (%s) is negative", *tag.LengthRef)
				}
			}
		}

		data, err := encodeField(rv.Field(i), tag)
		if err != nil {
			return nil, fmt.Errorf("failed to convert field [%s] to bytes: %w", rt.Field(i).Name, err)
		}

		result = append(result, data...)

		if tag.Skip != nil {
			result = append(result, make([]byte, *tag.Skip)...)
		}
	}

	return result, nil
}

func encodeField(v reflect.Value, tag Tag) ([]byte, error) {
	switch v.Kind() {
	case reflect.Struct:
		return encode(v.Interface())

	case reflect.Slice:
		buf := []byte{}
		// The length in the tag dictates the amount of elements in the slice.
		length := v.Len()
		padding := 0
		if tag.Length != nil {
			if *tag.Length < length {
				length = *tag.Length
			} else if *tag.Length > length {
				padding = *tag.Length - length
			}
		}

		tag.Length = tag.ElementLength
		for i := range length {
			data, err := encodeField(v.Index(i), tag)
			if err != nil {
				return nil, fmt.Errorf("failed to convert slice element to bytes: %w", err)
			}
			buf = append(buf, data...)
		}

		for range padding {
			fieldBytes, err := encodeField(reflect.Zero(v.Type().Elem()), Tag{})
			if err != nil {
				return nil, fmt.Errorf("failed to convert slice padding to bytes: %w", err)
			}
			buf = append(buf, fieldBytes...)
		}

		return buf, nil

	case reflect.Bool:
		if v.Bool() {
			return []byte{1}, nil
		} else {
			return []byte{0}, nil
		}

	case reflect.Uint8:
		return []byte{byte(v.Uint())}, nil

	case reflect.Uint16:
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(v.Uint()))
		return buf, nil

	case reflect.Uint32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(v.Uint()))
		return buf, nil

	case reflect.Uint64:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, v.Uint())
		return buf, nil

	case reflect.Int8:
		return []byte{byte(v.Int())}, nil

	case reflect.Int16:
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(v.Int()))
		return buf, nil

	case reflect.Int32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(v.Int()))
		return buf, nil

	case reflect.Int64:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, uint64(v.Int()))
		return buf, nil

	case reflect.Float32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(v.Float())))
		return buf, nil

	case reflect.Float64:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, math.Float64bits(v.Float()))
		return buf, nil

	case reflect.String:
		str := v.String()

		var buf []byte
		if tag.Length == nil {
			buf = []byte(str)
		} else if len(str) > *tag.Length {
			buf = []byte(str[:*tag.Length])
		} else {
			buf = make([]byte, *tag.Length)
			copy(buf, []byte(str)) // Pad with zeroes if too short
		}

		if tag.ZeroTerminated {
			if tag.Length == nil {
				buf = append(buf, 0) //nolint:makezero
			} else if *tag.Length <= len(str) {
				buf[len(buf)-1] = 0
			} else {
				buf = buf[:len(str)+1]
			}
		}
		return buf, nil

	case reflect.Interface:
		return encodeField(v.Elem(), tag)
	}

	return nil, fmt.Errorf("unsupported field type: %s", v.Type())
}
