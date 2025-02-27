// package m1binary provides a way to m1binary binary data from an m1 source into a struct and vice versa.
// The struct fields are annotated with a struct tag that specifies how the data should be unpacked.
// Available functionality is catered towards the needs of the project and may not be suitable for all use cases.
package m1binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/ysmilda/m1-go/internals/ptr"
)

var ErrInvalidDataLength = errors.New("invalid data length")

var (
	timeType     = reflect.TypeOf(time.Time{})
	durationType = reflect.TypeOf(time.Duration(0))
)

type Decoder interface {
	DecodeM1([]byte) (int, error)
}

// Decode reads the data bytes and unpacks them into the given struct.
func Decode(data []byte, v any) (int, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		panic("m1binary may only be called on pointer to value")
	}

	return decode(data, rv)
}

func decode(data []byte, v reflect.Value) (int, error) {
	d, pv := indirect(v)
	if d != nil {
		return d.DecodeM1(data)
	}
	v = pv

	// If it is not a struct parse the field directly.
	if v.Kind() != reflect.Struct {
		return decodeField(v, data, Tag{
			Length: ptr.For(len(data)),
		})
	}

	// Loop through all the structs fields and parse them.
	index := 0
	rt := v.Type()
	for i := range rt.NumField() {
		tag := parseTag(rt.Field(i).Tag.Get(structTag))

		// If the length is given by reference, fetch the actual
		if tag.LengthRef != nil {
			ref := v.FieldByName(*tag.LengthRef)
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
			}
		}

		n, err := decodeField(v.Field(i), data[index:], tag)
		if err != nil {
			return 0, fmt.Errorf("failed to set field %s: %w", rt.Field(i).Name, err)
		}

		index += n
		if tag.Skip != nil {
			index += *tag.Skip
		}
	}

	return index, nil
}

func decodeField(v reflect.Value, data []byte, tag Tag) (int, error) { //nolint: gocyclo
	d, pv := indirect(v)
	if d != nil {
		return d.DecodeM1(data)
	}
	v = pv

	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			if !v.CanSet() {
				panic(fmt.Errorf("cannot set embedded pointer to unexported struct: %v", v.Type().Elem()))
			}
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}

	// This is for the special cases like time.Duration and time.Time
	switch v.Type() {
	case timeType, durationType:
		if tag.Length == nil {
			panic("missing length tag for time.Duration field")
		} else if *tag.Length > len(data) {
			return 0, fmt.Errorf("%w for time", ErrInvalidDataLength)
		}

		duration := time.Duration(0)
		switch *tag.Length {
		case 1:
			duration = time.Duration(int(data[0]))
		case 2:
			duration = time.Duration(int(binary.LittleEndian.Uint16(data)))
		case 4:
			duration = time.Duration(int(binary.LittleEndian.Uint32(data)))
		case 8:
			duration = time.Duration(int(binary.LittleEndian.Uint64(data)))
		default:
			panic(fmt.Errorf("invalid length for time.Duration field: %d", *tag.Length))
		}

		if tag.TimeUnit != nil {
			switch *tag.TimeUnit {
			case "milliseconds":
				duration *= time.Millisecond
			case "seconds":
				duration *= time.Second
			case "minutes":
				duration *= time.Minute
			case "hours":
				duration *= time.Hour
			case "days":
				duration *= time.Hour * 24
			default:
				panic(fmt.Errorf("invalid unit for time.* field: %s", *tag.TimeUnit))
			}
		} else {
			duration *= time.Millisecond
		}

		switch v.Type() {
		case timeType:
			v.Set(reflect.ValueOf(
				time.UnixMicro(duration.Microseconds()),
			))

		case durationType:
			v.Set(reflect.ValueOf(
				duration,
			))
		}

		return *tag.Length, nil
	}

	switch v.Kind() {
	case reflect.Struct:
		return decode(data, v)

	case reflect.Interface:
		return decodeField(v, data, tag)

	case reflect.Slice:
		if tag.TillEnd {
			// When reading till the end we need to divide the remaining data by the size of one element.
			tag.Length = ptr.For(len(data) / sizeOf(reflect.Zero(v.Type().Elem())))
		} else if tag.Length == nil {
			panic("missing length tag for slice field")
		} else if *tag.Length > len(data) {
			return 0, fmt.Errorf("%w for slice", ErrInvalidDataLength)
		}

		length := *tag.Length
		n := 0

		v.Set(reflect.MakeSlice(v.Type(), length, length))
		tag.Length = tag.ElementLength
		for i := range length {
			c, err := decodeField(v.Index(i), data[n:], tag)
			if err != nil {
				return 0, fmt.Errorf("failed to set slice field: %w", err)
			}

			n += c

			if tag.Allign4 {
				n = (n + 3) & 0xfffffffc
			}
		}

		return n, nil

	case reflect.Uint8:
		if len(data) < 1 {
			return 0, fmt.Errorf("%w for uint8", ErrInvalidDataLength)
		}
		v.SetUint(uint64(data[0]))
		return 1, nil

	case reflect.Uint16:
		if len(data) < 2 {
			return 0, fmt.Errorf("%w for uint16", ErrInvalidDataLength)
		}
		v.SetUint(uint64(binary.LittleEndian.Uint16(data)))
		return 2, nil

	case reflect.Uint32:
		if len(data) < 4 {
			return 0, fmt.Errorf("%w for uint32", ErrInvalidDataLength)
		}
		v.SetUint(uint64(binary.LittleEndian.Uint32(data)))
		return 4, nil

	case reflect.Uint64:
		if len(data) < 8 {
			return 0, fmt.Errorf("%w for uint64", ErrInvalidDataLength)
		}
		v.SetUint(binary.LittleEndian.Uint64(data))
		return 8, nil

	case reflect.Int8:
		if len(data) < 1 {
			return 0, fmt.Errorf("%w for int8", ErrInvalidDataLength)
		}
		v.SetInt(int64(data[0]))
		return 1, nil

	case reflect.Int16:
		if len(data) < 2 {
			return 0, fmt.Errorf("%w for int16", ErrInvalidDataLength)
		}
		v.SetInt(int64(int16(binary.LittleEndian.Uint16(data))))
		return 2, nil

	case reflect.Int32:
		if len(data) < 4 {
			return 0, fmt.Errorf("%w for int32", ErrInvalidDataLength)
		}
		v.SetInt(int64(int32(binary.LittleEndian.Uint32(data))))
		return 4, nil

	case reflect.Int64:
		if len(data) < 8 {
			return 0, fmt.Errorf("%w for int64", ErrInvalidDataLength)
		}
		v.SetInt(int64(binary.LittleEndian.Uint64(data)))
		return 8, nil

	case reflect.Bool:
		if len(data) < 1 {
			return 0, fmt.Errorf("%w for bool", ErrInvalidDataLength)
		}
		v.SetBool(data[0] != 0)
		return 1, nil

	case reflect.Float32:
		if len(data) < 4 {
			return 0, fmt.Errorf("%w for float32", ErrInvalidDataLength)
		}
		v.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(data))))
		return 4, nil

	case reflect.Float64:
		if len(data) < 8 {
			return 0, fmt.Errorf("%w for float64", ErrInvalidDataLength)
		}
		v.SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(data)))
		return 8, nil

	case reflect.String:
		if tag.Length == nil {
			if tag.ZeroTerminated {
				tag.Length = ptr.For(len(data))
			} else {
				panic("missing length tag for string field")
			}
		}

		for i, c := range data[:*tag.Length] {
			if c == 0 {
				v.SetString(string(data[:i]))

				if tag.ZeroTerminated {
					return i + 1, nil
				} else {
					return *tag.Length, nil
				}
			}
		}

		v.SetString(string(data[:*tag.Length]))
		return *tag.Length, nil

	default:
		panic(fmt.Errorf("unsupported field type: %s", v.Type()))
	}
}

// indirect walks down v allocating pointers as needed,
// until it gets to a non-pointer.
// If it encounters an Unmarshaler, indirect stops and returns that.
// Adapted from encoding/json.indirect.
func indirect(v reflect.Value) (Decoder, reflect.Value) {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	if v.Kind() != reflect.Pointer && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Pointer && !e.IsNil() {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Pointer {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.Type().NumMethod() > 0 && v.CanInterface() {
			if u, ok := v.Interface().(Decoder); ok {
				return u, reflect.Value{}
			}
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return nil, v
}
