package m1errors

// Error sources.
const (
	SourceUnknown  uint32 = 0x00000000 // Unknown source
	SourceSVI      uint32 = 0x00010000 // SVI Functions
	SourceSMI      uint32 = 0x00020000 // SMI Functions
	SourceRES      uint32 = 0x00030000 // Resource Handler
	SourceMIO      uint32 = 0x00040000 // IO Handler
	SourceVHD      uint32 = 0x00050000 // Vis Handler
	SourceINF      uint32 = 0x00060000 // Info Handler
	SourcePLC      uint32 = 0x00070000 // PLC Runtime system
	SourceMOD      uint32 = 0x00080000 // Module Handler
	SourceCAN      uint32 = 0x00090000 // CAN Handler
	SourcePF       uint32 = 0x000A0000 // Profile Functions
	SourceSYS      uint32 = 0x000F0000 // System MSys
	SourceCORE     uint32 = 0x00100000 // System MCore
	SourceEHD      uint32 = 0x00110000 // Error Handler
	SourcePB       uint32 = 0x00120000 // Profibus Handler
	SourceDBG      uint32 = 0x00130000 // Debug Handler
	SourceDN       uint32 = 0x00140000 // DeviceNet Handler
	SourceRFS      uint32 = 0x00150000 // Remote File Server
	SourceSLC      uint32 = 0x00160000 // SLC Server
	SourceDMW      uint32 = 0x00170000 // Drive Middleware
	SourceSEM201   uint32 = 0x00180000 // SERCOS driver
	SourceUFB      uint32 = 0x00190000 // Unified fieldbus components
	SourcePN       uint32 = 0x001A0000 // Profinet
	SourceEC       uint32 = 0x001B0000 // EtherCAT
	SourceBCR      uint32 = 0x001C0000 // BCR2xx
	SourceST       uint32 = 0x001D0000 // Self Test Module
	SourceC_TDLL   uint32 = 0x01010000 // TMANw32-DLL
	SourceC_PLCCOM uint32 = 0x01020000 // PLCCOM-DLL
	SourceC_TCONF  uint32 = 0x01030000 // TargetManager EXE
	SourceC_PLCHWM uint32 = 0x01040000 // PLCHWM-DLL
	SourceC_TVIEW  uint32 = 0x01050000 // TMAN-View
	SourceC_MMAN   uint32 = 0x01060000 // M-Manager
	SourceC_MPLC   uint32 = 0x01070000 // M-PLC
	SourceC_MIF    uint32 = 0x01080000 // M-Interface
	SourceAHD      uint32 = 0x02010000 // AHD System SW Module
	SourceLOGGER   uint32 = 0x02020000 // Logger System SW Module
	SourceEVT      uint32 = 0x02030000 // Event Logger
	SourceETCP     uint32 = 0x02040000 // 60870-5
	SourceDNP3     uint32 = 0x02050000 // DNP3
	SourceATEC     uint32 = 0x02060000 // ATEC
	SourceM1C      uint32 = 0x81100000 // M1 Core
)

type source uint32

func (s source) SVI() bool {
	return s.check(SourceSVI)
}

func (s source) SMI() bool {
	return s.check(SourceSMI)
}

func (s source) RES() bool {
	return s.check(SourceRES)
}

func (s source) MIO() bool {
	return s.check(SourceMIO)
}

func (s source) VHD() bool {
	return s.check(SourceVHD)
}

func (s source) INF() bool {
	return s.check(SourceINF)
}

func (s source) PLC() bool {
	return s.check(SourcePLC)
}

func (s source) MOD() bool {
	return s.check(SourceMOD)
}

func (s source) CAN() bool {
	return s.check(SourceCAN)
}

func (s source) PF() bool {
	return s.check(SourcePF)
}

func (s source) SYS() bool {
	return s.check(SourceSYS)
}

func (s source) CORE() bool {
	return s.check(SourceCORE)
}

func (s source) EHD() bool {
	return s.check(SourceEHD)
}

func (s source) PB() bool {
	return s.check(SourcePB)
}

func (s source) DBG() bool {
	return s.check(SourceDBG)
}

func (s source) DN() bool {
	return s.check(SourceDN)
}

func (s source) check(code uint32) bool {
	return (uint32(s)&0xFFFF0000)&code != 0
}

func (s source) Error() string { //nolint: gocyclo
	switch uint32(s) & 0xFFFF0000 {
	default:
		return "unknown"
	case SourceSVI:
		return "SVI functions"
	case SourceSMI:
		return "SMI functions"
	case SourceRES:
		return "resource handler"
	case SourceMIO:
		return "IO handler"
	case SourceVHD:
		return "vis handler"
	case SourceINF:
		return "info handler"
	case SourcePLC:
		return "PLC runtime system"
	case SourceMOD:
		return "module handler"
	case SourceCAN:
		return "CAN handler"
	case SourcePF:
		return "profile functions"
	case SourceSYS:
		return "system MSys"
	case SourceCORE:
		return "system MCore"
	case SourceEHD:
		return "error handler"
	case SourcePB:
		return "profibus handler"
	case SourceDBG:
		return "debug handler"
	case SourceDN:
		return "DeviceNet handler"
	case SourceRFS:
		return "remote file server"
	case SourceSLC:
		return "SLC server"
	case SourceDMW:
		return "drive middleware"
	case SourceSEM201:
		return "SERCOS driver"
	case SourceUFB:
		return "unified fieldbus components"
	case SourcePN:
		return "profinet"
	case SourceEC:
		return "EtherCAT"
	case SourceBCR:
		return "BCR2xx"
	case SourceST:
		return "self test module"
	case SourceC_TDLL:
		return "TMANw32-DLL"
	case SourceC_PLCCOM:
		return "PLCCOM-DLL"
	case SourceC_TCONF:
		return "TargetManager EXE"
	case SourceC_PLCHWM:
		return "PLCHWM-DLL"
	case SourceC_TVIEW:
		return "TMAN-View"
	case SourceC_MMAN:
		return "M-Manager"
	case SourceC_MPLC:
		return "M-PLC"
	case SourceC_MIF:
		return "M-Interface"
	case SourceAHD:
		return "AHD system module"
	case SourceLOGGER:
		return "logger system module"
	case SourceEVT:
		return "event logger"
	case SourceETCP:
		return "60870-5"
	case SourceDNP3:
		return "DNP3"
	case SourceATEC:
		return "ATEC"
	case SourceM1C:
		return "M1 Core"
	}
}
