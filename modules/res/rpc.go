//nolint:unused
package res

import (
	"net"

	"github.com/ysmilda/m1-go/internals/rpc"
)

var Module = ModuleNumber{
	Number: (0x20000000 | 0x00001000),
	Port:   3000,
}

type (
	ModuleInfoCall struct {
		Name string `binary:"length:12"`
	}

	ModuleInfoReply struct {
		rpc.ReturnCode
		ModuleInfo
	}

	ListModuleInfoCall struct {
		rpc.PaginatedCallFirstLast
	}

	ListModuleInfoReply struct {
		rpc.ReturnCode
		rpc.PaginatedReplyCount[ModuleInfo]
	}

	ExtendedModuleInfoCall struct {
		Name string `binary:"length:12"`
	}

	ExtendedModuleInfoReply struct {
		rpc.ReturnCode
		ExtendedModuleInfo
	}

	ExtendedModuleInfoListCall struct {
		rpc.PaginatedCallFirstLast
	}

	ExtendedModuleInfoListReply struct {
		rpc.ReturnCode
		rpc.PaginatedReplyCount[ExtendedModuleInfo]
	}

	ModuleChildCall struct {
		ParentTaskID uint32
	}

	ModuleChildReply struct {
		rpc.ReturnCode
		ParentTask ModuleTaskInfo
		Count      uint32
		Children   []ModuleTaskInfo `binary:"lengthRef:Count"`
	}

	ModuleTaskCall struct {
		Appname string `binary:"length:12"`
	}

	ModuleTaskReply struct {
		rpc.ReturnCode
		Count       uint32           `binary:"skip:16"`
		ModuleTasks []ModuleTaskInfo `binary:"lengthRef:Count"`
	}

	ModuleAccessCall struct {
		IPAddress    net.IP `binary:"length:4"` // IP address of the caller
		ModuleNumber uint32 // Module number of the caller
		Appname      string `binary:"length:12"`
	}

	ModuleAccessReply struct {
		rpc.ReturnCode
		ModuleNumber  uint32
		UDPPortNumber uint16
		TCPPortNumber uint16
	}

	ModuleFreeCall struct {
		IPAddress    net.IP `binary:"length:4"` // IP address of the caller
		ModuleNumber uint32 // Module number of the caller
		Appname      string `binary:"length:12"`
	}

	ModuleFreeReply struct {
		rpc.ReturnCode
	}

	ModuleNumberCall struct {
		Name string `binary:"length:12"`
	}

	ModuleNumberReply struct {
		rpc.ReturnCode
		ModuleNumber
	}

	SystemInfoCall struct {
		Parameter   uint32 // Must be zero
		MainVersion uint32 // Tool main version, 0 = not relevant
		SubVersion  uint32 // Tool sub version, 0 = not relevant
		Toolname    string `binary:"length:12"`
	}

	SystemInfoReply struct {
		rpc.ReturnCode
		SystemInfo
	}

	LoginCall struct {
		SystemInfoCall
		Username string `binary:"length:20"`
		Password string `binary:"length:16"`
	}

	LoginReply struct {
		rpc.ReturnCode
		Login
	}

	Login2Call struct {
		SystemInfoCall
		Username string `binary:"length:64"`
		Password string `binary:"length:32"`
		Local    bool   `binary:"skip:19"`
	}

	Login2Reply struct {
		rpc.ReturnCode
		Login2
	}

	LogoutCall struct {
		Parameter uint32 // Must be zero
		AuthEnt   []byte `binary:"length:128"` // Deprecated
	}

	LogoutReply struct {
		rpc.ReturnCode
	}

	OpenCall struct {
		RequestedSessionTimeout        uint32
		RequestedSessionLifetime       uint32
		RequestedSMISize               uint32
		RequestedSessionIdlePrevention bool `binary:"skip:127"`
	}

	OpenReply struct {
		rpc.ReturnCode
		Open
	}

	CloseCall struct {
		Spare []byte `binary:"length:8"`
	}

	CloseReply struct {
		rpc.ReturnCode
	}

	RenewCall struct {
		RequestAuthenticationRenewal bool `binary:"skip:147"`
	}

	RenewReply struct {
		rpc.ReturnCode
		Renew
	}

	ExtPingCall struct {
		IPMask   net.IPMask `binary:"length:4"`
		Filter   Filter
		Reply    ReplyType
		Mode     ReplyMode
		Reserved []byte `binary:"length:12"`
	}

	ExtPingReply struct {
		rpc.ReturnCode
		ExtPing
	}

	ExtPingReply2 struct {
		rpc.ReturnCode
		ExtPing2
	}

	FlashLEDCall struct {
		SerialNumber string `binary:"length:12,skip:16"`
	}

	FlashLEDReply struct {
		rpc.ReturnCode `binary:"skip:16"`
	}
)
