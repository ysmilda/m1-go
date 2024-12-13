package vhd

// ElementAddress is a generic struct to represent different types of addresses.
// The type of the element should match the given AddressType. It is up to the user to ensure this.
type ElementAddress struct {
	Type    AddressType
	Element AddressI
}

// AddressI acts as a marker interface for the different types of addresses.
// It ensures that the ElementAddress struct can only be used with the correct types.
type AddressI interface {
	noop() uint64
}

type ElementAddressSVISingleBaseType struct {
	AddressI
	Address uint64
}

type ElementAddressSVISVIListBaseType struct {
	AddressI
	NumberOfElements uint32
	Elements         []ElementAddressSVISingleBaseType `unpack:"length=NumberOfElements"`
}

type ElementAddressSVIConsecutiveBaseType struct {
	AddressI
	NumberOfElements uint32
	Address          uint64
}

type ElementAddressSVISingleCompoundType struct {
	AddressI
	ElementLength uint32
	Address       uint64
}

type ElementAddressIndexBaseType struct {
	AddressI
	Address uint64
}

type ElementAddressIndexListBaseType struct {
	AddressI
	NumberOfElements uint32
	Elements         []ElementAddressIndexBaseType `unpack:"length=NumberOfElements"`
}

type ElementAddressIndexConsecutiveBaseType struct {
	AddressI
	NumberOfElements uint32
	Address          uint64
}

type ElementAddressIndexSingleCompoundType struct {
	AddressI
	ElementLength uint32
	Address       uint64
}

// ElementValue is a generic struct to represent different types of values.
// The type of the element should match the given AddressType. It is up to the user to ensure this.
type ElementValue struct {
	Type  AddressType
	Value ValueI
}

type ValueI interface {
	noop() uint64
}

type ElementValueSVISingleBaseType struct {
	ValueI
	Address uint64
	Value   uint32
}

type ElementValueSVIListBaseType struct {
	ValueI
	NumberOfElements uint32
	Elements         []ElementValueSVISingleBaseType `unpack:"length=NumberOfElements"`
}

type ElementValueSVIConsecutiveBaseTypeValue struct {
	ValueI
	NumberOfElements uint32
	Address          uint64
	Values           []uint32 `unpack:"length=NumberOfElements"`
}

type ElementValueSVISingleCompoundType struct {
	ValueI
	ElementLength uint32
	Address       uint64
	Value         []byte `unpack:"length=ElementLength"`
}

type ElementValueIndexBaseType struct {
	ValueI
	Index uint64
	Value uint32
}

type ElementValueIndexListBaseType struct {
	ValueI
	NumberOfElements uint32
	Elements         []ElementValueIndexBaseType `unpack:"length=NumberOfElements"`
}

type ElementValueIndexConsecutiveBaseType struct {
	ValueI
	NumberOfElements uint32
	Index            uint64
	Values           []uint32 `unpack:"length=NumberOfElements"`
}

type ElementValueIndexSingleCompoundType struct {
	ValueI
	ElementLength uint32
	Index         uint64
	Value         []byte `unpack:"length=ElementLength"`
}
