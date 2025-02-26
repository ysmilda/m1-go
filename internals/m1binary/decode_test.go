package m1binary_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ysmilda/m1-go/internals/m1binary"
)

func TestBaseTypesDecode(t *testing.T) {
	t.Parallel()

	var bool_t bool
	_, err := m1binary.Decode([]byte{0x01}, &bool_t)
	assert.NoError(t, err)
	assert.Equal(t, true, bool_t)

	var uint8_t uint8
	_, err = m1binary.Decode([]byte{0x7B}, &uint8_t)
	assert.NoError(t, err)
	assert.Equal(t, uint8(123), uint8_t)

	var int8_t int8
	_, err = m1binary.Decode([]byte{0x85}, &int8_t)
	assert.NoError(t, err)
	assert.Equal(t, int8(-123), int8_t)

	var uint16_t uint16
	_, err = m1binary.Decode([]byte{0x39, 0x30}, &uint16_t)
	assert.NoError(t, err)
	assert.Equal(t, uint16(12345), uint16_t)

	var int16_t int16
	_, err = m1binary.Decode([]byte{0xC7, 0xCF}, &int16_t)
	assert.NoError(t, err)
	assert.Equal(t, int16(-12345), int16_t)

	var uint32_t uint32
	_, err = m1binary.Decode([]byte{0x15, 0xCD, 0x5B, 0x07}, &uint32_t)
	assert.NoError(t, err)
	assert.Equal(t, uint32(123456789), uint32_t)

	var int32_t int32
	_, err = m1binary.Decode([]byte{0xEB, 0x32, 0xA4, 0xF8}, &int32_t)
	assert.NoError(t, err)
	assert.Equal(t, int32(-123456789), int32_t)

	var uint64_t uint64
	_, err = m1binary.Decode([]byte{0x79, 0xDF, 0x0D, 0x86, 0x48, 0x70, 0x00, 0x00}, &uint64_t)
	assert.NoError(t, err)
	assert.Equal(t, uint64(123456789012345), uint64_t)

	var int64_t int64
	_, err = m1binary.Decode([]byte{0x87, 0x20, 0xF2, 0x79, 0xB7, 0x8F, 0xFF, 0xFF}, &int64_t)
	assert.NoError(t, err)
	assert.Equal(t, int64(-123456789012345), int64_t)

	var float32_t float32
	_, err = m1binary.Decode([]byte{0x0C, 0x91, 0xE0, 0x56}, &float32_t)
	assert.NoError(t, err)
	assert.Equal(t, float32(123456789012345), float32_t)

	var float64_t float64
	_, err = m1binary.Decode([]byte{0x40, 0xDE, 0x77, 0x83, 0x21, 0x12, 0xDC, 0xC2}, &float64_t)
	assert.NoError(t, err)
	assert.Equal(t, float64(-123456789012345), float64_t)

	var interface_t any = &float64_t
	_, err = m1binary.Decode([]byte{0x40, 0xDE, 0x77, 0x83, 0x21, 0x12, 0xDC, 0xC2}, &interface_t)
	assert.NoError(t, err)
	assert.Equal(t, float64(-123456789012345), *interface_t.(*float64))
}

func TestStringDecode(t *testing.T) {
	t.Parallel()

	var strTarget string
	_, err := m1binary.Decode([]byte{'t', 'e', 's', 't'}, &strTarget)
	assert.NoError(t, err)
	assert.Equal(t, "test", strTarget)

	var structTarget struct {
		Str string `m1binary:"length:4"`
	}
	_, err = m1binary.Decode([]byte{'t', 'e', 's', 't'}, &structTarget)
	assert.NoError(t, err)
	assert.Equal(t, "test", structTarget.Str)

	var structTargetInsufficientLength struct {
		Str string `m1binary:"length:3"`
	}
	_, err = m1binary.Decode([]byte{'t', 'e', 's', 't'}, &structTargetInsufficientLength)
	assert.NoError(t, err)
	assert.Equal(t, "tes", structTargetInsufficientLength.Str)

	var structTargetZeroes struct {
		Str string `m1binary:"length:6"`
	}
	_, err = m1binary.Decode([]byte{'t', 'e', 's', 't', 0, 0}, &structTargetZeroes)
	assert.NoError(t, err)
	assert.Equal(t, "test", structTargetZeroes.Str)

	var structTargetZeroTerminated struct {
		Str string `m1binary:"length:6,zeroTerminated"`
	}
	n, err := m1binary.Decode([]byte{'t', 'e', 's', 't', 0, 0}, &structTargetZeroTerminated)
	assert.NoError(t, err)
	assert.Equal(t, "test", structTargetZeroTerminated.Str)
	assert.Equal(t, 5, n)

	var structTargetSlice struct {
		Strs []string `m1binary:"length:2,elementLength:5"`
	}
	_, err = m1binary.Decode([]byte{'t', 'e', 's', 't', 0, 'o', 't', 'h', 'e', 'r'}, &structTargetSlice)
	assert.NoError(t, err)
	assert.Equal(t, []string{"test", "other"}, structTargetSlice.Strs)
	assert.Equal(t, 5, n)

	var structTargetSliceZeroTerminated struct {
		Strs []string `m1binary:"length:2,elementLength:5,zeroTerminated"`
	}
	_, err = m1binary.Decode([]byte{'t', 'e', 's', 't', 0, 'o', 't', 'h', 'e', 'r'}, &structTargetSliceZeroTerminated)
	assert.NoError(t, err)
	assert.Equal(t, []string{"test", "other"}, structTargetSliceZeroTerminated.Strs)

	var structTargetSliceZeroTerminatedWithoutLength struct {
		Strs []string `m1binary:"length:2,zeroTerminated"`
	}
	_, err = m1binary.Decode(
		[]byte{'t', 'e', 's', 't', 0, 'o', 't', 'h', 'e', 'r'},
		&structTargetSliceZeroTerminatedWithoutLength,
	)
	assert.NoError(t, err)
	assert.Equal(t, []string{"test", "other"}, structTargetSliceZeroTerminatedWithoutLength.Strs)
}

func TestSliceDecode(t *testing.T) {
	t.Parallel()

	var structTargetSliceReadTillEnd struct {
		Slice []byte `m1binary:"tillEnd"`
	}
	_, err := m1binary.Decode([]byte{'t', 'e', 's', 't', 0, 'o', 't', 'h', 'e', 'r'}, &structTargetSliceReadTillEnd)
	assert.NoError(t, err)
	assert.Equal(t, []byte{'t', 'e', 's', 't', 0, 'o', 't', 'h', 'e', 'r'}, structTargetSliceReadTillEnd.Slice)
}

func TestTimeDecode(t *testing.T) {
	t.Parallel()

	var timeUint8 struct {
		Timestamp time.Time `m1binary:"length:1"`
	}
	_, err := m1binary.Decode([]byte{123}, &timeUint8)
	assert.NoError(t, err)
	assert.Equal(t, time.UnixMicro((123 * time.Second).Milliseconds()), timeUint8.Timestamp)

	var timeUint16Minutes struct {
		Timestamp time.Time `m1binary:"length:2,unit:minutes"`
	}
	_, err = m1binary.Decode([]byte{0x39, 0x30}, &timeUint16Minutes)
	assert.NoError(t, err)
	assert.Equal(t, time.UnixMicro((12345 * time.Minute).Microseconds()), timeUint16Minutes.Timestamp)

	var timeUint32Seconds struct {
		Timestamp time.Time `m1binary:"length:4,unit:seconds"`
	}
	_, err = m1binary.Decode([]byte{0x15, 0xCD, 0x5B, 0x07}, &timeUint32Seconds)
	assert.NoError(t, err)
	assert.Equal(t, time.UnixMicro((123456789 * time.Second).Microseconds()), timeUint32Seconds.Timestamp)

	var timeUint64Milliseconds struct {
		Timestamp time.Time `m1binary:"length:8,unit:milliseconds"`
	}
	_, err = m1binary.Decode([]byte{0xCB, 0x04, 0xFB, 0x71, 0x1F, 0x01, 0x00, 0x00}, &timeUint64Milliseconds)
	assert.NoError(t, err)
	assert.Equal(t, time.UnixMicro((1234567890123 * time.Millisecond).Microseconds()), timeUint64Milliseconds.Timestamp)

	var durationUint8Hours struct {
		Duration time.Duration `m1binary:"length:1,unit:hours"`
	}
	_, err = m1binary.Decode([]byte{123}, &durationUint8Hours)
	assert.NoError(t, err)
	assert.Equal(t, 123*time.Hour, durationUint8Hours.Duration)

	var durationUint16Days struct {
		Duration time.Duration `m1binary:"length:2,unit:days"`
	}
	_, err = m1binary.Decode([]byte{0x39, 0x30}, &durationUint16Days)
	assert.NoError(t, err)
	assert.Equal(t, 12345*time.Hour*24, durationUint16Days.Duration)
}

func TestErrorsDecode(t *testing.T) {
	t.Parallel()

	var bool_t bool
	_, err := m1binary.Decode([]byte{}, &bool_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var uint8_t uint8
	_, err = m1binary.Decode([]byte{}, &uint8_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var int8_t int8
	_, err = m1binary.Decode([]byte{}, &int8_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var uint16_t uint16
	_, err = m1binary.Decode([]byte{0x39}, &uint16_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var int16_t int16
	_, err = m1binary.Decode([]byte{0xC7}, &int16_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var uint32_t uint32
	_, err = m1binary.Decode([]byte{0x15, 0xCD, 0x5B}, &uint32_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var int32_t int32
	_, err = m1binary.Decode([]byte{0xEB, 0x32, 0xA4}, &int32_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var uint64_t uint64
	_, err = m1binary.Decode([]byte{0x79, 0xDF, 0x0D, 0x86, 0x48, 0x70, 0x00}, &uint64_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var int64_t int64
	_, err = m1binary.Decode([]byte{0x87, 0x20, 0xF2, 0x79, 0xB7, 0x8F, 0xFF}, &int64_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var float32_t float32
	_, err = m1binary.Decode([]byte{0x0C, 0x91, 0xE0}, &float32_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var float64_t float64
	_, err = m1binary.Decode([]byte{0x40, 0xDE, 0x77, 0x83, 0x21, 0x12, 0xDC}, &float64_t)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)

	var structTargetSliceNotEnoughData struct {
		Slice []uint8 `m1binary:"length:1"`
	}

	_, err = m1binary.Decode([]byte{}, &structTargetSliceNotEnoughData)
	assert.ErrorIs(t, err, m1binary.ErrInvalidDataLength)
}

func TestPanicsDecode(t *testing.T) {
	var structTargetSliceMissingLengthTag struct {
		Slice []uint8
	}
	assert.Panics(t, func() {
		_, _ = m1binary.Decode([]byte{}, &structTargetSliceMissingLengthTag)
	})

	var structTargetStringMissingLengthTag struct {
		String string
	}
	assert.Panics(t, func() {
		_, _ = m1binary.Decode([]byte{}, &structTargetStringMissingLengthTag)
	})

	var unsupportedField complex128
	assert.Panics(t, func() {
		_, _ = m1binary.Decode([]byte{}, &unsupportedField)
	})
}

type DecoderM1Test struct {
	A uint32
	B uint32
}

func (d *DecoderM1Test) DecodeM1(data []byte) (int, error) {
	temp := struct {
		B uint32
		A uint32
	}{}
	n, err := m1binary.Decode(data, &temp)
	if err != nil {
		return n, err
	}

	d.A = temp.A
	d.B = temp.B

	return n, nil
}

func TestDecodeM1Interface(t *testing.T) {
	var temp DecoderM1Test
	_, err := m1binary.Decode([]byte{0x15, 0xCD, 0x5B, 0x07, 0o7, 0x5B, 0xCD, 0x15}, &temp)
	assert.NoError(t, err)
	assert.Equal(t, uint32(365779719), temp.A)
	assert.Equal(t, uint32(123456789), temp.B)
}
