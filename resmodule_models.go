package m1

import (
	"time"

	"github.com/ysmilda/m1-go/pkg/buffer"
)

// ExtPingResponse contains the response of the ExtPing procedure.
type ExtPingResponse struct {
	SerialNumber    string
	MSysVersion     Version
	TargetName      string
	RequestMode     uint32
	DHCPEnabled     bool
	OperatingSystem uint8 // 0 = vxworks, 1 = linux, 2 = windows
	CPUType         uint32
	CPUVariant      uint32
}

func (e *ExtPingResponse) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	e.SerialNumber, _ = in.ReadString(_MIO_ProductNumberLength)
	e.MSysVersion.parse(in)
	e.TargetName, _ = in.ReadString(_BootParameterLength)
	return returnCode
}

// ModuleNumberResponse contains the response of the ModuleInfo procedure.
type ModuleNumberResponse struct {
	ModuleNumber uint32
	UDPPort      uint16
	TCPPort      uint16
}

func (m *ModuleNumberResponse) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	m.ModuleNumber, _ = in.LittleEndian.ReadUint32()
	m.UDPPort, _ = in.LittleEndian.ReadUint16()
	m.TCPPort, _ = in.LittleEndian.ReadUint16()
	return returnCode
}

// SystemInfoResponse contains the response of the SystemInfo procedure.
type SystemInfoResponse struct {
	CPUSwitch       uint32
	ProcessorNumber uint32
	RestartCount    uint32
	IPAddress       uint32
	CPUID           CPUType
	MCoreVersion    Version
	MSysVersion     Version
	ApplicationName string
	SecurityLevel   uint32
	LoginRequired   bool
	DateTime        DateTime
	LoginChecker    bool
	AccessControl   bool
	AccessHandler   bool
	UserManagement  bool
	// 1 = run, 2 = error, 3 = stop, 4 = init, 5 = deinit, 6 = eoi, 7 = reset, 8 = warning, 9 = error_smart
	AppState uint32
	// 1 = run, 2 = error, 3 = stop, 4 = init, 5 = deinit, 6 = eoi, 7 = reset, 8 = warning, 9 = error_smart
	SysState        uint32
	SessionLifetime uint32
	TimeBase        uint16
	Timezone        string
	TimeOffset      int32
	TimeUTC         time.Time
}

func (r *SystemInfoResponse) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	r.CPUSwitch, _ = in.LittleEndian.ReadUint32()
	r.ProcessorNumber, _ = in.LittleEndian.ReadUint32()
	r.RestartCount, _ = in.LittleEndian.ReadUint32()
	r.IPAddress, _ = in.LittleEndian.ReadUint32()
	r.CPUID.parse(in)
	r.MCoreVersion.parse(in)
	r.MSysVersion.parse(in)
	r.ApplicationName, _ = in.ReadString(_AppNameLength)
	r.SecurityLevel, _ = in.LittleEndian.ReadUint32()
	r.LoginRequired, _ = in.ReadBool()
	in.Skip(3) // spare
	r.DateTime.parse(in)
	r.LoginChecker, _ = in.ReadBool()
	in.Skip(3) // spare
	r.AccessControl, _ = in.ReadBool()
	r.AccessHandler, _ = in.ReadBool()
	r.UserManagement, _ = in.ReadBool()
	in.Skip(1) // spare
	r.AppState, _ = in.LittleEndian.ReadUint32()
	r.SysState, _ = in.LittleEndian.ReadUint32()
	r.SessionLifetime, _ = in.LittleEndian.ReadUint32()
	r.TimeBase, _ = in.LittleEndian.ReadUint16()
	in.Skip(2) // spare
	r.Timezone, _ = in.ReadString(_TimezoneLength)
	r.TimeOffset, _ = in.LittleEndian.ReadInt32()
	sec, _ := in.LittleEndian.ReadUint32()
	nsec, _ := in.LittleEndian.ReadUint32()
	r.TimeUTC = time.Unix(int64(sec), int64(nsec)).UTC()
	in.Skip(8) // spare
	return returnCode
}

// Login2Response contains the response of the Login2 procedure.
type Login2Response struct {
	SecurityLevel          uint32
	UserAccount            UserAccess
	AuthLen                uint32
	Auth                   []byte
	UserData               []byte
	LockoutWaitTime        time.Duration
	PasswordMinimumClasses uint32
	PasswordMinimumLength  uint32
}

func (r *Login2Response) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.SecurityLevel, _ = in.LittleEndian.ReadUint32()
	in.Skip(4) // spare
	r.UserAccount.parse(in)
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	r.UserData, _ = in.ReadBytes(128)
	lockoutWaitTime, _ := in.LittleEndian.ReadUint32()
	r.LockoutWaitTime = time.Duration(lockoutWaitTime) * time.Second
	r.PasswordMinimumClasses, _ = in.LittleEndian.ReadUint32()
	r.PasswordMinimumLength, _ = in.LittleEndian.ReadUint32()
	in.Skip(116) // spare
	return returncode
}

// LoginResponse contains the response of the Login procedure.
type LoginResponse struct {
	SecurityLevel uint32
	Permissions   int64
	AuthLen       uint32
	Auth          []byte
}

func (r *LoginResponse) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.SecurityLevel, _ = in.LittleEndian.ReadUint32()
	r.Permissions, _ = in.LittleEndian.ReadInt64()
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	return returncode
}

// ExtLoginResponse contains the response of the ExtLogin procedure.
type ExtLoginResponse struct {
	SecurityLevel uint32
	AuthLen       uint32
	Auth          []byte
	UserData      []byte
}

func (r *ExtLoginResponse) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.SecurityLevel, _ = in.LittleEndian.ReadUint32()
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	r.UserData, _ = in.ReadBytes(128)
	return returncode
}

// OpenResponse contains the response of the Open procedure.
type OpenResponse struct {
	SessionTimeout        int32
	SessionLifetime       int32
	SMIMessageSize        int32
	SessionIdlePrevention bool
	AuthLen               uint32
	Auth                  []byte
}

func (r *OpenResponse) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.SessionTimeout, _ = in.LittleEndian.ReadInt32()
	r.SessionLifetime, _ = in.LittleEndian.ReadInt32()
	r.SMIMessageSize, _ = in.LittleEndian.ReadInt32()
	r.SessionIdlePrevention, _ = in.ReadBool()
	in.Skip(127) // spare
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	return returncode
}

// RenewResponse contains the response of the Renew procedure.
type RenewResponse struct {
	RevisedAuthenticationRenewal bool
	RestartCount                 uint32
	// 1 = run, 2 = error, 3 = stop, 4 = init, 5 = deinit, 6 = eoi, 7 = reset, 8 = warning, 9 = error_smart
	ApplicationState uint32 // TODO: Maybe create a type for this?
	// 1 = run, 2 = error, 3 = stop, 4 = init, 5 = deinit, 6 = eoi, 7 = reset, 8 = warning, 9 = error_smart
	SystemState uint32
	AuthLen     uint32
	Auth        []byte
}

func (r *RenewResponse) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.RevisedAuthenticationRenewal, _ = in.ReadBool()
	in.Skip(19) // spare
	r.RestartCount, _ = in.LittleEndian.ReadUint32()
	r.ApplicationState, _ = in.LittleEndian.ReadUint32()
	r.SystemState, _ = in.LittleEndian.ReadUint32()
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	return returncode
}
