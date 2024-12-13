package msys

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

type TickSource uint32

const (
	TickSourceTimer TickSource = 0
	TickSourceRTC   TickSource = 1
)

func (t TickSource) String() string {
	switch t {
	case 0:
		return "Timer"
	case 1:
		return "RTC"
	}
	return "Unknown"
}
