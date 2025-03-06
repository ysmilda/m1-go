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

func (p *Procedures) GetModuleInfo(c ModuleInfoCall) (*ModuleInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ModuleInfoCall, ModuleInfoReply](104, rpc.VersionDefault, c))
}

func (p *Procedures) RequestModuleAccess(c ModuleAccessCall) (*ModuleAccessReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ModuleAccessCall, ModuleAccessReply](108, rpc.VersionDefault, c))
}

func (p *Procedures) ReleaseModuleAccess(c ModuleFreeCall) (*ModuleFreeReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ModuleFreeCall, ModuleFreeReply](110, rpc.VersionDefault, c))
}

func (p *Procedures) GetModuleNumber(c ModuleNumberCall) (*ModuleNumberReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ModuleNumberCall, ModuleNumberReply](112, rpc.VersionDefault, c))
}

func (p *Procedures) GetExtendedModuleInfo(c ExtendedModuleInfoCall) (*ExtendedModuleInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ExtendedModuleInfoCall, ExtendedModuleInfoReply](116, rpc.VersionDefault, c))
}

func (p *Procedures) ListModuleChildTasks(c ModuleChildCall) (*ModuleChildReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ModuleChildCall, ModuleChildReply](120, rpc.VersionDefault, c))
}

func (p *Procedures) ListModuleTasks(c ModuleTaskCall) (*ModuleTaskReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ModuleTaskCall, ModuleTaskReply](122, rpc.VersionDefault, c))
}

// TODO: Add missing procedures

func (p *Procedures) GetSystemInfo(c SystemInfoCall) (*SystemInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[SystemInfoCall, SystemInfoReply](282, rpc.VersionRES, c))
}

func (p *Procedures) Login(c LoginCall) (*LoginReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[LoginCall, LoginReply](284, rpc.VersionRES, c))
}

func (p *Procedures) Logout(c LogoutCall) (*LogoutReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[LogoutCall, LogoutReply](286, rpc.VersionRES, c))
}

// TODO: Add missing procedures

func (p *Procedures) Login2(c Login2Call) (*Login2Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[Login2Call, Login2Reply](304, rpc.VersionRES, c))
}

func (p *Procedures) OpenConnection(c OpenCall) (*OpenReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[OpenCall, OpenReply](306, rpc.VersionRES, c))
}

func (p *Procedures) CloseConnection(c CloseCall) (*CloseReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[CloseCall, CloseReply](308, rpc.VersionRES, c))
}

func (p *Procedures) RenewConnection(c RenewCall) (*RenewReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[RenewCall, RenewReply](310, rpc.VersionRES, c))
}

// TODO: Add missing procedures

func (p *Procedures) ExtPing(c ExtPingCall) (*ExtPingReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[ExtPingCall, ExtPingReply](320, rpc.VersionDefault, c))
}

// TODO: Add missing procedures

func (p *Procedures) FlashLED(c FlashLEDCall) (*FlashLEDReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewProcedure[FlashLEDCall, FlashLEDReply](324, rpc.VersionDefault, c))
}

// --------------
// Paginated procedures
// --------------

func (p *Procedures) ModuleInfo(c *ListModuleInfoCall, pageSize uint32) ([]ModuleInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewPaginatedProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply](106, rpc.VersionDefault, c), pageSize)
}

func (p *Procedures) ExtendedModuleInfo(c *ExtendedModuleInfoListCall, pageSize uint32) ([]ExtendedModuleInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: Module.ModuleNumber, Port: Module.Port}, rpc.NewPaginatedProcedure[ExtendedModuleInfo, *ExtendedModuleInfoListCall, *ExtendedModuleInfoListReply](118, rpc.VersionDefault, c), pageSize)
}
