package rpc

type Procedure[C any, R ReturnCoder] struct {
	procedure uint32
	version   Version
	Call      C
}

func NewProcedure[C any, R ReturnCoder](procedure uint32, version Version, call C) Procedure[C, R] {
	return Procedure[C, R]{
		procedure: procedure,
		version:   version,
		Call:      call,
	}
}

func (p Procedure[C, R]) Procedure() uint32 {
	return p.procedure
}

func (p Procedure[C, R]) RPCVersion() Version {
	return p.version
}

type PaginatedProcedure[T any, C PaginatedCaller, R PaginatedReplier[T]] struct {
	procedure uint32
	version   Version
	Call      C
}

func NewPaginatedProcedure[T any, C PaginatedCaller, R PaginatedReturnCoder[T]](
	procedure uint32, version Version, call C,
) PaginatedProcedure[T, C, R] {
	return PaginatedProcedure[T, C, R]{
		procedure: procedure,
		version:   version,
		Call:      call,
	}
}

func (p PaginatedProcedure[T, C, R]) Procedure() uint32 {
	return p.procedure
}

func (p PaginatedProcedure[T, C, R]) RPCVersion() Version {
	return p.version
}
