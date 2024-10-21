package m1

import (
	"crypto/md5" //nolint:gosec
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"time"

	"github.com/ysmilda/m1-go/pkg/rpc"
)

const (
	_MODULE_RES = (0x20000000 | 0x00001000)

	_VERSION_DEFAULT = 2
	_VERSION_RES     = 3

	_PROCEDURE_RES_LIST_MODULE_INFO  = 106
	_PROCEDURE_RES_GET_MODULE_NUMBER = 112
	_PROCEDURE_RES_SYSTEM_INFO       = 282
	_PROCEDURE_RES_LOGIN             = 284
	_PROCEDURE_RES_LOGOUT            = 286
	_PROCEDURE_RES_XLOGIN            = 290
	_PROCEDURE_RES_XLOGOUT           = 292
	_PROCEDURE_RES_LOGIN2            = 304
	_PROCEDURE_RES_OPEN              = 306
	_PROCEDURE_RES_CLOSE             = 308
	_PROCEDURE_RES_RENEW             = 310
	_PROCEDURE_RES_EXT_PING          = 320
	_PROCEDURE_RES_FLASH_LED         = 324
)

// ResModule wraps the RES module of the M1 controller.
// It should not be created directly, but by using the RES field of the Target struct.
type ResModule struct {
	client *client
}

func newResModule(client *client) *ResModule {
	return &ResModule{client: client}
}

// FlashLed flashes the led of the target for approximately 5 seconds.
// The function returns immediately after sending the request.
func (r *ResModule) FlashLed() error {
	resp, err := r.ExtPing(net.IPv4(255, 255, 255, 255))
	if err != nil {
		return err
	}

	conn := r.client.getConnection()
	err = conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}

	err = rpc.CallWithoutRead(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_DEFAULT,
			Procedure: _PROCEDURE_RES_FLASH_LED,
			Auth:      r.client.auth,
		},
		rpc.NewString(resp.SerialNumber, _MIO_ProductNumberLength),
		rpc.NewSpare(16),
	)
	return err
}

// ExtPing sends an extended ping to the target.
func (r *ResModule) ExtPing(ipMask net.IP) (*ExtPingResponse, error) {
	mask := binary.LittleEndian.Uint32(ipMask.To4())
	buf, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_DEFAULT,
			Procedure: _PROCEDURE_RES_EXT_PING,
			Auth:      r.client.auth,
		},
		mask, uint32(_RES_CompareEqual), uint32(_RES_ReplyWithIP), uint32(_RES_ReplyNormal), rpc.NewSpare(12),
	)
	if err != nil {
		return nil, err
	}

	reply := &ExtPingResponse{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to ext ping: %w", err)
	}

	return reply, nil
}

// GetModuleNumber returns the module number of the target.
// The ModuleNumber contains the module number of the target application and it's ports.
// These are necessary for making RPC calls against that module.
func (r *ResModule) GetModuleNumber(module string) (*ModuleNumberResponse, error) {
	buf, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_DEFAULT,
			Procedure: _PROCEDURE_RES_GET_MODULE_NUMBER,
			Auth:      r.client.auth,
		},
		rpc.NewString(module, _ModuleNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &ModuleNumberResponse{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to get module number: %w", err)
	}

	return reply, nil
}

// GetSystemInfo returns the system information of the target.
func (r *ResModule) GetSystemInfo() (*SystemInfoResponse, error) {
	buf, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_SYSTEM_INFO,
		},
		uint32(0), uint32(0), uint32(0), rpc.NewString("M1Com", _ModuleNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &SystemInfoResponse{}
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
func (r *ResModule) Login(user, password, toolName string, loginChecker bool) (*LoginResponse, error) {
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
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_LOGIN,
			Auth:      r.client.auth,
		},
		data...,
	)
	if err != nil {
		return nil, err
	}

	reply := &LoginResponse{}
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
) (*Login2Response, error) {
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
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_LOGIN2,
			Auth:      r.client.auth,
		},
		data...,
	)
	if err != nil {
		return nil, err
	}

	reply := &Login2Response{}
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
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_LOGOUT,
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
func (r *ResModule) ExtLogin(user, password, module string, userParam uint32) (*ExtLoginResponse, error) {
	buf, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_XLOGIN,
			Auth:      r.client.auth,
		},
		userParam, rpc.NewString(user, _UserNameLength),
		rpc.NewString(password, _PasswordLength), rpc.NewString(module, _ModuleNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &ExtLoginResponse{}
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
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_XLOGOUT,
			Auth:      r.client.auth,
		},
		userParam, rpc.NewString(user, _UserNameLength),
		rpc.NewString(password, _PasswordLength), rpc.NewString(module, _ModuleNameLength),
		r.client.auth, userData,
	)
	if err != nil {
		return err
	}

	reply := &ExtLoginResponse{}
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
func (r *ResModule) Open() (*OpenResponse, error) {
	res, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_OPEN,
		},
		uint32(0), uint32(0), uint32(math.MaxUint32), rpc.NewSpare(32*4),
	)
	if err != nil {
		return nil, err
	}

	reply := &OpenResponse{}
	returnCode := reply.parse(res)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}

	return reply, nil
}

// Close closes the connection to the RES module on the target.
func (r *ResModule) Close() error {
	buf, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_CLOSE,
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
func (r *ResModule) Renew() (*RenewResponse, error) {
	buf, err := rpc.Call(
		r.client.getConnectionWithTimeout(),
		rpc.Header{
			Module:    _MODULE_RES,
			Version:   _VERSION_RES,
			Procedure: _PROCEDURE_RES_RENEW,
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

	reply := &RenewResponse{}
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
	for i := 0; i < 5; i++ {
		start := moduleCount
		end := start + modulesPerCall - 1

		buf, err := rpc.Call(
			r.client.getConnectionWithTimeout(),
			rpc.Header{
				Module:    _MODULE_RES,
				Version:   _VERSION_DEFAULT,
				Procedure: _PROCEDURE_RES_LIST_MODULE_INFO,
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
