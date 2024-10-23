package m1

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/ysmilda/m1-go/pkg/rpc"
)

const (
	_VHD_UserNameLength = ((40 + 1 + 3) & 0xfffffffc)

	_VHD_Procedure_StartSession   = 102
	_VHD_Procedure_StopSession    = 104
	_VHD_Procedure_ResetSession   = 106
	_VHD_Procedure_GetSessionInfo = 108
	_VHD_Procedure_GetXAddress    = 134

	_VHD_SessionMode_Polling = 0
)

type VhdModule struct {
	client *client
	info   ModuleInfo

	sessionName string
	userID      uint32
}

// newVhdModule creates a new session for the VHD module. Make sure to have logged in on the client before calling
// this function.
func newVhdModule(client *client, info ModuleInfo) (*VhdModule, error) {
	vhd := &VhdModule{
		client:      client,
		info:        info,
		sessionName: fmt.Sprintf("m1c-%d-%d", rand.Uint32()%1000, time.Now().UnixNano()),
	}

	err := client.addConnection(info)
	if err != nil {
		return nil, fmt.Errorf("failed to add connection: %w", err)
	}

	session, err := vhd.GetSessionInfo()
	if err != nil && !strings.Contains(err.Error(), "failed to get session info") {
		return nil, err // The error is on the transport layer, not on the application layer.
	}

	// If there is an active session but the session name is different, we should reset it to avoid conflicts.
	// Otherwise, we should create a new session.
	if err == nil && session.Name != vhd.sessionName {
		// The returned session name is different from the one we set. We should reset the session.
		err := vhd.ResetSession(session.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to reset session: %w", err)
		}
	} else {
		// The session does not exist or the session name is the same. We should create a new session.
		err := vhd.StartSession()
		if err != nil {
			return nil, fmt.Errorf("failed to start session: %w", err)
		}
	}

	return vhd, nil
}

func (v *VhdModule) Close() error {
	if v.userID != 0 {
		err := v.StopSession()
		if err != nil {
			return fmt.Errorf("failed to stop session: %w", err)
		}
	}
	return nil
}

func (v *VhdModule) StartSession() error {
	delay := 0

	buf, err := rpc.Call(
		v.client.getConnection(v.info),
		rpc.Header{
			Module:    v.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _VHD_Procedure_StartSession,
			Auth:      v.client.auth,
		},
		rpc.NewString(v.sessionName, _VHD_UserNameLength), uint32(delay), uint32(_VHD_SessionMode_Polling),
	)
	if err != nil {
		return err
	}

	returnCode, _ := buf.LittleEndian.ReadUint32()
	if err := parseReturnCode(returnCode); err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}

	v.userID, _ = buf.LittleEndian.ReadUint32()
	return nil
}

func (v *VhdModule) StopSession() error {
	buf, err := rpc.Call(
		v.client.getConnection(v.info),
		rpc.Header{
			Module:    v.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _VHD_Procedure_StopSession,
			Auth:      v.client.auth,
		},
		v.userID,
	)
	if err != nil {
		return err
	}

	returnCode, _ := buf.LittleEndian.ReadUint32()
	if err := parseReturnCode(returnCode); err != nil {
		return fmt.Errorf("failed to stop session: %w", err)
	}

	v.userID = 0

	return nil
}

func (v *VhdModule) ResetSession(userID uint32) error {
	buf, err := rpc.Call(
		v.client.getConnection(v.info),
		rpc.Header{
			Module:    v.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _VHD_Procedure_ResetSession,
			Auth:      v.client.auth,
		},
		userID,
	)
	if err != nil {
		return err
	}

	returnCode, _ := buf.LittleEndian.ReadUint32()
	if err := parseReturnCode(returnCode); err != nil {
		return fmt.Errorf("failed to reset session: %w", err)
	}

	v.userID = userID
	return nil
}

func (v *VhdModule) GetSessionInfo() (*SessionInfo, error) {
	buf, err := rpc.Call(
		v.client.getConnection(v.info),
		rpc.Header{
			Module:    v.info.ModuleNumber,
			Version:   _RPC_VersionDefault,
			Procedure: _VHD_Procedure_GetSessionInfo,
			Auth:      v.client.auth,
		},
		rpc.NewSpare(4), rpc.NewString(v.sessionName, _VHD_UserNameLength),
	)
	if err != nil {
		return nil, err
	}

	reply := &SessionInfo{}
	returnCode := reply.parse(buf)
	if err := parseReturnCode(returnCode); err != nil {
		return nil, fmt.Errorf("failed to get session info: %w", err)
	}

	return reply, nil
}

// InitVariables initializes the variables on the target. It returns the number of variables that were successfully
// initialized. If an error occurs during the communication with the target it will return an error and stop the processing.
// If the response of a single variable is invalid the Error field of the variable will be set and the variable will be
// reset. If the returned count does not equal the number of variables, scan the variables for errors.
func (v *VhdModule) InitVariables(variables []*SviVariable) (initializedVariables int, err error) {
	defer func() {
		if err != nil {
			// TODO: Should this also set the variables to an error state?
			resetVariables(variables)
		}
	}()

	maxCallLength := v.client.getMaximumCallLength()
	maximumEntriesPerCall := (maxCallLength - 8) / 12

	buffer := []*SviVariable{}
	length := 8

	// Loop through all the variables and add them to the buffer.
	// If we reach the maximum number of entries per call or the maximum call length, we send the buffer.
	// If at any point we get an error, we stop and return the error.
	for i, variable := range variables {
		length += len(variable.Name) + 1

		last := i == len(variables)-1

		if len(buffer) == maximumEntriesPerCall || length > maxCallLength || last {
			data := []any{v.userID, uint32(len(buffer))}
			for _, b := range buffer {
				data = append(data, rpc.NewString(b.Name, len(b.Name)+1))
			}

			// We to the buffer after we have done our checks, so we would otherwise miss the last variable.
			if last {
				data = append(data, rpc.NewString(variable.Name, len(variable.Name)+1))
			}

			buf, err := rpc.Call(
				v.client.getConnection(v.info),
				rpc.Header{
					Module:    v.info.ModuleNumber,
					Version:   _RPC_VersionDefault,
					Procedure: _VHD_Procedure_GetXAddress,
					Auth:      v.client.auth,
				},
				data,
			)
			if err != nil {
				return 0, err
			}

			// Verify and parse the response for the variables in the buffer.
			returnCode, _ := buf.LittleEndian.ReadUint32()
			if err := parseReturnCode(returnCode); err != nil {
				// TODO: Set the variables in the call to an error state? That way we can see which variables failed.
				return 0, fmt.Errorf("failed to init variables: %w", err)
			}

			count, _ := buf.LittleEndian.ReadUint32()
			if count != uint32(len(buffer)) {
				return 0, fmt.Errorf("failed to init variables: expected %d variables, got %d", len(buffer), count)
			}

			for _, b := range buffer {
				b.parse(buf)

				if int64(b.Address) == -1 {
					b.Error = fmt.Errorf("failed to init: not found on target")
				} else if b.Length == 0 {
					b.Error = fmt.Errorf("failed to init: has a length of 0")
				}

				if b.Error != nil {
					b.Address = 0
					b.Format = 0
					b.Length = 0
				} else {
					b.Error = nil
					b.initialized = true
					initializedVariables++
				}
			}

			// Reset the counters
			length = 8 + len(variable.Name) + 1
			buffer = []*SviVariable{}
		}

		buffer = append(buffer, variable)
	}

	return initializedVariables, nil
}

func resetVariables(variables []*SviVariable) {
	for _, variable := range variables {
		variable.initialized = false
	}
}
