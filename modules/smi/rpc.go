package smi

import (
	"net"
	"time"

	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

type (
	// ModuleNameCall is used by procedures that only require a module name.
	ModuleNameCall struct {
		ModuleName string `binary:"length:12"`
	}

	InitialiseCall      = ModuleNameCall
	DeInitialiseCall    = ModuleNameCall
	ResetCall           = ModuleNameCall
	NewConfigCall       = ModuleNameCall
	GetInfoCall         = ModuleNameCall
	EndOfInitialiseCall = ModuleNameCall
	StopCall            = ModuleNameCall
	RunCall             = ModuleNameCall

	InitialiseReply struct {
		rpc.ReturnCode
	}

	DeInitialiseReply struct {
		rpc.ReturnCode
	}

	ResetReply struct {
		rpc.ReturnCode
	}

	NewConfigReply struct {
		rpc.ReturnCode
	}

	GetInfoReply struct {
		rpc.ReturnCode
		ModuleName  string `binary:"length:12"`
		Description string `binary:"length:80"`
		Version     msys.Version
		State       res.ResourceState
		DebugMode   uint32
	}

	GetMappedInfoListCall struct {
		StartIndex uint32 `binary:"skip:16"`
	}

	GetMappedInfoListReply struct {
		rpc.ReturnCode
		NumberOfInputChannels  uint32
		NumberOfOutputChannels uint32
		Count                  uint32                        `binary:"skip:16"`
		Channels               []ComponentManagerChannelInfo `binary:"lengthRef:Count"`
	}

	EndOfInitialiseReply struct {
		rpc.ReturnCode
	}

	SetDebugModeCall struct {
		ModuleName string `binary:"length:12"`
		DebugMode  uint32
	}

	SetDebugModeReply struct {
		rpc.ReturnCode
	}

	StopReply struct {
		rpc.ReturnCode
	}

	RunReply struct {
		rpc.ReturnCode
	}

	AliveCall struct {
		IPAddress      net.IP `binary:"length:4"`
		LoginSessionID uint32
		Internal       bool   `binary:"skip:3"`
		Username       string `binary:"length:20"`
	}

	AliveReply struct {
		rpc.ReturnCode `binary:"skip:8"`
	}

	ExtendedLoginCall struct {
		UserParameter uint32
		IPAddress     net.IP `binary:"length:4"`
		SessionID     uint32 `binary:"skip:4"`
		Username      string `binary:"length:20"`
		Password      string `binary:"length:16"`
	}

	ExtendedLoginReply struct {
		rpc.ReturnCode
		Permissions res.Permissions
		UserData    []byte `binary:"length:128"`
	}

	ExtendedLogin2Call struct {
		UserParameter uint32
		UserAuth      res.UserAuth
		Password      string `binary:"length:32,skip:8"`
	}

	ExtendedLogin2Reply struct {
		rpc.ReturnCode
		UserAccess res.UserAccess
		DelayTime  time.Duration `binary:"length:4,unit:milliseconds,skip:4"` // 0=auto
		UserData   []byte        `binary:"length:128"`
	}

	ExtendedRequestAccessCall struct {
		UserParameter uint32 `binary:"skip:12"`
	}

	ExtendedRequestAccessReply struct {
		rpc.ReturnCode `binary:"skip:16"`
	}

	ExtendedReleaseAccessCall struct {
		UserParameter uint32 `binary:"skip:12"`
	}

	ExtendedReleaseAccessReply struct {
		rpc.ReturnCode `binary:"skip:16"`
	}

	SetSVIAccessCall struct {
		ModuleName string `binary:"length:12,skip:16"`
	}

	SetSVIAccessReply struct {
		rpc.ReturnCode `binary:"skip:16"`
	}

	LicenseEventCall struct {
		Event LicenseEvent `binary:"skip:16"`
	}

	LicenseEventReply struct {
		rpc.ReturnCode `binary:"skip:16"`
	}

	ExtendedLogoutCall struct {
		UserParameter  uint32
		IPAddress      net.IP `binary:"length:4"`
		LoginSessionID uint32 `binary:"skip:152"`
	}

	ExtendedLogoutReply struct {
		rpc.ReturnCode `binary:"skip:8"`
	}
)
