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
