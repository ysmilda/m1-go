package rpc

// --------------
// ListCaller
// --------------

type PaginatedCaller interface {
	SetStart(uint32)
	SetCount(uint32)
}

type Start uint32

func (f *Start) SetStart(start uint32) {
	*f = Start(start)
}

type Count uint32

func (c *Count) SetCount(count uint32) {
	*c = Count(count)
}

func (c Count) GetCount() uint32 {
	return uint32(c)
}

var (
	_ PaginatedCaller = &PaginatedCallStartCount{}
	_ PaginatedCaller = &PaginatedCallCountStart{}
	_ PaginatedCaller = &PaginatedCallFirstLast{}
)

type PaginatedCallStartCount struct {
	Start
	Count
}

type PaginatedCallCountStart struct {
	Count
	Start
}

type PaginatedCallFirstLast struct {
	Start
	last uint32
}

func (l *PaginatedCallFirstLast) SetCount(count uint32) {
	l.last = uint32(l.Start) + count
}

type PaginatedReplier[T any] interface {
	GetCount() uint32
	Done(uint32) bool
	GetValues() []T
}

// --------------
// ListReplier
// --------------

type Values[T any] []T

func (v Values[T]) GetValues() []T {
	return v
}

var (
	_ PaginatedReplier[any] = &PaginatedReplyCount[any]{}
	_ PaginatedReplier[any] = &PaginatedReplyContinuationCount[any]{}
)

type PaginatedReplyCount[T any] struct {
	Count
	Values[T] `m1binary:"lengthRef:Count"`
}

func (l PaginatedReplyCount[T]) Done(step uint32) bool {
	return uint32(l.Count) < step
}

type PaginatedReplyContinuationCount[T any] struct {
	ContinuationPoint uint32 `m1binary:"skip:12"`
	Count
	Values[T] `m1binary:"lengthRef:Count,allign4"`
}

func (l PaginatedReplyContinuationCount[T]) Done(step uint32) bool {
	return l.ContinuationPoint == 0
}
