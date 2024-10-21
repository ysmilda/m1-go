package buffer

import (
	"encoding/binary"
)

type bigEndian struct {
	buf *Buffer
}

func (b *bigEndian) ReadUint16() (uint16, error) {
	line, err := b.buf.readSlice(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(line), nil
}

func (b *bigEndian) ReadUint32() (uint32, error) {
	line, err := b.buf.readSlice(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(line), nil
}

func (b *bigEndian) ReadUint64() (uint64, error) {
	line, err := b.buf.readSlice(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(line), nil
}

func (b *bigEndian) WriteUint16(i uint16) {
	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, i)
	_, _ = b.buf.Write(temp)
}

func (b *bigEndian) WriteUint32(i uint32) {
	temp := make([]byte, 4)
	binary.BigEndian.PutUint32(temp, i)
	_, _ = b.buf.Write(temp)
}

func (b *bigEndian) WriteUint64(i uint64) {
	temp := make([]byte, 8)
	binary.BigEndian.PutUint64(temp, i)
	_, _ = b.buf.Write(temp)
}
