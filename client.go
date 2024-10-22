// Package m1 contains an implementation for communicating with the Bachmann M1 PLC platform.
package m1

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
	"time"
)

type client struct {
	conn    net.Conn
	timeout time.Duration

	isTCP bool

	auth      []byte
	sessionID uint32
}

func newClient(conn net.Conn, timeout time.Duration, isTCP bool) *client {
	return &client{
		conn:    conn,
		timeout: timeout,
		isTCP:   isTCP,
	}
}

func (c client) Read(p []byte) (n int, err error) {
	if c.timeout != 0 {
		_ = c.conn.SetReadDeadline(time.Now().Add(c.timeout))
	}
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

func (c client) Write(p []byte) (n int, err error) {
	if c.isTCP {
		recordMarking := 0x80000000 | uint32(len(p))
		_ = binary.Write(c.conn, binary.BigEndian, recordMarking)
	}
	if c.timeout != 0 {
		_ = c.conn.SetWriteDeadline(time.Now().Add(c.timeout))
	}
	return c.conn.Write(p)
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
