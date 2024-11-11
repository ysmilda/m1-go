package m1

import (
	"fmt"
	"reflect"

	"github.com/ysmilda/m1-go/pkg/buffer"
)

type Variable struct {
	Name  string
	Error error

	Address uint64
	Format  uint16 // TODO: We need to parse this to a Go type.
	Length  uint16

	initialized bool
}

// NewVariable creates a new variable with the given module and name.
// By default it contains no error and is not initialized. Initialisation is done via the VHD module on the target.
func NewVariable(module, name string) *Variable {
	return &Variable{
		Name: fmt.Sprintf("%s/%s", module, name),
	}
}

func (s Variable) IsInitialized() bool {
	return s.initialized
}

func (s Variable) IsBlock() bool {
	return s.Format&_FormatBlock != 0
}

func (s Variable) GetArrayLength() int {
	if !s.IsBlock() {
		return 1
	}
	return int(s.Length) / s.getDataTypeLength()
}

func (s Variable) IsReadable() bool {
	return s.Format&_FormatOut != 0
}

func (s Variable) IsWritable() bool {
	return s.Format&_FormatIn != 0
}

func (s Variable) GetGoDataType() any { //nolint:gocyclo
	if s.IsBlock() {
		switch s.Format & _FormatElementaryTypeMask {
		case _FormatUint1, _FormatBool:
			return []bool{}
		case _FormatSint8:
			return []int8{}
		case _FormatUint16:
			return []uint16{}
		case _FormatSint16:
			return []int16{}
		case _FormatUint32:
			return []uint32{}
		case _FormatSint32:
			return []int32{}
		case _FormatReal32:
			return []float32{}
		case _FormatUint64:
			return []uint64{}
		case _FormatSint64:
			return []int64{}
		case _FormatReal64:
			return []float64{}
		case _FormatUint8:
			return []byte{}
		case _FormatChar8, _FormatChar16:
			return string("")
		case _FormatMixed:
			return []byte{}
		default:
			return nil
		}
	}

	switch s.Format & _FormatElementaryTypeMask {
	case _FormatUint1, _FormatBool:
		return bool(false)
	case _FormatSint8:
		return int8(0)
	case _FormatUint16:
		return uint16(0)
	case _FormatSint16:
		return int16(0)
	case _FormatUint32:
		return uint32(0)
	case _FormatSint32:
		return int32(0)
	case _FormatReal32:
		return float32(0)
	case _FormatUint64:
		return uint64(0)
	case _FormatSint64:
		return int64(0)
	case _FormatReal64:
		return float64(0)
	case _FormatUint8:
		return byte(0)
	case _FormatChar8, _FormatChar16:
		return string("")
	case _FormatMixed:
		return []byte{}
	default:
		return nil
	}
}

func (s Variable) String() string {
	t := s.GetGoDataType()

	if s.IsBlock() {
		return fmt.Sprintf("%s: %T (%d)", s.Name, t, s.GetArrayLength())
	}
	return fmt.Sprintf("%s: %T", s.Name, t)
}

func (s Variable) getBufferLength() int {
	return int(s.Length) + 3&_Align
}

func (s *Variable) getDataTypeLength() int {
	switch s.Format & _FormatTypeMask {
	case _FormatUint1, _FormatUint8, _FormatSint8, _FormatChar8,
		_FormatBool, _FormatMixed, _FormatString, _FormatStringListBase:
		return 1

	case _FormatUint16, _FormatSint16, _FormatChar16, _FormatUnicodeStringListBase:
		return 2

	case _FormatUint32, _FormatSint32, _FormatReal32:
		return 4

	case _FormatUint64, _FormatSint64, _FormatReal64:
		return 8

	default:
		return 1
	}
}

func (s *Variable) infoFromBuffer(buf *buffer.Buffer) {
	s.Address, _ = buf.LittleEndian.ReadUint64()
	s.Format, _ = buf.LittleEndian.ReadUint16()
	s.Length, _ = buf.LittleEndian.ReadUint16()
}

// valueToBuffer writes the given value to the buffer. The type and length of the value must match the format of the
// variable. Otherwise an error is returned.
func (s *Variable) valueToBuffer(value any) ([]byte, error) { //nolint:gocyclo
	buf := buffer.NewBuffer(nil)

	// Check if the type of the given value matches the format.
	t := s.GetGoDataType()
	if reflect.TypeOf(value) != reflect.TypeOf(t) {
		return nil, fmt.Errorf("expected %T, got %T", t, value)
	}

	// If it is a slice or string, check if the length matches the length of the variable.
	vt := reflect.TypeOf(value)
	switch vt.Kind() {
	case reflect.Slice:
		if vt.Len() != s.GetArrayLength() {
			return nil, fmt.Errorf("expected %d values, got %d", s.Length, vt.Len())
		}

	case reflect.String:
		if len(value.(string)) != int(s.Length) {
			return nil, fmt.Errorf("expected string with length %d, got %d", s.Length, len(value.(string)))
		}
	}

	// Write the value to the buffer.
	switch val := value.(type) {
	case []bool:
		for _, v := range val {
			_ = buf.WriteBool(v)
		}

	case []int8:
		for _, v := range val {
			_ = buf.WriteByte(byte(v))
		}

	case []uint16:
		for _, v := range val {
			buf.LittleEndian.WriteUint16(v)
		}

	case []int16:
		for _, v := range val {
			buf.LittleEndian.WriteInt16(v)
		}

	case []uint32:
		for _, v := range val {
			buf.LittleEndian.WriteUint32(v)
		}

	case []int32:
		for _, v := range val {
			buf.LittleEndian.WriteInt32(v)
		}

	case []float32:
		for _, v := range val {
			buf.LittleEndian.WriteFloat32(v)
		}

	case []uint64:
		for _, v := range val {
			buf.LittleEndian.WriteUint64(v)
		}

	case []int64:
		for _, v := range val {
			buf.LittleEndian.WriteInt64(v)
		}

	case []float64:
		for _, v := range val {
			buf.LittleEndian.WriteFloat64(v)
		}

	case []byte:
		_, _ = buf.Write(val)

	case string:
		_, _ = buf.WriteString(val)

	case bool:
		_ = buf.WriteBool(val)

	case int8:
		_ = buf.WriteByte(byte(val))

	case uint16:
		buf.LittleEndian.WriteUint16(val)

	case int16:
		buf.LittleEndian.WriteInt16(val)

	case uint32:
		buf.LittleEndian.WriteUint32(val)

	case int32:
		buf.LittleEndian.WriteInt32(val)

	case float32:
		buf.LittleEndian.WriteFloat32(val)

	case uint64:
		buf.LittleEndian.WriteUint64(val)

	case int64:
		buf.LittleEndian.WriteInt64(val)

	case float64:
		buf.LittleEndian.WriteFloat64(val)

	case byte:
		_ = buf.WriteByte(val)
	}

	return buf.Bytes(), nil
}

func (s *Variable) valueFromBuffer(in *buffer.Buffer) any { //nolint:gocyclo
	// Read the value from the buffer. And create an intermediate buffer for reading the value.
	// Depending on the type of the variable, we may use one or the other.
	temp, _ := in.ReadBytes(int(s.Length))
	buffer := buffer.NewBuffer(temp)

	value := s.GetGoDataType()
	switch val := value.(type) {
	case []bool:
		for i, v := range temp {
			val[i] = v == 1
		}
		return val

	case []int8:
		for i, v := range temp {
			val[i] = int8(v)
		}
		return val

	case []uint16:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadUint16()
		}
		return val

	case []int16:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadInt16()
		}
		return val

	case []uint32:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadUint32()
		}
		return val

	case []int32:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadInt32()
		}
		return val

	case []float32:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadFloat32()
		}
		return val

	case []uint64:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadUint64()
		}
		return val

	case []int64:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadInt64()
		}
		return val

	case []float64:
		for i := range s.GetArrayLength() {
			val[i], _ = buffer.LittleEndian.ReadFloat64()
		}
		return val

	case []byte:
		copy(val, temp)
		return val

	case string:
		return string(temp)

	case bool:
		return temp[0] == 1

	case int8:
		return int8(temp[0])

	case uint16:
		val, _ = buffer.LittleEndian.ReadUint16()
		return val

	case int16:
		val, _ = buffer.LittleEndian.ReadInt16()
		return val

	case uint32:
		val, _ = buffer.LittleEndian.ReadUint32()
		return val

	case int32:
		val, _ = buffer.LittleEndian.ReadInt32()
		return val

	case float32:
		val, _ = buffer.LittleEndian.ReadFloat32()
		return val

	case uint64:
		val, _ = buffer.LittleEndian.ReadUint64()
		return val

	case int64:
		val, _ = buffer.LittleEndian.ReadInt64()
		return val

	case float64:
		val, _ = buffer.LittleEndian.ReadFloat64()
		return val

	case byte:
		return temp[0]

	default:
		s.Error = fmt.Errorf("unknown format: %d", s.Format)
	}

	return value
}
