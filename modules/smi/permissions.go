package smi

type Permissions uint64

const (
	PermissionsNone                 Permissions = 0x0000000000000000
	PermissionsAll                  Permissions = 0xFFFFFFFFFFFFFFFF
	PermissionsQuerySystemInfo      Permissions = 0x0000000000000001
	PermissionsMeasureRuntime       Permissions = 0x0000000000000002
	PermissionsEditSoftwareModule   Permissions = 0x0000000000000040
	PermissionsSetDateAndTime       Permissions = 0x0000000000000080
	PermissionsReadFiles            Permissions = 0x0000000000000100
	PermissionsWriteFiles           Permissions = 0x0000000000000200
	PermissionsReadSVI              Permissions = 0x0000000000000400
	PermissionsWriteSVI             Permissions = 0x0000000000000800
	PermissionsReadMIOInputs        Permissions = 0x0000000000004000
	PermissionsWriteMIOOutputs      Permissions = 0x0000000000008000
	PermissionsUpdateSystemSoftware Permissions = 0x0000000000020000
	PermissionsUpdateFirmware       Permissions = 0x0000000000040000
	PermissionsReboot               Permissions = 0x0000000000080000
	PermissionsReadMConfig          Permissions = 0x0000000001000000
	PermissionsWriteMConfig         Permissions = 0x0000000002000000
	PermissionsReadConsole          Permissions = 0x0000000004000000
	PermissionsWriteConsole         Permissions = 0x0000000008000000
	PermissionsWriteToBootDevice    Permissions = 0x0000000010000000
	PermissionsFormat               Permissions = 0x0000000020000000
	PermissionsQueryUserInfo        Permissions = 0x0000000040000000
	PermissionsBrowseSVI            Permissions = 0x0000000080000000
	PermissionsTelnet               Permissions = 0x0000000100000000
	PermissionsWebserver            Permissions = 0x0000000200000000
	PermissionsSSH                  Permissions = 0x0000000400000000
	PermissionsUserManagement       Permissions = 0x0000000800000000
	PermissionsAlarmMethods         Permissions = 0x0000001000000000
)
