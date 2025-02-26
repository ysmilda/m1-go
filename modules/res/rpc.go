//nolint:unused
package res

import (
	"net"

	"github.com/ysmilda/m1-go/internals/rpc"
)

var Module = ModuleNumber{
	ModuleNumber: (0x20000000 | 0x00001000),
	Port:         3000,
}

type (
	ModuleInfoCall struct {
		Name string `m1binary:"length:12"`
	}

	ModuleInfoReply struct {
		rpc.ReturnCode
		ModuleInfo
	}

	ListModuleInfoCall struct {
		First uint32
		Last  uint32
	}

	ListModuleInfoReply struct {
		rpc.ReturnCode
		Count   uint32
		Modules []ModuleInfo `m1binary:"lengthRef:Count"`
	}

	ExtendedModuleInfoCall struct {
		Name string `m1binary:"length:12"`
	}

	ExtendedModuleInfoReply struct {
		rpc.ReturnCode
		ExtendedModuleInfo
	}

	ExtendedModuleInfoListCall struct {
		First uint32
		Last  uint32
	}

	ExtendedModuleInfoListReply struct {
		rpc.ReturnCode
		Count   uint32
		Modules []ExtendedModuleInfo `m1binary:"lengthRef:Count"`
	}

	ModuleChildListCall struct {
		ParentTaskID uint32
	}

	ModuleChildListReply struct {
		rpc.ReturnCode
		ParentTask ModuleTaskInfo
		Count      uint32
		Children   []ModuleTaskInfo `m1binary:"lengthRef:Count"`
	}

	ModuleTaskListCall struct {
		Appname string `m1binary:"length:12"`
	}

	ModuleTaskListReply struct {
		rpc.ReturnCode
		Count       uint32           `m1binary:"skip:16"`
		ModuleTasks []ModuleTaskInfo `m1binary:"lengthRef:couCountnt"`
	}

	ModuleAccessCall struct {
		IPAddress    net.IP `m1binary:"length:4"` // IP address of the caller
		ModuleNumber uint32 // Module number of the caller
		Appname      string `m1binary:"length:12"`
	}

	ModuleAccessReply struct {
		rpc.ReturnCode
		ModuleNumber  uint32
		UDPPortNumber uint16
		TCPPortNumber uint16
	}

	ModuleFreeCall struct {
		IPAddress    net.IP `m1binary:"length:4"` // IP address of the caller
		ModuleNumber uint32 // Module number of the caller
		Appname      string `m1binary:"length:12"`
	}

	ModuleFreeReply struct {
		rpc.ReturnCode
	}

	ModuleNumberCall struct {
		Name string `m1binary:"length:12"`
	}

	ModuleNumberReply struct {
		rpc.ReturnCode
		ModuleNumber
	}

	SystemInfoCall struct {
		parameter   uint32 // Must be zero
		MainVersion uint32 // Tool main version, 0 = not relevant
		SubVersion  uint32 // Tool sub version, 0 = not relevant
		Toolname    string `m1binary:"length:12"`
	}

	SystemInfoReply struct {
		rpc.ReturnCode
		SystemInfo
	}

	LoginCall struct {
		SystemInfoCall
		Username string `m1binary:"length:20"`
		Password string `m1binary:"length:16"`
	}

	LoginReply struct {
		rpc.ReturnCode
		Login
	}

	Login2Call struct {
		SystemInfoCall
		Username string `m1binary:"length:64"`
		Password string `m1binary:"length:32"`
		Local    bool   `m1binary:"skip:19"`
	}

	Login2Reply struct {
		rpc.ReturnCode
		Login2
	}

	LogoutCall struct {
		parameter uint32 // Must be zero
		authEnt   []byte `m1binary:"length:128"` // Deprecated
	}

	LogoutReply struct {
		rpc.ReturnCode
	}

	OpenCall struct {
		RequestedSessionTimeout        uint32
		RequestedSessionLifetime       uint32
		RequestedSMISize               uint32
		RequestedSessionIdlePrevention bool `m1binary:"skip:127"`
	}

	OpenReply struct {
		rpc.ReturnCode
		Open
	}

	CloseCall struct {
		Spare []byte `m1binary:"length:8"`
	}

	CloseReply struct {
		rpc.ReturnCode
	}

	RenewCall struct {
		RequestAuthenticationRenewal bool `m1binary:"skip:147"`
	}

	RenewReply struct {
		rpc.ReturnCode
		Renew
	}

	ExtPingCall struct {
		IPMask   net.IPMask `m1binary:"length:4"`
		Filter   Filter
		Reply    ReplyType
		Mode     ReplyMode
		Reserved []byte `m1binary:"length:12"`
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
		SerialNumber string `m1binary:"length:12,skip:16"`
	}

	FlashLEDReply struct {
		rpc.ReturnCode `m1binary:"skip:16"`
	}
)
