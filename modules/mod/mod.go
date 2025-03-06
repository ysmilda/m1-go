// package mod contains the interface definition of the Module Handler.
//
//nolint:lll
package mod

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

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) InstallModule(module res.ModuleNumber, c InstallModuleCall) (*InstallModuleReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[InstallModuleCall, InstallModuleReply](100, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) InstallJavaModule(module res.ModuleNumber, c InstallJavaModuleCall) (*InstallJavaModuleReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[InstallJavaModuleCall, InstallJavaModuleReply](102, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) RemoveModule(module res.ModuleNumber, c RemoveModuleCall) (*RemoveModuleReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[RemoveModuleCall, RemoveModuleReply](106, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsWriteMConfig.
func (p *Procedures) ChangeMConfig(module res.ModuleNumber, c ChangeMConfigCall) (*ChangeMConfigReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ChangeMConfigCall, ChangeMConfigReply](110, rpc.VersionDefault, c))
}

// Copy the MConfig.ini from the device to /ram0 or vice versa.
func (p *Procedures) CopyMConfig(module res.ModuleNumber, c CopyMConfigCall) (*CopyMConfigReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CopyMConfigCall, CopyMConfigReply](112, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsReadConsole and smi.PermissionsWriteConsole.
func (p *Procedures) LockObject(module res.ModuleNumber, c LockObjectCall) (*LockObjectReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LockObjectCall, LockObjectReply](114, rpc.VersionDefault, c))
}

func (p *Procedures) CopyFile(module res.ModuleNumber, c CopyFileCall) (*CopyFileReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CopyFileCall, CopyFileReply](116, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsSetDateTime.
func (p *Procedures) SetTime(module res.ModuleNumber, c SetTimeCall) (*SetTimeReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetTimeCall, SetTimeReply](122, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsSetDateTime.
func (p *Procedures) SetTimezone(module res.ModuleNumber, c SetTimezoneCall) (*SetTimezoneReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetTimezoneCall, SetTimezoneReply](124, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsSetDateTime.
func (p *Procedures) SetDate(module res.ModuleNumber, c SetDateCall) (*SetDateReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetDateCall, SetDateReply](126, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsWriteMConfig.
func (p *Procedures) GetBootParameters(module res.ModuleNumber, c GetBootParametersCall) (*GetBootParametersReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetBootParametersCall, GetBootParametersReply](128, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsWriteMConfig.
func (p *Procedures) SetBootParameters(module res.ModuleNumber, c SetBootParametersCall) (*SetBootParametersReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetBootParametersCall, SetBootParametersReply](130, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsFormat.
func (p *Procedures) ResetNVRam(module res.ModuleNumber, c ResetNVRamCall) (*ResetNVRamReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ResetNVRamCall, ResetNVRamReply](132, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsReboot.
func (p *Procedures) Reboot(module res.ModuleNumber, c RebootCall) (*RebootReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[RebootCall, RebootReply](134, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsFormat and in case of targeting the bootdevice smi.PermissionsWriteToBootDevice.
func (p *Procedures) Format(module res.ModuleNumber, c FormatCall) (*FormatReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[FormatCall, FormatReply](136, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsUpdateFirmware.
func (p *Procedures) UpdateFirmware(module res.ModuleNumber, c UpdateFirmwareCall) (*UpdateFirmwareReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[UpdateFirmwareCall, UpdateFirmwareReply](138, rpc.VersionDefault, c))
}

func (p *Procedures) GetFileInfo(module res.ModuleNumber, c GetFileInfoCall) (*GetFileInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetFileInfoCall, GetFileInfoReply](140, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) ResetAllModules(module res.ModuleNumber, c ResetAllModulesCall) (*ResetAllModulesReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ResetAllModulesCall, ResetAllModulesReply](142, rpc.VersionDefault, c))
}

func (p *Procedures) Progress(module res.ModuleNumber, c ProgressCall) (*ProgressReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ProgressCall, ProgressReply](144, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsWriteMConfig.
func (p *Procedures) SetMConfigPath(module res.ModuleNumber, c SetMConfigPathCall) (*SetMConfigPathReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetMConfigPathCall, SetMConfigPathReply](146, rpc.VersionDefault, c))
}

func (p *Procedures) CheckFilename(module res.ModuleNumber, c CheckFilenameCall) (*CheckFilenameReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CheckFilenameCall, CheckFilenameReply](148, rpc.VersionDefault, c))
}

// Copy the MConfig.ini from the device to /ram0 or vice versa for multi user.
func (p *Procedures) CopyMConfig2(module res.ModuleNumber, c CopyMConfig2Call) (*CopyMConfig2Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CopyMConfig2Call, CopyMConfig2Reply](150, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsFormat and in case of targeting the bootdevice smi.PermissionsWriteToBootDevice.
func (p *Procedures) Format64(module res.ModuleNumber, c Format64Call) (*Format64Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[Format64Call, Format64Reply](154, rpc.VersionDefault, c))
}

func (p *Procedures) Progress64(module res.ModuleNumber, c Progress64Call) (*Progress64Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[Progress64Call, Progress64Reply](156, rpc.VersionDefault, c))
}

func (p *Procedures) GetDiskPartitionInfo(module res.ModuleNumber, c GetDiskPartitionInfoCall) (*GetBootParametersReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetDiskPartitionInfoCall, GetBootParametersReply](158, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsFormat and in case of targeting the bootdevice smi.PermissionsWriteToBootDevice.
func (p *Procedures) PartitionDisk(module res.ModuleNumber, c PartitionDiskCall) (*PartitionDiskReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[PartitionDiskCall, PartitionDiskReply](160, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsUpdateFirmware.
func (p *Procedures) UpdatePackage(module res.ModuleNumber, c UpdatePackageCall) (*UpdatePackageReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[UpdatePackageCall, UpdatePackageReply](162, rpc.VersionDefault, c))
}

func (p *Procedures) GetDosFileSystemInfo(module res.ModuleNumber, c GetDosFileSystemInfoCall) (*GetDosFileSystemInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetDosFileSystemInfoCall, GetDosFileSystemInfoReply](164, rpc.VersionDefault, c))
}

func (p *Procedures) CopyFile2(module res.ModuleNumber, c CopyFile2Call) (*CopyFile2Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[CopyFile2Call, CopyFile2Reply](166, rpc.VersionDefault, c))
}

func (p *Procedures) ExtendedInstallModule(module res.ModuleNumber, c ExtendedInstallModuleCall) (*ExtendedInstallModuleReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedInstallModuleCall, ExtendedInstallModuleReply](168, rpc.VersionDefault, c))
}
