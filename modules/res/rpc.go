package res

type ModuleInfoCall struct {
	Name string `unpack:"length=12"`
}

type ModInfoReply struct {
	ReturnCode
	ModuleInfo
}

type ModuleInfoListCall struct {
	First uint32
	Last  uint32
}

type ModuleInfoListReply struct {
	ReturnCode
	NumberOfModules uint32
	Modules         []ModuleInfo `unpack:"length=NumberOfModules"`
}

type ModuleXInfoCall struct {
	Name string `unpack:"length=12"`
}

type ModuleXInfoReply struct {
	ReturnCode
	ModuleXInfo
}

type ModuleXInfoListCall struct {
	First uint32
	Last  uint32
}

type ModuleXInfoListReply struct {
	ReturnCode
	NumberOfModules uint32
	Modules         []ModuleXInfo `unpack:"length=NumberOfModules"`
}

type ModuleChildListCall struct {
	ParentTaskID uint32
}

type ModuleChildListReply struct {
	ReturnCode
	ParentTask       ModuleTaskInfo
	NumberOfChildren uint32
	Children         []ModuleTaskInfo `unpack:"length=NumberOfChildren"`
}

type ModuleTaskListCall struct {
	Appname string `unpack:"length=12"`
}

type ModuleTaskListReply struct {
	ReturnCode
	NumberOfModuleTasks uint32
	Spare               []uint32         `unpack:"length=4"`
	ModuleTasks         []ModuleTaskInfo `unpack:"length=NumberOfModuleTasks"`
}

type ModuleAccessCall struct {
	IPAddress    uint32 // IP address of the caller
	ModuleNumber uint32 // Module number of the caller
	Appname      string `unpack:"length=12"`
}

type ModuleAccessReply struct {
	ReturnCode
	ModuleNumber  uint32
	UDPPortNumber uint16
	TCPPortNumber uint16
}

type ModuleFreeCall struct {
	IPAddress    uint32 // IP address of the caller
	ModuleNumber uint32 // Module number of the caller
	Appname      string `unpack:"length=12"`
}

type ModuleFreeReply struct {
	ReturnCode
}

type ModuleNumberCall struct {
	Name string `unpack:"length=12"`
}

type ModuleNumberReply struct {
	ReturnCode
	ModuleNumber
}

type UserInfoCall struct {
	Appname string `unpack:"length=12"`
}

type UserInfoReply struct {
	ReturnCode
	NumberOfUsers uint32
	Users         []UserInfo `unpack:"length=NumberOfUsers"`
}

type SystemInfoCall struct {
	Parameter   uint32 // Not used, must be zero
	MainVersion uint32 // Tool main version, 0 means not relevant
	SubVersion  uint32 // Tool sub version, 0 means not relevant
	Toolname    string `unpack:"length=12"`
}

type SystemInfoReply struct {
	ReturnCode
	SystemInfo
}

type LoginCall struct {
	SystemInfoCall
	Username string `unpack:"length=20"`
	Password string `unpack:"length=16"`
}

type LoginReply struct {
	ReturnCode
	Login
}

type Login2Call struct {
	SystemInfoCall
	Username string `unpack:"length=64"`
	Password string `unpack:"length=32"`
	Local    bool
	Spare    []byte `unpack:"length=19"`
}

type Login2Reply struct {
	ReturnCode
	Login2
}

type LogoutCall struct {
	Parameter uint32 // Not used, must be zero
	Authent   []byte `unpack:"length=128"` // Deprecated
}

type LogoutReply struct {
	ReturnCode
}

type OpenCall struct {
	RequestedSessionTimeout        uint32
	RequestedSessionLifetime       uint32
	RequestedSMISize               uint32
	RequestedSessionIdlePrevention bool
	Spare                          []byte `unpack:"length=127"`
}

type OpenReply struct {
	ReturnCode
	Open
}

type CloseCall struct {
	Spare []byte `unpack:"length=8"`
}

type CloseReply struct {
	ReturnCode
}

type RenewCall struct {
	RequestAuthenticationRenewal bool
	Spare                        []byte `unpack:"length=147"`
}

type RenewReply struct {
	ReturnCode
	Renew
}

type ExtPingCall struct {
	IpMask   uint32 // May be constructed by using binary.LittleEndian.Uint32(net.IPv4(255, 255, 255, 255).To4())
	Filter   Filter
	Reply    ReplyType
	Mode     ReplyMode
	Reserved []byte `unpack:"length=12"`
}

type ExtPingReply struct {
	ReturnCode
	ExtPing
}

type ExtPingReply2 struct {
	ReturnCode
	ExtPing2
}

type FlashLEDCall struct {
	SerialNumber string `unpack:"length=12"`
	Spare        []byte `unpack:"length=16"`
}

type ReturnCode uint32

func (r ReturnCode) GetReturnCode() uint32 { return uint32(r) }
