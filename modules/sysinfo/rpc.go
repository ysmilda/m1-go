//nolint:unused
package sysinfo

import (
	"net"

	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/msys"
)

type (
	CardInfoCall struct {
		rpc.PaginatedCallFirstLast
		Filter uint32 // Must be zero
	}

	CardInfoReply struct {
		rpc.ReturnCode
		rpc.PaginatedReplyLast[CardInfo]
	}

	GetCPUAddressesCall struct {
		Parameter uint32 // Must be zero
	}

	GetCPUAddressesReply struct {
		rpc.ReturnCode
		CurrentIndex uint32
		Count        uint32
		Addresses    []CPUAddress `binary:"lengthRef:Count"`
	}

	CPUInfoCall struct {
		Parameter uint32 // Must be zero
	}

	CPUInfoReply struct {
		rpc.ReturnCode
		CPUInfo
	}

	GetSystemObjectInfoCall struct {
		Parameter uint32 // Must be zero
	}

	GetSystemObjectInfoReply struct {
		rpc.ReturnCode
		Count  uint32
		Object []SystemObjectInfo `binary:"lengthRef:Count"`
	}

	IODriverInfoCall struct {
		rpc.PaginatedCallFirstLast
	}

	IODriverInfoReply struct {
		rpc.ReturnCode
		rpc.PaginatedReplyCount[IODriverInfo]
	}

	LogInfoCall struct {
		Parameter uint32 // Must be zero
	}

	LogInfoReply struct {
		rpc.ReturnCode
		NumberOfEntries   uint32
		LogFileSize       uint32
		SizeOfOneLogEntry uint32
		Filename          string `binary:"length:84"`
	}

	TaskInfoCall struct {
		rpc.PaginatedCallFirstLast
		Spare uint32 // Must be zero
	}

	TaskInfoReply struct {
		rpc.ReturnCode
		Last bool `binary:"skip:3"`
		// TODO: With the current PaginatedCall implementation these values are thrown away, is there a way to keep them?
		TimeTotal int64  // clockcycles since power up
		TimeUnits uint32 // clockcyles per microsecond
		rpc.PaginatedReplyCount[TaskInfo]
	}

	ListExtendedTaskInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32
		Spare      uint32 // Must be zero
	}

	ListExtendedTaskInfoReply struct {
		rpc.ReturnCode
		Last          bool   `binary:"skip:3"`
		TimeTotal     int64  // clockcycles since power up
		TimeUnits     uint32 // clockcyles per microsecond
		NumberOfCores uint32 `binary:"skip:16"`
		Count         uint32
		Objects       []TaskInfo `binary:"lengthRef:Count"`
	}

	BootInfoCall struct {
		Parameter uint32 // Must be zero
	}

	BootInfoReply struct {
		rpc.ReturnCode
		msys.BootInfo
	}

	BootInfo2Call struct {
		Parameter uint32 // Must be zero
	}

	BootInfo2Reply struct {
		rpc.ReturnCode
		BootDevice         string `binary:"length:20"`
		HostName           string `binary:"length:20"`
		CorePath           string `binary:"length:80"`
		MsysPath           string `binary:"length:80"`
		ConfigPath         string `binary:"length:80"`
		CoreFallbackUsed   bool   `binary:"skip:3"`
		MsysFallbackUsed   bool   `binary:"skip:3"`
		ConfigFallbackUsed bool   `binary:"skip:3"`
		ConfigBackupUsed   bool   `binary:"skip:503"`
	}

	TimeMeasurementOnOffCall struct {
		Setting TimeOnOffSetting
	}

	TimeMeasurementOnOffReply struct {
		rpc.ReturnCode
		Enabled bool `binary:"skip:3"`
	}

	ApplicationNameCall struct {
		Parameter uint32 // Must be zero
	}

	ApplicationNameReply struct {
		rpc.ReturnCode
		Name string `binary:"length:24,skip:80"`
	}

	ConsoleReadCall struct {
		BufferSize uint32
		Spare      uint32 // Must be zero
	}

	ConsoleReadReply struct {
		rpc.ReturnCode
		Count uint32 `binary:"skip:4"`
		Text  string `binary:"lengthRef:Count"`
	}

	ConsoleCommandCall struct {
		Command string `binary:"length:256,zeroTerminated"`
	}

	ConsoleCommandReply struct {
		rpc.ReturnCode
	}

	AliveCall struct {
		CallerIP  net.IP `binary:"length:4"`
		SessionID uint32
		// Currently only support NORMALMODE (e.g. I am alive) command.
		// The others return a value instead of the returncode, which is not supported.
		Mode           uint32
		LoginSessionID uint32
	}

	AliveReply struct {
		rpc.ReturnCode
	}

	CPUUsageCall struct {
		Mode CPUUsageMeasurementMode `binary:"skip:8"`
	}

	CPUUsageReply struct {
		rpc.ReturnCode
		TimeUnits            uint32 // Clockcycles per microseconds
		Usage                CoreUsage
		TaskGroupMaxRuntime  []uint32 `binary:"length:10"`
		TaskGroupAllowedTime []uint32 `binary:"length:10,skip:16"`
	}

	CPUUsageMeasurementOnOffCall struct {
		Setting CPUUsageOnOffSetting
	}

	CPUUsageMeasurementOnOffReply struct {
		rpc.ReturnCode
		Enabled bool `binary:"skip:3"`
	}

	ListObjectInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32
	}

	ListObjectInfoReply struct {
		rpc.ReturnCode
		Last    bool         `binary:"skip:3"`
		Count   uint32       `binary:"skip:12"`
		Objects []ObjectInfo `binary:"lengthRef:Count"`
	}

	ListLibraryInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32 `binary:"skip:8"`
	}

	ListLibraryInfoReply struct {
		rpc.ReturnCode
		Count   uint32
		Objects []LibraryInfo `binary:"lengthRef:Count"`
	}

	ListServiceInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32 `binary:"skip:4"`
	}

	ListServiceInfoReply struct {
		rpc.ReturnCode
		Count   uint32
		Objects []ServiceInfo `binary:"lengthRef:Count"`
	}

	SystemConfigCall struct {
		Spare uint32 `binary:"skip:16"`
	}

	SystemConfigReply struct {
		rpc.ReturnCode
		LongFilenames                 bool
		ApplicationDevelopmentAllowed bool   `binary:"skip:8"`
		SystemDirectory               string `binary:"length:84"`
		DriverDirectory               string `binary:"length:84"`
		ServiceDirectory              string `binary:"length:84"`
		ApplicationDirectory          string `binary:"length:84"`
		ClassesDirectory              string `binary:"length:84"` // deprecated since MSys v4.00
		NonVolatileDirectory          string `binary:"length:84"`
		TemporaryDirectory            string `binary:"length:84"`
		KeysDirectory                 string `binary:"length:84"`
		ConfigurationDirectory        string `binary:"length:84"`
		KeysBackupDirectory           string `binary:"length:84,skip:232"`
	}

	CoreUsageCall struct {
		Spare uint32 `binary:"skip:8"`
	}

	CoreUsageReply struct {
		rpc.ReturnCode
		TimeUnits uint32 `binary:"skip:20"` // Clockcycles per microseconds
		Count     uint32
		Cores     []CoreUsage `binary:"lengthRef:Count"`
	}

	ListConfigurationInfoCall struct {
		StartIndex uint32 `binary:"skip:40"`
	}

	ListConfigurationInfoReply struct {
		rpc.ReturnCode
		TotalConfigurations uint32
		HasMore             bool
		Count               uint32              `binary:"skip:40"`
		Configurations      []ConfigurationInfo `binary:"lengthRef:Count"`
	}
)

func (t TaskInfoReply) Done(uint32) bool {
	return t.Last
}
