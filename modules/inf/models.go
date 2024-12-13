package inf

import (
	"time"

	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

type CPUAddressList struct {
	CurrentIndex uint32
	Addresses    []CPUAddress
}

type CPUInfo struct {
	CPUSwitch             uint32 // Position of CPU-ID
	ProcessorNumber       uint32
	ClockSpeed            uint32 // Clock speed in MHz
	TickRate              uint32 // Tick rate in Hz
	TimeSlice             uint32 // Time slice in ticks
	TotalTicks            uint32 // Total ticks since boot
	AuxiliaryClockRate    uint32 // Auxiliary clock rate in ticks/sec
	RestartCounter        uint32
	DebugMode             uint32
	DateTime              msys.DateTime
	DRAMSize              uint32
	NVRAMSize             uint32
	CpuID                 CPUID
	PccProgrammingEnabled bool // No PC card write protection if true
	BootMode              uint32
	DRAMStartAddress      uint32
	NVRAMStartAddress     uint32
	SystemState           res.ResourceState
	SlotNumber            uint32
	FPGAVersion           uint32
	MCoreType             uint32
	TickSource            msys.TickSource
	HardwareVariant       uint8
	OSVariant             OSVariant
	AssemblyCode          uint16
	FPGAVersion2          uint32
	AmountOfCores         uint32
	MMPEnabled            bool // M1 Multi Processing enables
	HyperThreadingEnabled bool
	UpTime                time.Duration // Time since boot
}

type LogInfo struct {
	NumberOfEntries uint32
	FileSize        uint32
	SizeOfLog       uint32
	Path            string // Path to the log file
}

type TaskInfoList struct {
	TotalTime int64  // Clocks since power up
	TimeUnits uint32 // Clock cycles per us
	Tasks     []TaskInfo
}

type BootInfo struct {
	msys.BootInfo
}
