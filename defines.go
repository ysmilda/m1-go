//nolint:lll
package m1

const (
	_Align = 0xFFFFFFFC
)

// RPC versions.
const (
	_RPC_VersionDefault = 2
	_RPC_VersionRES     = 3
)

// System Info Handler.
const (
	_BootParameterLength = 20
)

// mio.h.
const (
	_MIO_ProductNumberLength = ((10 + 1 + 3) & _Align)
)

// res.h.
const (
	_RES_CompareEqual = 1
	_RES_ReplyWithIP  = 2
	_RES_ReplyNormal  = 0
)

// msys.h.
const (
	_ModuleNameLength = ((8 + 1 + 3) & _Align)
	_AppNameLength    = ((20 + 1 + 3) & _Align)
	_UserNameLength   = 20
	_PasswordLength   = 16
	_UserNameLength2  = 64
	_PasswordLength2  = 32

	_TimezoneLength = 36
)

// svi.h.
const (
	_FormatIn                    = 0x80  // type is in (seen from the server)
	_FormatOut                   = 0x40  // type is out (seen from the server)
	_FormatInOut                 = 0xc0  // type is in and out
	_FormatBlock                 = 0x20  // block data type
	_FormatHidden                = 0x100 // hidden service variable
	_FormatUnknown               = 0x00  // unknown data type
	_FormatUint1                 = 0x01  // 1 bit unsigned integer
	_FormatUint8                 = 0x02  // 8 bit unsigned integer
	_FormatSint8                 = 0x03  // 8 bit signed integer
	_FormatUint16                = 0x04  // 16 bit unsigned integer
	_FormatSint16                = 0x05  // 16 bit signed integer
	_FormatUint32                = 0x06  // 32 bit unsigned integer
	_FormatSint32                = 0x07  // 32 bit signed integer
	_FormatReal32                = 0x08  // 32 bit floating point
	_FormatBool                  = 0x09  // boolean
	_FormatChar8                 = 0x0a  // 8 bit character
	_FormatMixed                 = 0x0b  // mixed data type; used with _FormatBlock
	_FormatUint64                = 0x0c  // 64 bit unsigned integer
	_FormatSint64                = 0x0d  // 64 bit signed integer
	_FormatReal64                = 0x0e  // 64 bit floating point
	_FormatChar16                = 0x0f  // 16 bit Unicode character
	_FormatStringListBase        = 0x10  // base value String list type
	_FormatUnicodeStringListBase = 0x11  // base value for Unicode String list type

	_FormatStringList        = (_FormatBlock | _FormatStringListBase)        // String list type
	_FormatUnicodeStringList = (_FormatBlock | _FormatUnicodeStringListBase) // Unicode String list type
	_FormatString            = (_FormatBlock | _FormatChar8)                 // String type
	_FormatUnicodeString     = (_FormatBlock | _FormatChar16)                // Unicode String type

	_FormatTypeMask           = 0x3f // mask for the data type
	_FormatElementaryTypeMask = 0x1f // mask for the elementary data type

	_SVI_Error_MultiBlockTransfer = (_SOURCE_SVI | _ERROR_MBTRANS)

	_SVI_NameLength = ((63 + 1 + 3) & _Align)
)
