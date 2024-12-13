package vhd

import (
	"time"

	"github.com/ysmilda/m1-go/modules/svi"
)

type StartSessionCall struct {
	UserName    string        `unpack:"length=44"`
	DelayTime   time.Duration `unpack:"length=4,unit=milliseconds"`
	SessionMode SessionMode
}

type StartSessionReply struct {
	ReturnCode
	UserID uint32
}

type StopSessionCall struct {
	UserID uint32
}

type StopSessionReply struct {
	ReturnCode
}

type ResetSessionCall struct {
	UserID uint32
}

type ResetSessionReply struct {
	ReturnCode
}

type GetSessionInfoCall struct {
	UserID   uint32
	UserName string `unpack:"length=44"`
}

type GetSessionInfoReply struct {
	ReturnCode
	SessionInfo
}

type SetValueCall struct {
	UserID           uint32
	ReportMode       ReportMode
	NumberOfElements uint32
	Elements         []ElementValue `unpack:"length=NumberOfElements"`
}

type SetValueReply struct {
	ReturnCode
	NumberOfElements uint32
	// Contains the adresses of the invalid values from the call
	InvalidAddresses []uint64 `unpack:"length=NumberOfElements"`
}

type SetXValueCall struct {
	UserID                       uint32
	ReportMode                   ReportMode
	NumberOfElements             uint32
	TotalTransferLength          uint32
	NumberOfMultiBlockElelements uint32
}

type SetXValueReply struct {
	ReturnCode
	BufferID uint32
}

type SetXValueErrReply struct {
	ReturnCode
	NumberOfElements uint32
	// Contains the adresses of the invalid values from the call
	InvalidAddresses []uint64 `unpack:"length=NumberOfElements"`
}

type GetValueCall struct {
	UserID           uint32
	ReportMode       ReportMode
	NumberOfElements uint32
	Elements         []byte `unpack:"readTillEnd"`
}

type GetValueReply struct {
	ReturnCode
	NumberOfElements uint32
	Elements         []byte `unpack:"readTillEnd"`
}

type GetXValueReply struct {
	ReturnCode
	Reserved               uint32
	NumberOfElements       uint32
	BufferID               uint32
	NumberOfBufferElements uint32
}

type AddListCall struct {
	UserID           uint32
	ObservationMode  ObservationMode
	UserParameter    uint32
	NumberOfElements uint32
	Elements         []ElementAddress `unpack:"length=NumberOfElements"`
}

type AddListReply struct {
	ReturnCode
	ListID           uint32
	NumberOfElements uint32
	Elements         []byte `unpack:"readTillEnd"`
}

type AddXListReply struct {
	ReturnCode
	Reserved         uint32
	ListID           uint32
	NumberOfElements uint32
	BufferID         uint32
}

type DeleteListCall struct {
	UserID      uint32
	NumberOfIDs uint32
	ListIDs     []uint32 `unpack:"length=NumberOfIDs"`
}

type DeleteListReply struct {
	ReturnCode
}

type ResetListCall struct {
	UserID      uint32
	NumberOfIDs uint32
	ListIDs     []uint32 `unpack:"length=NumberOfIDs"`
}

type ResetListReply struct {
	ReturnCode
}

type StartListCall struct {
	UserID      uint32
	NumberOfIDs uint32
	ListIDs     []uint32 `unpack:"length=NumberOfIDs"`
}

type StartListReply struct {
	ReturnCode
}

type StopListCall struct {
	UserID      uint32
	NumberOfIDs uint32
	ListIDs     []uint32 `unpack:"length=NumberOfIDs"`
}

type StopListReply struct {
	ReturnCode
}

type StartElementCall struct {
	UserID uint32
	ListID uint32
	Index  uint32 // Index of element in list
}

type StartElementReply struct {
	ReturnCode
}

type StopElementCall struct {
	UserID uint32
	ListID uint32
	Index  uint32 // Index of element in list
}

type StopElementReply struct {
	ReturnCode
}

type GetUpdateCall struct {
	UserID uint32
	ListID uint32
}

type GetUpdateReply struct {
	ReturnCode
	NumberOfElements uint32
	Elements         []byte `unpack:"readTillEnd"`
}

type GetXUpdateReply struct {
	ReturnCode
	Reserved         uint32
	NumberOfElements uint32
	BufferID         uint32
}

type NewUpdateCall struct {
	ListID           uint32
	UserParameter    uint32
	NumberOfElements uint32
	Elements         []ElementAddress `unpack:"length=NumberOfElements"`
}

type NewUpdateReply struct {
	ReturnCode
	UserID uint32
	ListID uint32
}

type GetAddressCall struct {
	UserID        uint32
	NumberOfPaths uint32
	// Zero-terminated strings
	Paths []string `unpack:"length=NumberOfPaths,zeroTerminate"`
}

type GetAddressReply struct {
	ReturnCode
	NumberOfAddresses uint32
	Addresses         []AddressInfo `unpack:"length=NumberOfAddresses"`
}

type GetXAddressCall struct {
	UserID        uint32
	NumberOfPaths uint32
	// Zero-terminated strings
	Paths []string `unpack:"length=NumberOfPaths,zeroTerminate"`
}

type GetXAddressReply struct {
	ReturnCode
	NumberOfAddresses uint32
	Addresses         []svi.Variable `unpack:"length=NumberOfAddresses"`
}

type GetCallbackInfoCall struct {
	UserID          uint32
	ServerListIndex uint32
	Spare           []uint32 `unpack:"length=2"`
}

type GetCallbackInfoReply struct {
	ReturnCode
	CallBackInfo
}

type ReturnCode uint32

func (r ReturnCode) GetReturnCode() uint32 { return uint32(r) }
