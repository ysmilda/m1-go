//nolint:unused
package sysinfo

import (
	"net"

	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/msys"
)

type (
	ListCardInfoCall struct {
		First  uint32
		Last   uint32
		filter uint32 // Must be zero
	}

	ListCardInfoReply struct {
		rpc.ReturnCode
		Last      bool `m1binary:"skip:3"`
		Count     uint32
		CardInfos []CardInfo `m1binary:"lengthRef:Count"`
	}

	ListCPUAddressesCall struct {
		parameter uint32 // Must be zero
	}

	ListCPUAddressesReply struct {
		rpc.ReturnCode
		CurrentIndex uint32
		Count        uint32
		Addresses    []CPUAddress `m1binary:"lengthRef:Count"`
	}

	CPUInfoCall struct {
		parameter uint32 // Must be zero
	}

	CPUInfoReply struct {
		rpc.ReturnCode
		CPUInfo
	}

	ListSystemObjectInfoCall struct {
		parameter uint32 // Must be zero
	}

	ListSystemObjectInfoReply struct {
		rpc.ReturnCode
		Count  uint32
		Object []SystemObjectInfo `m1binary:"lengthRef:Count"`
	}

	ListIODriverInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32
	}

	ListIODriverInfoReply struct {
		rpc.ReturnCode
		Count   uint32
		Objects []IODriverInfo `m1binary:"lengthRef:Count"`
	}

	LogInfoCall struct {
		parameter uint32 // Must be zero
	}

	LogInfoReply struct {
		rpc.ReturnCode
		NumberOfEntries   uint32
		LogFileSize       uint32
		SizeOfOneLogEntry uint32
		Filename          string `m1binary:"length:84"`
	}

	ListTaskInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32
		spare      uint32 // Must be zero
	}

	ListTaskInfoReply struct {
		rpc.ReturnCode
		Last      bool   `m1binary:"skip:3"`
		TimeTotal int64  // clockcycles since power up
		TimeUnits uint32 // clockcyles per microsecond
		Count     uint32
		Objects   []TaskInfo `m1binary:"lengthRef:Count"`
	}

	ListExtendedTaskInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32
		spare      uint32 // Must be zero
	}

	ListExtendedTaskInfoReply struct {
		rpc.ReturnCode
		Last          bool   `m1binary:"skip:3"`
		TimeTotal     int64  // clockcycles since power up
		TimeUnits     uint32 // clockcyles per microsecond
		NumberOfCores uint32 `m1binary:"skip:16"`
		Count         uint32
		Objects       []TaskInfo `m1binary:"lengthRef:Count"`
	}

	BootInfoCall struct {
		parameter uint32 // Must be zero
	}

	BootInfoReply struct {
		rpc.ReturnCode
		msys.BootInfo
	}

	BootInfo2Call struct {
		parameter uint32 // Must be zero
	}

	BootInfo2Reply struct {
		rpc.ReturnCode
		BootDevice         string `m1binary:"length:20"`
		HostName           string `m1binary:"length:20"`
		CorePath           string `m1binary:"length:80"`
		MsysPath           string `m1binary:"length:80"`
		ConfigPath         string `m1binary:"length:80"`
		CoreFallbackUsed   bool   `m1binary:"skip:3"`
		MsysFallbackUsed   bool   `m1binary:"skip:3"`
		ConfigFallbackUsed bool   `m1binary:"skip:3"`
		ConfigBackupUsed   bool   `m1binary:"skip:503"`
	}

	TimeMeasurementOnOffCall struct {
		Setting TimeOnOffSetting
	}

	TimeMeasurementOnOffReply struct {
		rpc.ReturnCode
		Enabled bool `m1binary:"skip:3"`
	}

	ApplicationNameCall struct {
		parameter uint32 // Must be zero
	}

	ApplicationNameReply struct {
		rpc.ReturnCode
		Name string `m1binary:"length:24,skip:80"`
	}

	ConsoleReadCall struct {
		BufferSize uint32
		spare      uint32 // Must be zero
	}

	ConsoleReadReply struct {
		rpc.ReturnCode
		Count uint32 `m1binary:"skip:4"`
		Text  string `m1binary:"lengthRef:Count"`
	}

	ConsoleCommandCall struct {
		Command string `m1binary:"length:256,zeroTarminated"`
	}

	ConsoleCommandReply struct {
		rpc.ReturnCode
	}

	AliveCall struct {
		CallerIP  net.IP `m1binary:"length:4"`
		SessionID uint32
		// Currently only support NORMALMODE (e.g. I am alive) command.
		// The others return a value instead of the returncode, which is not supported.
		mode           uint32
		LoginSessionID uint32
	}

	AliveReply struct {
		rpc.ReturnCode
	}

	CPUUsageCall struct {
		Mode CPUUsageMeasurementMode `m1binary:"skip:8"`
	}

	CPUUsageReply struct {
		rpc.ReturnCode
		TimeUnits            uint32 // Clockcycles per microseconds
		Usage                CoreUsage
		TaskGroupMaxRuntime  []uint32 `m1binary:"length:10"`
		TaskGroupAllowedTime []uint32 `m1binary:"length:10,skip:16"`
	}

	CPUUsageMeasurementOnOffCall struct {
		Setting CPUUsageOnOffSetting
	}

	CPUUsageMeasurementOnOffReply struct {
		rpc.ReturnCode
		Enabled bool `m1binary:"skip:3"`
	}

	ListObjectInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32 `m1binary:"length:24"`
	}

	ListObjectInfoReply struct {
		rpc.ReturnCode
		Last    bool         `m1binary:"skip:3"`
		Count   uint32       `m1binary:"skip:12"`
		Objects []ObjectInfo `m1binary:"lengthRef:Count"`
	}

	ListLibraryInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32 `m1binary:"skip:8"`
	}

	ListLibraryInfoReply struct {
		rpc.ReturnCode
		Count   uint32
		Objects []LibraryInfo `m1binary:"lengthRef:Count"`
	}

	ListServiceInfoCall struct {
		FirstIndex uint32
		LastIndex  uint32 `m1binary:"skip:4"`
	}

	ListServiceInfoReply struct {
		rpc.ReturnCode
		Count   uint32
		Objects []ServiceInfo `m1binary:"lengthRef:Count"`
	}

	SystemConfigCall struct {
		spare uint32 `m1binary:"skip:16"`
	}

	SystemConfigReply struct {
		rpc.ReturnCode
		LongFilenames                 bool
		ApplicationDevelopmentAllowed bool   `m1binary:"skip:8"`
		SystemDirectory               string `m1binary:"length:84"`
		DriverDirectory               string `m1binary:"length:84"`
		ServiceDirectory              string `m1binary:"length:84"`
		ApplicationDirectory          string `m1binary:"length:84"`
		ClassesDirectory              string `m1binary:"length:84"` // deprecated since MSys v4.00
		NonVolatileDirectory          string `m1binary:"length:84"`
		TemporaryDirectory            string `m1binary:"length:84"`
		KeysDirectory                 string `m1binary:"length:84"`
		ConfigurationDirectory        string `m1binary:"length:84"`
		KeysBackupDirectory           string `m1binary:"length:84,skip:232"`
	}

	CoreUsageCall struct {
		spare uint32 `m1binary:"skip:8"`
	}

	CoreUsageReply struct {
		rpc.ReturnCode
		TimeUnits uint32 `m1binary:"skip:20"` // Clockcycles per microseconds
		Count     uint32
		Cores     []CoreUsage `m1binary:"lengthRef:Count"`
	}

	ListConfigurationInfoCall struct {
		StartIndex uint32 `m1binary:"skip:40"`
	}

	ListConfigurationInfoReply struct {
		rpc.ReturnCode
		TotalConfigurations uint32
		HasMore             bool
		Count               uint32              `m1binary:"skip:40"`
		Configurations      []ConfigurationInfo `m1binary:"lengthRef:Count"`
	}
)
