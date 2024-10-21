// Package rpc contains the RPC client implementation.
package rpc

import (
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/ysmilda/m1-go/pkg/buffer"
)

const (
	_RPC_CALL_INDICATOR    = 0
	_RPC_REPLY_INDICATOR   = 1
	_RPC_MESSAGE_VERSION   = 2
	_MESSAGE_ACCEPTED      = 0
	_PROGRAM_MISMATCH      = 2
	_PROCEDURE_UNAVAILABLE = 3
	_RPC_MISMATCH          = 0
	_AUTH_ERROR            = 1

	bufferSize = 2048
)

// Header contains the header of an RPC call.
type Header struct {
	xID uint32

	Module    uint32
	Version   uint32
	Procedure uint32
	Auth      []byte
}

// Call sends an RPC call to the target with the given header and data.
// It returns the response from the target or an error if the call failed.
// The response has been read up until the start of the user data.
func Call(rw io.ReadWriter, header Header, data ...any) (*buffer.Buffer, error) {
	body := buffer.NewBuffer(nil)
	body.Reset()
	header.xID = uint32(rand.Int32()) // nolint:gosec

	writeHeader(body, header)
	writeData(body, data...)

	_, err := rw.Write(body.Bytes())
	if err != nil {
		return nil, err
	}

	buf := make([]byte, bufferSize)
	n, err := rw.Read(buf)
	if err != nil {
		return nil, err
	}
	body.Reset()
	_, _ = body.Write(buf[:n])

	// TODO: The original code has a retry mechanism here. Maybe add it later.
	// Check if the response has the correct xID.
	if val, _ := body.BigEndian.ReadUint32(); val != header.xID {
		return nil, ErrInvalidXID
	}

	// Check if the response is valid.
	err = verifyResponse(body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// CallWithoutRead sends an RPC call to the target with the given header and data.
// It returns an error if the call failed.
// The response from the target is not read.
func CallWithoutRead(w io.Writer, header Header, data ...any) error {
	body := buffer.NewBuffer(nil)
	body.Reset()
	header.xID = uint32(rand.Int32()) // nolint:gosec

	writeHeader(body, header)
	writeData(body, data...)

	_, err := w.Write(body.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// writeHeader writes the header of the RPC call.
// This should always be the first thing written to the buffer.
func writeHeader(body *buffer.Buffer, header Header) {
	body.BigEndian.WriteUint32(header.xID)
	body.BigEndian.WriteUint32(_RPC_CALL_INDICATOR)
	body.BigEndian.WriteUint32(_RPC_MESSAGE_VERSION)
	body.BigEndian.WriteUint32(header.Module)
	body.BigEndian.WriteUint32(header.Version)
	body.BigEndian.WriteUint32(header.Procedure)

	if header.Auth == nil {
		body.BigEndian.WriteUint32(0)
		body.BigEndian.WriteUint32(0)
	} else {
		_, _ = body.Write(header.Auth)
	}

	body.BigEndian.WriteUint32(0)
	body.BigEndian.WriteUint32(0)
}

// writeData writes data to the buffer.
// The data can be of type uint8, uint16, uint32, uint64, []byte, string or rpc.String.
// This should always be called after writing the header.
func writeData(body *buffer.Buffer, data ...any) {
	for _, d := range data {
		switch v := d.(type) {
		case uint8:
			_ = body.WriteByte(v)
		case uint16:
			body.BigEndian.WriteUint16(v)
		case uint32:
			body.BigEndian.WriteUint32(v)
		case uint64:
			body.BigEndian.WriteUint64(v)
		case []byte:
			_, _ = body.Write(v)
		case string:
			_, _ = body.WriteString(v)
		case String:
			body.WriteStringTillLength(v.value, v.length)
		case Spare:
			_, _ = body.Write(make([]byte, v.length))

		default:
			panic("unsupported type")
		}
	}
}

// verifyResponse checks if the response is valid.
// This should always be called after verifying the xID.
// If the response is invalid, an error is returned.
func verifyResponse(body *buffer.Buffer) error {
	if code, _ := body.BigEndian.ReadUint32(); code != _RPC_REPLY_INDICATOR {
		fmt.Println("Invalid response", code)
		return ErrNoReplyFrame
	}

	switch code, _ := body.BigEndian.ReadUint32(); code {
	case _MESSAGE_ACCEPTED:
		body.Skip(8) // dummies

		switch code, _ := body.BigEndian.ReadUint32(); code {
		case _PROGRAM_MISMATCH:
			return ErrProgramMismatch
		case _PROCEDURE_UNAVAILABLE:
			return ErrProcedureUnavailable
		}

	default:
		switch code, _ := body.BigEndian.ReadUint32(); code {
		case _RPC_MISMATCH:
			return ErrRPCMismatch
		case _AUTH_ERROR:
			return ErrAuthError
		default:
			return fmt.Errorf("%w: unknown error code (%d)", ErrInvalidResponse, code)
		}
	}
	return nil
}
