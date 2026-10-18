package msys

type BootInfo struct {
	BootDevice               string `binary:"length:20"`
	HostName                 string `binary:"length:20"`
	TargetName               string `binary:"length:20"`
	EthernetInternetAddress  string `binary:"length:30"`
	BackplaneInternetAddress string `binary:"length:30"`
	HostInternetAddress      string `binary:"length:30"`
	GatewayInternetAddress   string `binary:"length:30"`
	BootFile                 string `binary:"length:80"`
	StartScript              string `binary:"length:80"`
	User                     string `binary:"length:20"`
	Password                 string `binary:"length:20"`
	Other                    string `binary:"length:80"`
	ProcessorNumber          int32
	Flags                    int32
}
