package m1binary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ysmilda/m1-go/internals/m1binary"
)

func TestSizeOf(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		value    any
		expected int
	}{
		{
			name:     "Single element",
			value:    bool(true),
			expected: 1,
		},
		{
			name: "Simple struct",
			value: struct {
				A uint8
				B uint16
				C uint32
			}{},
			expected: 7,
		},
		{
			name: "Struct in struct",
			value: struct {
				A uint64
				B int8
				C struct {
					A int16
					B int32
					C int64
				}
			}{},
			expected: 23,
		},
		{
			name: "Slice",
			value: struct {
				A []float32
				B []float64
			}{
				A: []float32{1, 2, 3, 4},
				B: []float64{4, 5},
			},
			expected: 32,
		},
		{
			name: "String",
			value: struct {
				A string
				B string
			}{
				A: "hello",
				B: "world",
			},
			expected: 10,
		},
		{
			name: "Interface",
			value: struct {
				A any
				B any
			}{
				A: uint32(0),
				B: uint64(0),
			},
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := m1binary.SizeOf(tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}
}
