package m1

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/modules/msys"
)

var defaultModules = map[string]struct{}{
	"RES": {},
	"VHD": {},
}

// Target represents a target device.
// It's methods are convenience methods around the various modules on the target.
// For more advanced usage, the modules can be used directly via the struct members.
// Make sure to follow the documentation of the modules when using them directly.
// The target should be closed after usage.
type Target struct {
	// RES - The resource handler manages software and hardware resources on the controller
	RES *ResModule

	Modules map[string]*Module

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

	res := newResModule(client)

	t := &Target{
		client:  client,
		RES:     res,
		Modules: make(map[string]*Module),
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
	err := t.RES.Close()
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
// The toolName is used to identify the tool in the target, it may be left empty.
func (t *Target) Login(user, password, toolName string) error {
	if !t.loginRequired {
		return nil
	}

	// If the version is newer than 3.70.8-release we need to use the login2 procedure.
	if t.msysVersion.Compare(msys.Version{Major: 3, Minor: 70, Patch: 8, ReleaseType: msys.Release}) >= 0 {
		_, err := t.RES.Login2(user, password, t.loginChecker, 0)
		return err
	}

	_, err := t.RES.Login(user, password, t.loginChecker)
	return err
}

// Logout logs out from the target.
func (t *Target) Logout() error {
	if !t.loginRequired {
		return nil
	}

	return t.RES.Logout()
}

// ConnectModule connects to a custom module on the target.
// It returns an error if the module is a default module or if the module can't be connected to.
// The module can be accessed via the Modules map on the target.
func (t *Target) ConnectModule(moduleName string) error {
	if _, ok := defaultModules[moduleName]; ok {
		return errors.New("use the target struct members for access to the default modules")
	}

	if _, ok := t.Modules[moduleName]; ok {
		return nil
	}

	info, err := t.RES.GetModuleNumber(moduleName)
	if err != nil {
		return fmt.Errorf("failed to get module number for %s: %w", moduleName, err)
	}

	t.Modules[moduleName] = newModule(t.client, moduleName, *info, t.msysVersion)
	return nil
}

// connect connects to the target and initializes the target.
func (t *Target) connect() error {
	info, err := t.RES.GetSystemInfo()
	if err != nil {
		return err
	}

	t.loginChecker = info.LoginChecker
	t.loginRequired = info.LoginRequired
	t.msysVersion = info.MSysVersion

	// At the moment we opened the connection to the RES module we didn't know the version yet.
	t.RES.msysVersion = t.msysVersion

	// If the version is newer than 3.95.0-release we need to open the RES module.
	if info.MSysVersion.Compare(msys.Version{Major: 3, Minor: 95, Patch: 0, ReleaseType: msys.Release}) >= 0 {
		settings, err := t.RES.Open()
		if err != nil {
			return err
		}

		t.sessionTimeout = settings.SessionTimeout
		t.sessionLifetime = settings.SessionLifetime

		t.client.SetMaximalCallLength(settings.SMIMessageSize)
		t.client.SetAuth(settings.Auth, settings.AuthLen)
	}

	return nil
}
