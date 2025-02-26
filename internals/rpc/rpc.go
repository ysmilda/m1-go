// Package rpc contains the RPC client implementation.
package rpc

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/ysmilda/m1-go/internals/m1binary"
	"github.com/ysmilda/m1-go/internals/m1errors"
)

type Version uint32

const (
	VersionDefault Version = 2
	VersionRES     Version = 3
)

const (
	callIndicator        uint32 = 0
	replyIndicator       uint32 = 1
	messageVersion       uint32 = 2
	messageAccepted      uint32 = 0
	programMismatch      uint32 = 2
	procedureUnavailable uint32 = 3
	mismatch             uint32 = 0
	authError            uint32 = 1

	bufferSize = 4096
)

// Header contains the header of an RPC call.
type Header struct {
	xID uint32

	Module    uint32
	Version   Version
	Procedure uint32
	Auth      []byte
}

// Call sends an RPC call to the target with the given header and procedure.
// It packs and unpacks the call and reply objects and checks for errors.
func Call[C any, R ReturnCoder](rw io.ReadWriter, header Header, procedure Procedure[C, R]) (*R, error) {
	body, err := m1binary.Encode(&procedure.Call)
	if err != nil {
		return nil, fmt.Errorf("unable to parse call: %w", err)
	}

	buf, err := call(rw, header, body)
	if err != nil {
		return nil, fmt.Errorf("unable to make rpc call: %w", err)
	}

	reply := new(R)
	_, err = m1binary.Decode(buf, reply)
	if err != nil {
		return nil, fmt.Errorf("unable to parse reply: %w", err)
	}

	if err := m1errors.ParseReturnCode((*reply).GetReturnCode()); err != nil {
		return nil, err
	}

	return reply, nil
}

// call sends an RPC call to the target with the given header and call data.
// It returns the response from the target or an error if the call failed.
// If the response object is nill the response won't be read.
func call(rw io.ReadWriter, header Header, data []byte) ([]byte, error) {
	body := []byte{}
	header.xID = uint32(rand.Int32())

	body = writeHeader(body, header)
	body = append(body, data...)

	_, err := rw.Write(body)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, bufferSize)
	n, err := rw.Read(buf)
	if err != nil {
		return nil, err
	}

	msg := buf[:n]
	// Check if the response has the correct xID.
	if val := readUint32(msg); val != header.xID {
		return nil, ErrInvalidXID
	}

	// Check if the response is valid.
	n, err = verifyResponse(msg[4:])
	if err != nil {
		return nil, err
	}

	return msg[4+n:], nil
}

// writeHeader writes the header of the RPC call.
// This should always be the first thing written to the buffer.
func writeHeader(buf []byte, header Header) []byte {
	buf, _ = binary.Append(buf, binary.BigEndian, header.xID)
	buf, _ = binary.Append(buf, binary.BigEndian, callIndicator)
	buf, _ = binary.Append(buf, binary.BigEndian, messageVersion)
	buf, _ = binary.Append(buf, binary.BigEndian, header.Module)
	buf, _ = binary.Append(buf, binary.BigEndian, header.Version)
	buf, _ = binary.Append(buf, binary.BigEndian, header.Procedure)

	if header.Auth == nil {
		buf = append(buf, make([]byte, 8)...)
	} else {
		buf = append(buf, header.Auth...)
	}

	buf = append(buf, make([]byte, 8)...)

	return buf
}

// verifyResponse checks if the response is valid.
// This should always be called after verifying the xID.
// If the response is invalid, an error is returned.
func verifyResponse(body []byte) (int, error) {
	if code := readUint32(body); code != replyIndicator {
		return 4, ErrNoReplyFrame
	}

	switch code := readUint32(body[4:]); code {
	case messageAccepted:
		switch code := readUint32(body[16:]); code {
		case programMismatch:
			return 20, ErrProgramMismatch
		case procedureUnavailable:
			return 20, ErrProcedureUnavailable
		default:
			return 20, nil
		}

	default:
		switch code := readUint32(body[8:]); code {
		case mismatch:
			return 12, ErrRPCMismatch
		case authError:
			return 12, ErrAuthError
		default:
			return 12, fmt.Errorf("%w: unknown error code (%d)", ErrInvalidResponse, code)
		}
	}
}

func readUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}
