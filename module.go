package m1

import (
	"fmt"

	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
	"github.com/ysmilda/m1-go/modules/svi"
)

// Module wraps a generic module of the M1 controller.
type Module struct {
	client *m1client.Client
	info   res.ModuleNumber
	name   string

	msysVersion msys.Version
}

// newModule creates a new module with the given module number.
func newModule(client *m1client.Client, name string, info res.ModuleNumber, msysVersion msys.Version) *Module {
	m := &Module{
		client:      client,
		name:        name,
		info:        info,
		msysVersion: msysVersion,
	}

	return m
}

// GetSVIServerInfo returns the server information of the SVI server.
func (m *Module) GetSVIServerInfo() (*svi.ServerInfo, error) {
	reply, err := call(
		m.client,
		m.info,
		svi.Procedures.GetServerInfo(svi.GetServerInfoCall{}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get SVI server info: %w", err)
	}

	return &reply.ServerInfo, nil
}

// GetVariableCount returns the number of variables of the module.
func (m *Module) GetVariableCount() (uint32, error) {
	info, err := m.GetSVIServerInfo()
	if err != nil {
		return 0, fmt.Errorf("failed to get variable count: %w", err)
	}

	return info.NumberOfProcessValues, nil
}

// ListVariables returns a list of all variables of the module.
// The returned variables are not initialized. To initialize them, use the VHD module on the target.
func (m *Module) ListVariables() ([]Variable, error) {
	version425 := msys.Version{Major: 4, Minor: 25, Patch: 0, ReleaseType: msys.Release}
	if m.msysVersion.Compare(version425) >= 0 {
		return m.listVariables2()
	} else {
		return m.listVariables()
	}
}

// listVariables2 returns a list of all variables of the module.
// This is the preferred implementation for newer versions of the M1 controller.
// It supports a maximum of 255 characters for the variable name.
func (m *Module) listVariables2() ([]Variable, error) {
	const variablesPerCall = uint32(1000)
	index := uint32(0)
	path := m.name

	result := make([]Variable, 0)

	for {
		reply, err := call(
			m.client,
			m.info,
			svi.Procedures.GetExtendedProcessValueInfo(svi.GetExtendedProcessValueInfoCall{
				NumberOfProcessValues: variablesPerCall,
				ContinuationPoint:     index,
				GetSubprocessValues:   true,
				PathLength:            1,
				Path:                  "", // Start from the root.
			}),
		)
		if err != nil {
			return nil, fmt.Errorf("unable to get variables: %w", err)
		}

		for _, pv := range reply.ProcessValues {
			if pv.Flag == svi.FlagTypeDirectory {
				path = fmt.Sprintf("%s/%s", m.name, pv.Name)
				continue
			}

			result = append(result, Variable{
				Name: fmt.Sprintf("%s/%s", path, pv.Name),
				Variable: svi.Variable{
					Format: pv.Format,
					Length: pv.Length,
				},
			})
		}

		index = reply.ContinuationPoint
		if index == 0 {
			break
		}
	}

	return result, nil
}

// listVariables returns a list of all variables of the module.
// This is a fallback implementation for older versions of the M1 controller.
// It supports a maximum of 64 characters for the variable name.
func (m *Module) listVariables() ([]Variable, error) {
	const variablesPerCall = uint32(29)
	index := uint32(0)

	result := make([]Variable, 0)

	for {
		reply, err := call(
			m.client,
			m.info,
			svi.Procedures.GetProcessValueInfo(svi.GetProcessValueInfoCall{
				StartIndex: variablesPerCall,
			}),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get variables: %w", err)
		}

		for _, pv := range reply.ProcessValues {
			result = append(result, Variable{
				Name: "RES/" + pv.Name,
				Variable: svi.Variable{
					Format: pv.Format,
					Length: pv.Length,
				},
			})
		}

		if uint32(len(reply.ProcessValues)) < variablesPerCall {
			break
		}
		index += variablesPerCall
	}

	return result, nil
}
