package msys

type BootInfo struct {
	BootDevice               string `m1binary:"length:20"`
	HostName                 string `m1binary:"length:20"`
	TargetName               string `m1binary:"length:20"`
	EthernetInternetAddress  string `m1binary:"length:30"`
	BackplaneInternetAddress string `m1binary:"length:30"`
	HostInternetAddress      string `m1binary:"length:30"`
	GatewayInternetAddress   string `m1binary:"length:30"`
	BootFile                 string `m1binary:"length:80"`
	StartScript              string `m1binary:"length:80"`
	User                     string `m1binary:"length:20"`
	Password                 string `m1binary:"length:20"`
	Other                    string `m1binary:"length:80"`
	ProcessorNumber          int32
	Flags                    int32
}
