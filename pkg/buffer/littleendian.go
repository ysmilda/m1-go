package buffer

import "encoding/binary"

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

func (l *littleEndian) WriteUint16(i uint16) {
	temp := make([]byte, 2)
	binary.LittleEndian.PutUint16(temp, i)
	_, _ = l.buf.Write(temp)
}

func (l *littleEndian) WriteUint32(i uint32) {
	temp := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp, i)
	_, _ = l.buf.Write(temp)
}

func (l *littleEndian) WriteUint64(i uint64) {
	temp := make([]byte, 8)
	binary.LittleEndian.PutUint64(temp, i)
	_, _ = l.buf.Write(temp)
}
