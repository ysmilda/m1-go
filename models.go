package m1

import (
	"fmt"
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

type SviVariable struct {
	Name  string
	Error error

	Value   any // After a read the value will be stored here.
	Address uint64
	Format  uint16 // TODO: We need to parse this to a Go type.
	Length  uint16

	initialized bool
}

func NewSviVariable(module, name string) *SviVariable {
	return &SviVariable{Name: fmt.Sprintf("%s/%s", module, name)}
}

func (s SviVariable) IsInitialized() bool {
	return s.initialized
}

func (s SviVariable) IsBlock() bool {
	return s.Format&_FormatBlock != 0
}

func (s SviVariable) GetArrayLength() int {
	if !s.IsBlock() {
		return 1
	}
	return int(s.Length)
}

func (s SviVariable) getBufferLength() int {
	return int(s.Length) + 3&_Align
}

func (s *SviVariable) getDataTypeLength() int {
	switch s.Format & _FormatTypeMask {
	case _FormatUint1, _FormatUint8, _FormatSint8, _FormatChar8,
		_FormatBool, _FormatMixed, _FormatString, _FormatStringListBase:
		return 1

	case _FormatUint16, _FormatSint16, _FormatChar16, _FormatUnicodeStringListBase:
		return 2

	case _FormatUint32, _FormatSint32, _FormatReal32:
		return 4

	case _FormatUint64, _FormatSint64, _FormatReal64:
		return 8

	default:
		return 1
	}
}

func (s *SviVariable) parse(in *buffer.Buffer) {
	s.Address, _ = in.LittleEndian.ReadUint64()
	s.Format, _ = in.LittleEndian.ReadUint16()
	s.Length, _ = in.LittleEndian.ReadUint16()
}

func (s *SviVariable) parseValue(in *buffer.Buffer) {
	if s.IsBlock() {
		switch s.Format & _FormatElementaryTypeMask {
		case _FormatUint1, _FormatBool:
			val, _ := in.ReadBytes(int(s.Length))
			out := make([]bool, s.Length)
			for i, v := range val {
				out[i] = v == 1
			}
			s.Value = out

		case _FormatUint8:
			s.Value, _ = in.ReadBytes(int(s.Length))

		case _FormatSint8:
			val, _ := in.ReadBytes(int(s.Length))
			out := make([]int8, s.Length)
			for i, v := range val {
				out[i] = int8(v)
			}
			s.Value = out

		case _FormatUint16:
			out := make([]uint16, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadUint16()
			}
			s.Value = out

		case _FormatSint16:
			out := make([]int16, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadInt16()
			}
			s.Value = out

		case _FormatUint32:
			out := make([]uint32, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadUint32()
			}
			s.Value = out

		case _FormatSint32:
			out := make([]int32, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadInt32()
			}
			s.Value = out

		case _FormatReal32:
			out := make([]float32, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadFloat32()
			}
			s.Value = out

		case _FormatUint64:
			out := make([]uint64, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadUint64()
			}
			s.Value = out

		case _FormatSint64:
			out := make([]int64, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadInt64()
			}
			s.Value = out

		case _FormatReal64:
			out := make([]float64, s.getDataTypeLength())
			for i := range out {
				out[i], _ = in.LittleEndian.ReadFloat64()
			}
			s.Value = out

		case _FormatChar8, _FormatChar16:
			s.Value, _ = in.ReadString(int(s.Length))

		case _FormatMixed:
			s.Value, _ = in.ReadBytes(int(s.Length))

			// TODO: Not sure how to support these, or what they are.
			// case _FormatStringList, _FormatUnicodeStringList:
		}
	}

	switch s.Format & _FormatElementaryTypeMask {
	case _FormatUint1:
		s.Value, _ = in.ReadBool()
	case _FormatUint8:
		val, _ := in.ReadByte()
		s.Value = val
	case _FormatSint8:
		val, _ := in.ReadByte()
		s.Value = int8(val)
	case _FormatUint16:
		s.Value, _ = in.LittleEndian.ReadUint16()
	case _FormatSint16:
		s.Value, _ = in.LittleEndian.ReadInt16()
	case _FormatUint32:
		s.Value, _ = in.LittleEndian.ReadUint32()
	case _FormatSint32:
		s.Value, _ = in.LittleEndian.ReadInt32()
	case _FormatReal32:
		s.Value, _ = in.LittleEndian.ReadFloat32()
	case _FormatBool:
		s.Value, _ = in.ReadBool()
	case _FormatUint64:
		s.Value, _ = in.LittleEndian.ReadUint64()
	case _FormatSint64:
		s.Value, _ = in.LittleEndian.ReadInt64()
	case _FormatReal64:
		s.Value, _ = in.LittleEndian.ReadFloat64()
	case _FormatChar8:
		s.Value, _ = in.ReadString(1)
	case _FormatMixed:
		s.Value, _ = in.ReadBytes(int(s.Length))

	default:
		s.Error = fmt.Errorf("unknown format: %d", s.Format)
	}
}
