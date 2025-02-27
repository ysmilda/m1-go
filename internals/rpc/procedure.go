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

type ListProcedure[T any, C ListCaller, R ListReplier[T]] struct {
	procedure uint32
	version   Version
	Call      C
}

func NewListProcedure[T any, C ListCaller, R ListReturnCoder[T]](
	procedure uint32, version Version, call C,
) ListProcedure[T, C, R] {
	return ListProcedure[T, C, R]{
		procedure: procedure,
		version:   version,
		Call:      call,
	}
}

func (p ListProcedure[T, C, R]) Procedure() uint32 {
	return p.procedure
}

func (p ListProcedure[T, C, R]) RPCVersion() Version {
	return p.version
}
