package svi

type AddressType uint32

const (
	AddressTypeA AddressType = 0x0001
	AddressTypeB AddressType = 0x0002
)

type FlagType uint16

const (
	FlagTypeDirectory FlagType = 0x0000
	FlagTypeVariable  FlagType = 0x0001
	FlagTypeMarker    FlagType = 0x0002
)

type VariableType uint8

const (
	VariableTypeSymbolicValue   VariableType = 0  // Symbolic value, any type
	VariableTypeMarkerBit       VariableType = 1  // 1 bit value
	VariableTypeMarkerByte      VariableType = 2  // 8 bit value
	VariableTypeMarkerDWord     VariableType = 3  // 32 bit integer
	VariableTypeMarkerFloat     VariableType = 4  // 32 bit float
	VariableTypeMarkerWord      VariableType = 5  // 16 bit integer
	VariableTypeMarkerLong      VariableType = 6  // 64 bit integer
	VariableTypeMarkerReal      VariableType = 7  // 64 bit float
	VariableTypeGlobalDWord     VariableType = 8  // 32 bit integer
	VariableTypeRetainDWord     VariableType = 9  // 32 bit integer
	VariableTypePLCRetainDword  VariableType = 19 // Retain DWord used in smi.pld
	VariableTypeMixedBlockArray VariableType = 21
	VariableTypeCNCDataMarker   VariableType = 22 // 32 bit float
	VariableTypeProfileMarker   VariableType = 23 // 32 bit float
	VariableTypeOptionalMarker  VariableType = 24 // 32 bit float
	VariableTypePLCGlobalDWord  VariableType = 35 // Global DWord used in smi.pld
)
