package m1

import (
	//nolint:gosec

	"crypto/md5"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/ysmilda/m1-go/internals/client"
	"github.com/ysmilda/m1-go/internals/rpc"
	"github.com/ysmilda/m1-go/internals/unpack"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

const (
	resProcedureModuleInfo      = 104
	resProcedureListModuleInfo  = 106
	resProcedureGetModuleNumber = 112
	resProcedureSystemInfo      = 282
	resProcedureLogin           = 284
	resProcedureLogout          = 286
	resProcedureLogin2          = 304
	resProcedureOpen            = 306
	resProcedureClose           = 308
	resProcedureRenew           = 310
	resProcedureExtPing         = 320
	resProcedureFlashLED        = 324

	toolName = "m1-go"
)

// ResModule is a wrapper around the RES module of the M1 controller.
// It provides functions to interact with the RES module.
type ResModule struct {
	*Module
}

func newResModule(client *client.Client) (*ResModule, error) {
	r, err := newModule(client, "RES", res.ModuleNumber{
		ModuleNumber: (0x20000000 | 0x00001000),
		UDPPort:      3000,
		TCPPort:      3500,
	}, msys.Version{})
	if err != nil {
		return nil, fmt.Errorf("failed to create res module: %w", err)
	}

	return &ResModule{r}, nil
}

func (r *ResModule) GetModuleInfo(module string) (*res.ModuleInfo, error) {
	c := &res.ModuleInfoCall{
		Name: module,
	}

	reply, err := call[res.ModInfoReply](r.client, r.info, resProcedureModuleInfo, rpc.VersionDefault, c)
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
		c := res.ModuleInfoListCall{
			First: index,
			Last:  index + amountPerCall - 1,
		}

		reply, err := call[res.ModuleInfoListReply](r.client, r.info, resProcedureListModuleInfo, rpc.VersionDefault, c)
		if err != nil {
			return nil, fmt.Errorf("failed to list modules: %w", err)
		}

		result = append(result, reply.Modules...)

		index += reply.NumberOfModules
		if reply.NumberOfModules != amountPerCall {
			break
		}
	}

	return result, nil
}

// GetModuleNumber returns the module number of the target.
// The ModuleNumber contains the module number of the target application and it's ports.
// These are necessary for making RPC calls against that module.
func (r *ResModule) GetModuleNumber(module string) (*res.ModuleNumber, error) {
	c := &res.ModuleNumberCall{
		Name: module,
	}

	reply, err := call[res.ModuleNumberReply](r.client, r.info, resProcedureGetModuleNumber, rpc.VersionDefault, c)
	if err != nil {
		return nil, fmt.Errorf("failed to get module number: %w", err)
	}
	return &reply.ModuleNumber, nil
}

// GetSystemInfo returns the system information of the target.
func (r *ResModule) GetSystemInfo() (*res.SystemInfo, error) {
	c := &res.SystemInfoCall{
		Toolname: toolName,
	}

	reply, err := call[res.SystemInfoReply](r.client, r.info, resProcedureSystemInfo, rpc.VersionRES, c)
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
		hash := md5.Sum([]byte(password)) //nolint:gosec
		c.Password = string(hash[:])
	}

	reply, err := call[res.LoginReply](r.client, r.info, resProcedureLogin, rpc.VersionRES, c)
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
		hash := md5.Sum([]byte(password)) //nolint:gosec
		c.Password = string(hash[:])
	}

	reply, err := call[res.Login2Reply](r.client, r.info, resProcedureLogin, rpc.VersionRES, c)
	if err != nil {
		return nil, err
	}

	r.client.SetAuth(reply.Auth, reply.AuthLen)

	return &reply.Login2, nil
}

// Logout logs out the user from the target.
func (r *ResModule) Logout() error {
	_, err := call[res.LogoutReply](r.client, r.info, resProcedureLogout, rpc.VersionRES, &res.LogoutCall{})
	return err
}

// Open opens a connection to the RES module on the target.
//
// The function returns an OpenResponse that contains the response from the target. This response contains the
// authentication data that is used for further communication with the target. It also contains the lifetime of the
// session and the timeout of the session. If the session is not renewed within the timeout, the session is closed.
func (r *ResModule) Open() (*res.Open, error) {
	c := &res.OpenCall{
		RequestedSMISize: 0x7FFFFFFF,
	}

	reply, err := call[res.OpenReply](r.client, r.info, resProcedureSystemInfo, rpc.VersionRES, c)
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}

	return &reply.Open, nil
}

// Close closes the connection to the RES module on the target.
func (r *ResModule) Close() error {
	_, err := call[res.CloseReply](r.client, r.info, resProcedureClose, rpc.VersionRES, res.CloseCall{})
	return err
}

// Renew renews the connection to the RES module on the target.
//
// The function returns a RenewResponse that contains the response from the target. This response contains the
// authentication data that is used for further communication with the target. It also contains the state of the
// application and the system.
func (r *ResModule) Renew() (*res.Renew, error) {
	c := &res.RenewCall{
		RequestAuthenticationRenewal: true,
	}

	reply, err := call[res.RenewReply](r.client, r.info, resProcedureRenew, rpc.VersionRES, c)
	if err != nil {
		return nil, err
	}

	return &reply.Renew, nil
}

// ExtPing sends an extended ping to the target.
func (r *ResModule) ExtPing(ipMask net.IP) (*res.ExtPing, error) {
	c := res.ExtPingCall{
		IpMask: binary.LittleEndian.Uint32(ipMask.To4()),
		Filter: res.FilterEqual,
		Reply:  res.ReplyTypeIP,
		Mode:   res.ReplyModeNormal,
	}

	reply, err := call[res.ExtPingReply](r.client, r.info, resProcedureExtPing, rpc.VersionDefault, c)
	if err != nil {
		return nil, err
	}

	return &reply.ExtPing, nil
}

// FlashLed flashes the led of the target for approximately 5 seconds.
// The function returns immediately after sending the request.
func (r *ResModule) FlashLed() error {
	resp, err := r.ExtPing(net.IPv4(255, 255, 255, 255))
	if err != nil {
		return err
	}

	call := &res.FlashLEDCall{
		SerialNumber: resp.SerialNumber,
	}

	body, err := unpack.Pack(binary.LittleEndian, call)
	if err != nil {
		return err
	}

	return rpc.CallWithoutRead(
		r.client.GetConnection(r.info.ModuleNumber),
		rpc.Header{
			Module:    r.info.ModuleNumber,
			Version:   rpc.VersionDefault,
			Procedure: resProcedureFlashLED,
			Auth:      r.client.GetAuth(),
		},
		body,
	)
}
