// package svi contains the interface definition of the Standard Variable Interface library.
//
//nolint:lll
package svi

import "github.com/ysmilda/m1-go/internals/rpc"

var Procedures procedures

type procedures struct{}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (procedures) GetValue(c GetValueCall) rpc.Procedure[GetValueCall, GetValueReply] {
	return rpc.NewProcedure[GetValueCall, GetValueReply](10000, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (procedures) SetValue(c SetValueCall) rpc.Procedure[SetValueCall, SetValueReply] {
	return rpc.NewProcedure[SetValueCall, SetValueReply](10002, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (procedures) GetValues(c GetValuesCall) rpc.Procedure[GetValuesCall, GetValuesReply] {
	return rpc.NewProcedure[GetValuesCall, GetValuesReply](10004, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (procedures) SetValues(c SetValuesCall) rpc.Procedure[SetValuesCall, SetValuesReply] {
	return rpc.NewProcedure[SetValuesCall, SetValuesReply](10006, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (procedures) GetBlock(c GetBlockCall) rpc.Procedure[GetBlockCall, GetBlockReply] {
	return rpc.NewProcedure[GetBlockCall, GetBlockReply](10008, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (procedures) SetBlock(c SetBlockCall) rpc.Procedure[SetBlockCall, SetBlockReply] {
	return rpc.NewProcedure[SetBlockCall, SetBlockReply](10010, rpc.VersionDefault, c)
}

func (procedures) GetAddress(c GetAddressCall) rpc.Procedure[GetAddressCall, GetAddressReply] {
	return rpc.NewProcedure[GetAddressCall, GetAddressReply](10012, rpc.VersionDefault, c)
}

func (procedures) GetProcessValueInfo(c GetProcessValueInfoCall) rpc.Procedure[GetProcessValueInfoCall, GetProcessValueInfoReply] {
	return rpc.NewProcedure[GetProcessValueInfoCall, GetProcessValueInfoReply](10014, rpc.VersionDefault, c)
}

func (procedures) GetExtendedProcessValueInfo(c GetExtendedProcessValueInfoCall) rpc.Procedure[GetExtendedProcessValueInfoCall, GetExtendedProcessValueInfoReply] {
	c.extendedCallIdentifier = PvInfoExtendedCallIdentifier
	return rpc.NewProcedure[GetExtendedProcessValueInfoCall, GetExtendedProcessValueInfoReply](10014, rpc.VersionDefault, c)
}

func (procedures) GetServerInfo(c GetServerInfoCall) rpc.Procedure[GetServerInfoCall, GetServerInfoReply] {
	return rpc.NewProcedure[GetServerInfoCall, GetServerInfoReply](10016, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsReadSVI or smi.PermissionsReadMIOInputs in case of MIO variables.
func (procedures) GetMultiBlock(c GetMultiBlockCall) rpc.Procedure[GetMultiBlockCall, GetMultiBlockReply] {
	return rpc.NewProcedure[GetMultiBlockCall, GetMultiBlockReply](10018, rpc.VersionDefault, c)
}

// Required permissions: smi.PermissionsWriteSVI or smi.PermissionsWriteMIOOutputs in case of MIO variables.
func (procedures) SetMultiBlock(c SetMultiBlockCall) rpc.Procedure[SetMultiBlockCall, SetMultiBlockReply] {
	return rpc.NewProcedure[SetMultiBlockCall, SetMultiBlockReply](10020, rpc.VersionDefault, c)
}

func (procedures) GetExtendedAddress(c GetExtendedAddressCall) rpc.Procedure[GetExtendedAddressCall, GetExtendedAddressReply] {
	return rpc.NewProcedure[GetExtendedAddressCall, GetExtendedAddressReply](10022, rpc.VersionDefault, c)
}
