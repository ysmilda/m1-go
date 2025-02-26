package msys

import (
	"fmt"
)

// Version contains information about a version.
// This is the general representation of a version within the M1 platform.
type Version struct {
	Major       uint32
	Minor       uint32
	Patch       uint32
	ReleaseType ReleaseType
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

	if s.ReleaseType.String() == ReleaseString {
		if v2.ReleaseType.String() == ReleaseString {
			return 0
		}
		return 1
	}

	if v2.ReleaseType.String() == ReleaseString {
		return -1
	}

	if s.Patch > v2.Patch {
		return 1
	} else if s.Patch < v2.Patch {
		return -1
	}

	return 0
}

type ReleaseType uint32

const (
	Alpha   ReleaseType = 1
	Beta    ReleaseType = 2
	Release ReleaseType = 3
)

const (
	AlphaString   string = "alpha"
	BetaString    string = "beta"
	ReleaseString string = "release"
)

func (t ReleaseType) String() string {
	switch t {
	case 1:
		return AlphaString
	case 2:
		return BetaString
	case 3:
		return ReleaseString
	}
	return "Unknown"
}
