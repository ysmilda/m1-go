//nolint:lll
package m1

import "fmt"

func parseReturnCode(code uint32) error {
	if code == _OK {
		return nil
	}

	return fmt.Errorf("%s - %s", getErrorSource(code), getErrorDescription(code))
}

func getErrorSource(code uint32) string { //nolint: gocyclo
	switch code & 0xFFFF0000 {
	case _SOURCE_NO:
		return "Unknown source"
	case _SOURCE_SVI:
		return "SVI Functions"
	case _SOURCE_SMI:
		return "SMI Functions"
	case _SOURCE_RES:
		return "Resource Handler"
	case _SOURCE_MIO:
		return "IO Handler"
	case _SOURCE_VHD:
		return "Vis Handler"
	case _SOURCE_INF:
		return "Info Handler"
	case _SOURCE_PLC:
		return "PLC Runtime system"
	case _SOURCE_MOD:
		return "Module Handler"
	case _SOURCE_CAN:
		return "CAN Handler"
	case _SOURCE_PF:
		return "Profile Functions"
	case _SOURCE_SYS:
		return "System MSys"
	case _SOURCE_CORE:
		return "System MCore"
	case _SOURCE_EHD:
		return "Error Handler"
	case _SOURCE_PB:
		return "Profibus Handler"
	case _SOURCE_DBG:
		return "Debug Handler"
	case _SOURCE_DN:
		return "DeviceNet Handler"
	case _SOURCE_RFS:
		return "Remote File Server"
	case _SOURCE_SLC:
		return "SLC Server"
	case _SOURCE_DMW:
		return "Drive Middleware"
	case _SOURCE_SEM201:
		return "SERCOS driver"
	case _SOURCE_UFB:
		return "Unified fieldbus components"
	case _SOURCE_PN:
		return "Profinet"
	case _SOURCE_EC:
		return "EtherCAT"
	case _SOURCE_BCR:
		return "BCR2xx"
	case _SOURCE_ST:
		return "Self Test Module"
	case _SOURCE_C_TDLL:
		return "TMANw32-DLL"
	case _SOURCE_C_PLCCOM:
		return "PLCCOM-DLL"
	case _SOURCE_C_TCONF:
		return "TargetManager EXE"
	case _SOURCE_C_PLCHWM:
		return "PLCHWM-DLL"
	case _SOURCE_C_TVIEW:
		return "TMAN-View"
	case _SOURCE_C_MMAN:
		return "M-Manager"
	case _SOURCE_C_MPLC:
		return "M-PLC"
	case _SOURCE_C_MIF:
		return "M-Interface"
	case _SOURCE__AHD:
		return "AHD System SW Module"
	case _SOURCE__LOGGER:
		return "Logger System SW Module"
	case _SOURCE_EVT:
		return "Event Logger"
	case _SOURCE_ETCP:
		return "60870-5"
	case _SOURCE_DNP3:
		return "DNP3"
	case _SOURCE_ATEC:
		return "ATEC"
	case _SOURCE_M1C:
		return "M1 Core"
	}
	return ""
}

func getErrorDescription(code uint32) string { //nolint: gocyclo
	switch code {
	case _OK:
		return "O.K., no error"
	case _ERROR_INPROGRESS:
		return "Still in progress, no error"
	case _ERROR_CANCELED:
		return "Canceled by user, no error"
	case _ERROR_FAILED:
		return "Unspecified error"
	case _ERROR_BADPROG:
		return "SMI-Program not valid"
	case _ERROR_BADPROC:
		return "SMI-Procedure not valid"
	case _ERROR_BADARGS:
		return "SMI-Arguments not valid"
	case _ERROR_BADRPC:
		return "SMI-Version not valid"
	case _ERROR_BADAUTH:
		return "SMI-Authentication failed"
	case _ERROR_BADVERS:
		return "SMI-Protocol version bad"
	case _ERROR_BADPERM:
		return "SMI-Permission not valid"
	case _ERROR_PARM:
		return "Parameter not valid"
	case _ERROR_NOFILE:
		return "File not found"
	case _ERROR_FILEBIG:
		return "File too big"
	case _ERROR_FILEEMPTY:
		return "File is empty"
	case _ERROR_NOSEC:
		return "Section not found"
	case _ERROR_NOGRP:
		return "Group not found"
	case _ERROR_NOKEY:
		return "Keyword not found"
	case _ERROR_ENDFILE:
		return "End of file reached"
	case _ERROR_NOSEMA:
		return "Problem with semTake"
	case _ERROR_NOPFSET:
		return "Set not found"
	case _ERROR_NOPFUNIT:
		return "Unit not found"
	case _ERROR_SMODE:
		return "Not allowed in this SysMode"
	case _ERROR_INSTALL:
		return "Module not installable"
	case _ERROR_NOMEM:
		return "Not enough system memory"
	case _ERROR_NOTSUPP:
		return "Function not supported"
	case _ERROR_NOMODNBR:
		return "No module number assigned"
	case _ERROR_TIMEOUT1:
		return "Timeout in reply queue"
	case _ERROR_TIMEOUT2:
		return "No answer from SW-Module"
	case _ERROR_TIMEOUT3:
		return "Timeout in reply queue"
	case _ERROR_NOMOD1:
		return "Module not found"
	case _ERROR_NOMOD2:
		return "Module does not respond"
	case _ERROR_NOMOD3:
		return "Module not present any more"
	case _ERROR_NODELTSK:
		return "Module as task not present"
	case _ERROR_NOVXWOBJ:
		return "Module in VxWorks not present"
	case _ERROR_NOOBJ:
		return "Object not available"
	case _ERROR_BADINDEX:
		return "Index in request not valid"
	case _ERROR_BADADDR:
		return "Address in request not valid"
	case _ERROR_USERID:
		return "User-Id not valid"
	case _ERROR_LISTID:
		return "List-Id not valid"
	case _ERROR_DUPUSER:
		return "User already present"
	case _ERROR_NBELEM:
		return "Too many list elements"
	case _ERROR_CBACK:
		return "Callback-Address not valid"
	case _ERROR_BADELEM:
		return "Object contains bad element"
	case _ERROR_BADNAME:
		return "Name not valid"
	case _ERROR_CARDNB:
		return "IO module number not valid"
	case _ERROR_DRVID:
		return "Driver-ID not valid"
	case _ERROR_DUPRES:
		return "Resource already present"
	case _ERROR_FULL:
		return "Resource is full"
	case _ERROR_NOLIC:
		return "No license for SW-Module"
	case _ERROR_NOLICEXP:
		return "License expired"
	case _ERROR_WRONGVERS:
		return "Bad object/library version"
	case _ERROR_NOREAD:
		return "No read permission"
	case _ERROR_NOWRITE:
		return "No write permission"
	case _ERROR_BADREAD:
		return "Read error in file"
	case _ERROR_BADWRITE:
		return "Write error in file"
	case _ERROR_BADSEEK:
		return "Search error in file"
	case _ERROR_BADCHECK:
		return "Checksum not valid"
	case _ERROR_BADVXWLD:
		return "Module not loadable under VxWorks"
	case _ERROR_BADMEMLD:
		return "Module not loadable into memory"
	case _ERROR_NOLIBREG:
		return "Library can not be registered"
	case _ERROR_EMPTY:
		return "Resource is empty"
	case _ERROR_BADMODE:
		return "Mode not valid"
	case _ERROR_BADOBJ:
		return "Object not allowed"
	case _ERROR_LOCKED:
		return "Access to object is locked"
	case _ERROR_BADIPADDR:
		return "Object in use by other client"
	case _ERROR_NOENTRY:
		return "Function ???_Init is missing"
	case _ERROR_NOREG:
		return "Module can not be registered"
	case _ERROR_BADINIT:
		return "Function ???_Init returns error"
	case _ERROR_DEVFULL:
		return "Block-Device is full"
	case _ERROR_BADCOPY:
		return "Copy error in file"
	case _ERROR_NOGLOBMEM:
		return "Not enough global memory"
	case _ERROR_NOAPPMEM:
		return "Not enough application memory"
	case _ERROR_SYSTEM1:
		return "Internal system error 1"
	case _ERROR_BADROUTE:
		return "Gateway not valid"
	case _ERROR_BADLOGIN:
		return "Login Name/Password not valid"
	case _ERROR_NOSIODEV:
		return "Serial device is missing"
	case _ERROR_BADSIODEV:
		return "Serial device not valid"
	case _ERROR_DEVINUSE:
		return "Device already in use"
	case _ERROR_DEVMISS:
		return "Device not present"
	case _ERROR_NOMODEM:
		return "Modem not present"
	case _ERROR_NOUNIT:
		return "No more unit available"
	case _ERROR_BADMODEM:
		return "Modem not valid"
	case _ERROR_BADPPP1:
		return "PPP init error"
	case _ERROR_BADPPP2:
		return "PPP establish error"
	case _ERROR_BADSLIP:
		return "SLIP init error"
	case _ERROR_BADPROTO:
		return "Protocol not valid"
	case _ERROR_LNAMEMISS:
		return "Local name is missing"
	case _ERROR_BADHOST:
		return "Host name can not be set"
	case _ERROR_NOHOST:
		return "Host does not respond"
	case _ERROR_BADGATE:
		return "Gateway name can not be set"
	case _ERROR_NOGATE:
		return "Gateway does not respond"
	case _ERROR_BADMOUNT:
		return "Mount is bad"
	case _ERROR_BADTYPE:
		return "Wrong data type"
	case _ERROR_BADUMOUNT:
		return "Unmount is bad"
	case _ERROR_NODELSYS:
		return "SYS-Module can not be deleted"
	case _ERROR_BADPATH:
		return "Path/Name not valid"
	case _ERROR_NOSYS1:
		return "File MCore is missing"
	case _ERROR_NOSYS2:
		return "File MSys is missing"
	case _ERROR_NOFTP:
		return "No FTP-Server on host"
	case _ERROR_IPINUSE:
		return "IP-Address already in use"
	case _ERROR_NOLOGIN:
		return "Login from server denied"
	case _ERROR_NOMCONF:
		return "No MConfig.ini on device"
	case _ERROR_UPDATE:
		return "Firmware update not allowed"
	case _ERROR_BADFPROG:
		return "Programming flash failed"
	case _ERROR_BADVERIFY:
		return "Verify flash failed"
	case _ERROR_BADRANGE:
		return "Range not valid"
	case _ERROR_BADCPU:
		return "Not for this CPU"
	case _ERROR_OLDADDR:
		return "Address in request is old"
	case _ERROR_NOBLOCKDEV:
		return "Block Device not present"
	case _ERROR_BADFORMAT:
		return "Error in formatting"
	case _ERROR_BADREOPEN:
		return "Error in reopen after formatting"
	case _ERROR_BADDOSINI:
		return "Error in DOS initialization"
	case _ERROR_BADBOOTSEC:
		return "Boot sector not settable"
	case _ERROR_NOFORMAT:
		return "Formatting not allowed"
	case _ERROR_OLDVERS1:
		return "Old software version MCore"
	case _ERROR_OLDVERS2:
		return "Old software version MSys"
	case _ERROR_OLDVERS3:
		return "Old software version IO-Driver"
	case _ERROR_BADLOGDEV:
		return "Log-Device not valid"
	case _ERROR_BADMEMINIT:
		return "Error in Memory-Partition init"
	case _ERROR_BADPNCINIT:
		return "Error in Panic-Handler init"
	case _ERROR_BADWDGINIT:
		return "Error Watchdog-Handler init"
	case _ERROR_NOERRTOL:
		return "SW-Module not error tolerant"
	case _ERROR_BADTCKRATE:
		return "Tick rate too low"
	case _ERROR_BADSPAWN:
		return "Error in task spawn"
	case _ERROR_WDOGON:
		return "HW-Watchdog must be off"
	case _ERROR_INVSTATE:
		return "Action in this state not allowed"
	case _ERROR_NODE:
		return "Bad network/fieldbus node"
	case _ERROR_TIMEOUT:
		return "Timeout in function call"
	case _ERROR_SWREBOOT:
		return "Software reboot"
	case _ERROR_WDGREBOOT:
		return "Watchdog reboot"
	case _ERROR_NOONLINE:
		return "Online change not allowed"
	case _ERROR_DRVMISS:
		return "Driver is missing"
	case _ERROR_BADDRV:
		return "Driver not loadable"
	case _ERROR_BADSLOT:
		return "Station or slot not valid"
	case _ERROR_BELOWMIN:
		return "Number stations below minimum"
	case _ERROR_ABOVEMAX:
		return "Number stations above maximum"
	case _ERROR_STATMISS:
		return "Minimum one station is missing"
	case _ERROR_BADSESS:
		return "Session ID not valid"
	case _ERROR_BADUSER:
		return "Login user name not valid"
	case _ERROR_NOPWORD:
		return "Login password not present"
	case _ERROR_BADPWORD:
		return "Login password not valid"
	case _ERROR_BADTIME:
		return "Login time frame not valid"
	case _ERROR_NOTASK:
		return "Task does not exist"
	case _ERROR_NODEBUG:
		return "Task can not be debugged"
	case _ERROR_NOATTACH:
		return "Task is not attached"
	case _ERROR_NOBKPT:
		return "Breakpoint not set"
	case _ERROR_NOSYM:
		return "Symbol not found"
	case _ERROR_TSKMISS:
		return "Task name is missing"
	case _ERROR_NOSSTEP:
		return "Task not on breakpoint"
	case _ERROR_HAVECBACK:
		return "Callback already attached"
	case _ERROR_NOCBACK:
		return "Callback not attached"
	case _ERROR_HAVEBKPT:
		return "Breakpoint already set"
	case _ERROR_ALLBKPT:
		return "All breakpoints in use"
	case _ERROR_VARINUSE:
		return "Variable already in use"
	case _ERROR_MBTRANS:
		return "Signalising MB transfer in error for compatibility reasons"
	case _ERROR_NOBUFF:
		return "No buffer found"
	case _ERROR_BADRESTORE:
		return "Restoring device failed"
	case _ERROR_WEAKPWR:
		return "Weak power supply on a station"
	case _ERROR_BADCONFIG:
		return "Configuration data not valid"
	case _ERROR_BADFWARE:
		return "Firmware not for this module"
	case _ERROR_DOWNGRADE:
		return "Firmware down grade error"
	case _ERROR_MODMISS:
		return "Module is old or missing"
	case _ERROR_TNETINUSE:
		return "Telnet in use, no redirect"
	case _ERROR_NOSHELL:
		return "Shell not running, no redirect"
	case _ERROR_BADSNTP:
		return "SNTP client not started"
	case _ERROR_NOAPP:
		return "Service not available in boot mode NOAPP"
	case _ERROR_NETNB:
		return "Fieldbus network number not valid"
	case _ERROR_NODEID:
		return "Fieldbus network node id, mac id not valid"
	case _ERROR_DSIZE:
		return "Data size not valid, exceeds max size"
	case _ERROR_CMDTAG:
		return "Command tag not valid"
	case _ERROR_CMDARG:
		return "Command argument(s) not valid"
	case _ERROR_QUEUEFULL:
		return "Command/Event Queue overrun"
	case _ERROR_IOBUSFAIL:
		return "I/O Bus access fails (bad station)"
	case _ERROR_SLCLOGFULL:
		return "The logbook of the SLC reached a critical level"
	case _ERROR_BROKENWIRE:
		return "SERCOS: Broken wire"
	case _ERROR_DRIVEMISS:
		return "SERCOS: Missing drive"
	case _ERROR_STARTUPTMO:
		return "SERCOS: Timeout while waiting for startup phase"
	case _ERROR_DEAD:
		return "Alive check failed. Module is dead"
	case _ERROR_FALLBACK:
		return "Fallback device was used for booting"
	case _ERROR_MCONFIGBAK:
		return "MConfig.ini not found, used .bak instead"
	case _ERROR_CONTIGFILE:
		return "Error creating contiguous file"
	case _ERROR_EXCEPTION:
		return "Exception signal occurred"
	case _ERROR_OFFLINE:
		return "No connection to device"
	case _ERROR_INITVAL:
		return "Init value error"
	case _ERROR_CHECKVAL:
		return "Check value error"
	case _ERROR_DRIVE:
		return "Drive specific error"
	case _ERROR_HARDWARE:
		return "Hardware error"
	case _ERROR_COMMUNIC:
		return "Communication error"
	case _ERROR_NOSETVAL:
		return "Cyclic setvalue monitoring error"
	case _ERROR_CYCLETIME:
		return "Cycle time monitoring error"
	case _ERROR_COLDCLIMATE:
		return "A module does not support cold climate"
	case _ERROR_ACCDENIED:
		return "Access denied"
	case _ERROR_BADUSERLVL:
		return "Weak user level (deprecated: replaced by M_E_ACCDENIED)"
	case _ERROR_XERR:
		return "VHD_LST_XERR code information"
	case _ERROR_REDULOCK:
		return "Write access allowed only on primary CPU"
	case _ERROR_RICONN:
		return "RI offline"
	case _ERROR_RINETREDU:
		return "RI not redundant"
	case _ERROR_LENGTH:
		return "Data length invalid"
	case _ERROR_APPNAME:
		return "Application name mismatch"
	case _ERROR_BLOCKNB:
		return "Number of blocks different"
	case _ERROR_BLOCKDIFF:
		return "Memory blocks different"
	case _ERROR_DEVREDU:
		return "Device not redundant"
	case _ERROR_DEVNETREDU:
		return "Device network not redundant"
	case _ERROR_APPERROR1:
		return "Application error1"
	case _ERROR_APPERROR2:
		return "Application error2"
	case _ERROR_APPERROR3:
		return "Application error3"
	case _ERROR_APPFATAL:
		return "Application fatal error"
	case _ERROR_APPDONE:
		return "Application done error"
	case _ERROR_APPOVLOAD:
		return "Application overload"
	case _ERROR_CYCLEOVLOAD:
		return "Cycle overload"
	case _ERROR_BCHCMD:
		return "BCH command error"
	case _ERROR_PRICONFLICT:
		return "Primary assignment conflict"
	case _ERROR_SWITCHOVER:
		return "Switchover command"
	case _ERROR_APPCS:
		return "Application checksum mismatch"
	case _ERROR_STSWITCHOVER:
		return "Self test switchover command"
	case _ERROR_STRUN:
		return "Self test run"
	case _ERROR_STERROR:
		return "Self test error"
	case _ERROR_PROGRAM:
		return "Programming error"
	case _ERROR_STID:
		return "Self test ID not valid"
	case _ERROR_STRAM:
		return "Self test RAM error"
	case _ERROR_STPLC:
		return "Self test PLC error"
	case _ERROR_STCOM:
		return "Self test COM error"
	case _ERROR_CRYPT_METADATA_METHOD:
		return "Cipher method not available"
	case _ERROR_CRYPT_METADATA_HASH:
		return "Couldn't create cipher hash"
	case _ERROR_CRYPT_CRYPT_EXISTS:
		return "No encryption layer found"
	case _ERROR_CRYPT_CRYPT_MASTERKEY:
		return "No masterkey available"
	case _ERROR_CRYPT_CRYPT_BOOTDEV:
		return "Boot partition can not be encrypted"
	case _ERROR_CRYPT_ALGORITHM_LOAD_LIB:
		return "Loading cipher library failed"
	case _ERROR_CRYPT_ALGORITHM_LIB_VERSION:
		return "Cipher library version"
	case _ERROR_CRYPT_ALGORITHM_INIT_LIB:
		return "Init cipher library failed"
	case _ERROR_CRYPT_ALGORITHM_KEYLENGTH:
		return "Cipher strength not valid"
	case _ERROR_CRYPT_ALGORITHM_INIT:
		return "Init cipher algorithm failed"
	case _ERROR_CRYPT_ALGORITHM_CRYPT:
		return "Error encrypting data"
	case _ERROR_CRYPT_ALGORITHM_DECRYPT:
		return "Error decrypting data"
	case _ERROR_CRYPT_KEY_RECOVERYFILE:
		return "Recover key file not found"
	case _ERROR_CRYPT_KEY_RECOVERYFILECRC:
		return "Recovery key file crc error"
	case _ERROR_CRYPT_KEY_RECOVERYPASSWORD:
		return "Recovery password wrong"
	case _ERROR_CRYPT_NOPROGRESS:
		return "No Progress available"
	case _ERROR_CRYPT_UNKNOWN:
		return "Detailed reason unknown"
	case _ERROR_BADSNTPSERVER:
		return "Configuration of SNTP server not valid"
	case _ERROR_NOONLINECFG:
		return "Online re-config not supported"
	case _ERROR_BADCORECATEG:
		return "Configuration of CoreCategory not valid"
	case _ERROR_PWDEXPIRED:
		return "Password expired"
	case _ERROR_BADLOGINLOCKED:
		return "User is disabled"
	case _ERROR_BUSY:
		return "Resource is currently busy"
	case _ERROR_TASK_WDOG:
		return "Error creating watchdog"
	case _ERROR_TASK_SEM:
		return "Error creating/deleting semaphore"
	case _ERROR_TASK_PRIO:
		return "Priority out of range (has to be between 1 and 255)"
	case _ERROR_TASK_OFFSET:
		return "Offset does not match tick/snyc rate or is greater than cycle time"
	case _ERROR_TASK_CYCDIFF:
		return "Next possible cycle time does differ more than 25% from configured cycle time"
	case _ERROR_TASK_CYCTIME:
		return "Cycle time is too small"
	case _ERROR_TASK_ATTACH:
		return "Error attaching aux/sync"
	case _ERROR_TASK_AUXOFF:
		return "Auxiliary clock is turned off"
	case _ERROR_TASK_SOURCE:
		return "Unknown task trigger source"
	case _ERROR_TASK_NOMAINFUNC:
		return "User main function is NULL"
	case _ERROR_TASK_NOPTP:
		return "PTP synchronization not activated"
	case _ERROR_TASK_CORECATEG:
		return "Core category not valid"
	case _ERROR_TASK_RUNNING:
		return "Task has not been stopped"
	case _ERROR_TASK_CHANGEATTR:
		return "Selected attributes must not be changed"
	case _ERROR_RTC:
		return "Error init RTC chip"
	case _ERROR_ETHER:
		return "Error init Ethernet controller"
	case _ERROR_IP_ADDR:
		return "IP-Address not valid"
	case _ERROR_TICKRATE:
		return "Tick rate not valid"
	case _ERROR_CPUCLOCK:
		return "CPU clock not valid"
	case _ERROR_NVMAKE:
		return "NV-Ram create failed"
	case _ERROR_NVMISS:
		return "NV-Ram not present"
	case _ERROR_NVEMPTY:
		return "NV-Ram is empty"
	case _ERROR_NVNODOS:
		return "NV-Ram no DOS file system"
	case _ERROR_RDMAKE:
		return "Ram-Disk create failed"
	case _ERROR_BATTLOW:
		return "Battery voltage too low"
	case _ERROR_SIO:
		return "Error init SIO driver"
	case _ERROR_DUPIP:
		return "Duplicate ip address"
	case _ERROR_EHDOFLOW:
		return "Too much EHD entries in MCore"
	case _ERROR_BOOTPANIC:
		return "Panic situation occurred during boot"
	case _ERROR_NVRAMSIG:
		return "Invalid nvram signature"
	case _ERROR_NATFAIL:
		return "NAT initialization failed"
	case _ERROR_SMART:
		return "SMART self-test on CFC failed"
	case _ERROR_FILESYSTEM:
		return "Error in DOS file system"
	case _ERROR_SHORTKEY:
		return "MConfig.ini keyword too short"
	case _ERROR_MSYSINVALID:
		return "MSys is invalid (checksum, not found, ...)"
	case _ERROR_SNTPSYNC:
		return "SNTP synchronization failure"
	case _ERROR_BADSIOHANDSHAKE:
		return "Hardware handshake is not permitted (probably device is used as console)"
	case _ERROR_ALLREADYEXIST:
		return "Object already exist"
	case _ERROR_NOOBJECT:
		return "Object not found"
	case _ERROR_SQLERROR:
		return "SQL error occurred"
	case _ERROR_NOAHD:
		return "SW-Module AHD is not installed"
	case _ERROR_NOSESSAV:
		return "No session available"
	default:
		return "Unknown error code"
	}
}

// Error sources.
const (
	_SOURCE_NO       uint32 = 0x00000000 // Unknown source
	_SOURCE_SVI      uint32 = 0x00010000 // SVI Functions
	_SOURCE_SMI      uint32 = 0x00020000 // SMI Functions
	_SOURCE_RES      uint32 = 0x00030000 // Resource Handler
	_SOURCE_MIO      uint32 = 0x00040000 // IO Handler
	_SOURCE_VHD      uint32 = 0x00050000 // Vis Handler
	_SOURCE_INF      uint32 = 0x00060000 // Info Handler
	_SOURCE_PLC      uint32 = 0x00070000 // PLC Runtime system
	_SOURCE_MOD      uint32 = 0x00080000 // Module Handler
	_SOURCE_CAN      uint32 = 0x00090000 // CAN Handler
	_SOURCE_PF       uint32 = 0x000A0000 // Profile Functions
	_SOURCE_SYS      uint32 = 0x000F0000 // System MSys
	_SOURCE_CORE     uint32 = 0x00100000 // System MCore
	_SOURCE_EHD      uint32 = 0x00110000 // Error Handler
	_SOURCE_PB       uint32 = 0x00120000 // Profibus Handler
	_SOURCE_DBG      uint32 = 0x00130000 // Debug Handler
	_SOURCE_DN       uint32 = 0x00140000 // DeviceNet Handler
	_SOURCE_RFS      uint32 = 0x00150000 // Remote File Server
	_SOURCE_SLC      uint32 = 0x00160000 // SLC Server
	_SOURCE_DMW      uint32 = 0x00170000 // Drive Middleware
	_SOURCE_SEM201   uint32 = 0x00180000 // SERCOS driver
	_SOURCE_UFB      uint32 = 0x00190000 // Unified fieldbus components
	_SOURCE_PN       uint32 = 0x001A0000 // Profinet
	_SOURCE_EC       uint32 = 0x001B0000 // EtherCAT
	_SOURCE_BCR      uint32 = 0x001C0000 // BCR2xx
	_SOURCE_ST       uint32 = 0x001D0000 // Self Test Module
	_SOURCE_C_TDLL   uint32 = 0x01010000 // TMANw32-DLL
	_SOURCE_C_PLCCOM uint32 = 0x01020000 // PLCCOM-DLL
	_SOURCE_C_TCONF  uint32 = 0x01030000 // TargetManager EXE
	_SOURCE_C_PLCHWM uint32 = 0x01040000 // PLCHWM-DLL
	_SOURCE_C_TVIEW  uint32 = 0x01050000 // TMAN-View
	_SOURCE_C_MMAN   uint32 = 0x01060000 // M-Manager
	_SOURCE_C_MPLC   uint32 = 0x01070000 // M-PLC
	_SOURCE_C_MIF    uint32 = 0x01080000 // M-Interface
	_SOURCE__AHD     uint32 = 0x02010000 // AHD System SW Module
	_SOURCE__LOGGER  uint32 = 0x02020000 // Logger System SW Module
	_SOURCE_EVT      uint32 = 0x02030000 // Event Logger
	_SOURCE_ETCP     uint32 = 0x02040000 // 60870-5
	_SOURCE_DNP3     uint32 = 0x02050000 // DNP3
	_SOURCE_ATEC     uint32 = 0x02060000 // ATEC
	_SOURCE_M1C      uint32 = 0x81100000 // M1 Core
)

// Error codes.
const (
	_OK                                uint32 = 0          // O.K., no error
	_ERROR_INPROGRESS                  uint32 = 1          // Still in progress, no error
	_ERROR_CANCELED                    uint32 = 2          // Canceled by user, no error
	_ERROR_FAILED                      uint32 = 0x80000100 // Unspecified error
	_ERROR_BADPROG                     uint32 = 0x80000101 // SMI-Program not valid
	_ERROR_BADPROC                     uint32 = 0x80000102 // SMI-Procedure not valid
	_ERROR_BADARGS                     uint32 = 0x80000103 // SMI-Arguments not valid
	_ERROR_BADRPC                      uint32 = 0x80000104 // SMI-Version not valid
	_ERROR_BADAUTH                     uint32 = 0x80000105 // SMI-Authentication failed
	_ERROR_BADVERS                     uint32 = 0x80000106 // SMI-Protocol version bad
	_ERROR_BADPERM                     uint32 = 0x80000107 // SMI-Permission not valid
	_ERROR_PARM                        uint32 = 0x80000110 // Parameter not valid
	_ERROR_NOFILE                      uint32 = 0x80000111 // File not found
	_ERROR_FILEBIG                     uint32 = 0x80000112 // File too big
	_ERROR_FILEEMPTY                   uint32 = 0x80000113 // File is empty
	_ERROR_NOSEC                       uint32 = 0x80000114 // Section not found
	_ERROR_NOGRP                       uint32 = 0x80000115 // Group not found
	_ERROR_NOKEY                       uint32 = 0x80000116 // Keyword not found
	_ERROR_ENDFILE                     uint32 = 0x80000117 // End of file reached
	_ERROR_NOSEMA                      uint32 = 0x80000118 // Problem with semTake
	_ERROR_NOPFSET                     uint32 = 0x80000119 // Set not found
	_ERROR_NOPFUNIT                    uint32 = 0x80000120 // Unit not found
	_ERROR_SMODE                       uint32 = 0x80000121 // Not allowed in this SysMode
	_ERROR_INSTALL                     uint32 = 0x80000122 // Module not installable
	_ERROR_NOMEM                       uint32 = 0x80000123 // Not enough system memory
	_ERROR_NOTSUPP                     uint32 = 0x80000124 // Function not supported
	_ERROR_NOMODNBR                    uint32 = 0x80000125 // No module number assigned
	_ERROR_TIMEOUT1                    uint32 = 0x80000126 // Timeout in reply queue
	_ERROR_TIMEOUT2                    uint32 = 0x80000127 // No answer from SW-Module
	_ERROR_TIMEOUT3                    uint32 = 0x80000128 // Timeout in reply queue
	_ERROR_NOMOD1                      uint32 = 0x80000129 // Module not found
	_ERROR_NOMOD2                      uint32 = 0x8000012A // Module does not respond
	_ERROR_NOMOD3                      uint32 = 0x8000012B // Module not present any more
	_ERROR_NODELTSK                    uint32 = 0x8000012C // Module as task not present
	_ERROR_NOVXWOBJ                    uint32 = 0x8000012D // Module in VxWorks not present
	_ERROR_NOOBJ                       uint32 = 0x8000012E // Object not available
	_ERROR_BADINDEX                    uint32 = 0x8000012F // Index in request not valid
	_ERROR_BADADDR                     uint32 = 0x80000130 // Address in request not valid
	_ERROR_USERID                      uint32 = 0x80000131 // User-Id not valid
	_ERROR_LISTID                      uint32 = 0x80000132 // List-Id not valid
	_ERROR_DUPUSER                     uint32 = 0x80000133 // User already present
	_ERROR_NBELEM                      uint32 = 0x80000134 // Too many list elements
	_ERROR_CBACK                       uint32 = 0x80000135 // Callback-Address not valid
	_ERROR_BADELEM                     uint32 = 0x80000136 // Object contains bad element
	_ERROR_BADNAME                     uint32 = 0x80000137 // Name not valid
	_ERROR_CARDNB                      uint32 = 0x80000138 // IO module number not valid
	_ERROR_DRVID                       uint32 = 0x80000139 // Driver-ID not valid
	_ERROR_DUPRES                      uint32 = 0x8000013A // Resource already present
	_ERROR_FULL                        uint32 = 0x8000013B // Resource is full
	_ERROR_NOLIC                       uint32 = 0x8000013C // No license for SW-Module
	_ERROR_NOLICEXP                    uint32 = 0x8000013D // License expired
	_ERROR_WRONGVERS                   uint32 = 0x8000013E // Bad object/library version
	_ERROR_NOREAD                      uint32 = 0x80000140 // No read permission
	_ERROR_NOWRITE                     uint32 = 0x80000141 // No write permission
	_ERROR_BADREAD                     uint32 = 0x80000142 // Read error in file
	_ERROR_BADWRITE                    uint32 = 0x80000143 // Write error in file
	_ERROR_BADSEEK                     uint32 = 0x80000144 // Search error in file
	_ERROR_BADCHECK                    uint32 = 0x80000145 // Checksum not valid
	_ERROR_BADVXWLD                    uint32 = 0x80000146 // Module not loadable under VxWorks
	_ERROR_BADMEMLD                    uint32 = 0x80000147 // Module not loadable into memory
	_ERROR_NOLIBREG                    uint32 = 0x80000148 // Library can not be registered
	_ERROR_EMPTY                       uint32 = 0x80000149 // Resource is empty
	_ERROR_BADMODE                     uint32 = 0x8000014A // Mode not valid
	_ERROR_BADOBJ                      uint32 = 0x8000014B // Object not allowed
	_ERROR_LOCKED                      uint32 = 0x8000014C // Access to object is locked
	_ERROR_BADIPADDR                   uint32 = 0x8000014D // Object in use by other client
	_ERROR_NOENTRY                     uint32 = 0x8000014E // Function ???_Init is missing
	_ERROR_NOREG                       uint32 = 0x8000014F // Module can not be registered
	_ERROR_BADINIT                     uint32 = 0x80000150 // Function ???_Init returns error
	_ERROR_DEVFULL                     uint32 = 0x80000151 // Block-Device is full
	_ERROR_BADCOPY                     uint32 = 0x80000152 // Copy error in file
	_ERROR_NOGLOBMEM                   uint32 = 0x80000153 // Not enough global memory
	_ERROR_NOAPPMEM                    uint32 = 0x80000154 // Not enough application memory
	_ERROR_SYSTEM1                     uint32 = 0x80000156 // Internal system error 1
	_ERROR_BADROUTE                    uint32 = 0x80000157 // Gateway not valid
	_ERROR_BADLOGIN                    uint32 = 0x80000158 // Login Name/Password not valid
	_ERROR_NOSIODEV                    uint32 = 0x80000159 // Serial device is missing
	_ERROR_BADSIODEV                   uint32 = 0x8000015A // Serial device not valid
	_ERROR_DEVINUSE                    uint32 = 0x8000015B // Device already in use
	_ERROR_DEVMISS                     uint32 = 0x8000015C // Device not present
	_ERROR_NOMODEM                     uint32 = 0x8000015D // Modem not present
	_ERROR_NOUNIT                      uint32 = 0x8000015E // No more unit available
	_ERROR_BADMODEM                    uint32 = 0x8000015F // Modem not valid
	_ERROR_BADPPP1                     uint32 = 0x80000160 // PPP init error
	_ERROR_BADPPP2                     uint32 = 0x80000161 // PPP establish error
	_ERROR_BADSLIP                     uint32 = 0x80000162 // SLIP init error
	_ERROR_BADPROTO                    uint32 = 0x80000163 // Protocol not valid
	_ERROR_LNAMEMISS                   uint32 = 0x80000164 // Local name is missing
	_ERROR_BADHOST                     uint32 = 0x80000165 // Host name can not be set
	_ERROR_NOHOST                      uint32 = 0x80000166 // Host does not respond
	_ERROR_BADGATE                     uint32 = 0x80000167 // Gateway name can not be set
	_ERROR_NOGATE                      uint32 = 0x80000168 // Gateway does not respond
	_ERROR_BADMOUNT                    uint32 = 0x80000169 // Mount is bad
	_ERROR_BADTYPE                     uint32 = 0x8000016A // Wrong data type
	_ERROR_BADUMOUNT                   uint32 = 0x80000170 // Unmount is bad
	_ERROR_NODELSYS                    uint32 = 0x80000171 // SYS-Module can not be deleted
	_ERROR_BADPATH                     uint32 = 0x80000172 // Path/Name not valid
	_ERROR_NOSYS1                      uint32 = 0x80000173 // File MCore is missing
	_ERROR_NOSYS2                      uint32 = 0x80000174 // File MSys is missing
	_ERROR_NOFTP                       uint32 = 0x80000175 // No FTP-Server on host
	_ERROR_IPINUSE                     uint32 = 0x80000176 // IP-Address already in use
	_ERROR_NOLOGIN                     uint32 = 0x80000177 // Login from server denied
	_ERROR_NOMCONF                     uint32 = 0x80000178 // No MConfig.ini on device
	_ERROR_UPDATE                      uint32 = 0x80000179 // Firmware update not allowed
	_ERROR_BADFPROG                    uint32 = 0x8000017A // Programming flash failed
	_ERROR_BADVERIFY                   uint32 = 0x8000017B // Verify flash failed
	_ERROR_BADRANGE                    uint32 = 0x8000017C // Range not valid
	_ERROR_BADCPU                      uint32 = 0x8000017D // Not for this CPU
	_ERROR_OLDADDR                     uint32 = 0x8000017E // Address in request is old
	_ERROR_NOBLOCKDEV                  uint32 = 0x8000017F // Block Device not present
	_ERROR_BADFORMAT                   uint32 = 0x80000180 // Error in formatting
	_ERROR_BADREOPEN                   uint32 = 0x80000181 // Error in reopen after formatting
	_ERROR_BADDOSINI                   uint32 = 0x80000182 // Error in DOS initialization
	_ERROR_BADBOOTSEC                  uint32 = 0x80000183 // Boot sector not settable
	_ERROR_NOFORMAT                    uint32 = 0x80000184 // Formatting not allowed
	_ERROR_OLDVERS1                    uint32 = 0x80000185 // Old software version MCore
	_ERROR_OLDVERS2                    uint32 = 0x80000186 // Old software version MSys
	_ERROR_OLDVERS3                    uint32 = 0x80000187 // Old software version IO-Driver
	_ERROR_BADLOGDEV                   uint32 = 0x80000188 // Log-Device not valid
	_ERROR_BADMEMINIT                  uint32 = 0x80000189 // Error in Memory-Partition init
	_ERROR_BADPNCINIT                  uint32 = 0x8000018A // Error in Panic-Handler init
	_ERROR_BADWDGINIT                  uint32 = 0x8000018B // Error Watchdog-Handler init
	_ERROR_NOERRTOL                    uint32 = 0x8000018C // SW-Module not error tolerant
	_ERROR_BADTCKRATE                  uint32 = 0x8000018D // Tick rate too low
	_ERROR_BADSPAWN                    uint32 = 0x8000018E // Error in task spawn
	_ERROR_WDOGON                      uint32 = 0x8000018F // HW-Watchdog must be off
	_ERROR_INVSTATE                    uint32 = 0x80000190 // Action in this state not allowed
	_ERROR_NODE                        uint32 = 0x80000191 // Bad network/fieldbus node
	_ERROR_TIMEOUT                     uint32 = 0x80000192 // Timeout in function call
	_ERROR_SWREBOOT                    uint32 = 0x80000193 // Software reboot
	_ERROR_WDGREBOOT                   uint32 = 0x80000194 // Watchdog reboot
	_ERROR_NOONLINE                    uint32 = 0x80000195 // Online change not allowed
	_ERROR_DRVMISS                     uint32 = 0x80000196 // Driver is missing
	_ERROR_BADDRV                      uint32 = 0x80000197 // Driver not loadable
	_ERROR_BADSLOT                     uint32 = 0x80000198 // Station or slot not valid
	_ERROR_BELOWMIN                    uint32 = 0x80000199 // Number stations below minimum
	_ERROR_ABOVEMAX                    uint32 = 0x8000019A // Number stations above maximum
	_ERROR_STATMISS                    uint32 = 0x8000019B // Minimum one station is missing
	_ERROR_BADSESS                     uint32 = 0x8000019C // Session ID not valid
	_ERROR_BADUSER                     uint32 = 0x8000019D // Login user name not valid
	_ERROR_NOPWORD                     uint32 = 0x8000019E // Login password not present
	_ERROR_BADPWORD                    uint32 = 0x8000019F // Login password not valid
	_ERROR_BADTIME                     uint32 = 0x800001A0 // Login time frame not valid
	_ERROR_NOTASK                      uint32 = 0x800001A1 // Task does not exist
	_ERROR_NODEBUG                     uint32 = 0x800001A2 // Task can not be debugged
	_ERROR_NOATTACH                    uint32 = 0x800001A3 // Task is not attached
	_ERROR_NOBKPT                      uint32 = 0x800001A4 // Breakpoint not set
	_ERROR_NOSYM                       uint32 = 0x800001A5 // Symbol not found
	_ERROR_TSKMISS                     uint32 = 0x800001A6 // Task name is missing
	_ERROR_NOSSTEP                     uint32 = 0x800001A7 // Task not on breakpoint
	_ERROR_HAVECBACK                   uint32 = 0x800001A8 // Callback already attached
	_ERROR_NOCBACK                     uint32 = 0x800001A9 // Callback not attached
	_ERROR_HAVEBKPT                    uint32 = 0x800001AA // Breakpoint already set
	_ERROR_ALLBKPT                     uint32 = 0x800001AB // All breakpoints in use
	_ERROR_VARINUSE                    uint32 = 0x800001AC // Variable already in use
	_ERROR_MBTRANS                     uint32 = 0x800001AD // Signalising MB transfer in error for compatibility reasons
	_ERROR_NOBUFF                      uint32 = 0x800001AE // No buffer found
	_ERROR_BADRESTORE                  uint32 = 0x800001AF // Restoring device failed
	_ERROR_WEAKPWR                     uint32 = 0x800001B0 // Weak power supply on a station
	_ERROR_BADCONFIG                   uint32 = 0x800001B1 // Configuration data not valid
	_ERROR_BADFWARE                    uint32 = 0x800001B2 // Firmware not for this module
	_ERROR_DOWNGRADE                   uint32 = 0x800001B3 // Firmware down grade error
	_ERROR_MODMISS                     uint32 = 0x800001B4 // Module is old or missing
	_ERROR_TNETINUSE                   uint32 = 0x800001B5 // Telnet in use, no redirect
	_ERROR_NOSHELL                     uint32 = 0x800001B6 // Shell not running, no redirect
	_ERROR_BADSNTP                     uint32 = 0x800001B7 // SNTP client not started
	_ERROR_NOAPP                       uint32 = 0x800001B8 // Service not available in boot mode NOAPP
	_ERROR_NETNB                       uint32 = 0x800001C0 // Fieldbus network number not valid
	_ERROR_NODEID                      uint32 = 0x800001C1 // Fieldbus network node id, mac id not valid
	_ERROR_DSIZE                       uint32 = 0x800001C2 // Data size not valid, exceeds max size
	_ERROR_CMDTAG                      uint32 = 0x800001C3 // Command tag not valid
	_ERROR_CMDARG                      uint32 = 0x800001C4 // Command argument(s) not valid
	_ERROR_QUEUEFULL                   uint32 = 0x800001C5 // Command/Event Queue overrun
	_ERROR_IOBUSFAIL                   uint32 = 0x800001C6 // I/O Bus access fails (bad station)
	_ERROR_SLCLOGFULL                  uint32 = 0x800001C7 // The logbook of the SLC reached a critical level
	_ERROR_BROKENWIRE                  uint32 = 0x800001C8 // SERCOS: Broken wire
	_ERROR_DRIVEMISS                   uint32 = 0x800001C9 // SERCOS: Missing drive
	_ERROR_STARTUPTMO                  uint32 = 0x800001CA // SERCOS: Timeout while waiting for startup phase
	_ERROR_DEAD                        uint32 = 0x800001CB // Alive check failed. Module is dead
	_ERROR_FALLBACK                    uint32 = 0x800001D0 // Fallback device was used for booting
	_ERROR_MCONFIGBAK                  uint32 = 0x800001D1 // MConfig.ini not found, used .bak instead
	_ERROR_CONTIGFILE                  uint32 = 0x800001D2 // Error creating contiguous file
	_ERROR_EXCEPTION                   uint32 = 0x800001D3 // Exception signal occurred
	_ERROR_OFFLINE                     uint32 = 0x800001E0 // No connection to device
	_ERROR_INITVAL                     uint32 = 0x800001E1 // Init value error
	_ERROR_CHECKVAL                    uint32 = 0x800001E2 // Check value error
	_ERROR_DRIVE                       uint32 = 0x800001E3 // Drive specific error
	_ERROR_HARDWARE                    uint32 = 0x800001E4 // Hardware error
	_ERROR_COMMUNIC                    uint32 = 0x800001E5 // Communication error
	_ERROR_NOSETVAL                    uint32 = 0x800001E6 // Cyclic setvalue monitoring error
	_ERROR_CYCLETIME                   uint32 = 0x800001E7 // Cycle time monitoring error
	_ERROR_COLDCLIMATE                 uint32 = 0x800001E8 // A module does not support cold climate
	_ERROR_ACCDENIED                   uint32 = 0x800001E9 // Access denied
	_ERROR_BADUSERLVL                  uint32 = 0x800001EA // Weak user level (deprecated: replaced by M_E_ACCDENIED)
	_ERROR_XERR                        uint32 = 0x800001EB // VHD_LST_XERR code information
	_ERROR_REDULOCK                    uint32 = 0x800001F0 // Write access allowed only on primary CPU
	_ERROR_RICONN                      uint32 = 0x800001F1 // RI offline
	_ERROR_RINETREDU                   uint32 = 0x800001F2 // RI not redundant
	_ERROR_LENGTH                      uint32 = 0x800001F3 // Data length invalid
	_ERROR_APPNAME                     uint32 = 0x800001F4 // Application name mismatch
	_ERROR_BLOCKNB                     uint32 = 0x800001F5 // Number of blocks different
	_ERROR_BLOCKDIFF                   uint32 = 0x800001F6 // Memory blocks different
	_ERROR_DEVREDU                     uint32 = 0x800001F7 // Device not redundant
	_ERROR_DEVNETREDU                  uint32 = 0x800001F8 // Device network not redundant
	_ERROR_APPERROR1                   uint32 = 0x800001F9 // Application error1
	_ERROR_APPERROR2                   uint32 = 0x800001FA // Application error2
	_ERROR_APPERROR3                   uint32 = 0x800001FB // Application error3
	_ERROR_APPFATAL                    uint32 = 0x800001FC // Application fatal error
	_ERROR_APPDONE                     uint32 = 0x800001FD // Application done error
	_ERROR_APPOVLOAD                   uint32 = 0x800001FE // Application overload
	_ERROR_CYCLEOVLOAD                 uint32 = 0x800001FF // Cycle overload
	_ERROR_BCHCMD                      uint32 = 0x80000200 // BCH command error
	_ERROR_PRICONFLICT                 uint32 = 0x80000201 // Primary assignment conflict
	_ERROR_SWITCHOVER                  uint32 = 0x80000202 // Switchover command
	_ERROR_APPCS                       uint32 = 0x80000203 // Application checksum mismatch
	_ERROR_STSWITCHOVER                uint32 = 0x80000204 // Self test switchover command
	_ERROR_STRUN                       uint32 = 0x80000205 // Self test run
	_ERROR_STERROR                     uint32 = 0x80000206 // Self test error
	_ERROR_PROGRAM                     uint32 = 0x80000207 // Programming error
	_ERROR_STID                        uint32 = 0x80000210 // Self test ID not valid
	_ERROR_STRAM                       uint32 = 0x80000211 // Self test RAM error
	_ERROR_STPLC                       uint32 = 0x80000212 // Self test PLC error
	_ERROR_STCOM                       uint32 = 0x80000213 // Self test COM error
	_ERROR_CRYPT_METADATA_METHOD       uint32 = 0x80000214 // Cipher method not available
	_ERROR_CRYPT_METADATA_HASH         uint32 = 0x80000215 // Couldn't create cipher hash
	_ERROR_CRYPT_CRYPT_EXISTS          uint32 = 0x80000216 // No encryption layer found
	_ERROR_CRYPT_CRYPT_MASTERKEY       uint32 = 0x80000217 // No masterkey available
	_ERROR_CRYPT_CRYPT_BOOTDEV         uint32 = 0x80000218 // Boot partition can not be encrypted
	_ERROR_CRYPT_ALGORITHM_LOAD_LIB    uint32 = 0x80000219 // Loading cipher library failed
	_ERROR_CRYPT_ALGORITHM_LIB_VERSION uint32 = 0x80000221 // Cipher library version
	_ERROR_CRYPT_ALGORITHM_INIT_LIB    uint32 = 0x80000222 // Init cipher library failed
	_ERROR_CRYPT_ALGORITHM_KEYLENGTH   uint32 = 0x80000223 // Cipher strength not valid
	_ERROR_CRYPT_ALGORITHM_INIT        uint32 = 0x80000224 // Init cipher algorithm failed
	_ERROR_CRYPT_ALGORITHM_CRYPT       uint32 = 0x80000225 // Error encrypting data
	_ERROR_CRYPT_ALGORITHM_DECRYPT     uint32 = 0x80000226 // Error decrypting data
	_ERROR_CRYPT_KEY_RECOVERYFILE      uint32 = 0x80000227 // Recover key file not found
	_ERROR_CRYPT_KEY_RECOVERYFILECRC   uint32 = 0x80000228 // Recovery key file crc error
	_ERROR_CRYPT_KEY_RECOVERYPASSWORD  uint32 = 0x80000229 // Recovery password wrong
	_ERROR_CRYPT_NOPROGRESS            uint32 = 0x80000230 // No Progress available
	_ERROR_CRYPT_UNKNOWN               uint32 = 0x80000231 // Detailed reason unknown
	_ERROR_BADSNTPSERVER               uint32 = 0x80000232 // Configuration of SNTP server not valid
	_ERROR_NOONLINECFG                 uint32 = 0x80000233 // Online re-config not supported
	_ERROR_BADCORECATEG                uint32 = 0x80000234 // Configuration of CoreCategory not valid
	_ERROR_PWDEXPIRED                  uint32 = 0x80000235 // Password expired
	_ERROR_BADLOGINLOCKED              uint32 = 0x80000236 // User is disabled
	_ERROR_BUSY                        uint32 = 0x80000237 // Resource is currently busy
	_ERROR_TASK_WDOG                   uint32 = 0x80000240 // Error creating watchdog
	_ERROR_TASK_SEM                    uint32 = 0x80000241 // Error creating/deleting semaphore
	_ERROR_TASK_PRIO                   uint32 = 0x80000242 // Priority out of range (has to be between 1 and 255)
	_ERROR_TASK_OFFSET                 uint32 = 0x80000243 // Offset does not match tick/snyc rate or is greater than cycle time
	_ERROR_TASK_CYCDIFF                uint32 = 0x80000244 // Next possible cycle time does differ more than 25% from configured cycle time
	_ERROR_TASK_CYCTIME                uint32 = 0x80000245 // Cycle time is too small
	_ERROR_TASK_ATTACH                 uint32 = 0x80000246 // Error attaching aux/sync
	_ERROR_TASK_AUXOFF                 uint32 = 0x80000247 // Auxiliary clock is turned off
	_ERROR_TASK_SOURCE                 uint32 = 0x80000248 // Unknown task trigger source
	_ERROR_TASK_NOMAINFUNC             uint32 = 0x80000249 // User main function is NULL
	_ERROR_TASK_NOPTP                  uint32 = 0x8000024A // PTP synchronization not activated
	_ERROR_TASK_CORECATEG              uint32 = 0x8000024B // Core category not valid
	_ERROR_TASK_RUNNING                uint32 = 0x8000024C // Task has not been stopped
	_ERROR_TASK_CHANGEATTR             uint32 = 0x8000024D // Selected attributes must not be changed
	_ERROR_RTC                         uint32 = 0x80000300 // Error init RTC chip
	_ERROR_ETHER                       uint32 = 0x80000301 // Error init Ethernet controller
	_ERROR_IP_ADDR                     uint32 = 0x80000302 // IP-Address not valid
	_ERROR_TICKRATE                    uint32 = 0x80000303 // Tick rate not valid
	_ERROR_CPUCLOCK                    uint32 = 0x80000304 // CPU clock not valid
	_ERROR_NVMAKE                      uint32 = 0x80000305 // NV-Ram create failed
	_ERROR_NVMISS                      uint32 = 0x80000306 // NV-Ram not present
	_ERROR_NVEMPTY                     uint32 = 0x80000307 // NV-Ram is empty
	_ERROR_NVNODOS                     uint32 = 0x80000308 // NV-Ram no DOS file system
	_ERROR_RDMAKE                      uint32 = 0x80000309 // Ram-Disk create failed
	_ERROR_BATTLOW                     uint32 = 0x8000030A // Battery voltage too low
	_ERROR_SIO                         uint32 = 0x8000030D // Error init SIO driver
	_ERROR_DUPIP                       uint32 = 0x8000030E // Duplicate ip address
	_ERROR_EHDOFLOW                    uint32 = 0x8000030F // Too much EHD entries in MCore
	_ERROR_BOOTPANIC                   uint32 = 0x80000310 // Panic situation occurred during boot
	_ERROR_NVRAMSIG                    uint32 = 0x80000311 // Invalid nvram signature
	_ERROR_NATFAIL                     uint32 = 0x80000312 // NAT initialization failed
	_ERROR_SMART                       uint32 = 0x80000313 // SMART self-test on CFC failed
	_ERROR_FILESYSTEM                  uint32 = 0x80000314 // Error in DOS file system
	_ERROR_SHORTKEY                    uint32 = 0x80000315 // MConfig.ini keyword too short
	_ERROR_MSYSINVALID                 uint32 = 0x80000316 // MSys is invalid (checksum, not found, ...)
	_ERROR_SNTPSYNC                    uint32 = 0x80000317 // SNTP synchronization failure
	_ERROR_BADSIOHANDSHAKE             uint32 = 0x80000318 // Hardware handshake is not permitted (probably device is used as console)
	_ERROR_ALLREADYEXIST               uint32 = 0x80002001 // Object already exist
	_ERROR_NOOBJECT                    uint32 = 0x80002002 // Object not found
	_ERROR_SQLERROR                    uint32 = 0x80002003 // SQL error occurred
	_ERROR_NOAHD                       uint32 = 0x80002004 // SW-Module AHD is not installed
	_ERROR_NOSESSAV                    uint32 = 0x80002005 // No session available
)
