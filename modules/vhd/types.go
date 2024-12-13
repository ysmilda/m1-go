package vhd

type ReportMode uint32

const (
	ReportModeNoSort             ReportMode = 0
	ReportModeExtendedError      ReportMode = 0x10
	ReportModeMultiBlockTransfer ReportMode = 0x20
)

type AddressType uint32

const (
	AddressTypeSVISingleBaseType        AddressType = 0
	AddressTypeSVIListBaseType          AddressType = 1
	AddressTypeSVIConsecutiveBaseType   AddressType = 2
	AddressTypeSVISingleCompoundType    AddressType = 3
	AddressTypeIndexBaseType            AddressType = 4
	AddressTypeIndexListBaseType        AddressType = 5
	AddressTypeIndexConsecutiveBaseType AddressType = 6
	AddressTypeIndexSingleCompoundType  AddressType = 7
)

type SessionMode uint32

const (
	SessionModeNormal    SessionMode = 0 // Normal polling mode
	SessionModeCallback  SessionMode = 1 // Callback mode
	SessionModeIdleWatch SessionMode = 2 // Normal polling mode + idle watch
)

type ProtocolType uint32

const (
	ProtocolTypeUDP ProtocolType = 1
	ProtocolTypeTCP ProtocolType = 2
)

type ObservationMode uint32

const (
	// Value changes will be reported to the client automatically.
	ObservationModeAuto ObservationMode = 1
	// Value changes will be reported to the client only when requested.
	ObservationModeRequest ObservationMode = 2
	// Value changes will be reported to the client automatically after a trigger is received.
	ObservationModeTrigger ObservationMode = 3
)
