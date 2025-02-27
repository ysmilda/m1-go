package m1

import (
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/res"
)

func Call[C any, R rpc.ReturnCoder](
	target *Target, info res.ModuleNumber, procedure rpc.Procedure[C, R],
) (*R, error) {
	return call(target.client, info, procedure)
}

func ListCall[T any, C rpc.ListCaller, R rpc.ListReturnCoder[T]](
	target *Target, info res.ModuleNumber, procedure rpc.ListProcedure[T, C, R], stepSize uint32,
) ([]T, error) {
	return listCall(target.client, info, procedure, stepSize)
}

// call is a helper function to call a procedure on the target.
func call[C any, R rpc.ReturnCoder](
	client *m1client.Client, info res.ModuleNumber, procedure rpc.Procedure[C, R],
) (*R, error) {
	return rpc.Call(client.GetConnection(info.Port), rpc.Header{
		Module:    info.ModuleNumber,
		Version:   procedure.RPCVersion(),
		Procedure: procedure.Procedure(),
		Auth:      client.GetAuth(),
	}, procedure)
}

// call is a helper function to call a procedure on the target.
func listCall[T any, C rpc.ListCaller, R rpc.ListReturnCoder[T]](
	client *m1client.Client, info res.ModuleNumber, procedure rpc.ListProcedure[T, C, R], stepSize uint32,
) ([]T, error) {
	return rpc.ListCall(client.GetConnection(info.Port), rpc.Header{
		Module:    info.ModuleNumber,
		Version:   procedure.RPCVersion(),
		Procedure: procedure.Procedure(),
		Auth:      client.GetAuth(),
	}, procedure, stepSize)
}
