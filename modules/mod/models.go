package mod

import (
	"time"
)

type (
	ModuleConfig struct {
		PathName        string `m1binary:"length:84"`
		TypeName        string `m1binary:"length:12"`
		AppName         string `m1binary:"length:12"`
		AppIndex        int32
		AppIdentity     uint32
		MemoryPartition uint32
		TaskPriority    int32
		DebugMode       uint32
		ProfileName     string `m1binary:"length:84"`
		LineNumber      int32  // Line number in config file, 0 = start of file
	}

	ModuleConfigExtended struct {
		ModuleConfig
		CoreCategory string `m1binary:"length:32,skip:32"`
	}

	FileHeader struct {
		FileType         FileType
		MainVersion      uint32
		SubVersion       uint32
		Timestamp        time.Time `m1binary:"length:4,unit:seconds"`
		Version          Version
		SizeChecksum     uint32
		Entries          uint32
		OffsetEntryTable uint32
		Checksum         uint32
	}

	ObjectEntry struct {
		Type    ObjectType
		Mode    ObjectMode
		Size    uint32
		Offset  uint32
		Version Version
		Name    string `m1binary:"length:14,skip:3"`
	}

	Attribute struct {
		Reentrant                    bool
		NoOnlineConfigurationAllowed bool
		NoOnlineInstallAllowed       bool
		NoOnlineDeInstallAllowed     bool
		ErrorTolerant                bool
		RetainInUse                  bool
		NoZeroRetain                 bool
		DebugCode                    bool
		DebugInfo                    bool
		NoInstallAtBoot              bool
		SystemModule                 bool
		JavaClassLoader              bool
		IncludesMIO                  bool
		HasStandardSymbols           bool
		IsComponent                  bool
		MMPReady                     bool `m1binary:"skip:16"`
	}

	DosFileSystemInfo struct {
		FatType           string `m1binary:"length:16"`
		BootVolumeLabel   string `m1binary:"length:16"`
		VolumeID          uint32
		TotalSectors      uint32
		BytesPerSector    uint32
		SectorsPerCluster uint32
		ReservedSectors   uint32
		SectorsPerFAT     uint32
		FATTables         uint32
		HiddenSectors     uint32
		DataStartSector   uint32
		Properties        uint32 `m1binary:"skip:46"`
	}

	DiskInfo struct {
		ModelName        string      `m1binary:"length:48"`
		SerialNumber     string      `m1binary:"length:32"`
		FirmwareRevision string      `m1binary:"length:16"`
		SmartStatus      SmartStatus `m1binary:"skip:156"`
	}

	PartitionInfo struct {
		Offset uint32
		Size   uint32
		Status PartitionStatus
		Type   PartitionType `m1binary:"skip:16"`
	}
)
