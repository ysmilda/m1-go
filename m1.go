// Package m1 contains an implementation for communicating with the Bachmann M1 PLC platform.
package m1

import (
	"encoding/binary"
	"net"
	"time"
)

type client struct {
	conn    net.Conn
	timeout time.Duration

	auth      []byte
	sessionID uint32
}

func (c client) getConnection() net.Conn {
	_ = c.conn.SetDeadline(time.Time{})
	return c.conn
}

func (c client) getConnectionWithTimeout() net.Conn {
	_ = c.conn.SetDeadline(time.Now().Add(c.timeout))
	return c.conn
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
