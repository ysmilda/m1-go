package res

import (
	"net"
	"time"

	"github.com/ysmilda/m1-go/modules/msys"
)

type (
	ModuleInfo struct {
		Type              string    `m1binary:"length:12"`
		Name              string    `m1binary:"length:12"`
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

	ExtendedModuleInfo struct {
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

	ModuleTaskInfo struct {
		TaskID      uint32
		Name        string `m1binary:"length:16"`
		Priority    uint32
		Status      uint32
		ErrorStatus uint32
		Affinity    uint32
	}

	// ModuleInfo contains the response of the ModuleInfo procedure.
	ModuleNumber struct {
		ModuleNumber uint32
		Port         uint16
	}

	// SystemInfo contains the response of the SystemInfo procedure.
	SystemInfo struct {
		CPUSwitch       uint32
		ProcessorNumber uint32
		RestartCount    uint32
		IPAddress       net.IP `m1binary:"length:4"`
		CPUID           CPUType
		MCoreVersion    msys.Version
		MSysVersion     msys.Version
		ApplicationName string `m1binary:"length:24"`
		SecurityLevel   uint32
		LoginRequired   bool `m1binary:"skip:3"`
		DateTime        msys.DateTime
		LoginChecker    bool `m1binary:"skip:3"`
		AccessControl   bool
		AccessHandler   bool
		UserManagement  bool `m1binary:"skip:1"`
		AppState        ResourceState
		SystemState     ResourceState
		SessionLifetime uint32
		TimeBase        uint16 `m1binary:"skip:2"`
		Timezone        string `m1binary:"length:36"`
		TimeOffset      int32
		TimeUTCSec      uint32
		TimeUTCNsec     uint32
	}

	// CPUType contains information about the CPU.
	CPUType struct {
		Type         uint32
		Variant      uint32
		SerialNumber string `m1binary:"length:12"`
	}

	// Login2 contains the response of the Login2 procedure.
	Login2 struct {
		SecurityLevel          uint32 `m1binary:"skip:4"`
		UserAccount            UserAccess
		AuthLen                uint32
		Auth                   []byte        `m1binary:"length:128"`
		UserData               []byte        `m1binary:"length:128"`
		LockoutWaitTime        time.Duration `m1binary:"length:4,unit:seconds"`
		PasswordMinimumClasses uint32
		PasswordMinimumLength  uint32
	}

	// Login contains the response of the Login procedure.
	Login struct {
		SecurityLevel uint32
		Permissions   Permissions
		AuthLen       uint32
		Auth          []byte `m1binary:"length:128"`
	}

	// ExtLogin contains the response of the ExtLogin procedure.
	ExtLogin struct {
		SecurityLevel uint32
		AuthLen       uint32
		Auth          []byte `m1binary:"length:128"`
		UserData      []byte `m1binary:"length:128"`
	}

	// Open contains the response of the Open procedure.
	Open struct {
		SessionTimeout        int32
		SessionLifetime       int32
		SMIMessageSize        int32
		SessionIdlePrevention bool `m1binary:"skip:127"`
		AuthLen               uint32
		Auth                  []byte `m1binary:"length:128"`
	}

	// Renew contains the response of the Renew procedure.
	Renew struct {
		RevisedAuthenticationRenewal bool `m1binary:"skip:19"`
		RestartCount                 uint32
		ApplicationState             ResourceState
		SystemState                  ResourceState `m1binary:"skip:128"`
		AuthLen                      uint32
		Auth                         []byte `m1binary:"length:128"`
	}

	// ExtPing contains the response of the ExtPing procedure.
	ExtPing struct {
		SerialNumber    string `m1binary:"length:12"`
		MSysVersion     msys.Version
		TargetName      string `m1binary:"length:20"`
		RequestMode     ReplyMode
		DHCPEnabled     bool
		OperatingSystem OSVariant `m1binary:"skip:2"`
		CPUType         uint32
		CPUVariant      uint32
	}

	ExtPing2 struct {
		ExtPing
		NumberOfApps uint32
		Apps         []ModName // TODO: Figure out how to m1binary this
	}

	ModName struct {
		ApplicationName string `m1binary:"length:12"`
		State           ResourceState
		Attributes      Attribute
	}
)
