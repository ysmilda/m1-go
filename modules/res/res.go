// package res contains the interface definition of the Resource Manager
//
//nolint:lll
package res

import (
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/internals/rpc"
)

func NewProcedures(client *m1client.Client) *Procedures {
	return &Procedures{
		client: client,
	}
}

type Procedures struct {
	client *m1client.Client
}

func (p *Procedures) GetModuleInfo(module ModuleNumber, c ModuleInfoCall) (*ModuleInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleInfoCall, ModuleInfoReply](104, rpc.VersionDefault, c))
}

func (p *Procedures) RequestModuleAccess(module ModuleNumber, c ModuleAccessCall) (*ModuleAccessReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleAccessCall, ModuleAccessReply](108, rpc.VersionDefault, c))
}

func (p *Procedures) ReleaseModuleAccess(module ModuleNumber, c ModuleFreeCall) (*ModuleFreeReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleFreeCall, ModuleFreeReply](110, rpc.VersionDefault, c))
}

func (p *Procedures) GetModuleNumber(module ModuleNumber, c ModuleNumberCall) (*ModuleNumberReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleNumberCall, ModuleNumberReply](112, rpc.VersionDefault, c))
}

func (p *Procedures) GetExtendedModuleInfo(module ModuleNumber, c ExtendedModuleInfoCall) (*ExtendedModuleInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedModuleInfoCall, ExtendedModuleInfoReply](116, rpc.VersionDefault, c))
}

func (p *Procedures) ListModuleChildTasks(module ModuleNumber, c ModuleChildCall) (*ModuleChildReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleChildCall, ModuleChildReply](120, rpc.VersionDefault, c))
}

func (p *Procedures) ListModuleTasks(module ModuleNumber, c ModuleTaskCall) (*ModuleTaskReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ModuleTaskCall, ModuleTaskReply](122, rpc.VersionDefault, c))
}

// TODO: Add missing procedures

func (p *Procedures) GetSystemInfo(module ModuleNumber, c SystemInfoCall) (*SystemInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SystemInfoCall, SystemInfoReply](282, rpc.VersionRES, c))
}

func (p *Procedures) Login(module ModuleNumber, c LoginCall) (*LoginReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LoginCall, LoginReply](284, rpc.VersionRES, c))
}

func (p *Procedures) Logout(module ModuleNumber, c LogoutCall) (*LogoutReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LogoutCall, LogoutReply](286, rpc.VersionRES, c))
}

// TODO: Add missing procedures

func (p *Procedures) Login2(module ModuleNumber, c Login2Call) (*Login2Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[Login2Call, Login2Reply](304, rpc.VersionRES, c))
}

func (p *Procedures) OpenConnection(module ModuleNumber, c OpenCall) (*OpenReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[OpenCall, OpenReply](306, rpc.VersionRES, c))
}

func (p *Procedures) CloseConnection(module ModuleNumber, c CloseCall) (*CloseReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CloseCall, CloseReply](308, rpc.VersionRES, c))
}

func (p *Procedures) RenewConnection(module ModuleNumber, c RenewCall) (*RenewReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[RenewCall, RenewReply](310, rpc.VersionRES, c))
}

// TODO: Add missing procedures

func (p *Procedures) ExtPing(module ModuleNumber, c ExtPingCall) (*ExtPingReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtPingCall, ExtPingReply](320, rpc.VersionDefault, c))
}

// TODO: Add missing procedures

func (p *Procedures) FlashLED(module ModuleNumber, c FlashLEDCall) (*FlashLEDReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[FlashLEDCall, FlashLEDReply](324, rpc.VersionDefault, c))
}

// --------------
// ListProcedures
// --------------

func (p *Procedures) ModuleInfo(module ModuleNumber, c *ListModuleInfoCall, stepSize uint32) ([]ModuleInfo, error) {
	return rpc.ListCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewListProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply](106, rpc.VersionDefault, c), stepSize)
}

func (p *Procedures) ExtendedModuleInfo(module ModuleNumber, c *ExtendedModuleInfoListCall, stepSize uint32) ([]ExtendedModuleInfo, error) {
	return rpc.ListCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewListProcedure[ExtendedModuleInfo, *ExtendedModuleInfoListCall, *ExtendedModuleInfoListReply](118, rpc.VersionDefault, c), stepSize)
}
