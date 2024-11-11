package m1

import (
	"errors"
	"fmt"
	"net"
	"time"
)

var defaultModules = map[string]struct{}{
	"RES":  {},
	"VHD":  {},
	"INFO": {},
}

// Target represents a target device.
// It's methods are convenience methods around the various modules on the target.
// For more advanced usage, the modules can be used directly via the struct members.
// Make sure to follow the documentation of the modules when using them directly.
// The target should be closed after usage.
type Target struct {
	// RES - This module offers access to basic information about the target as well as the ability to log in and out.
	RES *ResModule

	// VHD - This module offers access to t.b.d.
	//
	// To use this module one must be logged in and have initialized the VHD module by calling target.InitializeVHD().
	VHD *VhdModule

	INFO *InfoModule

	Modules map[string]*Module

	client *client

	sessionTimeout  int32
	sessionLifetime int32
	loginChecker    bool
	loginRequired   bool

	msysVersion Version
}

// NewTarget creates a new target with the given IP address, protocol and timeout.
// Protocol must be either "tcp" or "udp". "tcp" is currently not fully functional.
func NewTarget(ip net.IP, protocol string, timeout time.Duration) (*Target, error) {
	client := newClient(ip, timeout, protocol)

	res, err := newResModule(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create RES module: %w", err)
	}

	t := &Target{
		client:  client,
		RES:     res,
		Modules: make(map[string]*Module),
	}

	// Setup the connection to the target.
	err = t.connect()
	if err != nil {
		return nil, err
	}

	err = t.initializeINFO()
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

	if t.VHD != nil {
		err = t.VHD.Close()
		if err != nil {
			return err
		}
	}

	return t.client.close()
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
	if t.msysVersion.Compare(Version{Major: 3, Minor: 70, Patch: 8, ReleaseType: "release"}) >= 0 {
		_, err := t.RES.Login2(user, password, toolName, t.loginChecker, 0)
		return err
	}

	_, err := t.RES.Login(user, password, toolName, t.loginChecker)
	return err
}

// Logout logs out from the target.
func (t *Target) Logout() error {
	if !t.loginRequired {
		return nil
	}

	return t.RES.Logout(0)
}

// ConnectModule connects to a custom module on the target.
// It returns an error if the module is a default module or if the module can't be connected to.
// The module can be accessed via the Modules map on the target.
func (t *Target) ConnectModule(moduleName string) error {
	if _, ok := defaultModules[moduleName]; ok {
		return errors.New("use the target struct members for the default modules")
	}

	if _, ok := t.Modules[moduleName]; ok {
		return nil
	}

	info, err := t.RES.GetModuleNumber(moduleName)
	if err != nil {
		return fmt.Errorf("failed to get module number for %s: %w", moduleName, err)
	}

	m, err := newModule(t.client, moduleName, *info, t.msysVersion)
	if err != nil {
		return fmt.Errorf("failed to create module %s: %w", moduleName, err)
	}

	t.Modules[moduleName] = m
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
	if info.MSysVersion.Compare(Version{Major: 3, Minor: 95, Patch: 0, ReleaseType: "release"}) >= 0 {
		settings, err := t.RES.Open()
		if err != nil {
			return err
		}

		t.sessionTimeout = settings.SessionTimeout
		t.sessionLifetime = settings.SessionLifetime

		t.client.maxCallLength = settings.SMIMessageSize
		t.client.setAuth(settings.Auth, settings.AuthLen)
	}

	return nil
}

func (t *Target) InitializeVHD() error {
	// Create a session for the VHD module.
	info, err := t.RES.GetModuleNumber("VHD")
	if err != nil {
		return fmt.Errorf("failed to get VHD module number: %w", err)
	}

	vhd, err := newVhdModule(t.client, *info)
	if err != nil {
		return fmt.Errorf("failed to create VHD module: %w", err)
	}

	t.VHD = vhd
	return nil
}

func (t *Target) initializeINFO() error {
	infoModuleInfo, err := t.RES.GetModuleNumber("INFO")
	if err != nil {
		return fmt.Errorf("failed to get INFO module number: %w", err)
	}
	info, err := newInfoModule(t.client, *infoModuleInfo, t.msysVersion)
	if err != nil {
		return fmt.Errorf("failed to create INFO module: %w", err)
	}
	t.INFO = info
	return nil
}
