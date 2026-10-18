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
		Values []Address `binary:"tillEnd"`
	}

	SetValuesCall struct {
		Amount    uint32
		Addresses []Address `binary:"lengthRef:Amount"`
		Values    []uint32  `binary:"lengthRef:Amount"`
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
		Data  []byte `binary:"lengthRef:Count"`
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
		Data             []byte `binary:"lengthRef:Count"`
	}

	SetBlockCall struct {
		Address Address
		Length  uint32
		Data    []byte `binary:"lengthRef:Length"`
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
		Elements         []byte `binary:"lengthRef:Count"`
	}

	GetAddressCall struct {
		Name string
	}

	GetAddressReply struct {
		rpc.ReturnCode
		Address Address
		Format  uint16
	}

	ListProcessValueInfoCall struct {
		rpc.PaginatedCallStartCount
	}

	ListProcessValueInfoReply struct {
		rpc.ReturnCode
		rpc.PaginatedReplyCount[ProcessValueInfo]
	}

	ListExtendedProcessValueInfoCall struct {
		ExtendedCallIdentifier uint32 // Is set to svi.PvInfoExtendedCallIdentifier
		rpc.PaginatedCallCountStart
		GetSubprocessValues bool `binary:"skip:11"`
		PathLength          uint32
		Path                string `binary:"lengthRef:PathLength"`
	}

	ListExtendedProcessValueInfoReply struct {
		rpc.ReturnCode `binary:"skip:4"`
		rpc.PaginatedReplyContinuationCount[ExtendedProcessValueInfo]
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
		Length  uint16 `binary:"skip:20"`
	}
)
