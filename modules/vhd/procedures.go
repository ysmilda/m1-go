package vhd

import "github.com/ysmilda/m1-go/internals/rpc"

var Procedures = procedures{}

type procedures struct{}

func (p procedures) StartSession(c StartSessionCall) rpc.Procedure[StartSessionCall, StartSessionReply] {
	return rpc.NewProcedure[StartSessionCall, StartSessionReply](102, rpc.VersionDefault, c)
}

func (p procedures) StopSession(c StopSessionCall) rpc.Procedure[StopSessionCall, StopSessionReply] {
	return rpc.NewProcedure[StopSessionCall, StopSessionReply](104, rpc.VersionDefault, c)
}

func (p procedures) ResetSession(c ResetSessionCall) rpc.Procedure[ResetSessionCall, ResetSessionReply] {
	return rpc.NewProcedure[ResetSessionCall, ResetSessionReply](106, rpc.VersionDefault, c)
}

func (p procedures) GetSessionInfo(c GetSessionInfoCall) rpc.Procedure[GetSessionInfoCall, GetSessionInfoReply] {
	return rpc.NewProcedure[GetSessionInfoCall, GetSessionInfoReply](108, rpc.VersionDefault, c)
}

func (p procedures) GetValue(c GetValueCall) rpc.Procedure[GetValueCall, GetValueReply] {
	return rpc.NewProcedure[GetValueCall, GetValueReply](110, rpc.VersionDefault, c)
}

func (p procedures) SetValue(c SetValueCall) rpc.Procedure[SetValueCall, SetValueReply] {
	return rpc.NewProcedure[SetValueCall, SetValueReply](112, rpc.VersionDefault, c)
}

func (p procedures) AddList(c AddListCall) rpc.Procedure[AddListCall, AddListReply] {
	return rpc.NewProcedure[AddListCall, AddListReply](114, rpc.VersionDefault, c)
}

func (p procedures) DeleteList(c DeleteListCall) rpc.Procedure[DeleteListCall, DeleteListReply] {
	return rpc.NewProcedure[DeleteListCall, DeleteListReply](116, rpc.VersionDefault, c)
}

func (p procedures) GetUpdate(c GetUpdateCall) rpc.Procedure[GetUpdateCall, GetUpdateReply] {
	return rpc.NewProcedure[GetUpdateCall, GetUpdateReply](118, rpc.VersionDefault, c)
}

func (p procedures) GetAddress(c GetAddressCall) rpc.Procedure[GetAddressCall, GetAddressReply] {
	return rpc.NewProcedure[GetAddressCall, GetAddressReply](120, rpc.VersionDefault, c)
}

func (p procedures) ResetList(c ResetListCall) rpc.Procedure[ResetListCall, ResetListReply] {
	return rpc.NewProcedure[ResetListCall, ResetListReply](122, rpc.VersionDefault, c)
}

func (p procedures) StartList(c StartListCall) rpc.Procedure[StartListCall, StartListReply] {
	return rpc.NewProcedure[StartListCall, StartListReply](124, rpc.VersionDefault, c)
}

func (p procedures) StopList(c StopListCall) rpc.Procedure[StopListCall, StopListReply] {
	return rpc.NewProcedure[StopListCall, StopListReply](126, rpc.VersionDefault, c)
}

func (p procedures) CallbackInfo(c GetCallbackInfoCall) rpc.Procedure[GetCallbackInfoCall, GetCallbackInfoReply] {
	return rpc.NewProcedure[GetCallbackInfoCall, GetCallbackInfoReply](128, rpc.VersionDefault, c)
}

func (p procedures) StartElement(c StartElementCall) rpc.Procedure[StartElementCall, StartElementReply] {
	return rpc.NewProcedure[StartElementCall, StartElementReply](130, rpc.VersionDefault, c)
}

func (p procedures) StopElement(c StopElementCall) rpc.Procedure[StopElementCall, StopElementReply] {
	return rpc.NewProcedure[StopElementCall, StopElementReply](132, rpc.VersionDefault, c)
}

func (p procedures) GetXAddress(c GetXAddressCall) rpc.Procedure[GetXAddressCall, GetXAddressReply] {
	return rpc.NewProcedure[GetXAddressCall, GetXAddressReply](134, rpc.VersionDefault, c)
}
