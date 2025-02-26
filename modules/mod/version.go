package mod

import (
	"github.com/ysmilda/m1-go/internals/m1binary"
	"github.com/ysmilda/m1-go/modules/msys"
)

type Version msys.Version

func (v *Version) DecodeM1(in []byte) (int, error) {
	code := struct {
		ReleaseType msys.ReleaseType
		Major       uint32
		Minor       uint32
		Patch       uint32
	}{}

	n, err := m1binary.Decode(in, &code)
	if err != nil {
		return n, err
	}

	v.ReleaseType = code.ReleaseType
	v.Major = code.Major
	v.Minor = code.Minor
	v.Patch = code.Patch
	return n, nil
}
