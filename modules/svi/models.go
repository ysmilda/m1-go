package svi

type ProcessValueInfo struct {
	Name   string `binary:"length:64"`
	Format uint16
	Length uint16
}

type ExtendedProcessValueInfo struct {
	Flag   FlagType `binary:"skip:2"`
	Format uint16
	Length uint16 `binary:"skip:4"`
	Name   string `binary:"zeroTerminated"`
}

type ServerInfo struct {
	Version               uint32
	AddressType           AddressType
	NumberOfProcessValues uint32 // Number of exported Process Value variables
}

type VariableInfo struct {
	StartIndexLow              uint32 // Start index of low matrix elements (low or 2nd index)
	StartIndexHigh             uint32 // Start index of high matrix elements (high or 1st index)
	AmountOfLowMatrixElements  uint32
	AmountOfHighMatrixElements uint32
	Reserved                   []uint32 `binary:"length:4"`
}
