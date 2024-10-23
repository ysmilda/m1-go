package m1

import (
	"time"

	"github.com/ysmilda/m1-go/pkg/buffer"
)

type SessionInfo struct {
	Name   string
	UserID uint32
	Start  time.Time
}

func (v *SessionInfo) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	v.Name, _ = in.ReadString(_VHD_UserNameLength)
	v.UserID, _ = in.LittleEndian.ReadUint32()
	start, _ := in.LittleEndian.ReadUint32()
	v.Start = time.Unix(int64(start), 0)
	return returnCode
}

type SviVariable struct {
	Name  string
	Error error

	Address uint64
	Format  uint16
	Length  uint16

	initialized bool
}

func (v SviVariable) IsBlock() bool {
	return v.Format&_FormatBlock != 0
}

func (s *SviVariable) parse(in *buffer.Buffer) {
	s.Address, _ = in.LittleEndian.ReadUint64()
	s.Format, _ = in.LittleEndian.ReadUint16()
	s.Length, _ = in.LittleEndian.ReadUint16()
}
