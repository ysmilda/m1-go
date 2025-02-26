package svi

type Variable struct {
	Address Address
	Format  uint16
	Length  uint16
}

func (v Variable) IsBlock() bool {
	return v.Format&Block != 0
}

func (v Variable) IsReadable() bool {
	return v.Format&Out != 0
}

func (v Variable) IsWritable() bool {
	return v.Format&In != 0
}

// TODO: Figure out a way to not need to use pointers here.
// This is needed for the indirection of the interface in m1binary.Decode.
func (v Variable) GetGoDataType() any {
	if v.IsBlock() {
		switch v.Format & ElementaryTypeMask {
		case Uint1, Bool:
			return new([]bool)
		case Sint8:
			return new([]int8)
		case Uint16:
			return new([]uint16)
		case Sint16:
			return new([]int16)
		case Uint32:
			return new([]uint32)
		case Sint32:
			return new([]int32)
		case Real32:
			return new([]float32)
		case Uint64:
			return new([]uint64)
		case Sint64:
			return new([]int64)
		case Real64:
			return new([]float64)
		case Uint8:
			return new([]byte)
		case Char8, Char16:
			return new(string)
		case Mixed:
			return new([]byte)
		default:
			return nil
		}
	}

	switch v.Format & ElementaryTypeMask {
	case Uint1, Bool:
		return new(bool)
	case Sint8:
		return new(int8)
	case Uint16:
		return new(uint16)
	case Sint16:
		return new(int16)
	case Uint32:
		return new(uint32)
	case Sint32:
		return new(int32)
	case Real32:
		return new(float32)
	case Uint64:
		return new(uint64)
	case Sint64:
		return new(int64)
	case Real64:
		return new(float64)
	case Uint8:
		return new(byte)
	case Char8, Char16:
		return new(string)
	case Mixed:
		return new([]byte)
	default:
		return nil
	}
}

func (v Variable) GetDataTypeLength() int {
	switch v.Format & TypeMask {
	case Uint1, Uint8, Sint8, Char8,
		Bool, Mixed, String, StringListBase:
		return 1

	case Uint16, Sint16, Char16, UnicodeStringListBase:
		return 2

	case Uint32, Sint32, Real32:
		return 4

	case Uint64, Sint64, Real64:
		return 8

	default:
		return 1
	}
}

func (v Variable) GetArrayLength() int {
	if !v.IsBlock() {
		return 1
	}
	return int(v.Length) / v.GetDataTypeLength()
}

func (v Variable) GetBufferLength() int {
	return int(v.Length) + 3&_Align
}
