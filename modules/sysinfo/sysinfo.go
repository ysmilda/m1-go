// package info contains the interface definition of the System Info Handler
// nolint:lll
package sysinfo

import (
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/res"
)

func NewProcedures(client *m1client.Client) *Procedures {
	return &Procedures{
		client: client,
	}
}

type Procedures struct {
	client *m1client.Client
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) GetCPUAddresses(module res.ModuleNumber, c GetCPUAddressesCall) (*GetCPUAddressesReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetCPUAddressesCall, GetCPUAddressesReply](100, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) CPUInfo(module res.ModuleNumber, c CPUInfoCall) (*CPUInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CPUInfoCall, CPUInfoReply](102, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) GetSystemObjectInfo(module res.ModuleNumber, c GetSystemObjectInfoCall) (*GetSystemObjectInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetSystemObjectInfoCall, GetSystemObjectInfoReply](104, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) LogInfo(module res.ModuleNumber, c LogInfoCall) (*LogInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LogInfoCall, LogInfoReply](108, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) BootInfo(module res.ModuleNumber, c BootInfoCall) (*BootInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[BootInfoCall, BootInfoReply](112, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsMeasureRuntime.
func (p *Procedures) TimeMeasurementOnOff(module res.ModuleNumber, c TimeMeasurementOnOffCall) (*TimeMeasurementOnOffReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[TimeMeasurementOnOffCall, TimeMeasurementOnOffReply](114, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) ApplicationName(module res.ModuleNumber, c ApplicationNameCall) (*ApplicationNameReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ApplicationNameCall, ApplicationNameReply](116, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsReadConsole.
func (p *Procedures) ConsoleRead(module res.ModuleNumber, c ConsoleReadCall) (*ConsoleReadReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ConsoleReadCall, ConsoleReadReply](122, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsWriteConsole and smi.PermissionsReboot in case of reboot command.
func (p *Procedures) ConsoleCommand(module res.ModuleNumber, c ConsoleCommandCall) (*ConsoleCommandReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ConsoleCommandCall, ConsoleCommandReply](124, rpc.VersionDefault, c))
}

func (p *Procedures) Alive(module res.ModuleNumber, c AliveCall) (*AliveReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[AliveCall, AliveReply](126, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) CPUUsage(module res.ModuleNumber, c CPUUsageCall) (*CPUUsageReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CPUUsageCall, CPUUsageReply](128, rpc.VersionDefault, c))
}

// Required permission: smi.PermissionsMeasureRuntime.
func (p *Procedures) CPUUsageMeasurementOnOff(module res.ModuleNumber, c CPUUsageMeasurementOnOffCall) (*CPUUsageMeasurementOnOffReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CPUUsageMeasurementOnOffCall, CPUUsageMeasurementOnOffReply](130, rpc.VersionDefault, c))
}

// TODO: Finish list

// --------------
// Paginated procedures
// --------------

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) IODriverInfo(module res.ModuleNumber, c *IODriverInfoCall, pageSize uint32) ([]IODriverInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewPaginatedProcedure[IODriverInfo, *IODriverInfoCall, *IODriverInfoReply](106, rpc.VersionDefault, c), pageSize)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) TaskInfo(module res.ModuleNumber, c *TaskInfoCall, pageSize uint32) ([]TaskInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewPaginatedProcedure[TaskInfo, *TaskInfoCall, *TaskInfoReply](110, rpc.VersionDefault, c), pageSize)
}

// Required permission: smi.PermissionsQuerySystemInfo.
func (p *Procedures) CardInfo(module res.ModuleNumber, c *CardInfoCall, pageSize uint32) ([]CardInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewPaginatedProcedure[CardInfo, *CardInfoCall, *CardInfoReply](120, rpc.VersionDefault, c), pageSize)
}
