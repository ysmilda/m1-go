package rpc

// --------------
// ListCaller
// --------------

type ListCaller interface {
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
	_ ListCaller = &ListCallStartCount{}
	_ ListCaller = &ListCallCountStart{}
	_ ListCaller = &ListCallFirstLast{}
)

type ListCallStartCount struct {
	Start
	Count
}

type ListCallCountStart struct {
	Count
	Start
}

type ListCallFirstLast struct {
	Start
	last uint32
}

func (l *ListCallFirstLast) SetCount(count uint32) {
	l.last = uint32(l.Start) + count
}

type ListReplier[T any] interface {
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
	_ ListReplier[any] = &ListReplyCount[any]{}
	_ ListReplier[any] = &ListReplyContinuationCount[any]{}
)

type ListReplyCount[T any] struct {
	Count
	Values[T] `m1binary:"lengthRef:Count"`
}

func (l ListReplyCount[T]) Done(step uint32) bool {
	return uint32(l.Count) < step
}

type ListReplyContinuationCount[T any] struct {
	ContinuationPoint uint32 `m1binary:"skip:12"`
	Count
	Values[T] `m1binary:"lengthRef:Count,allign4"`
}

func (l ListReplyContinuationCount[T]) Done(step uint32) bool {
	return l.ContinuationPoint == 0
}
