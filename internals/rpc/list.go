package rpc

type ListCaller interface {
	SetFirst(uint32)
	SetCount(uint32)
}

var (
	_ ListCaller = &ListCallStartCount{}
	_ ListCaller = &ListCallCountStart{}
	_ ListCaller = &ListCallFirstLast{}
)

type ListCallStartCount struct {
	start uint32
	count uint32
}

func (l *ListCallStartCount) SetFirst(first uint32) {
	l.start = first
}

func (l *ListCallStartCount) SetCount(count uint32) {
	l.count = count
}

type ListCallCountStart struct {
	count uint32
	start uint32
}

func (l *ListCallCountStart) SetFirst(first uint32) {
	l.start = first
}

func (l *ListCallCountStart) SetCount(count uint32) {
	l.count = count
}

type ListCallFirstLast struct {
	first uint32
	last  uint32
}

func (l *ListCallFirstLast) SetFirst(first uint32) {
	l.first = first
}

func (l *ListCallFirstLast) SetCount(count uint32) {
	l.last = l.first + count
}

type ListReturnCoder[T any] interface {
	ReturnCoder
	ListReplier[T]
}

type ListReplier[T any] interface {
	GetCount() uint32
	Done(uint32) bool
	GetValues() []T
}

var (
	_ ListReplier[any] = &ListReplyCount[any]{}
	_ ListReplier[any] = &ListReplyContinuationCount[any]{}
)

type ListReplyCount[T any] struct {
	Count  uint32
	Values []T `m1binary:"lengthRef:Count"`
}

func (l ListReplyCount[T]) GetCount() uint32 {
	return l.Count
}

func (l ListReplyCount[T]) Done(step uint32) bool {
	return l.Count < step
}

func (l *ListReplyCount[T]) GetValues() []T {
	return l.Values
}

type ListReplyContinuationCount[T any] struct {
	ContinuationPoint uint32 `m1binary:"skip:12"`
	Count             uint32
	Values            []T `m1binary:"lengthRef:Count,allign4"`
}

func (l ListReplyContinuationCount[T]) GetCount() uint32 {
	return l.Count
}

func (l ListReplyContinuationCount[T]) Done(step uint32) bool {
	return l.ContinuationPoint == 0
}

func (l *ListReplyContinuationCount[T]) GetValues() []T {
	return l.Values
}
