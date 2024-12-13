package rpc

type ReturnCoder interface {
	GetReturnCode() uint32
}

type Procedure[C any, R ReturnCoder] struct {
	procedure uint32
	version   Version
	Call      C
	Reply     R
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
