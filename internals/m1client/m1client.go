// Package m1 contains an implementation for communicating with the Bachmann M1 PLC platform.
package m1client

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// client is the main struct that holds the connection to the target.
// It is intended solely for internal use and should not be used directly.
// On improper usage, the client may panic.
// It is used to manage the connections to the target and to store the authentication information.
type Client struct {
	// overall configuration
	ip      net.IP
	timeout time.Duration

	// all connections for all modules. Every module has its own connection.
	connections map[uint16]*conn

	auth          []byte
	sessionID     uint32
	maxCallLength int32
}

// NewClient creates a new client with the given IP address and timeout.
func NewClient(ip net.IP, timeout time.Duration) *Client {
	return &Client{
		ip:            ip,
		timeout:       timeout,
		connections:   make(map[uint16]*conn),
		maxCallLength: 2048,
	}
}

// Close closes all connections to the target.
func (c Client) Close() error {
	for key, conn := range c.connections {
		err := conn.Close()
		if err != nil {
			return err
		}
		delete(c.connections, key)
	}
	return nil
}

// GetConnection returns the connection with the given ID.
func (c *Client) GetConnection(port uint16) *conn {
	connection, ok := c.connections[port]
	if ok {
		return connection
	}

	nc, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.ip, port))
	if err != nil {
		return nil
	}

	connection = &conn{
		conn:    nc,
		timeout: c.timeout,
	}
	c.connections[port] = connection
	return connection
}

// SetAuth sets the authentication information for the client.
func (c *Client) SetAuth(auth []byte, length uint32) {
	c.auth = auth[:length]
	if length >= 12 {
		flavor := binary.BigEndian.Uint32(auth[:4])

		if flavor == 101 || flavor == 102 {
			c.sessionID = binary.BigEndian.Uint32(auth[8:12])
		}
	} else {
		c.sessionID = 0
	}
}

// GetAuth returns the authentication information for the client.
func (c *Client) GetAuth() []byte {
	return c.auth
}

// SetMaximalCallLength sets the maximal call length for the client.
func (c *Client) SetMaximalCallLength(length int32) {
	c.maxCallLength = length
}

// GetMaximumCallLength returns the maximal call length for the client.
func (c *Client) GetMaximumCallLength() int {
	if len(c.auth) == 0 {
		return int(c.maxCallLength - 40)
	}
	return int(c.maxCallLength - 40 - int32(len(c.auth)))
}

// conn is a wrapper around a net.Conn that adds some functionality for the M1 protocol.
type conn struct {
	conn    net.Conn
	timeout time.Duration
}

// Close closes the underlying connection.
func (c conn) Close() error {
	return c.conn.Close()
}

// Read reads data from the connection and sets the read deadline if a timeout is set.
func (c conn) Read(p []byte) (n int, err error) {
	// Set the read deadline if a timeout is set.
	if c.timeout != 0 {
		_ = c.conn.SetReadDeadline(time.Now().Add(c.timeout))
	}

	return c.conn.Read(p)
}

// Write writes data to the connection and sets the write deadline if a timeout is set.
func (c conn) Write(p []byte) (n int, err error) {
	// Set the write deadline if a timeout is set.
	if c.timeout != 0 {
		_ = c.conn.SetWriteDeadline(time.Now().Add(c.timeout))
	}
	return c.conn.Write(p)
}
