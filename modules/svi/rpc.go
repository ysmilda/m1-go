package svi

import "github.com/ysmilda/m1-go/internals/rpc"

type (
	GetValueCall struct {
		Address Address
	}

	GetValueReply struct {
		rpc.ReturnCode
		Value uint32
	}

	SetValueCall struct {
		Address Address
		Value   uint32
	}

	SetValueReply struct {
		rpc.ReturnCode
	}

	GetValuesCall struct {
		Amount    uint32
		Addresses []Address
	}

	GetValuesReply struct {
		rpc.ReturnCode
		Values []Address `m1binary:"tillEnd"`
	}

	SetValuesCall struct {
		Amount    uint32
		Addresses []Address `m1binary:"lengthRef:Amount"`
		Values    []uint32  `m1binary:"lengthRef:Amount"`
	}

	SetValuesReply struct {
		rpc.ReturnCode
	}

	GetBlockCall struct {
		Address Address
		Length  uint32 // Amount of bytes to read from address
	}

	GetBlockReply struct {
		rpc.ReturnCode
		Count uint32
		Data  []byte `m1binary:"lengthRef:Count"`
	}

	// Multiblock reply.
	GetXBlockReply struct {
		rpc.ReturnCode
		Length   uint32
		BufferID uint32
	}

	GetMultiBlockCall struct {
		Offset   uint32 // Offset from start of block
		BufferID uint32
	}

	GetMultiBlockReply struct {
		rpc.ReturnCode
		Offset           uint32
		Count            uint32
		NumberOfElements uint32
		Data             []byte `m1binary:"lengthRef:Count"`
	}

	SetBlockCall struct {
		Address Address
		Length  uint32
		Data    []byte `m1binary:"lengthRef:Length"`
	}

	SetBlockReply struct {
		rpc.ReturnCode
	}

	SetXBlockReply struct {
		rpc.ReturnCode
		BufferID uint32
	}

	SetMultiBlockCall struct {
		Offset         uint32
		BlockLength    uint32
		RemainingBytes uint32
		BufferID       uint32
		Buffer         []byte
	}

	SetMultiBlockReply struct {
		rpc.ReturnCode
		Type             uint32
		Offset           uint32
		Count            uint32
		NumberOfElements uint32
		Elements         []byte `m1binary:"lengthRef:Count"`
	}

	GetAddressCall struct {
		Name string
	}

	GetAddressReply struct {
		rpc.ReturnCode
		Address Address
		Format  uint16
	}

	GetProcessValueInfoCall struct {
		StartIndex            uint32
		NumberOfProcessValues uint32
	}

	GetProcessValueInfoReply struct {
		rpc.ReturnCode
		Count         uint32
		ProcessValues []ProcessValueInfo `m1binary:"lengthRef:Count"`
	}

	GetExtendedProcessValueInfoCall struct {
		extendedCallIdentifier uint32 // Is set to svi.PvInfoExtendedCallIdentifier
		NumberOfProcessValues  uint32
		ContinuationPoint      uint32 // 0 = start
		GetSubprocessValues    bool   `m1binary:"skip:11"`
		PathLength             uint32
		Path                   string `m1binary:"lengthRef:PathLength"`
	}

	GetExtendedProcessValueInfoReply struct {
		rpc.ReturnCode    `m1binary:"skip:4"`
		ContinuationPoint uint32 `m1binary:"skip:12"` // 0 = end
		Count             uint32
		ProcessValues     []ExtendedProcessValueInfo `m1binary:"lengthRef:Count,allign4"`
	}

	GetServerInfoCall struct{}

	GetServerInfoReply struct {
		rpc.ReturnCode
		ServerInfo
	}

	GetExtendedAddressCall struct {
		Name string
	}

	GetExtendedAddressReply struct {
		rpc.ReturnCode
		Address Address
		Format  uint16
		Length  uint16 `m1binary:"skip:20"`
	}
)
