package m1

import (
	"fmt"
	"time"

	"github.com/ysmilda/m1-go/pkg/buffer"
)

// CPUType contains information about the CPU.
type CPUType struct {
	Type         uint32
	Variant      uint32
	SerialNumber string
}

func (c *CPUType) parse(in *buffer.Buffer) {
	c.Type, _ = in.LittleEndian.ReadUint32()
	c.Variant, _ = in.LittleEndian.ReadUint32()
	c.SerialNumber, _ = in.ReadString(12)
}

// Version contains information about a version.
// This is the general represantation of a version within the M1 platform.
type Version struct {
	Major       uint32
	Minor       uint32
	Patch       uint32
	ReleaseType string
}

func (s *Version) parse(in *buffer.Buffer) {
	s.Major, _ = in.LittleEndian.ReadUint32()
	s.Minor, _ = in.LittleEndian.ReadUint32()
	s.Patch, _ = in.LittleEndian.ReadUint32()
	switch code, _ := in.LittleEndian.ReadUint32(); code {
	case 1:
		s.ReleaseType = "alpha"
	case 2:
		s.ReleaseType = "beta"
	case 3:
		s.ReleaseType = "release"
	}
}

func (s *Version) String() string {
	return fmt.Sprintf("V%d.%d.%d-%s", s.Major, s.Minor, s.Patch, s.ReleaseType)
}

// Compare compares a Version (v1) with another Version (v2).
// 1 = v1 > v2
// 0 = v1 == v2
// -1 = v1 < v2.
func (s *Version) Compare(v2 Version) int {
	if s.Major > v2.Major {
		return 1
	} else if s.Major < v2.Major {
		return -1
	}

	if s.Minor > v2.Minor {
		return 1
	} else if s.Minor < v2.Minor {
		return -1
	}

	if s.ReleaseType == "release" {
		if v2.ReleaseType == "release" {
			return 0
		}
		return 1
	}

	if v2.ReleaseType == "release" {
		return -1
	}

	if s.Patch > v2.Patch {
		return 1
	} else if s.Patch < v2.Patch {
		return -1
	}

	return 0
}

// DateTime contains information about the date and time.
type DateTime struct {
	Second   uint32
	Minute   uint32
	Hour     uint32
	MonthDay uint32
	Month    uint32
	Year     uint32
	WeekDay  uint32
	YearDay  uint32
	IsDST    uint32
}

func (s *DateTime) parse(in *buffer.Buffer) {
	s.Second, _ = in.LittleEndian.ReadUint32()
	s.Minute, _ = in.LittleEndian.ReadUint32()
	s.Hour, _ = in.LittleEndian.ReadUint32()
	s.MonthDay, _ = in.LittleEndian.ReadUint32()
	s.Month, _ = in.LittleEndian.ReadUint32()
	s.Year, _ = in.LittleEndian.ReadUint32()
	s.WeekDay, _ = in.LittleEndian.ReadUint32()
	s.YearDay, _ = in.LittleEndian.ReadUint32()
	s.IsDST, _ = in.LittleEndian.ReadUint32()
}

// UserAccess contains information about the user access.
type UserAccess struct {
	Group                      uint8
	Level                      uint8
	Priority                   uint8
	SystemPermissions          int64
	ApplicationPermissions     int64
	AppData                    int32
	PasswordValidityTime       time.Duration // Number of days that the password is valid (0 - 365)
	DaysTillPasswordExpiration time.Duration // Number of days till the
}

func (u *UserAccess) parse(in *buffer.Buffer) {
	u.Group, _ = in.ReadByte()
	u.Level, _ = in.ReadByte()
	u.Priority, _ = in.ReadByte()
	in.Skip(1) // spare
	u.SystemPermissions, _ = in.LittleEndian.ReadInt64()
	u.ApplicationPermissions, _ = in.LittleEndian.ReadInt64()
	u.AppData, _ = in.LittleEndian.ReadInt32()
	passwordValidityTime, _ := in.LittleEndian.ReadUint32()
	u.PasswordValidityTime = time.Duration(passwordValidityTime) * time.Hour * 24
	daysTillPasswordExpiration, _ := in.LittleEndian.ReadUint32()
	u.DaysTillPasswordExpiration = time.Duration(daysTillPasswordExpiration) * time.Hour * 24
	in.Skip(4) // spare
}
