package m1binary

import (
	"strconv"
	"strings"

	"github.com/ysmilda/m1-go/internals/ptr"
)

const (
	structTag = "m1binary"
)

// Tag contains all available tag values. Depending on the type of value some may need to be set.
type Tag struct {
	// The length of an array, string or time.Duration/time.Time field. May not be used when lengthRef is used.
	// Used as length:number in the struct tag.
	Length *int
	// The length of an array element, used in combination with Length.
	// Used as elementLength:number in the struct tag.
	ElementLength *int
	// Alternative to length for dynamic length fields. Refers to the value of another struct field. May not be used
	// when length is used.
	// Used as lengthRef:otherVariableName in the struct tag.
	LengthRef *string
	// The "unit" of a time.Duration/time.Time field. If the field is not present the value defaults to nanoseconds.
	// Used as unit:[milliseconds|seconds|minutes|hours|days] in the struct tag.
	TimeUnit *string
	// The amount of bytes to Skip after reading the field.
	// Used as skip:number in the struct tag.
	Skip *int
	// If set a string will be read till the first null termination or the set length, whichever comes first. The next
	// entry will than be read from that index.s
	// Used as zeroTerminated in the struct tag.
	ZeroTerminated bool
	TillEnd        bool
	Allign4        bool
}

func parseTag(s string) Tag {
	if s == "" {
		return Tag{}
	}

	t := Tag{
		TimeUnit:       mustFindString(s, "unit"),
		Skip:           mustFindInt(s, "skip"),
		Length:         mustFindInt(s, "length"),
		ElementLength:  mustFindInt(s, "elementLength"),
		LengthRef:      mustFindString(s, "lengthRef"),
		ZeroTerminated: mustFindBool(s, "zeroTerminated"),
		TillEnd:        mustFindBool(s, "tillEnd"),
		Allign4:        mustFindBool(s, "allign4"),
	}

	if t.Length != nil && t.LengthRef != nil {
		panic("only one of length and lengtRef may be set")
	}

	return t
}

func mustFindString(s string, tag string) *string {
	idx := strings.Index(s, tag+":")
	if idx == -1 {
		return nil
	}
	idx += len(tag) + 1
	endIdx := idx + 1
	for ; endIdx < len(s); endIdx++ {
		if s[endIdx] == ',' {
			break
		}
	}
	return ptr.For(s[idx:endIdx])
}

func mustFindInt(s string, tag string) *int {
	str := mustFindString(s, tag)
	if str == nil {
		return nil
	}
	i, err := strconv.Atoi(*str)
	if err != nil {
		panic(err)
	}
	return ptr.For(i)
}

func mustFindBool(s string, tag string) bool {
	return strings.Contains(s, tag)
}
