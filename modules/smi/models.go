package smi

type (
	ComponentManagerChannelInfo struct {
		State              uint32
		Type               ChannelType
		VariableName       string `m1binary:"length:64"`
		MappedVariableName string `m1binary:"length:64"`
	}
)
