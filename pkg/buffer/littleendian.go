package buffer

import (
	"encoding/binary"
	"math"
)

type littleEndian struct {
	buf *Buffer
}

func (l *littleEndian) ReadUint16() (uint16, error) {
	line, err := l.buf.readSlice(2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(line), nil
}

func (l *littleEndian) ReadInt16() (int16, error) {
	val, err := l.ReadUint16()
	return int16(val), err
}

func (l *littleEndian) ReadUint32() (uint32, error) {
	line, err := l.buf.readSlice(4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(line), nil
}

func (l *littleEndian) ReadInt32() (int32, error) {
	val, err := l.ReadUint32()
	return int32(val), err
}

func (l *littleEndian) ReadUint64() (uint64, error) {
	line, err := l.buf.readSlice(8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(line), nil
}

func (l *littleEndian) ReadInt64() (int64, error) {
	val, err := l.ReadUint64()
	return int64(val), err
}

func (l *littleEndian) ReadFloat32() (float32, error) {
	val, err := l.ReadUint32()
	return math.Float32frombits(val), err
}

func (l *littleEndian) ReadFloat64() (float64, error) {
	val, err := l.ReadUint64()
	return math.Float64frombits(val), err
}

func (l *littleEndian) WriteUint16(i uint16) {
	temp := make([]byte, 2)
	binary.LittleEndian.PutUint16(temp, i)
	_, _ = l.buf.Write(temp)
}

func (l *littleEndian) WriteInt16(i int16) {
	l.WriteUint16(uint16(i))
}

func (l *littleEndian) WriteUint32(i uint32) {
	temp := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp, i)
	_, _ = l.buf.Write(temp)
}

func (l *littleEndian) WriteInt32(i int32) {
	l.WriteUint32(uint32(i))
}

func (l *littleEndian) WriteUint64(i uint64) {
	temp := make([]byte, 8)
	binary.LittleEndian.PutUint64(temp, i)
	_, _ = l.buf.Write(temp)
}

func (l *littleEndian) WriteInt64(i int64) {
	l.WriteUint64(uint64(i))
}

func (l *littleEndian) WriteFloat32(f float32) {
	i := math.Float32bits(f)
	l.WriteUint32(i)
}

func (l *littleEndian) WriteFloat64(f float64) {
	i := math.Float64bits(f)
	l.WriteUint64(i)
}
