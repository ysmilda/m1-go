package msys

import "time"

// DateTime contains information about the date and time.
// TODO: Verify the definition of `struct tm` on VxWorks.
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

func (d DateTime) ToTime() time.Time {
	return time.Date(
		int(d.Year), time.Month(d.Month), int(d.MonthDay), int(d.Hour), int(d.Minute), int(d.Second), 0, time.UTC,
	)
}

type Timestamp struct {
	Seconds     uint32
	NanoSeconds int32
}

func (t Timestamp) ToTime() time.Time {
	return time.Unix(int64(t.Seconds), int64(t.NanoSeconds))
}

func (t *Timestamp) FromTime(in time.Time) {
	t.Seconds = uint32(in.Second())
	t.NanoSeconds = int32(in.Nanosecond())
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
