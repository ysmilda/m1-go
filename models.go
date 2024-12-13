package m1

import (
	"net"

	"github.com/ysmilda/m1-go/modules/res"
)

type SVIServerInfo struct {
	Version           uint32 // Version of supported SVI
	AddressType       uint32 // The type of address class used by the server
	NumberOfVariables uint32 // The number of exported SVI variables
}

type TargetInfo struct {
	res.ExtPing
	IPAddress net.IP
}
