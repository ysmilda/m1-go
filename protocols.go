package m1

// Protocols contains the supported protocols for communicating with the target.
var Protocols = struct {
	TCP protocol
	UDP protocol
}{
	TCP: protocol{value: "tcp"}, // TCP is untested
	UDP: protocol{value: "udp"},
}

type protocol struct {
	value string
}
