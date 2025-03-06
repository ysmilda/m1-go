package m1

import (
	"crypto/md5"
	"fmt"
	"net"
	"time"

	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/modules/mod"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
	"github.com/ysmilda/m1-go/modules/smi"
	"github.com/ysmilda/m1-go/modules/svi"
	"github.com/ysmilda/m1-go/modules/sysinfo"
)

const (
	toolName = "m1-go"
)

// Target represents a target device.
// It's methods are convenience methods around the various modules on the target.
// For more advanced usage, the modules can be used directly via the struct members.
// Make sure to follow the documentation of the modules when using them directly.
// The target should be closed after usage.
type Target struct {
	// Res - The resource handler manages software and hardware resources on the controller
	Res     *res.Procedures
	Mod     *mod.Procedures
	SVI     *svi.Procedures
	SMI     *smi.Procedures
	SysInfo *sysinfo.Procedures

	client *m1client.Client

	sessionTimeout  int32
	sessionLifetime int32
	loginChecker    bool
	loginRequired   bool

	msysVersion msys.Version
}

// NewTarget creates a new target for the given IP address and timeout.
func NewTarget(ip net.IP, timeout time.Duration) (*Target, error) {
	client := m1client.NewClient(ip, timeout)

	t := &Target{
		client: client,

		Res:     res.NewProcedures(client),
		Mod:     mod.NewProcedures(client),
		SVI:     svi.NewProcedures(client),
		SMI:     smi.NewProcedures(client),
		SysInfo: sysinfo.NewProcedures(client),
	}

	// Setup the connection to the target.
	err := t.connect()
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Close closes the target connection.
// After closing the target, the target should not be used anymore.
func (t *Target) Close() error {
	_, err := t.Res.CloseConnection(res.CloseCall{})
	if err != nil {
		return err
	}

	return t.client.Close()
}

// GetClient returns the client used to connect to the target.
func (t *Target) GetClient() *m1client.Client {
	return t.client
}

// Login logs in to the target with the given user and password.
//
// This is a helper function that calls the RES.Login or RES.Login2 procedure depending on the target version.
// When using a custom module for authentication use the RES.ExtLogin and RES.ExtLogout functions.
//
// If the login is not required, this function does nothing.
func (t *Target) Login(user, password string) error {
	if !t.loginRequired {
		return nil
	}

	// If the version is newer than 3.70.8-release we need to use the login2 procedure.
	if t.msysVersion.Compare(msys.Version{Major: 3, Minor: 70, Patch: 8, ReleaseType: msys.Release}) >= 0 {
		_, err := t.login2(user, password, t.loginChecker)
		return err
	}

	_, err := t.login(user, password, t.loginChecker)
	return err
}

func (t *Target) login(user, password string, loginChecker bool) (*res.Login, error) {
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

	reply, err := t.Res.Login(c)
	if err != nil {
		return nil, err
	}

	t.client.SetAuth(reply.Auth, reply.AuthLen)

	return &reply.Login, nil
}

func (t *Target) login2(
	user, password string, loginChecker bool,
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

	reply, err := t.Res.Login2(c)
	if err != nil {
		return nil, err
	}

	t.client.SetAuth(reply.Auth, reply.AuthLen)

	return &reply.Login2, nil
}

// Logout logs out from the target.
func (t *Target) Logout() error {
	if !t.loginRequired {
		return nil
	}

	_, err := t.Res.Logout(res.LogoutCall{})
	return err
}

// ListVariables returns a list of all variables of the module.
// The returned variables are not initialized. To initialize them, use the VHD module on the target.
func (t *Target) ListVariables(module string) ([]Variable, error) {
	number, err := t.Res.GetModuleNumber(res.ModuleNumberCall{
		Name: module,
	})
	if err != nil {
		return nil, err
	}

	version425 := msys.Version{Major: 4, Minor: 25, Patch: 0, ReleaseType: msys.Release}
	if t.msysVersion.Compare(version425) >= 0 {
		return t.listVariables2(number.ModuleNumber, module)
	} else {
		return t.listVariables(number.ModuleNumber)
	}
}

// listVariables2 returns a list of all variables of the module.
// This is the preferred implementation for newer versions of the M1 controller.
// It supports a maximum of 255 characters for the variable name.
func (t *Target) listVariables2(module res.ModuleNumber, name string) ([]Variable, error) {
	reply, err := t.SVI.ListExtendedProcessValueInfo(module, &svi.ListExtendedProcessValueInfoCall{
		GetSubprocessValues: true,
		PathLength:          1,
		Path:                "", // Start from the root.
	}, 1000)
	if err != nil {
		return nil, err
	}

	path := name
	result := []Variable{}
	for _, value := range reply {
		if value.Flag == svi.FlagTypeDirectory {
			path = fmt.Sprintf("%s/%s", name, value.Name)
			continue
		}

		result = append(result, Variable{
			Name: fmt.Sprintf("%s/%s", path, value.Name),
			Variable: svi.Variable{
				Format: value.Format,
				Length: value.Length,
			},
		})
	}

	return result, nil
}

// listVariables returns a list of all variables of the module.
// This is a fallback implementation for older versions of the M1 controller.
// It supports a maximum of 64 characters for the variable name.
func (t *Target) listVariables(module res.ModuleNumber) ([]Variable, error) {
	reply, err := t.SVI.ListProcessValueInfo(module, &svi.ListProcessValueInfoCall{}, 29)
	if err != nil {
		return nil, err
	}

	result := []Variable{}
	for _, value := range reply {
		result = append(result, Variable{
			Name: "RES/" + value.Name,
			Variable: svi.Variable{
				Format: value.Format,
				Length: value.Length,
			},
		})
	}

	return result, nil
}

// connect connects to the target and initializes the target.
func (t *Target) connect() error {
	info, err := t.Res.GetSystemInfo(res.SystemInfoCall{
		Toolname: toolName,
	})
	if err != nil {
		return err
	}

	t.loginChecker = info.LoginChecker
	t.loginRequired = info.LoginRequired
	t.msysVersion = info.MSysVersion

	// If the version is newer than 3.95.0-release we need to open the RES module.
	if info.MSysVersion.Compare(msys.Version{Major: 3, Minor: 95, Patch: 0, ReleaseType: msys.Release}) >= 0 {
		reply, err := t.Res.OpenConnection(res.OpenCall{
			RequestedSMISize: 0x7FFFFFFF,
		})
		if err != nil {
			return fmt.Errorf("failed to open: %w", err)
		}

		t.sessionTimeout = reply.SessionTimeout
		t.sessionLifetime = reply.SessionLifetime

		t.client.SetMaximalCallLength(reply.SMIMessageSize)
		t.client.SetAuth(reply.Auth, reply.AuthLen)
	}

	return nil
}
