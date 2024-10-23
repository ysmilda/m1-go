package rpc

// String is a wrapper around a string that also contains the length of the string.
// This is used for writing strings to the buffer that need to be padded to a certain length.
type String struct {
	value  string
	length int
}

// NewString creates a new String with the given value and length.
// The length is the number of bytes that should be written to the buffer.
// If the value is longer than the length, it will be truncated.
// If the value is shorter than the length, it will be padded with zeros.
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
// The length is the number of bytes that should be written to the buffer.
func NewSpare(length uint) Spare {
	return Spare{
		length: length,
	}
}
