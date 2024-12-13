package vhd

import (
	"net"
	"time"
)

type SessionInfo struct {
	Name   string `unpack:"length=44"`
	UserID uint32
	Start  time.Time `unpack:"length=4,unit=seconds"`
}

type AddressInfo struct {
	Address uint64
	Format  uint16
}

type CallBackInfo struct {
	IPAddress       net.IP `unpack:"length=4"`
	ProgramNumber   uint32
	ProtocolVersion uint32
	Protocol        ProtocolType
	Procedure       uint32
	UserParameter   uint32
	CallBackID      uint32
	Spare           []uint32 `unpack:"length=4"`
}
