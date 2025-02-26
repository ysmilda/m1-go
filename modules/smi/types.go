package smi

type ChannelType uint32

const (
	ChannelTypeOut ChannelType = 0
	ChannelTypeIn  ChannelType = 1
)

type LicenseEvent uint32

const (
	LicenseEventPlug LicenseEvent = 0
	LicenseEventPull LicenseEvent = 1
)
