package m1

import (
	"time"

	"github.com/ysmilda/m1-go/pkg/buffer"
)

// ExtPing contains the response of the ExtPing procedure.
type ExtPing struct {
	SerialNumber    string
	MSysVersion     Version
	TargetName      string
	RequestMode     uint32
	DHCPEnabled     bool
	OperatingSystem uint8 // 0 = vxworks, 1 = linux, 2 = windows
	CPUType         uint32
	CPUVariant      uint32
}

func (e *ExtPing) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	e.SerialNumber, _ = in.ReadString(_MIO_ProductNumberLength)
	e.MSysVersion.parse(in)
	e.TargetName, _ = in.ReadString(_BootParameterLength)
	return returnCode
}

// ModuleInfo contains the response of the ModuleInfo procedure.
type ModuleInfo struct {
	ModuleNumber uint32
	UDPPort      uint16
	TCPPort      uint16
}

func (m *ModuleInfo) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	m.ModuleNumber, _ = in.LittleEndian.ReadUint32()
	m.UDPPort, _ = in.LittleEndian.ReadUint16()
	m.TCPPort, _ = in.LittleEndian.ReadUint16()
	return returnCode
}

// SystemInfo contains the response of the SystemInfo procedure.
type SystemInfo struct {
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

func (r *SystemInfo) parse(in *buffer.Buffer) uint32 {
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

// Login2 contains the response of the Login2 procedure.
type Login2 struct {
	SecurityLevel          uint32
	UserAccount            UserAccess
	AuthLen                uint32
	Auth                   []byte
	UserData               []byte
	LockoutWaitTime        time.Duration
	PasswordMinimumClasses uint32
	PasswordMinimumLength  uint32
}

func (r *Login2) parse(in *buffer.Buffer) uint32 {
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

// Login contains the response of the Login procedure.
type Login struct {
	SecurityLevel uint32
	Permissions   int64
	AuthLen       uint32
	Auth          []byte
}

func (r *Login) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.SecurityLevel, _ = in.LittleEndian.ReadUint32()
	r.Permissions, _ = in.LittleEndian.ReadInt64()
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	return returncode
}

// ExtLogin contains the response of the ExtLogin procedure.
type ExtLogin struct {
	SecurityLevel uint32
	AuthLen       uint32
	Auth          []byte
	UserData      []byte
}

func (r *ExtLogin) parse(in *buffer.Buffer) uint32 {
	returncode, _ := in.LittleEndian.ReadUint32()
	r.SecurityLevel, _ = in.LittleEndian.ReadUint32()
	r.AuthLen, _ = in.LittleEndian.ReadUint32()
	r.Auth, _ = in.ReadBytes(128)
	r.UserData, _ = in.ReadBytes(128)
	return returncode
}

// Open contains the response of the Open procedure.
type Open struct {
	SessionTimeout        int32
	SessionLifetime       int32
	SMIMessageSize        int32
	SessionIdlePrevention bool
	AuthLen               uint32
	Auth                  []byte
}

func (r *Open) parse(in *buffer.Buffer) uint32 {
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

// Renew contains the response of the Renew procedure.
type Renew struct {
	RevisedAuthenticationRenewal bool
	RestartCount                 uint32
	// 1 = run, 2 = error, 3 = stop, 4 = init, 5 = deinit, 6 = eoi, 7 = reset, 8 = warning, 9 = error_smart
	ApplicationState uint32 // TODO: Maybe create a type for this?
	// 1 = run, 2 = error, 3 = stop, 4 = init, 5 = deinit, 6 = eoi, 7 = reset, 8 = warning, 9 = error_smart
	SystemState uint32
	AuthLen     uint32
	Auth        []byte
}

func (r *Renew) parse(in *buffer.Buffer) uint32 {
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

type SessionInfo struct {
	Name   string
	UserID uint32
	Start  time.Time
}

func (v *SessionInfo) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	v.Name, _ = in.ReadString(_VHD_UserNameLength)
	v.UserID, _ = in.LittleEndian.ReadUint32()
	start, _ := in.LittleEndian.ReadUint32()
	v.Start = time.Unix(int64(start), 0)
	return returnCode
}

type SVIServerInfo struct {
	Version           uint32 // Version of supported SVI
	AddressType       uint32 // The type of address class used by the server
	NumberOfVariables uint32 // The number of exported SVI variables
}

func (v *SVIServerInfo) parse(in *buffer.Buffer) uint32 {
	returnCode, _ := in.LittleEndian.ReadUint32()
	v.Version, _ = in.LittleEndian.ReadUint32()
	v.AddressType, _ = in.LittleEndian.ReadUint32()
	v.NumberOfVariables, _ = in.LittleEndian.ReadUint32()
	return returnCode
}
