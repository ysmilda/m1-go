// package res contains the interface definition of the Resource Manager
//
//nolint:lll
package res

import (
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/internals/rpc"
)

var Procedures = procedures{}

type procedures struct{}

func (procedures) GetModuleInfo(client *m1client.Client, module ModuleNumber, c ModuleInfoCall) (*ModuleInfoReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleInfoCall, ModuleInfoReply](104, rpc.VersionDefault, c))
}

func (procedures) RequestModuleAccess(client *m1client.Client, module ModuleNumber, c ModuleAccessCall) (*ModuleAccessReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleAccessCall, ModuleAccessReply](108, rpc.VersionDefault, c))
}

func (procedures) ReleaseModuleAccess(client *m1client.Client, module ModuleNumber, c ModuleFreeCall) (*ModuleFreeReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleFreeCall, ModuleFreeReply](110, rpc.VersionDefault, c))
}

func (procedures) GetModuleNumber(client *m1client.Client, module ModuleNumber, c ModuleNumberCall) (*ModuleNumberReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleNumberCall, ModuleNumberReply](112, rpc.VersionDefault, c))
}

func (procedures) GetExtendedModuleInfo(client *m1client.Client, module ModuleNumber, c ExtendedModuleInfoCall) (*ExtendedModuleInfoReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedModuleInfoCall, ExtendedModuleInfoReply](116, rpc.VersionDefault, c))
}

func (procedures) ListModuleChildTasks(client *m1client.Client, module ModuleNumber, c ModuleChildCall) (*ModuleChildReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleChildCall, ModuleChildReply](120, rpc.VersionDefault, c))
}

func (procedures) ListModuleTasks(client *m1client.Client, module ModuleNumber, c ModuleTaskCall) (*ModuleTaskReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleTaskCall, ModuleTaskReply](122, rpc.VersionDefault, c))
}

// TODO: Add missing procedures

func (procedures) GetSystemInfo(client *m1client.Client, module ModuleNumber, c SystemInfoCall) (*SystemInfoReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SystemInfoCall, SystemInfoReply](282, rpc.VersionRES, c))
}

func (procedures) Login(client *m1client.Client, module ModuleNumber, c LoginCall) (*LoginReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LoginCall, LoginReply](284, rpc.VersionRES, c))
}

func (procedures) Logout(client *m1client.Client, module ModuleNumber, c LogoutCall) (*LogoutReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LogoutCall, LogoutReply](286, rpc.VersionRES, c))
}

// TODO: Add missing procedures

func (procedures) Login2(client *m1client.Client, module ModuleNumber, c Login2Call) (*Login2Reply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[Login2Call, Login2Reply](304, rpc.VersionRES, c))
}

func (procedures) OpenConnection(client *m1client.Client, module ModuleNumber, c OpenCall) (*OpenReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[OpenCall, OpenReply](306, rpc.VersionRES, c))
}

func (procedures) CloseConnection(client *m1client.Client, module ModuleNumber, c CloseCall) (*CloseReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CloseCall, CloseReply](308, rpc.VersionRES, c))
}

func (procedures) RenewConnection(client *m1client.Client, module ModuleNumber, c RenewCall) (*RenewReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[RenewCall, RenewReply](310, rpc.VersionRES, c))
}

// TODO: Add missing procedures

func (procedures) ExtPing(client *m1client.Client, module ModuleNumber, c ExtPingCall) (*ExtPingReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtPingCall, ExtPingReply](320, rpc.VersionDefault, c))
}

// TODO: Add missing procedures

func (procedures) FlashLED(client *m1client.Client, module ModuleNumber, c FlashLEDCall) (*FlashLEDReply, error) {
	return rpc.Call(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[FlashLEDCall, FlashLEDReply](324, rpc.VersionDefault, c))
}

// --------------
// ListProcedures
// --------------

func (procedures) ModuleInfo(client *m1client.Client, module ModuleNumber, c *ListModuleInfoCall, stepSize uint32) ([]ModuleInfo, error) {
	return rpc.ListCall(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewListProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply](106, rpc.VersionDefault, c), stepSize)
}

func (procedures) ExtendedModuleInfo(client *m1client.Client, module ModuleNumber, c *ExtendedModuleInfoListCall, stepSize uint32) ([]ExtendedModuleInfo, error) {
	return rpc.ListCall(client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewListProcedure[ExtendedModuleInfo, *ExtendedModuleInfoListCall, *ExtendedModuleInfoListReply](118, rpc.VersionDefault, c), stepSize)
}
