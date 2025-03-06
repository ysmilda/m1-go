// package smi contains the interface definition of the Standard Module Interface library.
//
//nolint:lll
package smi

import (
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/modules/res"
)

func NewProcedures(client *m1client.Client) *Procedures {
	return &Procedures{
		client: client,
	}
}

type Procedures struct {
	client *m1client.Client
}

func (p *Procedures) Initialise(module res.ModuleNumber, c InitialiseCall) (*InitialiseReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[InitialiseCall, InitialiseReply](2, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) DeInitialise(module res.ModuleNumber, c DeInitialiseCall) (*DeInitialiseReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[DeInitialiseCall, DeInitialiseReply](4, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) Reset(module res.ModuleNumber, c ResetCall) (*ResetReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ResetCall, ResetReply](6, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) NewConfig(module res.ModuleNumber, c NewConfigCall) (*NewConfigReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[NewConfigCall, NewConfigReply](8, rpc.VersionDefault, c))
}

func (p *Procedures) GetInfo(module res.ModuleNumber, c GetInfoCall) (*GetInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetInfoCall, GetInfoReply](10, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) EndOfInitialisation(module res.ModuleNumber, c EndOfInitialiseCall) (*EndOfInitialiseReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[EndOfInitialiseCall, EndOfInitialiseReply](14, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) SetDebugMode(module res.ModuleNumber, c SetDebugModeCall) (*SetDebugModeReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetDebugModeCall, SetDebugModeReply](16, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) Stop(module res.ModuleNumber, c StopCall) (*StopReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[StopCall, StopReply](18, rpc.VersionDefault, c))
}

// Required permission smi.PermissionsEditSoftwareModule.
func (p *Procedures) Run(module res.ModuleNumber, c RunCall) (*RunReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[RunCall, RunReply](20, rpc.VersionDefault, c))
}

func (p *Procedures) GetMappedInfoList(module res.ModuleNumber, c GetMappedInfoListCall) (*GetMappedInfoListReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetMappedInfoListCall, GetMappedInfoListReply](78, rpc.VersionDefault, c))
}

func (p *Procedures) LicenseEvent(module res.ModuleNumber, c LicenseEventCall) (*LicenseEventReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[LicenseEventCall, LicenseEventReply](80, rpc.VersionDefault, c))
}

func (p *Procedures) SetSVIAccess(module res.ModuleNumber, c SetSVIAccessCall) (*SetSVIAccessReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetSVIAccessCall, SetSVIAccessReply](82, rpc.VersionDefault, c))
}

func (p *Procedures) ExtendedLogin2(module res.ModuleNumber, c ExtendedLogin2Call) (*ExtendedLogin2Reply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedLogin2Call, ExtendedLogin2Reply](84, rpc.VersionDefault, c))
}

func (p *Procedures) ExtendedRequestAccess(module res.ModuleNumber, c ExtendedRequestAccessCall) (*ExtendedRequestAccessReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedRequestAccessCall, ExtendedRequestAccessReply](86, rpc.VersionDefault, c))
}

func (p *Procedures) ExtendedReleaseAccess(module res.ModuleNumber, c ExtendedReleaseAccessCall) (*ExtendedReleaseAccessReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedReleaseAccessCall, ExtendedReleaseAccessReply](88, rpc.VersionDefault, c))
}

func (p *Procedures) Alive(module res.ModuleNumber, c AliveCall) (*AliveReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[AliveCall, AliveReply](92, rpc.VersionDefault, c))
}

func (p *Procedures) ExtendedLogin(module res.ModuleNumber, c ExtendedLoginCall) (*ExtendedLoginReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedLoginCall, ExtendedLoginReply](84, rpc.VersionDefault, c))
}

func (p *Procedures) ExtendedLogout(module res.ModuleNumber, c ExtendedLogoutCall) (*ExtendedLogoutReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[ExtendedLogoutCall, ExtendedLogoutReply](96, rpc.VersionDefault, c))
}
