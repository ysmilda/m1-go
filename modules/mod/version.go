package mod

import (
	"github.com/ysmilda/gobinary"
	"github.com/ysmilda/m1-go/modules/msys"
)

type Version msys.Version

func (v *Version) BinaryDecode(in []byte) (int, error) {
	code := struct {
		ReleaseType msys.ReleaseType
		Major       uint32
		Minor       uint32
		Patch       uint32
	}{}

	n, err := gobinary.Decode(in, &code)
	if err != nil {
		return n, err
	}

	v.ReleaseType = code.ReleaseType
	v.Major = code.Major
	v.Minor = code.Minor
	v.Patch = code.Patch
	return n, nil
}
