package m1

import (
	"crypto/md5" //nolint:gosec
	"encoding/binary"
	"fmt"
	"net"

	"github.com/ysmilda/m1-go/pkg/rpc"
)

const (
	_RES_Procedure_ListModuleInfo  = 106
	_RES_Procedure_GetModuleNumber = 112
	_RES_Procedure_SystemInfo      = 282
	_RES_Procedure_Login           = 284
	_RES_Procedure_Logout          = 286
	_RES_Procedure_XLogin          = 290
	_RES_Procedure_XLogout         = 292
	_RES_Procedure_Login2          = 304
	_RES_Procedure_Open            = 306
	_RES_Procedure_Close           = 308
	_RES_Procedure_Renew           = 310
	_RES_Procedure_ExtPing         = 320
	_RES_Procedure_FlashLed        = 324
)

// ResModule is a wrapper around the RES module of the M1 controller.
// It provides functions to interact with the RES module.
type ResModule struct {
	*Module
}

func newResModule(client *client) (*ResModule, error) {
	r, err := newModule(client, "RES", ModuleInfo{
		ModuleNumber: (0x20000000 | 0x00001000),
		UDPPort:      3000,
		TCPPort:      3500,
	}, Version{})
	if err != nil {
		return nil, fmt.Errorf("failed to create res module: %w", err)
	}

	return &ResModule{r}, nil
}

// FlashLed flashes the led of the target for approximately 5 seconds.
// The function returns immediately after sending the request.
func (r *ResModule) FlashLed() error {
	resp, err := r.ExtPing(net.IPv4(255, 255, 255, 255))
	if err != nil {
		return err
	}

	err = rpc.CallWithoutRead(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _RES_Procedure_FlashLed,
			Auth:      r.client.auth,
		},
		rpc.NewString(resp.SerialNumber, _MIO_ProductNumberLength),
		rpc.NewSpare(16),
	)
	return err
}

// ExtPing sends an extended ping to the target.
func (r *ResModule) ExtPing(ipMask net.IP) (*ExtPing, error) {
	mask := binary.LittleEndian.Uint32(ipMask.To4())
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _RES_Procedure_ExtPing,
			Auth:      r.client.auth,
		},
		mask, uint32(_RES_CompareEqual), uint32(_RES_ReplyWithIP), uint32(_RES_ReplyNormal), rpc.NewSpare(12),
	)
	if err != nil {
		return nil, err
	}

	reply := &ExtPing{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to ext ping: %w", err)
	}

	return reply, nil
}

// GetModuleNumber returns the module number of the target.
// The ModuleNumber contains the module number of the target application and it's ports.
// These are necessary for making RPC calls against that module.
func (r *ResModule) GetModuleNumber(module string) (*ModuleInfo, error) {
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _RES_Procedure_GetModuleNumber,
			Auth:      r.client.auth,
		},
		rpc.NewString(module, _ModuleNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &ModuleInfo{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to get module number: %w", err)
	}

	return reply, nil
}

// GetSystemInfo returns the system information of the target.
func (r *ResModule) GetSystemInfo() (*SystemInfo, error) {
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_SystemInfo,
		},
		uint32(0), uint32(0), uint32(0), rpc.NewString("M1Com", _ModuleNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &SystemInfo{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to get system info: %w", err)
	}

	return reply, nil
}

// Login logs in the user to the target.
//
// Deprecated: The loginChecker parameter indicates how the password is communicated. The value of this parameter
// can be found in the GetSystemInfo response. The toolName parameter is the name of the tool that is logging in.
// This may be left empty.
func (r *ResModule) Login(user, password, toolName string, loginChecker bool) (*Login, error) {
	data := []any{}
	data = append(data,
		uint32(0), uint32(0), uint32(0), rpc.NewString(toolName, _ModuleNameLength), rpc.NewString(user, _UserNameLength),
	)

	if loginChecker {
		data = append(data, rpc.NewString(password, _PasswordLength))
	} else {
		hash := md5.Sum([]byte(password)) //nolint:gosec
		data = append(data, rpc.NewString(string(hash[:]), _PasswordLength))
	}

	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_Login,
			Auth:      r.client.auth,
		},
		data...,
	)
	if err != nil {
		return nil, err
	}

	reply := &Login{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	r.client.setAuth(reply.Auth, reply.AuthLen)

	return reply, nil
}

// Login2 logs in the user to the target.
//
// The userParameter parameter is a user-defined parameter that is passed to the target. The loginChecker parameter
// indicates how the password is communicated. The value of this parameter can be found in the GetSystemInfo response.
// The toolName parameter is the name of the tool that is logging in. This may be left empty.
func (r *ResModule) Login2(
	user, password, toolName string, loginChecker bool, userParameter uint32,
) (*Login2, error) {
	data := []any{}
	data = append(data,
		userParameter, uint32(0), uint32(0),
		rpc.NewString(toolName, _ModuleNameLength), rpc.NewString(user, _UserNameLength2),
	)

	if loginChecker {
		data = append(data, rpc.NewString(password, _PasswordLength2))
	} else {
		hash := md5.Sum([]byte(password)) //nolint:gosec
		data = append(data, rpc.NewString(string(hash[:]), _PasswordLength2))
	}

	data = append(data, uint32(0), uint32(0), uint32(0), uint32(0), uint32(0))

	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_Login2,
			Auth:      r.client.auth,
		},
		data...,
	)
	if err != nil {
		return nil, err
	}

	reply := &Login2{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	r.client.setAuth(reply.Auth, reply.AuthLen)

	return reply, nil
}

// Logout logs out the user from the target.
func (r *ResModule) Logout(userParam uint32) error {
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_Logout,
			Auth:      r.client.auth,
		},
		userParam,
		r.client.auth,
	)
	if err != nil {
		return err
	}

	returnCode, _ := buf.LittleEndian.ReadUint32()
	if err := parseReturnCode(returnCode); err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	return nil
}

// ExtLogin logs in the user to the target via an installed module. This is used when a custom login application is
// installed on the target.
//
// The userParameter parameter is a user-defined parameter that is passed to the target. The module parameter is the
// name of the module that is used for the login. The user and password parameters are the credentials of the user.
func (r *ResModule) ExtLogin(user, password, module string, userParam uint32) (*ExtLogin, error) {
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_XLogin,
			Auth:      r.client.auth,
		},
		userParam, rpc.NewString(user, _UserNameLength),
		rpc.NewString(password, _PasswordLength), rpc.NewString(module, _ModuleNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &ExtLogin{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to ext login: %w", err)
	}

	return reply, nil
}

// ExtLogout logs out the user from the target via an installed module. This is used when a custom login application is
// installed on the target.
//
// The userParameter parameter is a user-defined parameter that is passed to the target. The module parameter is the
// name of the module that is used for the login. The user and password parameters are the credentials of the user.
// The userData parameter is a 128-byte array that is passed to the target.
func (r *ResModule) ExtLogout(user, password, module string, userParam uint32, userData []byte) error {
	if len(userData) != 128 {
		return fmt.Errorf("user data must be 128 bytes long")
	}

	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_XLogout,
			Auth:      r.client.auth,
		},
		userParam, rpc.NewString(user, _UserNameLength),
		rpc.NewString(password, _PasswordLength), rpc.NewString(module, _ModuleNameLength),
		r.client.auth, userData,
	)
	if err != nil {
		return err
	}

	reply := &ExtLogin{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return fmt.Errorf("failed to ext logout: %w", err)
	}

	return nil
}

// Open opens a connection to the RES module on the target.
//
// The function returns an OpenResponse that contains the response from the target. This response contains the
// authentication data that is used for further communication with the target. It also contains the lifetime of the
// session and the timeout of the session. If the session is not renewed within the timeout, the session is closed.
func (r *ResModule) Open() (*Open, error) {
	res, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_Open,
		},
		uint32(0), uint32(0), uint32(0x7FFFFFFF), rpc.NewSpare(32*4),
	)
	if err != nil {
		return nil, err
	}

	reply := &Open{}
	returnCode := reply.parse(res)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}

	return reply, nil
}

// Close closes the connection to the RES module on the target.
func (r *ResModule) Close() error {
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_Close,
			Auth:      r.client.auth,
		},
	)
	if err != nil {
		return err
	}

	returnCode, _ := buf.LittleEndian.ReadUint32()
	if err := parseReturnCode(returnCode); err != nil {
		return fmt.Errorf("failed to close: %w", err)
	}

	return nil
}

// Renew renews the connection to the RES module on the target.
//
// The function returns a RenewResponse that contains the response from the target. This response contains the
// authentication data that is used for further communication with the target. It also contains the state of the
// application and the system.
func (r *ResModule) Renew() (*Renew, error) {
	buf, err := rpc.Call(
		r.client.getConnection(r.info),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   _RPC_VersionRES,
			Procedure: _RES_Procedure_Renew,
			Auth:      r.client.auth,
		},
		// This indicates we want to renew the connection.
		// The original has this configurable, but why would you not want to renew?
		uint32(1),
		rpc.NewSpare(19+128),
	)
	if err != nil {
		return nil, err
	}

	reply := &Renew{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to renew: %w", err)
	}

	return reply, nil
}

// ListModules lists all modules installed on the target.
func (r *ResModule) ListModules() ([]string, error) {
	moduleCount := uint32(0)
	modulesPerCall := uint32(30)
	result := []string{}

	// The target has a limit of 30 modules per call, so we need to call the procedure multiple times.
	// An arbitray limit of 5 calls is set here to prevent infinite loops. It seems unlikely that a
	// target would have more than 150 modules.
	for range 5 {
		start := moduleCount
		end := start + modulesPerCall - 1

		buf, err := rpc.Call(
			r.client.getConnection(r.info),
			rpc.Header{
				Module:    r.info.ModuleNumber,
				Version:   _RPC_VersionDefault,
				Procedure: _RES_Procedure_ListModuleInfo,
				Auth:      r.client.auth,
			},
			start, end,
		)
		if err != nil {
			return nil, err
		}

		returnCode, _ := buf.LittleEndian.ReadUint32()
		if err := parseReturnCode(returnCode); err != nil {
			return nil, fmt.Errorf("failed to list modules: %w", err)
		}

		modulesInResult, _ := buf.LittleEndian.ReadUint32()
		for range modulesInResult {
			buf.Skip(12) // Skip type name
			module, _ := buf.ReadString(12)
			buf.Skip(40) // Skip the rest of the data

			result = append(result, module)
		}

		moduleCount += modulesInResult
		if modulesInResult != modulesPerCall {
			break
		}
	}

	return result, nil
}
