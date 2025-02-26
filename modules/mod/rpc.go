//nolint:unused
package mod

import (
	"net"
	"time"

	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/msys"
)

type (
	InstallModuleCall struct {
		Config ModuleConfig
	}

	InstallModuleReply struct {
		rpc.ReturnCode
	}

	ExtendedInstallModuleCall struct {
		Config ModuleConfigExtended
	}

	ExtendedInstallModuleReply struct {
		rpc.ReturnCode
	}

	InstallJavaModuleCall struct {
		Config    ModuleConfig
		MainClass string `m1binary:"length:84"`
	}

	InstallJavaModuleReply struct {
		rpc.ReturnCode
	}

	RemoveModuleCall struct {
		AppName string `m1binary:"length:12"`
	}

	RemoveModuleReply struct {
		rpc.ReturnCode
	}

	ChangeMConfigCall struct {
		Section    string `m1binary:"length:16"`
		Group      string `m1binary:"length:16"`
		Status     SystemConfigStatus
		LineNumber int32 // Line number in config file, 0 = start of file
	}

	ChangeMConfigReply struct {
		rpc.ReturnCode
		Reboot     RebootMode
		LineNumber int32 // Line number in config file, 0 = start of file
	}

	CopyMConfigCall struct {
		Mode CopyConfigMode
	}

	CopyMConfigReply struct {
		rpc.ReturnCode
	}

	CopyMConfig2Call struct {
		Mode        CopyConfigMode
		HandleID    uint32
		ReferenceID uint32 `m1binary:"skip:4"`
		Filename    string `m1binary:"length:84,skip:92"`
	}

	CopyMConfig2Reply struct {
		rpc.ReturnCode
		HandleID       uint32
		ReferenceID    uint32
		Progress       uint8 // 0-100%
		ConfigFileLock bool
		InOperation    bool   `m1binary:"skip:1"`
		Filename       string `m1binary:"length:84,skip:92"`
	}

	CopyFileCall struct {
		Mode            CopyFileMode
		SourceName      string `m1binary:"length:84"`
		DestinationName string `m1binary:"length:84"`
	}

	CopyFileReply struct {
		rpc.ReturnCode
		Filesize int32
	}

	CopyFile2Call struct {
		Mode            CopyFileMode
		SourceName      string `m1binary:"length:84"`
		DestinationName string `m1binary:"length:84"`
		ID              uint32
	}

	CopyFile2Reply struct {
		rpc.ReturnCode
		ID       uint32
		Filesize int32
		Progress int8 `m1binary:"skip:7"` // 0 - 100%
	}

	LockObjectCall struct {
		Mode      LockObjectMode
		Object    int32
		CallerIP  net.IP `m1binary:"length:4"`
		SessionID uint32
	}

	LockObjectReply struct {
		rpc.ReturnCode
	}

	SetTimeCall struct {
		Timestamp msys.Timestamp `m1binary:"skip:128"`
	}

	SetTimeReply struct {
		rpc.ReturnCode
	}

	SetTimezoneCall struct {
		// Fixed offset using GMT/UTC prefix + offset to UTC ("UTC+01:00") or Olson timezone format ("Europe/Vienna").
		Timezone string `m1binary:"length:36"`
		Remanent bool   `m1binary:"skip:127"`
	}

	SetTimezoneReply struct {
		rpc.ReturnCode
	}

	SetDateCall struct {
		Date      msys.DateTime
		GMTOffset uint64 // Offset from UTC in seconds
	}

	SetDateReply struct {
		rpc.ReturnCode
	}

	GetBootParametersCall struct {
		mode int32 // Must be 0
	}

	GetBootParametersReply struct {
		rpc.ReturnCode
		Parameters msys.BootInfo
	}

	SetBootParametersCall struct {
		Parameters msys.BootInfo
	}

	SetBootParametersReply struct {
		rpc.ReturnCode
		Reboot RebootMode
	}

	ResetNVRamCall struct {
		Mode NVRamResetMode
	}

	ResetNVRamReply struct {
		rpc.ReturnCode
		Reboot RebootMode
	}

	RebootCall struct {
		mode int32 // Must be 0
	}

	RebootReply struct {
		rpc.ReturnCode
	}

	FormatCall struct {
		DeviceName string `m1binary:"length:84"`
		DeviceSize uint32 // 0 = default
		Mode       FormatMode
	}

	FormatReply struct {
		rpc.ReturnCode
		Reboot     RebootMode
		DeviceSize uint32
	}

	Format64Call struct {
		DeviceName string     `m1binary:"length:84"`
		DeviceSize uint32     // 0 = default
		Mode       FormatMode `m1binary:"skip:128"`
	}

	Format64Reply struct {
		rpc.ReturnCode
		Reboot     RebootMode
		DeviceSize uint32
	}

	GetDosFileSystemInfoCall struct {
		VolumeName string `m1binary:"length:84,skip:172"`
	}

	GetDosFileSystemInfoReply struct {
		rpc.ReturnCode
		Info DosFileSystemInfo `m1binary:"skip:252"`
	}

	GetDiskPartitionInfoCall struct {
		DeviceName string `m1binary:"length:84"`
		// Specify level of detail in reply.
		// 0 = disk size
		// 1 = additional disk name, serial, revision
		// 2 >= additional partitionin information
		DetailLevel uint32 `m1binary:"skip:168"`
	}

	GetDiskPartitionInfoReply struct {
		rpc.ReturnCode
		Sectors             uint32
		SectorSize          uint32 `m1binary:"skip:64"`
		DiskInfo            DiskInfo
		NumberOfPartEntries uint32
		PartEntries         []PartitionInfo `m1binary:"length:16"`
	}

	PartitionDiskCall struct {
		DeviceName         string `m1binary:"length:84"`
		NumberOfPartitions uint32 `m1binary:"skip:4"` // Number of partitions to create (1-4)
		SizeOfPartition2   uint32 // Size of partition 2 in % (0-99)
		SizeOfPartition3   uint32 // Size of partition 3 in % (0-99)
		SizeOfPartition4   uint32 // Size of partition 4 in % (0-99)
		TypePartition1     PartitionType
		TypePartition2     PartitionType
		TypePartition3     PartitionType
		TypePartition4     PartitionType
		ForcePartitioning  bool `m1binary:"skip:391"`
	}

	PartitionDiskReply struct {
		rpc.ReturnCode `m1binary:"skip:252"`
	}

	UpdateFirmwareCall struct {
		cardNumber uint32 // Must be 0.
		Object     FirmwareObjectType
		FileName   string `m1binary:"length:84"` // Absolute path of new software
	}

	UpdateFirmwareReply struct {
		rpc.ReturnCode
		Reboot RebootMode
	}

	UpdatePackageCall struct {
		Mode       UpdatePackageMode
		FileName   string `m1binary:"length:84"` // Absolute path of package file
		CardNumber uint32 `m1binary:"skip:164"`
	}

	UpdatePackageReply struct {
		rpc.ReturnCode
		Status            UpdatePackageStatus
		ProcessedSize     uint64
		TotalSize         uint64
		ProcessedSizeSub  uint64
		TotalSizeSub      uint64
		CompletedElements uint32
		TotalElements     uint32
		ElementName       string `m1binary:"length:16,skip:192"` // Filename/object currently in progress
	}

	GetFileInfoCall struct {
		Mode       GetFileInfoMode
		RemotePath string `m1binary:"length:84"`
		RemoteName string `m1binary:"length:167"`
	}

	GetFileInfoReply struct {
		rpc.ReturnCode
		Timestamp time.Time `m1binary:"length:4,unit:seconds"`
		FileSize  uint32    `m1binary:"skip:8"`
	}

	ResetAllModulesCall struct {
		replyRequested uint32 // As we can't handle no reply, leave this at 0
	}

	ResetAllModulesReply struct {
		rpc.ReturnCode
	}

	ProgressCall struct {
		DeviceName string `m1binary:"length:84,skip:4"`
	}

	ProgressReply struct {
		rpc.ReturnCode
		Reboot        RebootMode
		ProcessedSize uint32
		TotalSize     uint32
	}

	Progress64Call struct {
		DeviceName string `m1binary:"length:84,skip:128"`
	}

	Progress64Reply struct {
		rpc.ReturnCode
		Reboot        RebootMode
		ProcessedSize uint64
		TotalSize     uint64 `m1binary:"skip:128"`
	}

	SetMConfigPathCall struct {
		DeviceName string `m1binary:"length:84,skip:4"`
	}

	SetMConfigPathReply struct {
		rpc.ReturnCode
	}

	CheckFilenameCall struct {
		mode     int32  // Must be 0
		FileName string `m1binary:"length:256,skip:16"`
	}

	CheckFilenameReply struct {
		rpc.ReturnCode
		FileName  string `m1binary:"length:256,skip:16"`
		ErrorCode int32  `m1binary:"skip:12"`
	}
)
