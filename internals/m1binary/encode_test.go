package m1binary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ysmilda/m1-go/internals/m1binary"
)

func TestBaseTypesEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected []byte
	}{
		{
			value:    bool(false),
			expected: []byte{0x00},
		},
		{
			value:    bool(true),
			expected: []byte{0x01},
		},
		{
			value:    uint8(123),
			expected: []byte{0x7B},
		},
		{
			value:    int8(-123),
			expected: []byte{0x85},
		},
		{
			value:    uint16(12345),
			expected: []byte{0x39, 0x30},
		},
		{
			value:    int16(-12345),
			expected: []byte{0xC7, 0xCF},
		},
		{
			value:    uint32(123456789),
			expected: []byte{0x15, 0xCD, 0x5B, 0x07},
		},
		{
			value:    int32(-123456789),
			expected: []byte{0xEB, 0x32, 0xA4, 0xF8},
		},
		{
			value:    uint64(123456789012345),
			expected: []byte{0x79, 0xDF, 0x0D, 0x86, 0x48, 0x70, 0x00, 0x00},
		},
		{
			value:    int64(-123456789012345),
			expected: []byte{0x87, 0x20, 0xF2, 0x79, 0xB7, 0x8F, 0xFF, 0xFF},
		},
		{
			value:    float32(123456789012345),
			expected: []byte{0x0C, 0x91, 0xE0, 0x56},
		},
		{
			value:    float64(-123456789012345),
			expected: []byte{0x40, 0xDE, 0x77, 0x83, 0x21, 0x12, 0xDC, 0xC2},
		},
	}

	for _, tc := range testCases {
		result, err := m1binary.Encode(&tc.value)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, result)
	}
}

func TestStringEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected []byte
	}{
		{
			name:     "base string",
			value:    "test",
			expected: []byte{'t', 'e', 's', 't'},
		},
		{
			name: "zero terminated",
			value: struct {
				val string `m1binary:"zeroTerminated"`
			}{"test"},
			expected: []byte{'t', 'e', 's', 't', 0},
		},
		{
			name: "tagged length shorter than actual length",
			value: struct {
				val string `m1binary:"length:3"`
			}{"test"},
			expected: []byte{'t', 'e', 's'},
		},
		{
			name: "tagged length shorter than actual length zero terminated",
			value: struct {
				val string `m1binary:"length:3,zeroTerminated"`
			}{"test"},
			expected: []byte{'t', 'e', 0},
		},
		{
			name: "tagged length longer than actual length (padding)",
			value: struct {
				val string `m1binary:"length:6"`
			}{"test"},
			expected: []byte{'t', 'e', 's', 't', 0, 0},
		},
		{
			name: "tagged length longer than actual length zero terminated",
			value: struct {
				val string `m1binary:"length:6,zeroTerminated"`
			}{"test"},
			expected: []byte{'t', 'e', 's', 't', 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := m1binary.Encode(&tc.value)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSliceEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected []byte
	}{
		{
			name:     "basic slice",
			value:    []uint16{12345, 54321},
			expected: []byte{0x39, 0x30, 0x31, 0xD4},
		},
		{
			name: "length matches slice length",
			value: struct {
				val []uint16 `m1binary:"length:2"`
			}{
				val: []uint16{12345, 54321},
			},
			expected: []byte{0x39, 0x30, 0x31, 0xD4},
		},
		{
			name: "length lower slice length",
			value: struct {
				val []uint16 `m1binary:"length:1"`
			}{
				val: []uint16{12345, 54321},
			},
			expected: []byte{0x39, 0x30},
		},
		{
			name: "length higher slice length",
			value: struct {
				val []uint16 `m1binary:"length:3"`
			}{
				val: []uint16{12345, 54321},
			},
			expected: []byte{0x39, 0x30, 0x31, 0xD4, 0x00, 0x00},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := m1binary.Encode(&tc.value)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestLengthRefEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected []byte
	}{
		{
			name: "lengthRef: amount matches slice length",
			value: struct {
				Amount byte
				Data   []byte `m1binary:"lengthRef:Amount"`
			}{
				Amount: 2,
				Data:   []byte{0xFF, 0xAA},
			},
			expected: []byte{0x02, 0xFF, 0xAA},
		},
		{
			name: "lengthRef: amount lower than slice length",
			value: struct {
				Amount byte
				Data   []byte `m1binary:"lengthRef:Amount"`
			}{
				Amount: 1,
				Data:   []byte{0xFF, 0xAA},
			},
			expected: []byte{0x01, 0xFF},
		},
		{
			name: "lengthRef: amount longer than slice length (padding)",
			value: struct {
				Amount byte
				Data   []byte `m1binary:"lengthRef:Amount"`
			}{
				Amount: 3,
				Data:   []byte{0xFF, 0xAA},
			},
			expected: []byte{0x03, 0xFF, 0xAA, 0x00},
		},
		{
			name: "lengthRef: amount longer than slice length (padding)",
			value: struct {
				amount int8
				Data   []uint16 `m1binary:"lengthRef:amount"`
			}{
				amount: 3,
				Data:   []uint16{0xFF, 0xAA},
			},
			expected: []byte{0x03, 0xFF, 0x00, 0xAA, 0x00, 0x00, 0x00},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := m1binary.Encode(&tc.value)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSkipEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected []byte
	}{
		{
			name: "lengthRef: amount matches slice length",
			value: struct {
				val bool `m1binary:"skip:3"`
			}{
				val: true,
			},
			expected: []byte{0x01, 0x00, 0x00, 0x00},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := m1binary.Encode(&tc.value)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPanicsEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		value any
	}{
		{
			name: "lengthRef: ref does not exist",
			value: struct {
				val bool `m1binary:"lengthRef:ref"`
			}{
				val: true,
			},
		},
		{
			name: "lengthRef: ref is non integer type",
			value: struct {
				ref float32
				val bool `m1binary:"lengthRef:ref"`
			}{
				val: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Panics(t, func() {
				_, _ = m1binary.Encode(&tc.value)
			})
		})
	}
}

func TestErrorsEncode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected string
	}{
		{
			name: "lengthRef: ref is negative",
			value: struct {
				ref int8
				val bool `m1binary:"lengthRef:ref"`
			}{
				ref: -1,
				val: true,
			},
			expected: "negative",
		},
		{
			name:     "unsupported field",
			value:    complex64(0),
			expected: "unsupported",
		},
		{
			name:     "unsupported struct field",
			value:    struct{ val complex64 }{val: 0},
			expected: "unsupported",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := m1binary.Encode(&tc.value)

			assert.ErrorContains(t, err, tc.expected)
		})
	}
}
