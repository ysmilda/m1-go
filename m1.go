package m1

import (
	"encoding/binary"
	"fmt"

	"github.com/ysmilda/m1-go/internals/client"
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/internals/unpack"
	"github.com/ysmilda/m1-go/modules/m1errors"
	"github.com/ysmilda/m1-go/modules/res"
)

func callProcedure[C any, R rpc.ReturnCoder](client *client.Client, info res.ModuleNumber, procedure rpc.Procedure[C, R]) error {
	return rpc.CallProcedure(client.GetConnection(info.ModuleNumber), rpc.Header{
		Module:    info.ModuleNumber,
		Version:   procedure.RPCVersion(),
		Procedure: procedure.Procedure(),
		Auth:      client.GetAuth(),
	}, procedure)
}

// call is a helper function to make a call to a module. It unpacks the call, sends it to the target and unpacks the
// reply. The reply is than checked for errors and returned.
func call[R rpc.ReturnCoder, C any](
	client *client.Client, info res.ModuleNumber, procedure uint32, version rpc.Version, call C,
) (*R, error) {
	body, err := unpack.Pack(binary.LittleEndian, call)
	if err != nil {
		return nil, fmt.Errorf("unable to parse call: %w", err)
	}

	buf, err := rpc.Call(
		client.GetConnection(info.ModuleNumber),
		rpc.Header{
			Module:    info.ModuleNumber,
			Version:   version,
			Procedure: procedure,
			Auth:      client.GetAuth(),
		},
		body,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to make rpc call: %w", err)
	}

	reply := *new(R)
	_, err = unpack.Unpack(buf, binary.LittleEndian, reply)
	if err != nil {
		return nil, fmt.Errorf("unable to parse reply: %w", err)
	}

	if err := m1errors.ParseReturnCode(reply.GetReturnCode()); err != nil {
		return nil, err
	}

	return &reply, nil
}
