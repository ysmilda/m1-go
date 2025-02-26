package svi

type Address uint64

func (a Address) TypeA() AddressClassA {
	return AddressClassA(a)
}

func (a Address) TypeB() AddressClassB {
	return AddressClassB(a)
}

type AddressClassA Address

func (a AddressClassA) Indexes() (uint16, uint16) {
	return uint16(a & 0xFFFF), uint16((a >> 16) & 0xFFFF)
}

func (a AddressClassA) VariableType() VariableType {
	return VariableType(uint8((a >> 32) & 0xFF))
}

func (a AddressClassA) Node() uint8 {
	return uint8((a >> 40) & 0xFF)
}

func (a AddressClassA) Net() uint8 {
	return uint8((a >> 48) & 0xFF)
}

func (a AddressClassA) Description() uint8 {
	return uint8((a >> 58) & 0x7F)
}

func (a AddressClassA) Class() bool {
	return ((a >> 63) & 0x01) == 1
}

type AddressClassB Address

func (a AddressClassB) Indexes() (uint16, uint16) {
	return uint16(a & 0xFFFF), uint16((a >> 16) & 0xFFFF)
}

func (a AddressClassB) Type() uint8 {
	return uint8(a>>32) & 0x3F
}

func (a AddressClassB) ServiceFlag() bool {
	return ((a >> 38) & 0x1) == 1
}

func (a AddressClassB) Flag3() bool {
	return ((a >> 39) & 0x1) == 1
}

func (a AddressClassB) Format() uint8 {
	return uint8((a >> 40) & 0xFF)
}

func (a AddressClassB) Description() uint8 {
	return uint8((a >> 48) & 0xFF)
}

func (a AddressClassB) Incarnation() uint8 {
	return uint8((a >> 56) & 0x3F)
}

func (a AddressClassB) Class() uint8 {
	return uint8((a >> 62) & 0x03)
}
