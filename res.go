package m1

import (
	"crypto/md5"
	"fmt"
	"net"

	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

const (
	toolName = "m1-go"
)

// ResModule is a wrapper around the RES module of the M1 controller.
// It provides functions to interact with the RES module.
type ResModule struct {
	*Module
}

func newResModule(client *m1client.Client) *ResModule {
	return &ResModule{newModule(client, "RES", res.Module, msys.Version{})}
}

func (r *ResModule) GetModuleInfo(module string) (*res.ModuleInfo, error) {
	reply, err := call(
		r.client,
		r.info,
		res.Procedures.GetModuleInfo(res.ModuleInfoCall{
			Name: module,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get module info: %w", err)
	}

	return &reply.ModuleInfo, nil
}

// ListModules lists all modules installed on the target.
func (r *ResModule) ListModules() ([]res.ModuleInfo, error) {
	const amountPerCall = uint32(30)
	index := uint32(0)
	result := []res.ModuleInfo{}

	for {
		reply, err := call(
			r.client,
			r.info,
			res.Procedures.ListModuleInfo(res.ListModuleInfoCall{
				First: index,
				Last:  index + amountPerCall - 1,
			}),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to list modules: %w", err)
		}

		result = append(result, reply.Modules...)

		length := uint32(len(reply.Modules))
		index += length
		if length != amountPerCall {
			break
		}
	}

	return result, nil
}

// GetModuleNumber returns the module number of the target.
// The ModuleNumber contains the module number of the target application and it's ports.
// These are necessary for making RPC calls against that module.
func (r *ResModule) GetModuleNumber(module string) (*res.ModuleNumber, error) {
	reply, err := call(
		r.client,
		r.info,
		res.Procedures.GetModuleNumber(res.ModuleNumberCall{
			Name: module,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get module number: %w", err)
	}
	return &reply.ModuleNumber, nil
}

// GetSystemInfo returns the system information of the target.
func (r *ResModule) GetSystemInfo() (*res.SystemInfo, error) {
	reply, err := call(
		r.client,
		r.info,
		res.Procedures.GetSystemInfo(res.SystemInfoCall{
			Toolname: toolName,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get system info: %w", err)
	}

	return &reply.SystemInfo, nil
}

// Login logs in the user to the target.
//
// Deprecated: The loginChecker parameter indicates how the password is communicated. The value of this parameter
// can be found in the GetSystemInfo response.
func (r *ResModule) Login(user, password string, loginChecker bool) (*res.Login, error) {
	c := res.LoginCall{
		SystemInfoCall: res.SystemInfoCall{
			Toolname: toolName,
		},
		Username: user,
		Password: password,
	}

	if !loginChecker {
		hash := md5.Sum([]byte(password))
		c.Password = string(hash[:])
	}

	reply, err := call(r.client, r.info, res.Procedures.Login(c))
	if err != nil {
		return nil, err
	}

	r.client.SetAuth(reply.Auth, reply.AuthLen)

	return &reply.Login, nil
}

// Login2 logs in the user to the target.
//
// The userParameter parameter is a user-defined parameter that is passed to the target. The loginChecker parameter
// indicates how the password is communicated. The value of this parameter can be found in the GetSystemInfo response.
// The toolName parameter is the name of the tool that is logging in. This may be left empty.
func (r *ResModule) Login2(
	user, password string, loginChecker bool, userParameter uint32,
) (*res.Login2, error) {
	c := res.Login2Call{
		SystemInfoCall: res.SystemInfoCall{
			Toolname: toolName,
		},
		Username: user,
		Password: password,
	}

	if !loginChecker {
		hash := md5.Sum([]byte(password))
		c.Password = string(hash[:])
	}

	reply, err := call(r.client, r.info, res.Procedures.Login2(c))
	if err != nil {
		return nil, err
	}

	r.client.SetAuth(reply.Auth, reply.AuthLen)

	return &reply.Login2, nil
}

// Logout logs out the user from the target.
func (r *ResModule) Logout() error {
	_, err := call(r.client, r.info, res.Procedures.Logout(res.LogoutCall{}))
	return err
}

// Open opens a connection to the RES module on the target.
//
// The function returns an OpenResponse that contains the response from the target. This response contains the
// authentication data that is used for further communication with the target. It also contains the lifetime of the
// session and the timeout of the session. If the session is not renewed within the timeout, the session is closed.
func (r *ResModule) Open() (*res.Open, error) {
	reply, err := call(
		r.client,
		r.info,
		res.Procedures.OpenConnection(res.OpenCall{
			RequestedSMISize: 0x7FFFFFFF,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}

	return &reply.Open, nil
}

// Close closes the connection to the RES module on the target.
func (r *ResModule) Close() error {
	_, err := call(r.client, r.info, res.Procedures.CloseConnection(res.CloseCall{}))
	return err
}

// Renew renews the connection to the RES module on the target.
//
// The function returns a RenewResponse that contains the response from the target. This response contains the
// authentication data that is used for further communication with the target. It also contains the state of the
// application and the system.
func (r *ResModule) Renew() (*res.Renew, error) {
	reply, err := call(
		r.client,
		r.info,
		res.Procedures.RenewConnection(res.RenewCall{
			RequestAuthenticationRenewal: true,
		}),
	)
	if err != nil {
		return nil, err
	}

	return &reply.Renew, nil
}

// ExtendedPing sends an extended ping to the target.
func (r *ResModule) ExtendedPing(ipMask net.IPMask) (*res.ExtPing, error) {
	reply, err := call(
		r.client,
		r.info,
		res.Procedures.ExtPing(res.ExtPingCall{
			IPMask: ipMask,
			Filter: res.FilterEqual,
			Reply:  res.ReplyTypeIP,
			Mode:   res.ReplyModeNormal,
		}),
	)
	if err != nil {
		return nil, err
	}

	return &reply.ExtPing, nil
}

// FlashLed flashes the led of the target for approximately 5 seconds.
// The function returns immediately after sending the request.
func (r *ResModule) FlashLed() error {
	reply, err := r.ExtendedPing(net.IPv4Mask(255, 255, 255, 255))
	if err != nil {
		return err
	}

	_, err = call(
		r.client,
		r.info,
		res.Procedures.FlashLED(res.FlashLEDCall{
			SerialNumber: reply.SerialNumber,
		}),
	)

	return err
}
