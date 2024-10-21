package rpc

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidXID is returned when an invalid xID is received.
	ErrInvalidXID = errors.New("invalid xID")
	// ErrInvalidResponse is returned when an invalid response is received.
	ErrInvalidResponse = errors.New("invalid response")
	// ErrNoReplyFrame is returned when no reply frame is received. It wraps ErrInvalidResponse.
	ErrNoReplyFrame = fmt.Errorf("%w: no reply frame received", ErrInvalidResponse)
	// ErrProgramMismatch is returned when the program does not match.
	ErrProgramMismatch = errors.New("program mismatch")
	// ErrProcedureUnavailable is returned when the procedure is unavailable.
	ErrProcedureUnavailable = errors.New("procedure unavailable")
	// ErrRPCMismatch is returned when the RPC does not match.
	ErrRPCMismatch = errors.New("rpc mismatch")
	// ErrAuthError is returned when an authentication error occurs.
	ErrAuthError = errors.New("auth error")
)
