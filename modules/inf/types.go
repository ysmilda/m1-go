package inf

import (
	"net"

	"github.com/ysmilda/m1-go/modules/mio"
	"github.com/ysmilda/m1-go/modules/msys"
)

type CPUAddress struct {
	ProcessorNumber int32
	IPAddress       net.IP
}

type CPUID struct {
	Type           uint32
	Variant        uint32
	Revision       uint32
	ProductionDate msys.DateTime // Only supported for MSys < 4.10. Later version will return 1.1.1970.
	SerialNumber   string
}

type SystemObject struct {
	Name         string
	Version      msys.Version
	Type         uint32
	StartAddress uint32
	Size         uint32 // Size in bytes
	State        uint32
}

type DriverInfo struct {
	Name    string
	Version msys.Version
	Type    uint32
	Size    uint32 // Size in bytes
	State   uint32
}

type TaskInfo struct {
	Name        string
	Status      string
	Priority    uint32
	DelayTicks  uint32 // If status = DELAY
	StackSize   uint32 // Size of the complete stack
	StackMargin uint32 // Size of the stack reserve
	ErrorStatus uint32
	Time        CPUTime // CPU time for this task
}

type CPUTime struct {
	RunTotal     int64  // Total running time in clock cycles
	RunMinimal   uint32 // Minimal run time in clock cycles
	RunMaximal   uint32 // Maximal run time in clock cycles
	RunAverage   uint32 // Average run time in clock cycles
	CycleMinimum uint32 // Minimal cycle time in clock cycles
	CycleMaximum uint32 // Maximal cycle time in clock cycles
	CycleAverage uint32 // Average cycle time in clock cycles
}

type CardInfo struct {
	Name          string
	CardNumber    uint32
	StationNumber uint32
	SlotNumber    uint32
	ChannelCount  uint32
	Mode          uint32       // Mode of the driver
	Type          mio.IOModule // Type of IO module
}

type OSVariant uint8

const (
	OSVariantVxWorks5 OSVariant = 1
	OSVariantVxWorks7 OSVariant = 2
)

func (o OSVariant) String() string {
	switch o {
	case 1:
		return "VxWorks 5"
	case 2:
		return "VxWorks 7"
	}
	return "Unknown"
}
