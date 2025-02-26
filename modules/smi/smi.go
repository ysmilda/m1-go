// package smi contains the interface definition of the Standard Module Interface library.
//
//nolint:lll
package smi

import "github.com/ysmilda/m1-go/internals/rpc"

var Procedures procedures

type procedures struct{}

func (procedures) Initialise(c InitialiseCall) rpc.Procedure[InitialiseCall, InitialiseReply] {
	return rpc.NewProcedure[InitialiseCall, InitialiseReply](2, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) DeInitialise(c DeInitialiseCall) rpc.Procedure[DeInitialiseCall, DeInitialiseReply] {
	return rpc.NewProcedure[DeInitialiseCall, DeInitialiseReply](4, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) Reset(c ResetCall) rpc.Procedure[ResetCall, ResetReply] {
	return rpc.NewProcedure[ResetCall, ResetReply](6, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) NewConfig(c NewConfigCall) rpc.Procedure[NewConfigCall, NewConfigReply] {
	return rpc.NewProcedure[NewConfigCall, NewConfigReply](8, rpc.VersionDefault, c)
}

func (procedures) GetInfo(c GetInfoCall) rpc.Procedure[GetInfoCall, GetInfoReply] {
	return rpc.NewProcedure[GetInfoCall, GetInfoReply](10, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) EndOfInitialisation(c EndOfInitialiseCall) rpc.Procedure[EndOfInitialiseCall, EndOfInitialiseReply] {
	return rpc.NewProcedure[EndOfInitialiseCall, EndOfInitialiseReply](14, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) SetDebugMode(c SetDebugModeCall) rpc.Procedure[SetDebugModeCall, SetDebugModeReply] {
	return rpc.NewProcedure[SetDebugModeCall, SetDebugModeReply](16, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) Stop(c StopCall) rpc.Procedure[StopCall, StopReply] {
	return rpc.NewProcedure[StopCall, StopReply](18, rpc.VersionDefault, c)
}

// Required permission smi.PermissionsEditSoftwareModule.
func (procedures) Run(c RunCall) rpc.Procedure[RunCall, RunReply] {
	return rpc.NewProcedure[RunCall, RunReply](20, rpc.VersionDefault, c)
}

func (procedures) GetMappedInfoList(c GetMappedInfoListCall) rpc.Procedure[GetMappedInfoListCall, GetMappedInfoListReply] {
	return rpc.NewProcedure[GetMappedInfoListCall, GetMappedInfoListReply](78, rpc.VersionDefault, c)
}

func (procedures) LicenseEvent(c LicenseEventCall) rpc.Procedure[LicenseEventCall, LicenseEventReply] {
	return rpc.NewProcedure[LicenseEventCall, LicenseEventReply](80, rpc.VersionDefault, c)
}

func (procedures) SetSVIAccess(c SetSVIAccessCall) rpc.Procedure[SetSVIAccessCall, SetSVIAccessReply] {
	return rpc.NewProcedure[SetSVIAccessCall, SetSVIAccessReply](82, rpc.VersionDefault, c)
}

func (procedures) ExtendedLogin2(c ExtendedLogin2Call) rpc.Procedure[ExtendedLogin2Call, ExtendedLogin2Reply] {
	return rpc.NewProcedure[ExtendedLogin2Call, ExtendedLogin2Reply](84, rpc.VersionDefault, c)
}

func (procedures) ExtendedRequestAccess(c ExtendedRequestAccessCall) rpc.Procedure[ExtendedRequestAccessCall, ExtendedRequestAccessReply] {
	return rpc.NewProcedure[ExtendedRequestAccessCall, ExtendedRequestAccessReply](86, rpc.VersionDefault, c)
}

func (procedures) ExtendedReleaseAccess(c ExtendedReleaseAccessCall) rpc.Procedure[ExtendedReleaseAccessCall, ExtendedReleaseAccessReply] {
	return rpc.NewProcedure[ExtendedReleaseAccessCall, ExtendedReleaseAccessReply](88, rpc.VersionDefault, c)
}

func (procedures) Alive(c AliveCall) rpc.Procedure[AliveCall, AliveReply] {
	return rpc.NewProcedure[AliveCall, AliveReply](92, rpc.VersionDefault, c)
}

func (procedures) ExtendedLogin(c ExtendedLoginCall) rpc.Procedure[ExtendedLoginCall, ExtendedLoginReply] {
	return rpc.NewProcedure[ExtendedLoginCall, ExtendedLoginReply](84, rpc.VersionDefault, c)
}

func (procedures) ExtendedLogout(c ExtendedLogoutCall) rpc.Procedure[ExtendedLogoutCall, ExtendedLogoutReply] {
	return rpc.NewProcedure[ExtendedLogoutCall, ExtendedLogoutReply](96, rpc.VersionDefault, c)
}
