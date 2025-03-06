// package svi contains the interface definition of the Standard Variable Interface library.
//
//nolint:lll
package svi

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

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (p *Procedures) GetValue(module res.ModuleNumber, c GetValueCall) (*GetValueReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetValueCall, GetValueReply](10000, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (p *Procedures) SetValue(module res.ModuleNumber, c SetValueCall) (*SetValueReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetValueCall, SetValueReply](10002, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (p *Procedures) GetValues(module res.ModuleNumber, c GetValuesCall) (*GetValuesReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetValuesCall, GetValuesReply](10004, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (p *Procedures) SetValues(module res.ModuleNumber, c SetValuesCall) (*SetValuesReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetValuesCall, SetValuesReply](10006, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (p *Procedures) GetBlock(module res.ModuleNumber, c GetBlockCall) (*GetBlockReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetBlockCall, GetBlockReply](10008, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (p *Procedures) SetBlock(module res.ModuleNumber, c SetBlockCall) (*SetBlockReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetBlockCall, SetBlockReply](10010, rpc.VersionDefault, c))
}

func (p *Procedures) GetAddress(module res.ModuleNumber, c GetAddressCall) (*GetAddressReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetAddressCall, GetAddressReply](10012, rpc.VersionDefault, c))
}

func (p *Procedures) GetServerInfo(module res.ModuleNumber, c GetServerInfoCall) (*GetServerInfoReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetServerInfoCall, GetServerInfoReply](10016, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (p *Procedures) GetMultiBlock(module res.ModuleNumber, c GetMultiBlockCall) (*GetMultiBlockReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetMultiBlockCall, GetMultiBlockReply](10018, rpc.VersionDefault, c))
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (p *Procedures) SetMultiBlock(module res.ModuleNumber, c SetMultiBlockCall) (*SetMultiBlockReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[SetMultiBlockCall, SetMultiBlockReply](10020, rpc.VersionDefault, c))
}

func (p *Procedures) GetExtendedAddress(module res.ModuleNumber, c GetExtendedAddressCall) (*GetExtendedAddressReply, error) {
	return rpc.Call(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewProcedure[GetExtendedAddressCall, GetExtendedAddressReply](10022, rpc.VersionDefault, c))
}

// --------------
// Paginated procedures
// --------------

func (p *Procedures) ListProcessValueInfo(module res.ModuleNumber, c *ListProcessValueInfoCall, pageSize uint32) ([]ProcessValueInfo, error) {
	return rpc.PaginatedCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewPaginatedProcedure[ProcessValueInfo, *ListProcessValueInfoCall, *ListProcessValueInfoReply](10014, rpc.VersionDefault, c), pageSize)
}

func (p *Procedures) ListExtendedProcessValueInfo(module res.ModuleNumber, c *ListExtendedProcessValueInfoCall, pageSize uint32) ([]ExtendedProcessValueInfo, error) {
	c.extendedCallIdentifier = PvInfoExtendedCallIdentifier
	return rpc.PaginatedCall(p.client, rpc.Module{Number: module.ModuleNumber, Port: module.Port}, rpc.NewPaginatedProcedure[ExtendedProcessValueInfo, *ListExtendedProcessValueInfoCall, *ListExtendedProcessValueInfoReply](10014, rpc.VersionDefault, c), pageSize)
}
