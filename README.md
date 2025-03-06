# M1-go

This packages provides an interface to interact with the Bachmann M1 controllers (PLC's). These controllers have a UDP interface which can be used to retrieve various information from the controller and to send commands to the controller. The goal of this package is to provide a complete interface to interact with the controller in native Go.

The implementation is based on the source code that is available with the Bachmann M1 SDK (Solution Center). The source code is written in C++ and is not directly usable in Go. The source code is used as a reference to implement the same functionality in Go.

The package is still under development and not all functionality is implemented yet. The package is tested with a Bachmann M1 controller running MSys version `4.50.99 Release`. An effort will be made to keep the package compatible with newer versions of the MSys software.

> [!NOTE]
> The package is currently versioned as pre release (e.g v0.x) due to the fact that I'm not yet fully happy with the naming of the functions. These are subject to change. However the implementation itself is production ready.

## Installation

The package can be installed using the `go get` command:

```bash
go get github.com/ysmilda/m1-go
```

## Usage

This package provides an interface to interact with the M1 controller. This interface is available through the `Target` struct. This contains communication instances of the supported modules as well as some helper functions for easy use.


The following example shows how to connect to the controller and retrieve the variables for a module:

```go
t, err := m1.NewTarget(net.IP{192, 168, 1, 1}, 1*time.Second)
if err != nil {
	// ...
}
defer t.Close()

err = t.Login("user", "password")
if err != nil {
	// ...
}

variables, err := t.ListVariables("Module")
if err != nil {
	// ...
}
```

The various communication instances need to be targeted at a specific module. To do this a `ModuleNumber` needs to be retrieved. 

```go
reply, err := t.Res.GetModuleNumber(res.ModuleNumberCall{
	Name: "Module",
})
if err != nil {
	// ...
}

serverInfo, err := t.SVI.GetServerInfo(reply.ModuleNumber, svi.GetServerInfoCall{})
if err != nil {
	// ...
}
```

If your application implements custom RPC procedures these can be accesed by using the `m1.Call` and `m1.PaginatedCall` functions. Take a look at the implementation of the various models for data layout and tagging.

The full documentation can be found at [pkg.go.dev](https://pkg.go.dev/github.com/ysmilda/m1-go).

  



