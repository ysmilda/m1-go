package m1

import (
	"fmt"
	"net"
	"time"
)

// Target represents a target device.
type Target struct {
	client *client

	RES *ResModule

	sessionTimeout  int32
	sessionLifetime int32
	loginChecker    bool
	loginRequired   bool

	msysVersion Version
}

// NewTarget creates a new target with the given IP address, protocol and timeout.
func NewTarget(ip net.IP, protocol protocol, timeout time.Duration) (*Target, error) {
	var (
		conn net.Conn
		err  error
	)

	switch protocol {
	case Protocols.UDP:
		conn, err = net.Dial("udp", fmt.Sprintf("%s:%d", ip.String(), 3000))

	case Protocols.TCP:
		conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", ip.String(), 3500))

	default:
		return nil, fmt.Errorf("unsupported protocol %s", protocol)
	}
	if err != nil {
		return nil, err
	}

	client := newClient(conn, timeout, protocol == Protocols.TCP)
	t := &Target{
		client: client,
		RES:    newResModule(client),
	}

	err = t.connect()
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

	return t.client.conn.Close()
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

// connect connects to the target and initializes the target.
func (t *Target) connect() error {
	info, err := t.RES.GetSystemInfo()
	if err != nil {
		return err
	}

	t.loginChecker = info.LoginChecker
	t.loginRequired = info.LoginRequired
	t.msysVersion = info.MSysVersion

	// If the version is newer than 3.95.0-release we need to open the RES module.
	if info.MSysVersion.Compare(Version{Major: 3, Minor: 95, Patch: 0, ReleaseType: "release"}) >= 0 {
		settings, err := t.RES.Open()
		if err != nil {
			return err
		}

		t.sessionTimeout = settings.SessionTimeout
		t.sessionLifetime = settings.SessionLifetime

		t.client.setAuth(settings.Auth, settings.AuthLen)
	}

	return nil
}
