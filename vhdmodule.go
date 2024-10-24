package m1

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/ysmilda/m1-go/pkg/buffer"
	"github.com/ysmilda/m1-go/pkg/rpc"
)

const (
	_VHD_UserNameLength = ((40 + 1 + 3) & _Align)

	_VHD_Procedure_StartSession   = 102
	_VHD_Procedure_StopSession    = 104
	_VHD_Procedure_ResetSession   = 106
	_VHD_Procedure_GetSessionInfo = 108
	_VHD_Procedure_GetValue       = 110
	_VHD_Procedure_GetXAddress    = 134

	_SVI_Procedure_GetMultiBlock = 10018

	_VHD_SessionMode_Polling = 0
	_VHD_ListNoSort          = 0
	_VHD_SingleRtype         = 7
	_VHD_SingleBtype         = 4
)

var ErrFailedToInitializeVariables = fmt.Errorf("failed to initialize variables")

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

// InitializeVariables initializes the variables on the target. It returns the number of variables that were successfully
// initialized. If an error occurs during the communication with the target it will return an error and stop the processing.
// If the response of a single variable is invalid the Error field of the variable will be set and the variable will be
// reset. If the returned count does not equal the number of variables, scan the variables for errors.
func (v *VhdModule) InitializeVariables(variables []*SviVariable) (initializedVariables int, err error) {
	defer func() {
		if err != nil {
			// TODO: Should this also set the variables to an error state?
			resetVariables(variables)
		}
	}()

	calls := v.splitInitialisationCall(variables)

	for _, call := range calls {
		data := []any{v.userID, uint32(len(call))}
		for _, v := range call {
			data = append(data, rpc.NewString(v.Name, len(v.Name)+1))
		}

		// We to the buffer after we have done our checks, so we would otherwise miss the last variable.

		buf, err := rpc.Call(
			v.client.getConnection(v.info),
			rpc.Header{
				Module:    v.info.ModuleNumber,
				Version:   _RPC_VersionDefault,
				Procedure: _VHD_Procedure_GetXAddress,
				Auth:      v.client.auth,
			},
			data...,
		)
		if err != nil {
			return 0, err
		}

		// Verify and parse the response for the variables in the buffer.
		returnCode, _ := buf.LittleEndian.ReadUint32()
		if err := parseReturnCode(returnCode); err != nil {
			// TODO: Set the variables in the call to an error state? That way we can see which variables failed.
			return 0, fmt.Errorf("%w: %w", ErrFailedToInitializeVariables, err)
		}

		count, _ := buf.LittleEndian.ReadUint32()
		if count != uint32(len(call)) {
			return 0, fmt.Errorf("%w: expected %d variables, got %d", ErrFailedToInitializeVariables, len(call), count)
		}

		for _, v := range call {
			v.parse(buf)

			if int64(v.Address) == -1 {
				v.Error = fmt.Errorf("%w: not found on target", ErrFailedToInitializeVariables)
			} else if v.Length == 0 {
				v.Error = fmt.Errorf("%w: has a length of 0", ErrFailedToInitializeVariables)
			}

			if v.Error != nil {
				v.Address = 0
				v.Format = 0
				v.Length = 0
			} else {
				v.Error = nil
				v.initialized = true
				initializedVariables++
			}
		}
	}

	return initializedVariables, nil
}

func (v *VhdModule) ReadVariables(variables []*SviVariable) (int, error) {
	// Split the variables into separate calls based on the maximum call length.
	calls := v.splitReadCall(variables)
	count := 0

	for _, call := range calls {
		data := []any{v.userID, uint32(_VHD_ListNoSort), uint32(len(call))}
		for _, variable := range call {
			if variable.IsBlock() {
				data = append(data, uint32(_VHD_SingleRtype), variable.Length)
			} else {
				data = append(data, uint32(_VHD_SingleBtype))
			}
			data = append(data, variable.Address)
		}

		buf, err := rpc.Call(
			v.client.getConnection(v.info),
			rpc.Header{
				Module:    v.info.ModuleNumber,
				Version:   _RPC_VersionDefault,
				Procedure: _VHD_Procedure_GetValue,
				Auth:      v.client.auth,
			},
			data...,
		)
		if err != nil {
			return 0, err
		}

		// In this case the return code indicates the type of reponse we got.
		returnCode, _ := buf.LittleEndian.ReadUint32()
		if returnCode == (_SVI_Error_MultiBlockTransfer | _SOURCE_VHD) { // Multiblock
			buf.Skip(2 * 4) // Ignore the sub return code and the spare
			bufferID, _ := buf.LittleEndian.ReadUint32()
			buf.Skip(4) // Ignore the spare

			err := v.readMultiBlock(variables, bufferID)
			if err != nil {
				return 0, fmt.Errorf("failed to read variables: %w", err)
			}
			count++
		} else if err := parseReturnCode(returnCode); err != nil { // Actual error
			return 0, fmt.Errorf("failed to read variables: %w", err)
		}

		count, _ := buf.LittleEndian.ReadUint32()
		if count != uint32(len(call)) {
			return 0, fmt.Errorf("failed to read variables: expected %d variables, got %d", len(call), count)
		}

		for _, variable := range call {
			err := v.readVariable(variable, buf)
			if err != nil {
				return 0, fmt.Errorf("failed to read variables: %w", err)
			}
			count++
		}
	}

	return count, nil
}

func (v *VhdModule) splitInitialisationCall(variables []*SviVariable) [][]*SviVariable {
	output := [][]*SviVariable{}
	buffer := []*SviVariable{}

	length := 8
	count := 0
	maxCallLength := v.client.getMaximumCallLength()
	maxCount := (maxCallLength - 8) / 12

	for _, variable := range variables {
		count++
		length += len(variable.Name) + 1

		if count == maxCount || length >= maxCallLength {
			output = append(output, buffer)
			buffer = []*SviVariable{}
			count = 0
			length = 8 + len(variable.Name) + 1
		}

		buffer = append(buffer, variable)
	}

	if len(buffer) > 0 {
		output = append(output, buffer)
	}

	return output
}

func (v *VhdModule) splitReadCall(variables []*SviVariable) [][]*SviVariable {
	maxCallLength := v.client.getMaximumCallLength()
	callLength := 12 // UserID + report mode + Number of elements (3 x sizeof(uint32))
	replyLength := 8 // Return code + number of elements (2 x sizeof(uint32))

	output := [][]*SviVariable{}
	buffer := []*SviVariable{}

	for _, variable := range variables {
		var (
			variableCallLength  int
			variableReplyLength int
		)

		if !variable.initialized {
			continue
		}

		if variable.IsBlock() {
			variableCallLength = 16                               // Type + byte length + address
			variableReplyLength = 12 + variable.getBufferLength() // type + index + buffer length
		} else {
			variableCallLength = 12                              // Type + address
			variableReplyLength = 8 + variable.getBufferLength() // type + index + buffer length
		}

		if callLength+variableCallLength > maxCallLength || replyLength+variableReplyLength > maxCallLength {
			output = append(output, buffer)
			buffer = []*SviVariable{}
			callLength = 12
			replyLength = 8
		}

		buffer = append(buffer, variable)
		callLength += variableCallLength
		replyLength += variableReplyLength
	}

	if len(buffer) > 0 {
		output = append(output, buffer)
	}

	return output
}

func (v *VhdModule) readMultiBlock(variables []*SviVariable, bufferID uint32) error {
	if len(variables) != 1 {
		return fmt.Errorf("failed to read multiblock: expected 1 variable, got %d", len(variables))
	}

	variable := variables[0]

	// The first variable in the buffer is the number of elements.
	offset := uint32(0)
	blockLength := uint32(1)
	headerRead := false
	buffer := make([]byte, 0, v.client.maxCallLength)

	for blockLength != 0 {
		res, err := rpc.Call(
			v.client.getConnection(v.info),
			rpc.Header{
				Module:    v.info.ModuleNumber,
				Version:   _RPC_VersionDefault,
				Procedure: _SVI_Procedure_GetMultiBlock,
				Auth:      v.client.auth,
			},
			offset, bufferID,
		)
		if err != nil {
			return err
		}

		returnCode, _ := res.LittleEndian.ReadUint32()
		if err := parseReturnCode(returnCode); err != nil {
			return fmt.Errorf("failed to read multiblock: %w", err)
		}

		offset, _ = res.LittleEndian.ReadUint32()
		blockLength, _ = res.LittleEndian.ReadUint32()
		res.Skip(4) // Skip the number of elements

		if blockLength == 0 {
			break
		}

		if !headerRead {
			res.Skip(2 * 4) // Skip type and length from the header
			index, _ := res.LittleEndian.ReadUint32()

			if int32(index) < 0 {
				variable.Error = fmt.Errorf("failed to read multiblock: invalid index")
				return variable.Error
			}
			headerRead = true

		}

		line, _ := res.ReadBytes(int(blockLength) - 12) // Total length - header length
		buffer = append(buffer, line...)
	}

	// TODO: Not quite sure what to do here. The multi block read should be parsed somehow...
	// For now we just set the value to the buffer.
	variable.Value = buffer

	return nil
}

func (v *VhdModule) readVariable(variable *SviVariable, buf *buffer.Buffer) error {
	_type, _ := buf.LittleEndian.ReadUint32()
	switch _type {
	case _VHD_SingleRtype:
		buf.Skip(4) // Skip the length, we don't need it
	case _VHD_SingleBtype:
	default:
		return fmt.Errorf("failed to read variables: invalid type %d", _type)
	}

	index, _ := buf.LittleEndian.ReadUint32()
	if int32(index) < 0 {
		variable.Error = fmt.Errorf("failed to read variables: invalid index")
		buf.Skip(uint(variable.getBufferLength())) // Skip the buffer
		return nil
	}

	variable.parse(buf)
	return nil
}

func resetVariables(variables []*SviVariable) {
	for _, variable := range variables {
		variable.initialized = false
	}
}
