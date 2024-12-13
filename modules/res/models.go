package res

import (
	"time"

	"github.com/ysmilda/m1-go/modules/msys"
)

type ModuleInfo struct {
	Type              string    `unpack:"length=12"`
	Name              string    `unpack:"length=12"`
	Index             uint32    // Index of the module instance
	Partition         uint32    // Memory partition the module runs in
	MinimumRPCVersion uint32    // Oldest supported RPC version
	MaximumRPCVersion uint32    // Newest supported RPC version
	MaximumUsers      uint32    // Maximum number of concurrent users
	Attributes        Attribute // Module attributes
	ModuleID          uint32    // Module ID as in VxWorks
	ActiveUsers       uint32    // Number of active users
	State             ResourceState
	ModuleNumber      uint32 // Module number for SMI
}

type ModuleXInfo struct {
	ModuleInfo
	ParentTaskID uint32 // Task ID of the parent module
	UDPPort      uint16
	TCPPort      uint16
	Checksum     uint32       // Checksum of the module
	Version      msys.Version // Version of the module
	Incarnation  uint32       // Incarnation ciybter of the module
	TaskID       uint32       // Task ID of the module
	Affinity     uint32       // Core affinity of the module
}

type ModuleTaskInfo struct {
	TaskID      uint32
	Name        string `unpack:"length=16"`
	Priority    uint32
	Status      uint32
	ErrorStatus uint32
	Affinity    uint32
}

// ModuleInfo contains the response of the ModuleInfo procedure.
type ModuleNumber struct {
	ModuleNumber uint32
	UDPPort      uint16
	TCPPort      uint16
}

type UserInfo struct {
	IPAddress    uint32
	ModuleNumber uint32
}

// SystemInfo contains the response of the SystemInfo procedure.
type SystemInfo struct {
	CPUSwitch       uint32
	ProcessorNumber uint32
	RestartCount    uint32
	IPAddress       uint32
	CPUID           CPUType
	MCoreVersion    msys.Version
	MSysVersion     msys.Version
	ApplicationName string `unpack:"length=24"`
	SecurityLevel   uint32
	LoginRequired   bool `unpack:"skip=3"`
	DateTime        msys.DateTime
	LoginChecker    bool `unpack:"skip=3"`
	AccessControl   bool
	AccessHandler   bool
	UserManagement  bool `unpack:"skip=1"`
	AppState        ResourceState
	SystemState     ResourceState
	SessionLifetime uint32
	TimeBase        uint16 `unpack:"skip=2"`
	Timezone        string `unpack:"length=36"`
	TimeOffset      int32
	TimeUTCSec      uint32
	TimeUTCNsec     uint32
}

// CPUType contains information about the CPU.
type CPUType struct {
	Type         uint32
	Variant      uint32
	SerialNumber string `unpack:"length=12"`
}

// UserAccess contains information about the user access.
type UserAccess struct {
	Group                      uint8
	Level                      uint8
	Priority                   uint8 `unpack:"skip=1"`
	SystemPermissions          int64
	ApplicationPermissions     int64
	AppData                    int32
	PasswordValidityTime       time.Duration `unpack:"length=4,unit=days"`
	DaysTillPasswordExpiration time.Duration `unpack:"length=4,unit=days"`
}

// Login2 contains the response of the Login2 procedure.
type Login2 struct {
	SecurityLevel          uint32 `unpack:"skip=4"`
	UserAccount            UserAccess
	AuthLen                uint32
	Auth                   []byte        `unpack:"length=128"`
	UserData               []byte        `unpack:"length=128"`
	LockoutWaitTime        time.Duration `unpack:"length=4,unit=seconds"`
	PasswordMinimumClasses uint32
	PasswordMinimumLength  uint32
}

// Login contains the response of the Login procedure.
type Login struct {
	SecurityLevel uint32
	Permissions   int64
	AuthLen       uint32
	Auth          []byte `unpack:"length=128"`
}

// ExtLogin contains the response of the ExtLogin procedure.
type ExtLogin struct {
	SecurityLevel uint32
	AuthLen       uint32
	Auth          []byte `unpack:"length=128"`
	UserData      []byte `unpack:"length=128"`
}

// Open contains the response of the Open procedure.
type Open struct {
	SessionTimeout        int32
	SessionLifetime       int32
	SMIMessageSize        int32
	SessionIdlePrevention bool `unpack:"skip=127"`
	AuthLen               uint32
	Auth                  []byte `unpack:"length=128"`
}

// Renew contains the response of the Renew procedure.
type Renew struct {
	RevisedAuthenticationRenewal bool `unpack:"skip=19"`
	RestartCount                 uint32
	ApplicationState             ResourceState
	SystemState                  ResourceState `unpack:"skip=128"`
	AuthLen                      uint32
	Auth                         []byte `unpack:"length=128"`
}

// ExtPing contains the response of the ExtPing procedure.
type ExtPing struct {
	SerialNumber    string `unpack:"length=12"`
	MSysVersion     msys.Version
	TargetName      string `unpack:"length=20"`
	RequestMode     ReplyMode
	DHCPEnabled     bool
	OperatingSystem OSVariant `unpack:"skip=2"`
	CPUType         uint32
	CPUVariant      uint32
}

type ExtPing2 struct {
	ExtPing
	NumberOfApps uint32
	Apps         []ModName // TODO: Figure out how to unpack this
}

type ModName struct {
	ApplicationName string `unpack:"length=12"`
	State           ResourceState
	Attributes      Attribute
}
