package svi

import "github.com/ysmilda/m1-go/internals/m1errors"

// Misc.
const (
	_Align = 0xFFFFFFFC

	PvInfoExtendedCallIdentifier = 0x7575abcd
)

// Errors.
const (
	ErrorMultiBlockTransfer = (m1errors.SourceSVI | m1errors.ErrorMBTRANS)
)

// SVI_F_.
const (
	In                    = 0x80  // type is in (seen from the server)
	Out                   = 0x40  // type is out (seen from the server)
	InOut                 = 0xc0  // type is in and out
	Block                 = 0x20  // block data type
	Hidden                = 0x100 // hidden service variable
	Unknown               = 0x00  // unknown data type
	Uint1                 = 0x01  // 1 bit unsigned integer
	Uint8                 = 0x02  // 8 bit unsigned integer
	Sint8                 = 0x03  // 8 bit signed integer
	Uint16                = 0x04  // 16 bit unsigned integer
	Sint16                = 0x05  // 16 bit signed integer
	Uint32                = 0x06  // 32 bit unsigned integer
	Sint32                = 0x07  // 32 bit signed integer
	Real32                = 0x08  // 32 bit floating point
	Bool                  = 0x09  // boolean
	Char8                 = 0x0a  // 8 bit character
	Mixed                 = 0x0b  // mixed data type; used with svi.Block
	Uint64                = 0x0c  // 64 bit unsigned integer
	Sint64                = 0x0d  // 64 bit signed integer
	Real64                = 0x0e  // 64 bit floating point
	Char16                = 0x0f  // 16 bit Unicode character
	StringListBase        = 0x10  // base value String list type
	UnicodeStringListBase = 0x11  // base value for Unicode String list type

	StringList        = (Block | StringListBase)        // String list type
	UnicodeStringList = (Block | UnicodeStringListBase) // Unicode String list type
	String            = (Block | Char8)                 // String type
	UnicodeString     = (Block | Char16)                // Unicode String type

	TypeMask           = 0x3f // mask for the data type
	ElementaryTypeMask = 0x1f // mask for the elementary data type
)
