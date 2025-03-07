// package info contains the interface definition of the System Info Handler
// nolint:lll
package sysinfo

import (
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/res"
)

func NewProcedures(client *m1client.Client, module res.ModuleNumber) *Procedures {
	return &Procedures{
		client: client,
		module: module,
	}
}

type Procedures struct {
	client *m1client.Client
	module res.ModuleNumber
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) GetCPUAddresses(c GetCPUAddressesCall) (*GetCPUAddressesReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[GetCPUAddressesCall, GetCPUAddressesReply](100, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) CPUInfo(c CPUInfoCall) (*CPUInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[CPUInfoCall, CPUInfoReply](102, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) GetSystemObjectInfo(c GetSystemObjectInfoCall) (*GetSystemObjectInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[GetSystemObjectInfoCall, GetSystemObjectInfoReply](104, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) LogInfo(c LogInfoCall) (*LogInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[LogInfoCall, LogInfoReply](108, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) BootInfo(c BootInfoCall) (*BootInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[BootInfoCall, BootInfoReply](112, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsMeasureRuntime.
func (p *Procedures) TimeMeasurementOnOff(c TimeMeasurementOnOffCall) (*TimeMeasurementOnOffReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[TimeMeasurementOnOffCall, TimeMeasurementOnOffReply](114, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) ApplicationName(c ApplicationNameCall) (*ApplicationNameReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[ApplicationNameCall, ApplicationNameReply](116, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsReadConsole.
func (p *Procedures) ConsoleRead(c ConsoleReadCall) (*ConsoleReadReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[ConsoleReadCall, ConsoleReadReply](122, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsWriteConsole and smi.PermissionsReboot in case of reboot command.
func (p *Procedures) ConsoleCommand(c ConsoleCommandCall) (*ConsoleCommandReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[ConsoleCommandCall, ConsoleCommandReply](124, rpc.VersionDefault, c))
}

func (p *Procedures) Alive(c AliveCall) (*AliveReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[AliveCall, AliveReply](126, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) CPUUsage(c CPUUsageCall) (*CPUUsageReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[CPUUsageCall, CPUUsageReply](128, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsMeasureRuntime.
func (p *Procedures) CPUUsageMeasurementOnOff(c CPUUsageMeasurementOnOffCall) (*CPUUsageMeasurementOnOffReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewProcedure[CPUUsageMeasurementOnOffCall, CPUUsageMeasurementOnOffReply](130, rpc.VersionDefault, c))
}

// TODO: Finish list

// --------------
// Paginated procedures
// --------------

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) IODriverInfo(c *IODriverInfoCall, pageSize uint32) ([]IODriverInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewPaginatedProcedure[IODriverInfo, *IODriverInfoCall, *IODriverInfoReply](106, rpc.VersionDefault, c), pageSize)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) TaskInfo(c *TaskInfoCall, pageSize uint32) ([]TaskInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewPaginatedProcedure[TaskInfo, *TaskInfoCall, *TaskInfoReply](110, rpc.VersionDefault, c), pageSize)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) CardInfo(c *CardInfoCall, pageSize uint32) ([]CardInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: p.module.ModuleNumber, Port: p.module.Port}, rpc.NewPaginatedProcedure[CardInfo, *CardInfoCall, *CardInfoReply](120, rpc.VersionDefault, c), pageSize)
}
