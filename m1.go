package m1

import (
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/res"
)

func Call[C any, R rpc.ReturnCoder](
	target *Target, info res.ModuleNumber, procedure rpc.Procedure[C, R],
) (*R, error) {
	return rpc.Call(target.client, rpc.Module{Number: info.ModuleNumber, Port: info.Port}, procedure)
}

func PaginatedCall[T any, C rpc.PaginatedCaller, R rpc.PaginatedReturnCoder[T]](
	target *Target, info res.ModuleNumber, procedure rpc.PaginatedProcedure[T, C, R], pageSize uint32,
) ([]T, error) {
	return rpc.PaginatedCall(target.client, rpc.Module{Number: info.ModuleNumber, Port: info.Port}, procedure, pageSize)
}
