package rpc

// String is a wrapper around a string that also contains the length of the string.
// This is used for writing strings to the buffer that need to be padded to a certain length.
type String struct {
	value  string
	length int
}

// NewString creates a new String with the given value and length.
func NewString(value string, length int) String {
	return String{
		value:  value,
		length: length,
	}
}

func (s String) String() string {
	return s.value
}

// Spare is a struct that is used to write padding to the buffer.
type Spare struct {
	length uint
}

// NewSpare creates a new Spare with the given length.
func NewSpare(length uint) Spare {
	return Spare{
		length: length,
	}
}
