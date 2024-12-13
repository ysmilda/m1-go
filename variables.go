package m1

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/ysmilda/m1-go/internals/unpack"
	"github.com/ysmilda/m1-go/modules/svi"
)

const (
	_Align = 0xFFFFFFFC
)

type Variable struct {
	Name  string
	Error error

	svi.Variable

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
	return s.Format&svi.Block != 0
}

func (s Variable) GetArrayLength() int {
	if !s.IsBlock() {
		return 1
	}
	return int(s.Length) / s.getDataTypeLength()
}

func (s Variable) IsReadable() bool {
	return s.Format&svi.Out != 0
}

func (s Variable) IsWritable() bool {
	return s.Format&svi.In != 0
}

func (s Variable) GetGoDataType() any { //nolint:gocyclo
	if s.IsBlock() {
		switch s.Format & svi.ElementaryTypeMask {
		case svi.Uint1, svi.Bool:
			return []bool{}
		case svi.Sint8:
			return []int8{}
		case svi.Uint16:
			return []uint16{}
		case svi.Sint16:
			return []int16{}
		case svi.Uint32:
			return []uint32{}
		case svi.Sint32:
			return []int32{}
		case svi.Real32:
			return []float32{}
		case svi.Uint64:
			return []uint64{}
		case svi.Sint64:
			return []int64{}
		case svi.Real64:
			return []float64{}
		case svi.Uint8:
			return []byte{}
		case svi.Char8, svi.Char16:
			return string("")
		case svi.Mixed:
			return []byte{}
		default:
			return nil
		}
	}

	switch s.Format & svi.ElementaryTypeMask {
	case svi.Uint1, svi.Bool:
		return bool(false)
	case svi.Sint8:
		return int8(0)
	case svi.Uint16:
		return uint16(0)
	case svi.Sint16:
		return int16(0)
	case svi.Uint32:
		return uint32(0)
	case svi.Sint32:
		return int32(0)
	case svi.Real32:
		return float32(0)
	case svi.Uint64:
		return uint64(0)
	case svi.Sint64:
		return int64(0)
	case svi.Real64:
		return float64(0)
	case svi.Uint8:
		return byte(0)
	case svi.Char8, svi.Char16:
		return string("")
	case svi.Mixed:
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
	switch s.Format & svi.TypeMask {
	case svi.Uint1, svi.Uint8, svi.Sint8, svi.Char8,
		svi.Bool, svi.Mixed, svi.String, svi.StringListBase:
		return 1

	case svi.Uint16, svi.Sint16, svi.Char16, svi.UnicodeStringListBase:
		return 2

	case svi.Uint32, svi.Sint32, svi.Real32:
		return 4

	case svi.Uint64, svi.Sint64, svi.Real64:
		return 8

	default:
		return 1
	}
}

// valueToBuffer writes the given value to the buffer. The type and length of the value must match the format of the
// variable. Otherwise an error is returned.
func (s *Variable) valueToBuffer(value any) ([]byte, error) { //nolint:gocyclo
	// Check if the type of the given value matches the format.
	t := s.GetGoDataType()
	if reflect.TypeOf(value) != reflect.TypeOf(t) {
		return nil, fmt.Errorf("expected %T, got %T", t, value)
	}

	return unpack.FieldToBytes(reflect.ValueOf(value), binary.LittleEndian)
}

func (s *Variable) valueFromBuffer(in []byte) (int, any) { //nolint:gocyclo
	value := s.GetGoDataType()
	return unpack.BytesToField(&value, in[:s.Length], binary.LittleEndian, s.GetArrayLength())
}
