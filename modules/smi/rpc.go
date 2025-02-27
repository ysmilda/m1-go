package smi

import (
	"net"
	"time"

	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

type (
	InitialiseCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	InitialiseReply struct {
		rpc.ReturnCode
	}

	DeInitialiseCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	DeInitialiseReply struct {
		rpc.ReturnCode
	}

	ResetCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	ResetReply struct {
		rpc.ReturnCode
	}

	NewConfigCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	NewConfigReply struct {
		rpc.ReturnCode
	}

	GetInfoCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	GetInfoReply struct {
		rpc.ReturnCode
		ModuleName  string `m1binary:"length:12"`
		Description string `m1binary:"length:80"`
		Version     msys.Version
		State       res.ResourceState
		DebugMode   uint32
	}

	GetMappedInfoListCall struct {
		StartIndex uint32 `m1binary:"skip:16"`
	}

	GetMappedInfoListReply struct {
		rpc.ReturnCode
		NumberOfInputChannels  uint32
		NumberOfOutputChannels uint32
		Count                  uint32                        `m1binary:"skip:16"`
		Channels               []ComponentManagerChannelInfo `m1binary:"lengthRef:Count"`
	}

	EndOfInitialiseCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	EndOfInitialiseReply struct {
		rpc.ReturnCode
	}

	SetDebugModeCall struct {
		ModuleName string `m1binary:"length:12"`
		DebugMode  uint32
	}

	SetDebugModeReply struct {
		rpc.ReturnCode
	}

	StopCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	StopReply struct {
		rpc.ReturnCode
	}

	RunCall struct {
		ModuleName string `m1binary:"length:12"`
	}

	RunReply struct {
		rpc.ReturnCode
	}

	AliveCall struct {
		IPAddress      net.IP `m1binary:"length:4"`
		LoginSessionID uint32
		Internal       bool   `m1binary:"skip:3"`
		Username       string `m1binary:"length:20"`
	}

	AliveReply struct {
		rpc.ReturnCode `m1binary:"skip:8"`
	}

	ExtendedLoginCall struct {
		UserParameter uint32
		IPAddress     net.IP `m1binary:"length:4"`
		SessionID     uint32 `m1binary:"skip:4"`
		Username      string `m1binary:"length:20"`
		Password      string `m1binary:"length:16"`
	}

	ExtendedLoginReply struct {
		rpc.ReturnCode
		Permissions res.Permissions
		UserData    []byte `m1binary:"length:128"`
	}

	ExtendedLogin2Call struct {
		UserParameter uint32
		UserAuth      res.UserAuth
		Password      string `m1binary:"length:32,skip:8"`
	}

	ExtendedLogin2Reply struct {
		rpc.ReturnCode
		UserAccess res.UserAccess
		DelayTime  time.Duration `m1binary:"length:4,unit:milliseconds,skip:4"` // 0=auto
		UserData   []byte        `m1binary:"length:128"`
	}

	ExtendedRequestAccessCall struct {
		UserParameter uint32 `m1binary:"skip:12"`
	}

	ExtendedRequestAccessReply struct {
		rpc.ReturnCode `m1binary:"skip:16"`
	}

	ExtendedReleaseAccessCall struct {
		UserParameter uint32 `m1binary:"skip:12"`
	}

	ExtendedReleaseAccessReply struct {
		rpc.ReturnCode `m1binary:"skip:16"`
	}

	SetSVIAccessCall struct {
		ModuleName string `m1binary:"length:12,skip:16"`
	}

	SetSVIAccessReply struct {
		rpc.ReturnCode `m1binary:"skip:16"`
	}

	LicenseEventCall struct {
		Event LicenseEvent `m1binary:"skip:16"`
	}

	LicenseEventReply struct {
		rpc.ReturnCode `m1binary:"skip:16"`
	}

	ExtendedLogoutCall struct {
		UserParameter  uint32
		IPAddress      net.IP `m1binary:"length:4"`
		LoginSessionID uint32 `m1binary:"skip:152"`
	}

	ExtendedLogoutReply struct {
		rpc.ReturnCode `m1binary:"skip:8"`
	}
)
