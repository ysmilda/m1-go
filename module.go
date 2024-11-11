package m1

import (
	"fmt"

	"github.com/ysmilda/m1-go/pkg/rpc"
)

const (
	_SVI_Procedure_GetVariableInfo = 10014
	_SVI_Procedure_GetServerInfo   = 10016

	_SVI_VariableInfoExtendedCall = 0x7575abcd

	_SVI_FlagDirectory = 0x0000
)

// Module wraps a generic module of the M1 controller.
type Module struct {
	client *client
	info   ModuleInfo
	name   string

	msysVersion Version
}

// newModule creates a new module with the given module number.
func newModule(client *client, name string, info ModuleInfo, msysVersion Version) (*Module, error) {
	m := &Module{
		client:      client,
		name:        name,
		info:        info,
		msysVersion: msysVersion,
	}

	err := client.addConnection(info)
	if err != nil {
		return nil, fmt.Errorf("failed to add connection: %w", err)
	}

	return m, nil
}

// GetSVIServerInfo returns the server information of the SVI server.
func (m *Module) GetSVIServerInfo() (*SVIServerInfo, error) {
	buf, err := rpc.Call(
		m.client.getConnection(m.info),
		rpc.Header{
			Module:    m.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _SVI_Procedure_GetServerInfo,
			Auth:      m.client.auth,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get SVI server info: %w", err)
	}

	reply := &SVIServerInfo{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to get SVI server info: %w", err)
	}

	return reply, nil
}

// GetVariableCount returns the number of variables of the module.
func (m *Module) GetVariableCount() (uint32, error) {
	info, err := m.GetSVIServerInfo()
	if err != nil {
		return 0, fmt.Errorf("failed to get variable count: %w", err)
	}

	return info.NumberOfVariables, nil
}

// ListVariables returns a list of all variables of the module.
// The returned variables are not initialized. To initialize them, use the VHD module on the target.
func (m *Module) ListVariables() ([]Variable, error) {
	version425 := Version{Major: 4, Minor: 25, Patch: 0, ReleaseType: "release"}
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
		buf, err := rpc.Call(
			m.client.getConnection(m.info),
			rpc.Header{
				Module:    m.info.ModuleNumber,
				Version:   _RPC_VersionDefault,
				Procedure: _SVI_Procedure_GetVariableInfo,
				Auth:      m.client.auth,
			},
			uint32(_SVI_VariableInfoExtendedCall), variablesPerCall, index,
			byte(1), rpc.NewSpare(11), uint32(1), rpc.NewString("", 1),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get variables: %w", err)
		}

		returnCode, _ := buf.LittleEndian.ReadUint32()
		if err := parseReturnCode(returnCode); err != nil {
			return nil, err
		}

		buf.Skip(4)                               // Number of PV (old, not used)
		index, _ = buf.LittleEndian.ReadUint32()  // Next index
		buf.Skip(3 * 4)                           // Spare
		count, _ := buf.LittleEndian.ReadUint32() // Number of returned variables

		for range count {
			buf.Align4()
			flags, _ := buf.LittleEndian.ReadUint16()
			buf.Skip(2)
			format, _ := buf.LittleEndian.ReadUint16()
			length, _ := buf.LittleEndian.ReadUint16()
			nameLength, _ := buf.LittleEndian.ReadUint32()
			name, _ := buf.ReadString(int(nameLength) + 1)

			if flags == _SVI_FlagDirectory {
				path = fmt.Sprintf("%s/%s", m.name, name)
				continue
			} else {
				name = fmt.Sprintf("%s/%s", path, name)
			}

			result = append(result, Variable{
				Name:   name,
				Format: format,
				Length: length,
			})
		}

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
		buf, err := rpc.Call(
			m.client.getConnection(m.info),
			rpc.Header{
				Module:    m.info.ModuleNumber,
				Version:   _RPC_VersionDefault,
				Procedure: _SVI_Procedure_GetVariableInfo,
				Auth:      m.client.auth,
			},
			index, variablesPerCall,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get variables: %w", err)
		}

		returnCode, _ := buf.LittleEndian.ReadUint32()
		if returnCode == (_SOURCE_SVI | _ERROR_FAILED) {
			// This is how the system announces that there are no more variables.
			break
		} else if err := parseReturnCode(returnCode); err != nil {
			return nil, err
		}

		count, _ := buf.LittleEndian.ReadUint32()
		for range count {
			name, _ := buf.ReadString(_SVI_NameLength)
			format, _ := buf.LittleEndian.ReadUint16()
			length, _ := buf.LittleEndian.ReadUint16()

			result = append(result, Variable{
				Name:   "RES/" + name,
				Format: format,
				Length: length,
			})
		}

		if count < variablesPerCall {
			break
		}
		index += variablesPerCall
	}

	return result, nil
}
