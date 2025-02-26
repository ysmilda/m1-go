package res

type MemoryType uint32

const (
	MemoryTypeUnknown         MemoryType = 0
	MemoryTypeDRAM            MemoryType = 1
	MemoryTypeSRAM            MemoryType = 2
	MemoryTypeFlash           MemoryType = 3
	MemoryTypeHardDisk        MemoryType = 4
	MemoryTypeFloppyDisk      MemoryType = 5
	MemoryTypeFTPHost         MemoryType = 6
	MemoryTypeNFSHost         MemoryType = 7
	MemoryTypeCFCard          MemoryType = 8
	MemoryTypeUSB             MemoryType = 9
	MemoryTypeTrueCryptVolume MemoryType = 10
	MemoryTypeSymlink         MemoryType = 11
	MemoryTypeDummy           MemoryType = 20
)

type FunctionInterface uint32

const (
	FunctionInterfaceFPU FunctionInterface = 1
)

type MemoryInfo uint32

const (
	MemoryInfoFile   MemoryInfo = 1
	MemoryInfoMemory MemoryInfo = 2
)

type ResourceState uint32

const (
	ResourceStateRun        ResourceState = 1
	ResourceStateError      ResourceState = 2
	ResourceStateStop       ResourceState = 3
	ResourceStateInit       ResourceState = 4
	ResourceStateDeinit     ResourceState = 5
	ResourceStateEndOfInit  ResourceState = 6
	ResourceStateReset      ResourceState = 7
	ResourceStateWarning    ResourceState = 8
	ResourceStateSmartError ResourceState = 9
	ResourceStateUnknown    ResourceState = 0
)

func (r ResourceState) String() string {
	switch r {
	case 1:
		return "Run"
	case 2:
		return "Error"
	case 3:
		return "Stop"
	case 4:
		return "Init"
	case 5:
		return "Deinit"
	case 6:
		return "End of init"
	case 7:
		return "Reset"
	case 8:
		return "Warning"
	case 9:
		return "S.M.A.R.T error"
	}
	return "Unknown"
}

type Filter uint32

const (
	FilterEqual    Filter = 1
	FilterNotEqual Filter = 2
)

type ReplyType uint32

const (
	ReplyTypeAuto      ReplyType = 1
	ReplyTypeIP        ReplyType = 2
	ReplyTypeBroadcast ReplyType = 3
)

type ReplyMode uint32

const (
	ReplyModeNormal ReplyMode = 0 // Reply ExtPingReply
	ReplyModeAllApp ReplyMode = 1 // Reply ExtPingReply2
	ReplyModePLCApp ReplyMode = 2 // Reply ExtPingReply2
)

type OSVariant uint32

const (
	OSVariantVxWorks OSVariant = 0
	OSVariantLinux   OSVariant = 1
	OSVariantWindows OSVariant = 2
)

func (o OSVariant) String() string {
	switch o {
	case 0:
		return "VxWorks"
	case 1:
		return "Linux"
	case 2:
		return "Windows"
	}
	return "Unknown"
}

type CallBack uint32

const (
	CallBackEHD CallBack = 0x001
	CallBackVHD CallBack = 0x002
	CallBackDBG CallBack = 0x004
	CallBackAll CallBack = 0xFFFF
)

type TaskState uint32

const (
	TaskStateRunning      TaskState = 1
	TaskStateNotRunning   TaskState = 2
	TaskStateInitializing TaskState = 3
)

type LicenseKeyType uint32

const (
	LicenseKeyTypeCFCSerial      LicenseKeyType = 0x01
	LicenseKeyTypeCPUSerial      LicenseKeyType = 0x02
	LicenseKeyTypeUSBFlashSerial LicenseKeyType = 0x04
	LicenseKeyTypeUSBSerial      LicenseKeyType = 0x08
	LicenseKeyTypeVersion        LicenseKeyType = 0x10
	LicenseKeyTypeDeadline       LicenseKeyType = 0x20
)

type Progress uint32

const (
	ProgressIdle        Progress = 0
	ProgressRunning     Progress = 1
	ProgressCancelled   Progress = 2
	ProgressFinished    Progress = 3
	ProgressCancellable Progress = 0x02
)

// TODO: Figure out structure.
type Permissions int64
