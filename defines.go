//nolint:lll
package m1

// System Info Handler.
const (
	_BootParameterLength = 20
)

// mio.h.
const (
	_MIO_ProductNumberLength = ((10 + 1 + 3) & 0xfffffffc)
)

// res.h.
const (
	_RES_CompareEqual = 1
	_RES_ReplyWithIP  = 2
	_RES_ReplyNormal  = 0
)

// msys.h.
const (
	_ModuleNameLength = ((8 + 1 + 3) & 0xfffffffc)
	_AppNameLength    = ((20 + 1 + 3) & 0xfffffffc)
	_UserNameLength   = 20
	_PasswordLength   = 16
	_UserNameLength2  = 64
	_PasswordLength2  = 32

	_TimezoneLength = 36
)
