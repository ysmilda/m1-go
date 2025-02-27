// package res contains the interface definition of the Resource Manager
//
//nolint:lll
package res

import "github.com/ysmilda/m1-go/internals/rpc"

var Procedures = procedures{}

type procedures struct{}

func (procedures) GetModuleInfo(c ModuleInfoCall) rpc.Procedure[ModuleInfoCall, ModuleInfoReply] {
	return rpc.NewProcedure[ModuleInfoCall, ModuleInfoReply](104, rpc.VersionDefault, c)
}

func (procedures) RequestModuleAccess(c ModuleAccessCall) rpc.Procedure[ModuleAccessCall, ModuleAccessReply] {
	return rpc.NewProcedure[ModuleAccessCall, ModuleAccessReply](108, rpc.VersionDefault, c)
}

func (procedures) ReleaseModuleAccess(c ModuleFreeCall) rpc.Procedure[ModuleFreeCall, ModuleFreeReply] {
	return rpc.NewProcedure[ModuleFreeCall, ModuleFreeReply](110, rpc.VersionDefault, c)
}

func (procedures) GetModuleNumber(c ModuleNumberCall) rpc.Procedure[ModuleNumberCall, ModuleNumberReply] {
	return rpc.NewProcedure[ModuleNumberCall, ModuleNumberReply](112, rpc.VersionDefault, c)
}

func (procedures) GetExtendedModuleInfo(
	c ExtendedModuleInfoCall,
) rpc.Procedure[ExtendedModuleInfoCall, ExtendedModuleInfoReply] {
	return rpc.NewProcedure[ExtendedModuleInfoCall, ExtendedModuleInfoReply](116, rpc.VersionDefault, c)
}

func (procedures) ListModuleChildTasks(
	c ModuleChildCall,
) rpc.Procedure[ModuleChildCall, ModuleChildReply] {
	return rpc.NewProcedure[ModuleChildCall, ModuleChildReply](120, rpc.VersionDefault, c)
}

func (procedures) ListModuleTasks(
	c ModuleTaskCall,
) rpc.Procedure[ModuleTaskCall, ModuleTaskReply] {
	return rpc.NewProcedure[ModuleTaskCall, ModuleTaskReply](122, rpc.VersionDefault, c)
}

// TODO: Add missing procedures

func (procedures) GetSystemInfo(c SystemInfoCall) rpc.Procedure[SystemInfoCall, SystemInfoReply] {
	return rpc.NewProcedure[SystemInfoCall, SystemInfoReply](282, rpc.VersionRES, c)
}

func (procedures) Login(c LoginCall) rpc.Procedure[LoginCall, LoginReply] {
	return rpc.NewProcedure[LoginCall, LoginReply](284, rpc.VersionRES, c)
}

func (procedures) Logout(c LogoutCall) rpc.Procedure[LogoutCall, LogoutReply] {
	return rpc.NewProcedure[LogoutCall, LogoutReply](286, rpc.VersionRES, c)
}

// TODO: Add missing procedures

func (procedures) Login2(c Login2Call) rpc.Procedure[Login2Call, Login2Reply] {
	return rpc.NewProcedure[Login2Call, Login2Reply](304, rpc.VersionRES, c)
}

func (procedures) OpenConnection(c OpenCall) rpc.Procedure[OpenCall, OpenReply] {
	return rpc.NewProcedure[OpenCall, OpenReply](306, rpc.VersionRES, c)
}

func (procedures) CloseConnection(c CloseCall) rpc.Procedure[CloseCall, CloseReply] {
	return rpc.NewProcedure[CloseCall, CloseReply](308, rpc.VersionRES, c)
}

func (procedures) RenewConnection(c RenewCall) rpc.Procedure[RenewCall, RenewReply] {
	return rpc.NewProcedure[RenewCall, RenewReply](310, rpc.VersionRES, c)
}

// TODO: Add missing procedures

func (procedures) ExtPing(c ExtPingCall) rpc.Procedure[ExtPingCall, ExtPingReply] {
	return rpc.NewProcedure[ExtPingCall, ExtPingReply](320, rpc.VersionDefault, c)
}

// TODO: Add missing procedures

func (procedures) FlashLED(c FlashLEDCall) rpc.Procedure[FlashLEDCall, FlashLEDReply] {
	return rpc.NewProcedure[FlashLEDCall, FlashLEDReply](324, rpc.VersionDefault, c)
}

// --------------
// ListProcedures
// --------------

var ListProcedures listProcedures

type listProcedures struct{}

func (listProcedures) ModuleInfo(c *ListModuleInfoCall) rpc.ListProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply] {
	return rpc.NewListProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply](106, rpc.VersionDefault, c)
}

func (listProcedures) ExtendedModuleInfo(
	c *ExtendedModuleInfoListCall,
) rpc.ListProcedure[ExtendedModuleInfo, *ExtendedModuleInfoListCall, *ExtendedModuleInfoListReply] {
	return rpc.NewListProcedure[ExtendedModuleInfo, *ExtendedModuleInfoListCall, *ExtendedModuleInfoListReply](118, rpc.VersionDefault, c)
}

func Test(c *ListModuleInfoCall) rpc.ListProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply] {
	return rpc.NewListProcedure[ModuleInfo, *ListModuleInfoCall, *ListModuleInfoReply](106, rpc.VersionDefault, c)
}
