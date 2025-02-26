package mod

type FileType uint32

const (
	FileTypeIODriver        FileType = 1
	FileTypeRegulator       FileType = 2
	FileTypeVxWorksModule   FileType = 2
	FileTypePLC1131         FileType = 3
	FileTypeOther           FileType = 4
	FileTypePLCMLibrary     FileType = 5
	FileTypeJavaModule      FileType = 6 // deprecated since MSys 4.00R
	FileTypeSoftwareService FileType = 7
	FileTypeVxWorksLibrary  FileType = 8
	FileTypeIOFirmware      FileType = 9
	FileTypeMTarget         FileType = 10
)

type ObjectType uint32

const (
	ObjectTypeOjectCode                     ObjectType = 1
	ObjectTypeConfig                        ObjectType = 2
	ObjectTypeLibrary                       ObjectType = 3
	ObjectTypeAttributes                    ObjectType = 4
	ObjectTypeBinaryConfigurationRule       ObjectType = 5
	ObjectTypeJAR                           ObjectType = 6 // deprecated since MSys 4.00R
	ObjectTypeHelpInfo                      ObjectType = 10
	ObjectTypeReadmeInfo                    ObjectType = 11
	ObjectTypeSource                        ObjectType = 12
	ObjectTypeInternalLibrary               ObjectType = 13
	ObjectTypePackedCProjectSource          ObjectType = 14
	ObjectTypeExternalLibrary               ObjectType = 15
	ObjectTypeSoftwareVersionRequirement    ObjectType = 16
	ObjectTypeSymbolDataDescription         ObjectType = 17
	ObjectTypeLogic                         ObjectType = 19
	ObjectTypeFirmware                      ObjectType = 20
	ObjectTypeExtendedValidationDescription ObjectType = 21
	ObjectTypeZippedCRU                     ObjectType = 22
	ObjectTypeMTargetProject                ObjectType = 23
	ObjectTypeZIP                           ObjectType = 24
	ObjectTypeXMLComponentInterface         ObjectType = 25
	ObjectTypeCMLCRUExtension               ObjectType = 26
	ObjectTypeCardDefinitionInformation     ObjectType = 100
	ObjectTypeUserDefinable                 ObjectType = 200
)

type ObjectMode uint32

const (
	ObjectModeNormal   ObjectMode = 1
	ObjectModeCompact  ObjectMode = 2
	ObjectModeExternal ObjectMode = 4
)

type SystemConfigStatus uint32

const (
	SystemConfigStatusChange SystemConfigStatus = 1
	SystemConfigStatusAdd    SystemConfigStatus = 2
	SystemConfigStatusDelete SystemConfigStatus = 3
)

type CopyConfigMode uint32

const (
	CopyConfigModeLoad     CopyConfigMode = 1
	CopyConfigModeStore    CopyConfigMode = 2
	CopyConfigModeProgress CopyConfigMode = 3
)

type CopyFileMode uint32

const (
	CopyFileModeCreate    CopyFileMode = 1
	CopyFileModeOverwrite CopyFileMode = 2
	CopyFileModeProgress  CopyFileMode = 3
	CopyFileModeMaxTasks  CopyFileMode = 32
)

type LockObjectMode uint32

const (
	LockObjectModeLock   LockObjectMode = 1
	LockObjectModeUnlock LockObjectMode = 2
)

type NVRamResetMode int32

const (
	NVRamResetModeNVRam0 NVRamResetMode = 0
	NVRamResetModeNVRam1 NVRamResetMode = 1
)

type FormatMode uint32

const (
	FormatModeNormal FormatMode = 1 // Run directly
	FormatModeTask   FormatMode = 2 // Run as task
)

type SmartStatus uint32

const (
	SmartStatusUnknown SmartStatus = 0
	SmartStatusPassed  SmartStatus = 1
	SmartStatusFailed  SmartStatus = 2
)

type PartitionStatus uint32

const (
	PartitionStatusIsBootable  = 0x80
	PartitionStatusNotBootable = 0x00
)

type PartitionType uint32

const (
	PartitionTypeFAT12   PartitionType = 0x01
	PartitionTypeFAT16   PartitionType = 0x06
	PartitionTypeFAT32   PartitionType = 0x0b
	PartitionTypeUnknown PartitionType = 0x0
)

type FirmwareObjectType uint32

const (
	FirmwareObjectTypeAny     FirmwareObjectType = 0
	FirmwareObjectTypeMBOOT1  FirmwareObjectType = 1
	FirmwareObjectTypeMBOOT2  FirmwareObjectType = 2
	FirmwareObjectTypeHXBIOS  FirmwareObjectType = 3
	FirmwareObjectTypeEXCORE  FirmwareObjectType = 4
	FirmwareObjectTypeEXBIOS  FirmwareObjectType = 5
	FirmwareObjectTypeMPCBOOT FirmwareObjectType = 6
	FirmwareObjectTypeMPCBIOS FirmwareObjectType = 7
	FirmwareObjectTypeMXBOOT  FirmwareObjectType = 8
	FirmwareObjectTypeMXBIOS  FirmwareObjectType = 9
	FirmwareObjectTypeTBOOT   FirmwareObjectType = 10
	FirmwareObjectTypeTBIOS   FirmwareObjectType = 11
	FirmwareObjectTypeMX2BIOS FirmwareObjectType = 12
	FirmwareObjectTypeMHBOOT  FirmwareObjectType = 13
	FirmwareObjectTypeMHBIOS  FirmwareObjectType = 14
	FirmwareObjectTypeMCBOOT  FirmwareObjectType = 15
	FirmwareObjectTypeMCBIOS  FirmwareObjectType = 16
	FirmwareObjectTypeMH2BIOS FirmwareObjectType = 17
	FirmwareObjectTypeMC2BIOS FirmwareObjectType = 18
	FirmwareObjectTypeMX3BIOS FirmwareObjectType = 19
)

type RebootMode int32

const (
	RebootModeOnline  RebootMode = 0 // Change applied immediately
	RebootModeReboot  RebootMode = 1 // Change active on next reboot
	RebootModeRestart RebootMode = 2 // Reboot initiated automatically
)

type UpdatePackageMode uint32

const (
	UpdatePackageModeProgram  UpdatePackageMode = 1
	UpdatePackageModeForce    UpdatePackageMode = 2
	UpdatePackageModeProgress UpdatePackageMode = 3
)

type GetFileInfoMode uint32

const (
	GetFileInfoModeGetAttributes          GetFileInfoMode = 1
	GetFileInfoModeEraseAttributes        GetFileInfoMode = 2
	GetFileInfoModeEraseAllFileAttributes GetFileInfoMode = 3
)

type UpdatePackageStatus int32

const (
	UpdatePackageStatusInProgress                        UpdatePackageStatus = 1
	UpdatePackageStatusDone                              UpdatePackageStatus = 0
	UpdatePackageStatusErrorFailed                       UpdatePackageStatus = -1
	UpdatePackageStatusErrorFileNotFound                 UpdatePackageStatus = -2
	UpdatePackageStatusErrorWrongModuleType              UpdatePackageStatus = -10
	UpdatePackageStatusErrorWrongModuleVariant           UpdatePackageStatus = -11
	UpdatePackageStatusErrorUnknownDevice                UpdatePackageStatus = -12
	UpdatePackageStatusErrorRunningOnReference           UpdatePackageStatus = -13
	UpdatePackageStatusErrorWrongID                      UpdatePackageStatus = -14
	UpdatePackageStatusErrorDowngrade                    UpdatePackageStatus = -15
	UpdatePackageStatusErrorNoM86Support                 UpdatePackageStatus = -16
	UpdatePackageStatusErrorMissingRequiredDriverVersion UpdatePackageStatus = -17
)
