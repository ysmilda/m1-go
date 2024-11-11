// Package m1 contains an implementation for communicating with the Bachmann M1 PLC platform.
package m1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

// client is the main struct that holds the connection to the target.
// It is intended solely for internal use and should not be used directly.
// On improper usage, the client may panic.
// It is used to manage the connections to the target and to store the authentication information.
type client struct {
	// overall configuration
	ip       net.IP
	timeout  time.Duration
	protocol string

	// all connections for all modules. Every module has its own connection.
	connections map[ModuleInfo]*conn

	auth          []byte
	sessionID     uint32
	maxCallLength int32
}

var validProtocols = map[string]bool{
	"tcp": true,
	"udp": true,
}

func newClient(ip net.IP, timeout time.Duration, protocol string) *client {
	if _, ok := validProtocols[protocol]; !ok {
		panic(fmt.Sprintf("invalid protocol %s", protocol))
	}

	return &client{
		ip:            ip,
		timeout:       timeout,
		protocol:      protocol,
		connections:   make(map[ModuleInfo]*conn),
		maxCallLength: 2048,
	}
}

func (c client) close() error {
	for _, conn := range c.connections {
		err := conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *client) addConnection(module ModuleInfo) error {
	var port uint16

	switch c.protocol {
	case "tcp":
		port = module.TCPPort
	case "udp":
		port = module.UDPPort
	}

	var err error
	nc, err := net.Dial(c.protocol, fmt.Sprintf("%s:%d", c.ip, port))
	if err != nil {
		return fmt.Errorf("failed to connect to %s:%d: %w", c.ip, port, err)
	}

	connection := &conn{
		conn:    nc,
		isTCP:   c.protocol == "tcp",
		timeout: c.timeout,
	}
	c.connections[module] = connection
	return nil
}

func (c *client) getConnection(module ModuleInfo) *conn {
	connection, ok := c.connections[module]
	if !ok {
		panic("module not initialized")
	}
	return connection
}

func (c *client) setAuth(auth []byte, length uint32) {
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

func (c *client) getMaximumCallLength() int {
	if len(c.auth) == 0 {
		return int(c.maxCallLength - 40)
	}
	return int(c.maxCallLength - 40 - int32(len(c.auth)))
}

// conn is a wrapper around a net.Conn that adds some functionality for the M1 protocol.
type conn struct {
	conn    net.Conn
	isTCP   bool
	timeout time.Duration
}

// Close closes the underlying connection.
func (c conn) Close() error {
	return c.conn.Close()
}

func (c conn) Read(p []byte) (n int, err error) {
	// Set the read deadline if a timeout is set.
	if c.timeout != 0 {
		_ = c.conn.SetReadDeadline(time.Now().Add(c.timeout))
	}

	// TCP encodes the length of the message in the first 4 bytes.
	// Verify the record marking and read the length of the message and remove them from the buffer.
	if c.isTCP {
		temp := make([]byte, 4)
		_, err = c.conn.Read(temp)
		if err != nil {
			return 0, err
		}

		recordMarking := binary.BigEndian.Uint32(temp)
		if recordMarking&0x80000000 == 0 {
			return 0, errors.New("invalid record marking")
		}

		length := recordMarking & 0x7FFFFFFF
		log.Println("length", length)
		if length > uint32(len(p)) {
			return 0, errors.New("buffer too small")
		}
		return c.conn.Read(p[:length])
	}

	return c.conn.Read(p)
}

func (c conn) Write(p []byte) (n int, err error) {
	// TCP adds a marking to the message to indicate the length of the message.
	if c.isTCP {
		recordMarking := 0x80000000 | uint32(len(p))
		_ = binary.Write(c.conn, binary.BigEndian, recordMarking)
	}
	// Set the write deadline if a timeout is set.
	if c.timeout != 0 {
		_ = c.conn.SetWriteDeadline(time.Now().Add(c.timeout))
	}
	return c.conn.Write(p)
}
