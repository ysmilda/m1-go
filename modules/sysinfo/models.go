package sysinfo

import (
	"net"
	"time"

	"github.com/ysmilda/m1-go/internals/m1binary"
	"github.com/ysmilda/m1-go/modules/mio"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

type (
	CardInfo struct {
		Name            string `m1binary:"length:24"`
		CardNumber      uint32
		StationNumber   uint32
		SlotNumber      uint32
		MaximumChannels uint32
		Mode            uint32       // Mode of the driver
		Type            mio.IOModule // Type of IO module
		State           uint32
		BusType         uint16 // 0 is Local, otherwise see mio.IOModule
		NetworkNumber   uint16
	}

	CPUAddress struct {
		ProcessorNumber int32
		IPAddress       net.IP `m1binary:"length:4"`
	}

	CPUID struct {
		Type           uint32
		Variant        uint32
		Revision       uint32
		ProductionDate msys.DateTime // Only supported for MSys < 4.10. Later version will return 1.1.1970.
		SerialNumber   string        `m1binary:"length:12"`
	}

	CPUInfo struct {
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
		CPUID                 CPUID
		PccProgrammingEnabled bool `m1binary:"skip:3"` // No PC card write protection if true
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
		MMPEnabled            bool          // M1 Multi Processing enables
		HyperThreadingEnabled bool          `m1binary:"skip:2"`
		UpTime                time.Duration // Time since boot
	}

	SystemObjectInfo struct {
		Name         string `m1binary:"length:16"`
		Version      msys.Version
		Type         uint32
		StartAddress uint32
		Size         uint32
		State        uint32
	}

	IODriverInfo struct {
		Name    string `m1binary:"length:12"`
		Version msys.Version
		Type    uint32
		Size    uint32
		State   uint32
	}

	CPUTime struct {
		RuntimeTotal     int64  // in clockcycles
		RuntimeMinimum   uint32 // in clockcycles
		RuntimeMaximum   uint32 // in clockcycles
		RuntimeAverage   uint32 // in clockcycles
		CycletimeMinimum uint32 // in clockcycles
		CycletimeMaximum uint32 // in clockcycles
		CycletimeAverage uint32 // in clockcycles
	}

	TaskInfo struct {
		Name         string `m1binary:"length:16"`
		Status       string `m1binary:"length:12"`
		Priority     uint32
		DelayTicks   uint32 // Set if status = DELAY
		StackSize    uint32
		StackReserve uint32
		ErrorStatus  uint32
		Time         CPUTime
	}

	ExtendedTaskInfo struct {
		TaskInfo
		LastCore  uint32 // The last core the task has run on
		UsedCores uint32 // Bitmask of the cores the task has ran on
		TaskID    int32  `m1binary:"skip:16"`
	}

	ObjectInfo struct {
		Name        string `m1binary:"length:16"`
		ModuleID    uint64
		Format      int32
		Group       int32
		TextAddress uint64
		DataAddress uint64
		BssAddress  uint64
		TextSize    uint32
		DataSize    uint32
		BssSize     uint32 `m1binary:"skip:12"`
	}

	LibraryInfo struct {
		Name    string `m1binary:"length:32"`
		Path    string `m1binary:"length:84"`
		Vendor  string `m1binary:"length:64"`
		Version msys.Version
		Size    uint32 `m1binary:"skip:4"` // in bytes
	}

	ServiceInfo struct {
		Name    string `m1binary:"length:16"`
		Version msys.Version
		Type    uint32            // Type of service
		Size    uint32            // in bytes
		State   res.ResourceState `m1binary:"skip:4"`
	}

	CoreUsage struct {
		UsageLast100s  int64
		UsageLast1000s int64
		UsageTotal     int64
		TimeTotal      int64 `m1binary:"skip:16"`
	}

	ConfigurationInfo struct {
		Name    string       `m1binary:"length:16"`
		Version msys.Version `m1binary:"skip:40"`
	}
)

func (c *CPUAddress) DecodeM1(data []byte) (int, error) {
	temp := struct {
		ProcessorNumber int32
		IPAddress       string `m1binary:"length:20"`
	}{}

	n, err := m1binary.Decode(data, &temp)
	if err != nil {
		return n, err
	}

	c.ProcessorNumber = temp.ProcessorNumber
	c.IPAddress = net.ParseIP(temp.IPAddress)

	return n, nil
}
