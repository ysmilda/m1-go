# M1-go

This packages provides an interface to interact with the Bachmann M1 controllers (PLC's). These controllers have a UDP/TCP interface which can be used to retrieve various information from the controller and to send commands to the controller. The goal of this package is to provide a complete interface to interact with the controller in native Go.

The implementation is based on the source code that is available with the Bachmann M1 SDK (Solution Center). The source code is written in C++ and is not directly usable in Go. The source code is used as a reference to implement the same functionality in Go.

The package is still under development and not all functionality is implemented yet. The package is tested with a Bachmann M1 controller running MSys version `4.50.99 Release`. An effort will be made to keep the package compatible with newer versions of the MSys software.

## Installation

The package can be installed using the `go get` command:

```bash
go get github.com/ysmilda/m1-go
```

## Usage

The package provides an interface to interact with the M1 controller. The interface is based on the `Target` struct. The `Target` struct provides methods to connect to the controller, retrieve information from the controller and to send commands to the controller. It is split up in to the various modules that exist on the controller. These can be accessed as members of the `Target` struct.

The following example shows how to connect to the controller and retrieve the controller information:

```go
    t, err := m1.NewTarget(net.IPv4(192, 168, 180, 91), m1.Protocols.UDP, 5*time.Second)
	if err != nil {
		...
	}

	info, err := t.RES.GetSystemInfo()
    if err != nil {
        ...
    }

    err = t.Login("user", "password", "m1-go")
	if err != nil {
        ...
	}

	modules, err := t.RES.ListModules()
	if err != nil {
        ...
	}
```

The full documentation can be found at [pkg.go.dev](https://pkg.go.dev/github.com/ysmilda/m1-go).

## Supported modules

The module support is still under development. The following modules are currently supported:

- [x] RES Module 
- [ ] VHD Module 
- [ ] RFS Module

These modules are most commonly used to interact with the controller. Other modules might be added at a later stage.
  



