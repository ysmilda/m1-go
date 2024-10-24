package main

import (
	"net"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ysmilda/m1-go"
)

func main() {
	t, err := m1.NewTarget(net.IPv4(192, 168, 180, 91), "udp", 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer t.Close()

	err = t.Login("smst", "sm20st07", "tool")
	if err != nil {
		panic(err)
	}

	err = t.InitializeVHD()
	if err != nil {
		panic(err)
	}

	variable := m1.NewSviVariable("RES", "Time_s")
	count, err := t.VHD.InitializeVariables([]*m1.SviVariable{variable})
	if err != nil {
		panic(err)
	}
	if count != 1 {
		panic("expected 1 variable to be initialized")
	}

	count, err = t.VHD.ReadVariables([]*m1.SviVariable{variable})
	if err != nil {
		panic(err)
	}
	if count != 1 {
		panic("expected 1 variable to be read")
	}

	spew.Dump(variable)
}
