package msys

type BootInfo struct {
	BootDevice               string
	HostName                 string
	TargetName               string
	EthernetInternetAddress  string
	BackplaneInternetAddress string
	HostInternetAddress      string
	GatewayInternetAddress   string
	BootFile                 string
	StartScript              string
	User                     string
	Password                 string
	Other                    string
	ProcessorNumber          uint32
	Flags                    uint32
}
