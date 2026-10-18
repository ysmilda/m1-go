package smi

type (
	ComponentManagerChannelInfo struct {
		State              uint32
		Type               ChannelType
		VariableName       string `binary:"length:64"`
		MappedVariableName string `binary:"length:64"`
	}
)
