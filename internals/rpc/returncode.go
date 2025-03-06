package rpc

type ReturnCode int32

func (r ReturnCode) GetReturnCode() uint32 {
	return uint32(r)
}

type ReturnCoder interface {
	GetReturnCode() uint32
}

type PaginatedReturnCoder[T any] interface {
	ReturnCoder
	PaginatedReplier[T]
}
