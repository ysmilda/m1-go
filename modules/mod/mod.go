// package mod contains the interface definition of the Module Handler.
//
//nolint:lll
package mod

import "github.com/ysmilda/m1-go/internals/rpc"

var Procedures procedures

type procedures struct{}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) InstallModule(c InstallModuleCall) rpc.Procedure[InstallModuleCall, InstallModuleReply] {
	return rpc.NewProcedure[InstallModuleCall, InstallModuleReply](100, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) InstallJavaModule(c InstallJavaModuleCall) rpc.Procedure[InstallJavaModuleCall, InstallJavaModuleReply] {
	return rpc.NewProcedure[InstallJavaModuleCall, InstallJavaModuleReply](102, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) RemoveModule(c RemoveModuleCall) rpc.Procedure[RemoveModuleCall, RemoveModuleReply] {
	return rpc.NewProcedure[RemoveModuleCall, RemoveModuleReply](106, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsWriteMConfig.
func (procedures) ChangeMConfig(c ChangeMConfigCall) rpc.Procedure[ChangeMConfigCall, ChangeMConfigReply] {
	return rpc.NewProcedure[ChangeMConfigCall, ChangeMConfigReply](110, rpc.VersionDefault, c)
}

// Copy the MConfig.ini from the device to /ram0 or vice versa.
func (procedures) CopyMConfig(c CopyMConfigCall) rpc.Procedure[CopyMConfigCall, CopyMConfigReply] {
	return rpc.NewProcedure[CopyMConfigCall, CopyMConfigReply](112, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsReadConsole and smi.PermissionsWriteConsole.
func (procedures) LockObject(c LockObjectCall) rpc.Procedure[LockObjectCall, LockObjectReply] {
	return rpc.NewProcedure[LockObjectCall, LockObjectReply](114, rpc.VersionDefault, c)
}

func (procedures) CopyFile(c CopyFileCall) rpc.Procedure[CopyFileCall, CopyFileReply] {
	return rpc.NewProcedure[CopyFileCall, CopyFileReply](116, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsSetDateTime.
func (procedures) SetTime(c SetTimeCall) rpc.Procedure[SetTimeCall, SetTimeReply] {
	return rpc.NewProcedure[SetTimeCall, SetTimeReply](122, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsSetDateTime.
func (procedures) SetTimezone(c SetTimezoneCall) rpc.Procedure[SetTimezoneCall, SetTimezoneReply] {
	return rpc.NewProcedure[SetTimezoneCall, SetTimezoneReply](124, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsSetDateTime.
func (procedures) SetDate(c SetDateCall) rpc.Procedure[SetDateCall, SetDateReply] {
	return rpc.NewProcedure[SetDateCall, SetDateReply](126, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsWriteMConfig.
func (procedures) GetBootParameters(c GetBootParametersCall) rpc.Procedure[GetBootParametersCall, GetBootParametersReply] {
	return rpc.NewProcedure[GetBootParametersCall, GetBootParametersReply](128, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsWriteMConfig.
func (procedures) SetBootParameters(c SetBootParametersCall) rpc.Procedure[SetBootParametersCall, SetBootParametersReply] {
	return rpc.NewProcedure[SetBootParametersCall, SetBootParametersReply](130, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsFormat.
func (procedures) ResetNVRam(c ResetNVRamCall) rpc.Procedure[ResetNVRamCall, ResetNVRamReply] {
	return rpc.NewProcedure[ResetNVRamCall, ResetNVRamReply](132, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsReboot.
func (procedures) Reboot(c RebootCall) rpc.Procedure[RebootCall, RebootReply] {
	return rpc.NewProcedure[RebootCall, RebootReply](134, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsFormat and in case of targeting the bootdevice smi.PermissionsWriteToBootDevice.
func (procedures) Format(c FormatCall) rpc.Procedure[FormatCall, FormatReply] {
	return rpc.NewProcedure[FormatCall, FormatReply](136, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsUpdateFirmware.
func (procedures) UpdateFirmware(c UpdateFirmwareCall) rpc.Procedure[UpdateFirmwareCall, UpdateFirmwareReply] {
	return rpc.NewProcedure[UpdateFirmwareCall, UpdateFirmwareReply](138, rpc.VersionDefault, c)
}

func (procedures) GetFileInfo(c GetFileInfoCall) rpc.Procedure[GetFileInfoCall, GetFileInfoReply] {
	return rpc.NewProcedure[GetFileInfoCall, GetFileInfoReply](140, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) ResetAllModules(c ResetAllModulesCall) rpc.Procedure[ResetAllModulesCall, ResetAllModulesReply] {
	return rpc.NewProcedure[ResetAllModulesCall, ResetAllModulesReply](142, rpc.VersionDefault, c)
}

func (procedures) Progress(c ProgressCall) rpc.Procedure[ProgressCall, ProgressReply] {
	return rpc.NewProcedure[ProgressCall, ProgressReply](144, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsWriteMConfig.
func (procedures) SetMConfigPath(c SetMConfigPathCall) rpc.Procedure[SetMConfigPathCall, SetMConfigPathReply] {
	return rpc.NewProcedure[SetMConfigPathCall, SetMConfigPathReply](146, rpc.VersionDefault, c)
}

func (procedures) CheckFilename(c CheckFilenameCall) rpc.Procedure[CheckFilenameCall, CheckFilenameReply] {
	return rpc.NewProcedure[CheckFilenameCall, CheckFilenameReply](148, rpc.VersionDefault, c)
}

// Copy the MConfig.ini from the device to /ram0 or vice versa for multi user.
func (procedures) CopyMConfig2(c CopyMConfig2Call) rpc.Procedure[CopyMConfig2Call, CopyMConfig2Reply] {
	return rpc.NewProcedure[CopyMConfig2Call, CopyMConfig2Reply](150, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsFormat and in case of targeting the bootdevice smi.PermissionsWriteToBootDevice.
func (procedures) Format64(c Format64Call) rpc.Procedure[Format64Call, Format64Reply] {
	return rpc.NewProcedure[Format64Call, Format64Reply](154, rpc.VersionDefault, c)
}

func (procedures) Progress64(c Progress64Call) rpc.Procedure[Progress64Call, Progress64Reply] {
	return rpc.NewProcedure[Progress64Call, Progress64Reply](156, rpc.VersionDefault, c)
}

func (procedures) GetDiskPartitionInfo(c GetDiskPartitionInfoCall) rpc.Procedure[GetDiskPartitionInfoCall, GetBootParametersReply] {
	return rpc.NewProcedure[GetDiskPartitionInfoCall, GetBootParametersReply](158, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsFormat and in case of targeting the bootdevice smi.PermissionsWriteToBootDevice.
func (procedures) PartitionDisk(c PartitionDiskCall) rpc.Procedure[PartitionDiskCall, PartitionDiskReply] {
	return rpc.NewProcedure[PartitionDiskCall, PartitionDiskReply](160, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsUpdateFirmware.
func (procedures) UpdatePackage(c UpdatePackageCall) rpc.Procedure[UpdatePackageCall, UpdatePackageReply] {
	return rpc.NewProcedure[UpdatePackageCall, UpdatePackageReply](162, rpc.VersionDefault, c)
}

func (procedures) GetDosFileSystemInfo(c GetDosFileSystemInfoCall) rpc.Procedure[GetDosFileSystemInfoCall, GetDosFileSystemInfoReply] {
	return rpc.NewProcedure[GetDosFileSystemInfoCall, GetDosFileSystemInfoReply](164, rpc.VersionDefault, c)
}

func (procedures) CopyFile2(c CopyFile2Call) rpc.Procedure[CopyFile2Call, CopyFile2Reply] {
	return rpc.NewProcedure[CopyFile2Call, CopyFile2Reply](166, rpc.VersionDefault, c)
}

func (procedures) ExtendedInstallModule(c ExtendedInstallModuleCall) rpc.Procedure[ExtendedInstallModuleCall, ExtendedInstallModuleReply] {
	return rpc.NewProcedure[ExtendedInstallModuleCall, ExtendedInstallModuleReply](168, rpc.VersionDefault, c)
}
