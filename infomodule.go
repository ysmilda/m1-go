package m1

import (
	"github.com/ysmilda/m1-go/internals/client"
	"github.com/ysmilda/m1-go/modules/msys"
	"github.com/ysmilda/m1-go/modules/res"
)

type InfoModule struct {
	*Module
}

const (
	_INFO_Procedure_ListCPUAddresses  = 100
	_INFO_Procedure_CPUInfo           = 102
	_INFO_Procedure_ListSystemObjects = 104
	_INFO_Procedure_ListDriverInfo    = 106
	_INFO_Procedure_GetLogInfo        = 108
	_INFO_Procedure_ListTaskInfo      = 110
	_INFO_Procedure_GetBootInfo       = 112
	_INFO_Procedure_EnableCPUTiming   = 114
	_INFO_Procedure_ApplicationName   = 116 // Not sure what this is supposed to do
	_INFO_Procedure_ListCardInfo      = 120
)

func newInfoModule(client *client.Client, info res.ModuleNumber, msysVersion msys.Version) (*InfoModule, error) {
	i, err := newModule(client, "INFO", info, msysVersion)
	if err != nil {
		return nil, err
	}

	return &InfoModule{i}, nil
}

// // ListCPUAddresses returns a list of CPU addresses that make up a CPU-alliance. A maximum of 3 CPUs can be used to
// // make up a network by multi-processing or networking.
// func (i *InfoModule) ListCPUAddresses() ([]inf.CPUAddress, error) {
// 	buf, err := rpc.Call(
// 		i.client.getConnection(i.info),
// 		rpc.Header{
// 			Module:    i.info.ModuleNumber,
// 			Version:   rpc.VersionDefault,
// 			Procedure: _INFO_Procedure_ListCPUAddresses,
// 			Auth:      i.client.auth,
// 		},
// 		uint32(0),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	addressList := &inf.CPUAddressList{}
// 	returnCode := addressList.Parse(buf)
// 	if err := msys.ParseReturnCode(returnCode); err != nil {
// 		return nil, err
// 	}

// 	return addressList.Addresses, nil
// }

// // CPUInfo returns info about the CPU.
// func (i *InfoModule) CPUInfo() (*inf.CPUInfo, error) {
// 	buf, err := rpc.Call(
// 		i.client.getConnection(i.info),
// 		rpc.Header{
// 			Module:    i.info.ModuleNumber,
// 			Version:   rpc.VersionDefault,
// 			Procedure: _INFO_Procedure_CPUInfo,
// 			Auth:      i.client.auth,
// 		},
// 		uint32(0),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	cpu := &inf.CPUInfo{}
// 	returnCode := cpu.Parse(buf)
// 	if err := msys.ParseReturnCode(returnCode); err != nil {
// 		return nil, err
// 	}

// 	return cpu, nil
// }

// // ListSystemObjects returns all available system objects such as MBios, MBoot, etc.
// func (i *InfoModule) ListSystemObjects() ([]inf.SystemObject, error) {
// 	buf, err := rpc.Call(
// 		i.client.getConnection(i.info),
// 		rpc.Header{
// 			Module:    i.info.ModuleNumber,
// 			Version:   rpc.VersionDefault,
// 			Procedure: _INFO_Procedure_ListSystemObjects,
// 			Auth:      i.client.auth,
// 		},
// 		uint32(0),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	returnCode, _ := buf.LittleEndian.ReadUint32()
// 	if err := msys.ParseReturnCode(returnCode); err != nil {
// 		return nil, err
// 	}

// 	count, _ := buf.LittleEndian.ReadUint32()
// 	output := make([]inf.SystemObject, count)
// 	for i := range count {
// 		output[i].Parse(buf)
// 	}

// 	return output, nil
// }

// // ListDriverInfo returns all available I/O drivers such as DI216, DO216, etc.
// func (i *InfoModule) ListDriverInfo() ([]inf.DriverInfo, error) {
// 	const amountPerCall = uint16(15)
// 	index := uint16(1)
// 	output := []inf.DriverInfo{}

// 	for {
// 		start := index
// 		end := index + amountPerCall

// 		buf, err := rpc.Call(
// 			i.client.getConnection(i.info),
// 			rpc.Header{
// 				Module:    i.info.ModuleNumber,
// 				Version:   rpc.VersionDefault,
// 				Procedure: _INFO_Procedure_ListDriverInfo,
// 				Auth:      i.client.auth,
// 			},
// 			start, end,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		returnCode, _ := buf.LittleEndian.ReadUint32()
// 		if err := msys.ParseReturnCode(returnCode); err != nil {
// 			return nil, err
// 		}

// 		driversInResult, _ := buf.LittleEndian.ReadUint32()
// 		for range driversInResult {
// 			entry := inf.DriverInfo{}
// 			entry.Parse(buf)
// 			output = append(output, entry)
// 		}

// 		index += uint16(driversInResult)
// 		if uint16(driversInResult) != amountPerCall {
// 			break
// 		}
// 	}

// 	return output, nil
// }

// // GetLogInfo returns information about the log book. The contents of the logbook can be retrieved via FTP-transfer.
// func (i *InfoModule) GetLogInfo() (*inf.LogInfo, error) {
// 	buf, err := rpc.Call(
// 		i.client.getConnection(i.info),
// 		rpc.Header{
// 			Module:    i.info.ModuleNumber,
// 			Version:   rpc.VersionDefault,
// 			Procedure: _INFO_Procedure_GetLogInfo,
// 			Auth:      i.client.auth,
// 		},
// 		uint32(0),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log := &inf.LogInfo{}
// 	returnCode := log.Parse(buf)
// 	if err := msys.ParseReturnCode(returnCode); err != nil {
// 		return nil, err
// 	}

// 	return log, nil
// }

// // ListTaskInfo returns a list of tasks that are currently running on the system.
// // The performance of the system can be negatively affected by querying this information.
// func (i *InfoModule) ListTaskInfo() (*inf.TaskInfoList, error) {
// 	const amountPerCall = uint32(15)
// 	index := uint32(1)
// 	output := &inf.TaskInfoList{}

// 	for {
// 		start := index
// 		end := index + amountPerCall

// 		buf, err := rpc.Call(
// 			i.client.getConnection(i.info),
// 			rpc.Header{
// 				Module:    i.info.ModuleNumber,
// 				Version:   rpc.VersionDefault,
// 				Procedure: _INFO_Procedure_ListTaskInfo,
// 				Auth:      i.client.auth,
// 			},
// 			start, end, uint32(0),
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		returnCode, _ := buf.LittleEndian.ReadUint32()
// 		if err := msys.ParseReturnCode(returnCode); err != nil {
// 			return nil, err
// 		}

// 		last, _ := buf.LittleEndian.ReadUint32()
// 		output.TotalTime, _ = buf.LittleEndian.ReadInt64()
// 		output.TimeUnits, _ = buf.LittleEndian.ReadUint32()

// 		count, _ := buf.LittleEndian.ReadUint32()
// 		for range count {
// 			entry := inf.TaskInfo{}
// 			entry.Parse(buf)
// 			output.Tasks = append(output.Tasks, entry)
// 		}

// 		index += count
// 		if last != 0 || count < amountPerCall {
// 			break
// 		}
// 	}

// 	return output, nil
// }

// // GetBootInfo return the current boot parameters that are used to boot up the M1 controller.
// func (i *InfoModule) GetBootInfo() (*inf.BootInfo, error) {
// 	buf, err := rpc.Call(
// 		i.client.getConnection(i.info),
// 		rpc.Header{
// 			Module:    i.info.ModuleNumber,
// 			Version:   rpc.VersionDefault,
// 			Procedure: _INFO_Procedure_GetBootInfo,
// 			Auth:      i.client.auth,
// 		},
// 		uint32(0),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &inf.BootInfo{}
// 	returnCode := result.Parse(buf)
// 	if err := msys.ParseReturnCode(returnCode); err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// // EnableCPUTiming enables or disables the measuring of the CPU run and cycle timing.
// func (i *InfoModule) EnableCPUTiming(enable bool) error {
// 	var value uint32
// 	if enable {
// 		value = 1
// 	}

// 	buf, err := rpc.Call(
// 		i.client.getConnection(i.info),
// 		rpc.Header{
// 			Module:    i.info.ModuleNumber,
// 			Version:   rpc.VersionDefault,
// 			Procedure: _INFO_Procedure_EnableCPUTiming,
// 			Auth:      i.client.auth,
// 		},
// 		value,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	returnCode, _ := buf.LittleEndian.ReadUint32()
// 	if err := msys.ParseReturnCode(returnCode); err != nil {
// 		return err
// 	}

// 	enabled, _ := buf.LittleEndian.ReadUint32()
// 	if enabled != value {
// 		return fmt.Errorf("failed to set CPU timing")
// 	}

// 	return nil
// }

// // TODO: Test
// func (i *InfoModule) ListCardInfo() ([]inf.CardInfo, error) {
// 	const amountPerCall = uint32(15)
// 	index := uint32(1)
// 	output := []inf.CardInfo{}

// 	for {
// 		start := index
// 		end := index + amountPerCall

// 		buf, err := rpc.Call(
// 			i.client.getConnection(i.info),
// 			rpc.Header{
// 				Module:    i.info.ModuleNumber,
// 				Version:   rpc.VersionDefault,
// 				Procedure: _INFO_Procedure_ListTaskInfo,
// 				Auth:      i.client.auth,
// 			},
// 			start, end, uint32(0),
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		returnCode, _ := buf.LittleEndian.ReadUint32()
// 		if err := msys.ParseReturnCode(returnCode); err != nil {
// 			return nil, err
// 		}

// 		last, _ := buf.LittleEndian.ReadUint32()
// 		count, _ := buf.LittleEndian.ReadUint32()
// 		for range count {
// 			entry := inf.CardInfo{}
// 			// entry.parse(buf)
// 			output = append(output, entry)
// 		}

// 		index += count
// 		if last != 0 || count < amountPerCall {
// 			break
// 		}
// 	}

// 	return output, nil
// }
