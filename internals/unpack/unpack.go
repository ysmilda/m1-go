// package unpack provides a way to unpack binary data into a struct and vice versa.
// The struct fields are annotated with a struct tag that specifies how the data should be unpacked.
// Available functionality is catered towards the needs of the project and may not be suitable for all use cases.
package unpack

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	structTag = "unpack:\""
)

type Tag struct {
	// The length of an array, string or time.Duration/time.Time field. May not be used with lengthRef.
	// Used as length=number in the struct tag.
	Length int
	// Alternative to length for dynamic length fields. Refers to another struct field. May not be used with length.
	// Used as lengthRef=otherVariableName in the struct tag.
	LengthRef string
	// The "unit" of a time.Duration/time.Time field. If the field is not present the value defaults to nanoseconds.
	// Used as unit=[milliseconds|seconds|minutes|hours|days] in the struct tag.
	Unit string
	// The amount of bytes to Skip after reading the field.
	// Used as skip=number in the struct tag.
	Skip int
	// Whether to zero terminate a string field.
	// Used as zeroTerminate in the struct tag.
	ZeroTerminate bool
	// Whether to read the slice till the end of the data. This should only be used on the last field.
	// Used as readTillEnd in the struct tag.
	ReadTillEnd bool
}

// Unpack reads the data bytes and unpacks them into the given struct.
func Unpack(data []byte, order binary.ByteOrder, datastruct any) (int, error) {
	return unpack(data, order, datastruct)
}

func Pack(order binary.ByteOrder, dataStruct any) ([]byte, error) {
	return pack(order, dataStruct)
}

func unpack(data []byte, order binary.ByteOrder, datastruct any) (int, error) {
	val := reflect.ValueOf(datastruct).Elem()
	typ := reflect.TypeOf(datastruct).Elem()

	if val.Kind() != reflect.Struct {
		return 0, errors.New("not supplied a struct")
	}

	index := 0

	for i := 0; i < val.NumField(); i++ {
		tag, err := getTag(string(typ.Field(i).Tag))
		if err != nil {
			return 0, err
		}

		n, err := bytesToField(val.Field(i), data[index:], order, tag)
		if err != nil {
			return 0, fmt.Errorf("failed to set field %s: %w", typ.Field(i).Name, err)
		}

		index += n
	}

	return index, nil
}

func pack(order binary.ByteOrder, dataStruct any) ([]byte, error) {
	v := reflect.ValueOf(dataStruct).Elem()
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, but got %s", v.Kind())
	}

	t := v.Type()

	// Determine the size of the byte array by finding the largest offset + length
	result := make([]byte, 0, 2048)

	for i := 0; i < t.NumField(); i++ {
		tag, err := getTag(string(t.Field(i).Tag))
		if err != nil {
			return nil, err
		}

		if tag.LengthRef != "" {
			lengthField, ok := t.FieldByName(tag.LengthRef)
			if !ok {
				return nil, fmt.Errorf("field %s not found", tag.LengthRef)
			}

			tag.Length = int(v.FieldByName(lengthField.Name).Uint())
		}

		fieldBytes, err := fieldToBytes(v.Field(i), order, tag)
		if err != nil {
			return nil, fmt.Errorf("failed to convert field [%s] to bytes: %w", t.Field(i).Name, err)
		}

		result = append(result, fieldBytes...)
	}

	return result, nil
}

func BytesToField(value any, data []byte, byteOrder binary.ByteOrder, length int) (int, error) {
	return bytesToField(reflect.ValueOf(value).Elem(), data, byteOrder, Tag{
		Length: length,
	})
}

func bytesToField( //nolint:gocyclo
	fieldValue reflect.Value, data []byte, byteOrder binary.ByteOrder, t Tag,
) (int, error) {
	// This is for the special cases like time.Duration
	switch fieldValue.Type() {
	case reflect.TypeOf(time.Duration(0)), reflect.TypeOf(time.Time{}):
		if t.Length == -1 {
			return 0, errors.New("missing length tag for time.Duration field")
		}
		dur := 0
		switch t.Length {
		case 1:
			dur = int(data[0])
		case 2:
			dur = int(byteOrder.Uint16(data))
		case 4:
			dur = int(byteOrder.Uint32(data))
		case 8:
			dur = int(byteOrder.Uint64(data))
		default:
			return 0, fmt.Errorf("invalid length for time.Duration field: %d", t.Length)
		}

		u := time.Duration(0)
		switch t.Unit {
		case "":
		case "milliseconds":
			u = time.Millisecond
		case "seconds":
			u = time.Second
		case "minutes":
			u = time.Minute
		case "hours":
			u = time.Hour
		case "days":
			u = time.Hour * 24
		default:
			return 0, fmt.Errorf("invalid unit for time.* field: %s", t.Unit)
		}

		switch fieldValue.Type() {
		case reflect.TypeOf(time.Time{}):
			fieldValue.Set(reflect.ValueOf(time.UnixMicro((time.Duration(dur) * u).Microseconds())))
		case reflect.TypeOf(time.Duration(0)):
			fieldValue.Set(reflect.ValueOf(time.Duration(dur) * u))
		}

		return t.Length + t.Skip, nil
	}

	n := 0

	switch fieldValue.Kind() {
	case reflect.Struct:
		c, err := unpack(data, byteOrder, fieldValue.Addr().Interface())
		if err != nil {
			return 0, fmt.Errorf("failed to unpack struct field: %w", err)
		}
		n = c

	case reflect.Slice:
		if t.ReadTillEnd {
			t.Length = len(data)
		} else if t.Length == -1 {
			return 0, errors.New("missing length tag for slice field")
		}

		for i := 0; i < fieldValue.Len(); i++ {
			c, err := bytesToField(fieldValue.Index(i), data[n:], byteOrder, Tag{})
			if err != nil {
				return 0, fmt.Errorf("failed to set slice field: %w", err)
			}
			n += c
		}

	case reflect.Uint8:
		if len(data) < 1 {
			return 0, fmt.Errorf("invalid data length for uint8")
		}
		val := data[0]
		fieldValue.SetUint(uint64(val))
		n = 1

	case reflect.Uint16:
		if len(data) < 2 {
			return 0, fmt.Errorf("invalid data length for uint16")
		}
		val := byteOrder.Uint16(data)
		fieldValue.SetUint(uint64(val))
		n = 2

	case reflect.Uint32:
		if len(data) < 4 {
			return 0, fmt.Errorf("invalid data length for uint32")
		}
		val := byteOrder.Uint32(data)
		fieldValue.SetUint(uint64(val))
		n = 4

	case reflect.Uint64:
		if len(data) < 8 {
			return 0, fmt.Errorf("invalid data length for uint64")
		}
		val := byteOrder.Uint64(data)
		fieldValue.SetUint(val)
		n = 8

	case reflect.Int8:
		if len(data) < 1 {
			return 0, fmt.Errorf("invalid data length for int8")
		}
		val := int8(data[0])
		fieldValue.SetInt(int64(val))
		n = 1

	case reflect.Int16:
		if len(data) < 2 {
			return 0, fmt.Errorf("invalid data length for int16")
		}
		val := int16(byteOrder.Uint16(data))
		fieldValue.SetInt(int64(val))
		n = 2

	case reflect.Int32:
		if len(data) < 4 {
			return 0, fmt.Errorf("invalid data length for int32")
		}
		val := int32(byteOrder.Uint32(data))
		fieldValue.SetInt(int64(val))
		n = 4

	case reflect.Int64:
		if len(data) < 8 {
			return 0, fmt.Errorf("invalid data length for int64")
		}
		val := int64(byteOrder.Uint64(data))
		fieldValue.SetInt(val)
		n = 8

	case reflect.Bool:
		if len(data) < 1 {
			return 0, fmt.Errorf("invalid data length for bool")
		}
		val := data[0] != 0
		fieldValue.SetBool(val)
		n = 1

	case reflect.String:
		for i, c := range data[:t.Length] {
			if c == 0 {
				fieldValue.SetString(string(data[:i]))
				break
			}
		}
		n = t.Length

	default:
		return 0, fmt.Errorf("unsupported field type: %s", fieldValue.Type())
	}

	return n + t.Skip, nil
}

func FieldToBytes(fieldValue reflect.Value, order binary.ByteOrder) ([]byte, error) {
	tag, err := getTag(string(fieldValue.Type().Field(0).Tag))
	if err != nil {
		return nil, err
	}

	return fieldToBytes(fieldValue, order, tag)
}

func fieldToBytes( //nolint:gocyclo
	fieldValue reflect.Value, order binary.ByteOrder, t Tag,
) ([]byte, error) {
	switch fieldValue.Kind() {
	case reflect.Struct:
		buf, err := pack(order, fieldValue.Addr().Interface())
		if err != nil {
			return nil, err
		}
		return buf, nil

	case reflect.Slice:
		buf := []byte{}
		// The length in the tag dictates the amount of elements in the slice.
		// If the length is -1, the length of the slice is used.
		length := fieldValue.Len()
		padding := 0
		if t.Length != -1 {
			if t.Length < length {
				length = t.Length
			} else if t.Length > length {
				padding = t.Length - length
			}
		}

		for i := range length {
			fieldBytes, err := fieldToBytes(fieldValue.Index(i), order, Tag{})
			if err != nil {
				return nil, fmt.Errorf("failed to convert slice element to bytes: %w", err)
			}
			buf = append(buf, fieldBytes...)
		}

		for range padding {
			fieldBytes, err := fieldToBytes(reflect.Zero(fieldValue.Type().Elem()), order, Tag{})
			if err != nil {
				return nil, fmt.Errorf("failed to convert slice padding to bytes: %w", err)
			}
			buf = append(buf, fieldBytes...)
		}

		return buf, nil

	case reflect.Bool:
		if fieldValue.Bool() {
			return []byte{1}, nil
		} else {
			return []byte{0}, nil
		}

	case reflect.Uint8:
		return []byte{byte(fieldValue.Uint())}, nil

	case reflect.Uint16:
		buf := make([]byte, 2)
		order.PutUint16(buf, uint16(fieldValue.Uint()))
		return buf, nil

	case reflect.Uint32:
		buf := make([]byte, 4)
		order.PutUint32(buf, uint32(fieldValue.Uint()))
		return buf, nil

	case reflect.Uint64:
		buf := make([]byte, 8)
		order.PutUint64(buf, fieldValue.Uint())
		return buf, nil

	case reflect.Int8:
		return []byte{byte(fieldValue.Int())}, nil

	case reflect.Int16:
		buf := make([]byte, 2)
		order.PutUint16(buf, uint16(fieldValue.Int()))
		return buf, nil

	case reflect.Int32:
		buf := make([]byte, 4)
		order.PutUint32(buf, uint32(fieldValue.Int()))
		return buf, nil

	case reflect.Int64:
		buf := make([]byte, 8)
		order.PutUint64(buf, uint64(fieldValue.Int()))
		return buf, nil

	case reflect.String:
		str := fieldValue.String()
		if t.ZeroTerminate {
			str += "\x00"
		}

		if t.Length == -1 {
			return []byte(str), nil
		}

		if len(str) > t.Length {
			str = str[:t.Length] // Trim if the string is too long
		}

		buf := make([]byte, t.Length)
		copy(buf, []byte(str)) // Pad with zeroes if too short
		return buf, nil
	}

	return nil, fmt.Errorf("unsupported field type: %s", fieldValue.Type())
}

func getTag(s string) (Tag, error) {
	start := strings.Index(s, structTag)
	if start == -1 {
		return Tag{}, nil
	}
	start += len(structTag)

	end := strings.Index(s[start:], "\"")
	if end == -1 {
		return Tag{}, errors.New("couldn't find end of unpack tag")
	}

	return parseUnpackTag(s[start : start+end]), nil
}

func parseUnpackTag(s string) Tag {
	t := Tag{
		Unit:          mustFindString(s, "unit="),
		Skip:          mustFindInt(s, "skip="),
		ZeroTerminate: mustFindBool(s, "zeroTerminate"),
		ReadTillEnd:   mustFindBool(s, "readTillEnd"),
	}

	if t.Skip == -1 {
		t.Skip = 0
	}

	length := mustFindString(s, "length=")
	if length == "" {
		t.Length = -1
	} else {
		l, err := strconv.Atoi(length)
		if err != nil {
			t.LengthRef = length
		} else {
			t.Length = l
		}
	}

	return t
}

func mustFindString(s string, tag string) string {
	idx := strings.Index(s, tag)
	if idx == -1 {
		return ""
	}
	idx += len(tag)
	endIdx := idx + 1
	for ; endIdx < len(s); endIdx++ {
		if s[endIdx] == ',' {
			break
		}
	}
	str := s[idx:endIdx]
	return str
}

func mustFindInt(s string, tag string) int {
	idx := strings.Index(s, tag)
	if idx == -1 {
		return -1
	}
	idx += len(tag)
	endIdx := idx + 1
	for ; endIdx < len(s); endIdx++ {
		if s[endIdx] == ',' {
			break
		}
	}
	i, err := strconv.Atoi(s[idx:endIdx])
	if err != nil {
		panic(err)
	}
	return i
}

func mustFindBool(s string, tag string) bool {
	idx := strings.Index(s, tag)
	return idx != -1
}
