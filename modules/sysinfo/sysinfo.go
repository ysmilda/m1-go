// package info contains the interface definition of the System Info Handler
// nolint:lll
package sysinfo

import "github.com/ysmilda/m1-go/internals/rpc"

var Procedures = procedures{}

type procedures struct{}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) ListCPUAddresses(c ListCPUAddressesCall) rpc.Procedure[ListCPUAddressesCall, ListCPUAddressesReply] {
	return rpc.NewProcedure[ListCPUAddressesCall, ListCPUAddressesReply](100, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) CPUInfo(c CPUInfoCall) rpc.Procedure[CPUInfoCall, CPUInfoReply] {
	return rpc.NewProcedure[CPUInfoCall, CPUInfoReply](102, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) ListSystemObjectInfo(c ListSystemObjectInfoCall) rpc.Procedure[ListSystemObjectInfoCall, ListSystemObjectInfoReply] {
	return rpc.NewProcedure[ListSystemObjectInfoCall, ListSystemObjectInfoReply](104, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) ListIODriverInfo(c ListIODriverInfoCall) rpc.Procedure[ListIODriverInfoCall, ListIODriverInfoReply] {
	return rpc.NewProcedure[ListIODriverInfoCall, ListIODriverInfoReply](106, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) LogInfo(c LogInfoCall) rpc.Procedure[LogInfoCall, LogInfoReply] {
	return rpc.NewProcedure[LogInfoCall, LogInfoReply](108, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) ListTaskInfo(c ListTaskInfoCall) rpc.Procedure[ListTaskInfoCall, ListTaskInfoReply] {
	return rpc.NewProcedure[ListTaskInfoCall, ListTaskInfoReply](110, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) BootInfo(c BootInfoCall) rpc.Procedure[BootInfoCall, BootInfoReply] {
	return rpc.NewProcedure[BootInfoCall, BootInfoReply](112, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsMeasureRuntime.
func (procedures) TimeMeasurementOnOff(c TimeMeasurementOnOffCall) rpc.Procedure[TimeMeasurementOnOffCall, TimeMeasurementOnOffReply] {
	return rpc.NewProcedure[TimeMeasurementOnOffCall, TimeMeasurementOnOffReply](114, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) ApplicationName(c ApplicationNameCall) rpc.Procedure[ApplicationNameCall, ApplicationNameReply] {
	return rpc.NewProcedure[ApplicationNameCall, ApplicationNameReply](116, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) ListCardInfo(c ListCardInfoCall) rpc.Procedure[ListCardInfoCall, ListCardInfoReply] {
	return rpc.NewProcedure[ListCardInfoCall, ListCardInfoReply](120, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsReadConsole.
func (procedures) ConsoleRead(c ConsoleReadCall) rpc.Procedure[ConsoleReadCall, ConsoleReadReply] {
	return rpc.NewProcedure[ConsoleReadCall, ConsoleReadReply](122, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsWriteConsole and smi.PermissionsReboot in case of reboot command.
func (procedures) ConsoleCommand(c ConsoleCommandCall) rpc.Procedure[ConsoleCommandCall, ConsoleCommandReply] {
	return rpc.NewProcedure[ConsoleCommandCall, ConsoleCommandReply](124, rpc.VersionDefault, c)
}

func (procedures) Alive(c AliveCall) rpc.Procedure[AliveCall, AliveReply] {
	return rpc.NewProcedure[AliveCall, AliveReply](126, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (procedures) CPUUsage(c CPUUsageCall) rpc.Procedure[CPUUsageCall, CPUUsageReply] {
	return rpc.NewProcedure[CPUUsageCall, CPUUsageReply](128, rpc.VersionDefault, c)
}

// Required permission: smi.PermissionsMeasureRuntime.
func (procedures) CPUUsageMeasurementOnOff(c CPUUsageMeasurementOnOffCall) rpc.Procedure[CPUUsageMeasurementOnOffCall, CPUUsageMeasurementOnOffReply] {
	return rpc.NewProcedure[CPUUsageMeasurementOnOffCall, CPUUsageMeasurementOnOffReply](130, rpc.VersionDefault, c)
}

// TODO: Finish list
