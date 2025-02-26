// nolint: gocyclo,lll
package m1errors

import "errors"

var (
	ErrDescriptionOK                                 = errors.New("OK")
	ErrInProgress                                    = errors.New("still in progress")
	ErrCanceled                                      = errors.New("canceled by user")
	ErrFailed                                        = errors.New("unspecified error")
	ErrInvalidProgram                                = errors.New("SMI-Program not valid")
	ErrInvalidProcedure                              = errors.New("SMI-Procedure not valid")
	ErrInvalidArgs                                   = errors.New("SMI-Arguments not valid")
	ErrInvalidRPC                                    = errors.New("SMI-Version not valid")
	ErrInvalidAuth                                   = errors.New("SMI-Authentication failed")
	ErrInvalidVersion                                = errors.New("SMI-Protocol version bad")
	ErrInvalidPermissions                            = errors.New("SMI-Permission not valid")
	ErrParameter                                     = errors.New("parameter not valid")
	ErrFileNotFound                                  = errors.New("file not found")
	ErrFileTooBig                                    = errors.New("file too big")
	ErrFileEmpty                                     = errors.New("file is empty")
	ErrSectionNotFound                               = errors.New("section not found")
	ErrGroupNotFound                                 = errors.New("group not found")
	ErrKeywordNotFound                               = errors.New("keyword not found")
	ErrEndOfFile                                     = errors.New("end of file reached")
	ErrSemTakeProblem                                = errors.New("problem with semTake")
	ErrSetNotFound                                   = errors.New("set not found")
	ErrUnitNotFound                                  = errors.New("unit not found")
	ErrNotAllowedInSysMode                           = errors.New("not allowed in this SysMode")
	ErrModuleNotInstallable                          = errors.New("module not installable")
	ErrNotEnoughSystemMemory                         = errors.New("not enough system memory")
	ErrFunctionNotSupported                          = errors.New("function not supported")
	ErrNoModuleNumberAssigned                        = errors.New("no module number assigned")
	ErrTimeoutInReplyQueue                           = errors.New("timeout in reply queue")
	ErrNoAnswerFromSWModule                          = errors.New("no answer from SW-Module")
	ErrModuleNotFound                                = errors.New("module not found")
	ErrModuleDoesNotRespond                          = errors.New("module does not respond")
	ErrModuleNotPresentAnymore                       = errors.New("module not present any more")
	ErrModuleAsTaskNotPresent                        = errors.New("module as task not present")
	ErrModuleInVxWorksNotPresent                     = errors.New("module in VxWorks not present")
	ErrObjectNotAvailable                            = errors.New("object not available")
	ErrIndexInRequestNotValid                        = errors.New("index in request not valid")
	ErrAddressInRequestNotValid                      = errors.New("address in request not valid")
	ErrUserIdNotValid                                = errors.New("user-Id not valid")
	ErrListIdNotValid                                = errors.New("list-Id not valid")
	ErrUserAlreadyPresent                            = errors.New("user already present")
	ErrTooManyListElements                           = errors.New("too many list elements")
	ErrCallbackAddressNotValid                       = errors.New("callback-Address not valid")
	ErrObjectContainsBadElement                      = errors.New("object contains bad element")
	ErrNameNotValid                                  = errors.New("name not valid")
	ErrIOModuleNumberNotValid                        = errors.New("IO module number not valid")
	ErrDriverIdNotValid                              = errors.New("driver-ID not valid")
	ErrResourceAlreadyPresent                        = errors.New("resource already present")
	ErrResourceIsFull                                = errors.New("resource is full")
	ErrNoLicenseForSWModule                          = errors.New("no license for SW-Module")
	ErrLicenseExpired                                = errors.New("license expired")
	ErrBadObjectLibraryVersion                       = errors.New("bad object/library version")
	ErrNoReadPermission                              = errors.New("no read permission")
	ErrNoWritePermission                             = errors.New("no write permission")
	ErrReadErrorInFile                               = errors.New("read error in file")
	ErrWriteErrorInFile                              = errors.New("write error in file")
	ErrSearchErrorInFile                             = errors.New("search error in file")
	ErrChecksumNotValid                              = errors.New("checksum not valid")
	ErrModuleNotLoadableUnderVxWorks                 = errors.New("module not loadable under VxWorks")
	ErrModuleNotLoadableIntoMemory                   = errors.New("module not loadable into memory")
	ErrLibraryCannotBeRegistered                     = errors.New("library can not be registered")
	ErrResourceIsEmpty                               = errors.New("resource is empty")
	ErrModeNotValid                                  = errors.New("mode not valid")
	ErrObjectNotAllowed                              = errors.New("object not allowed")
	ErrAccessToObjectIsLocked                        = errors.New("access to object is locked")
	ErrObjectInUseByOtherClient                      = errors.New("object in use by other client")
	ErrFunctionInitIsMissing                         = errors.New("function ???_Init is missing")
	ErrModuleCannotBeRegistered                      = errors.New("module can not be registered")
	ErrFunctionInitReturnsError                      = errors.New("function ???_Init returns error")
	ErrBlockDeviceIsFull                             = errors.New("block-Device is full")
	ErrCopyErrorInFile                               = errors.New("copy error in file")
	ErrNotEnoughGlobalMemory                         = errors.New("not enough global memory")
	ErrNotEnoughApplicationMemory                    = errors.New("not enough application memory")
	ErrInternalSystemError1                          = errors.New("internal system error 1")
	ErrGatewayNotValid                               = errors.New("gateway not valid")
	ErrLoginNamePasswordNotValid                     = errors.New("login Name/Password not valid")
	ErrSerialDeviceIsMissing                         = errors.New("serial device is missing")
	ErrSerialDeviceNotValid                          = errors.New("serial device not valid")
	ErrDeviceAlreadyInUse                            = errors.New("device already in use")
	ErrDeviceNotPresent                              = errors.New("device not present")
	ErrModemNotPresent                               = errors.New("modem not present")
	ErrNoMoreUnitAvailable                           = errors.New("no more unit available")
	ErrModemNotValid                                 = errors.New("modem not valid")
	ErrPPPInitError                                  = errors.New("PPP init error")
	ErrPPPEstablishError                             = errors.New("PPP establish error")
	ErrSLIPInitError                                 = errors.New("SLIP init error")
	ErrProtocolNotValid                              = errors.New("protocol not valid")
	ErrLocalNameIsMissing                            = errors.New("local name is missing")
	ErrHostNameCannotBeSet                           = errors.New("host name can not be set")
	ErrHostDoesNotRespond                            = errors.New("host does not respond")
	ErrGatewayNameCannotBeSet                        = errors.New("gateway name can not be set")
	ErrGatewayDoesNotRespond                         = errors.New("gateway does not respond")
	ErrMountIsBad                                    = errors.New("mount is bad")
	ErrWrongDataType                                 = errors.New("wrong data type")
	ErrUnmountIsBad                                  = errors.New("unmount is bad")
	ErrSYSModuleCannotBeDeleted                      = errors.New("SYS-Module can not be deleted")
	ErrPathNameNotValid                              = errors.New("path/Name not valid")
	ErrFileMCoreIsMissing                            = errors.New("file MCore is missing")
	ErrFileMSysIsMissing                             = errors.New("file MSys is missing")
	ErrNoFTPServerOnHost                             = errors.New("no FTP-Server on host")
	ErrIPAddressAlreadyInUse                         = errors.New("IP-Address already in use")
	ErrLoginFromServerDenied                         = errors.New("login from server denied")
	ErrNoMConfigIniOnDevice                          = errors.New("no MConfig.ini on device")
	ErrFirmwareUpdateNotAllowed                      = errors.New("firmware update not allowed")
	ErrProgrammingFlashFailed                        = errors.New("programming flash failed")
	ErrVerifyFlashFailed                             = errors.New("verify flash failed")
	ErrRangeNotValid                                 = errors.New("range not valid")
	ErrNotForThisCPU                                 = errors.New("not for this CPU")
	ErrAddressInRequestIsOld                         = errors.New("address in request is old")
	ErrBlockDeviceNotPresent                         = errors.New("block Device not present")
	ErrErrorInFormatting                             = errors.New("error in formatting")
	ErrErrorInReopenAfterFormatting                  = errors.New("error in reopen after formatting")
	ErrErrorInDOSInitialization                      = errors.New("error in DOS initialization")
	ErrBootSectorNotSettable                         = errors.New("boot sector not settable")
	ErrFormattingNotAllowed                          = errors.New("formatting not allowed")
	ErrOldSoftwareVersionMCore                       = errors.New("old software version MCore")
	ErrOldSoftwareVersionMSys                        = errors.New("old software version MSys")
	ErrOldSoftwareVersionIODriver                    = errors.New("old software version IO-Driver")
	ErrLogDeviceNotValid                             = errors.New("log-Device not valid")
	ErrErrorInMemoryPartitionInit                    = errors.New("error in Memory-Partition init")
	ErrErrorInPanicHandlerInit                       = errors.New("error in Panic-Handler init")
	ErrErrorWatchdogHandlerInit                      = errors.New("error Watchdog-Handler init")
	ErrSWModuleNotErrorTolerant                      = errors.New("SW-Module not error tolerant")
	ErrTickRateTooLow                                = errors.New("tick rate too low")
	ErrErrorInTaskSpawn                              = errors.New("error in task spawn")
	ErrHWWatchdogMustBeOff                           = errors.New("HW-Watchdog must be off")
	ErrActionInThisStateNotAllowed                   = errors.New("action in this state not allowed")
	ErrBadNetworkFieldbusNode                        = errors.New("bad network/fieldbus node")
	ErrTimeoutInFunctionCall                         = errors.New("timeout in function call")
	ErrSoftwareReboot                                = errors.New("software reboot")
	ErrWatchdogReboot                                = errors.New("watchdog reboot")
	ErrOnlineChangeNotAllowed                        = errors.New("online change not allowed")
	ErrDriverIsMissing                               = errors.New("driver is missing")
	ErrDriverNotLoadable                             = errors.New("driver not loadable")
	ErrStationOrSlotNotValid                         = errors.New("station or slot not valid")
	ErrNumberStationsBelowMinimum                    = errors.New("number stations below minimum")
	ErrNumberStationsAboveMaximum                    = errors.New("number stations above maximum")
	ErrMinimumOneStationIsMissing                    = errors.New("minimum one station is missing")
	ErrSessionIdNotValid                             = errors.New("session ID not valid")
	ErrLoginUserNameNotValid                         = errors.New("login user name not valid")
	ErrLoginPasswordNotPresent                       = errors.New("login password not present")
	ErrLoginPasswordNotValid                         = errors.New("login password not valid")
	ErrLoginTimeFrameNotValid                        = errors.New("login time frame not valid")
	ErrTaskDoesNotExist                              = errors.New("task does not exist")
	ErrTaskCannotBeDebugged                          = errors.New("task can not be debugged")
	ErrTaskIsNotAttached                             = errors.New("task is not attached")
	ErrBreakpointNotSet                              = errors.New("breakpoint not set")
	ErrSymbolNotFound                                = errors.New("symbol not found")
	ErrTaskNameIsMissing                             = errors.New("task name is missing")
	ErrTaskNotOnBreakpoint                           = errors.New("task not on breakpoint")
	ErrCallbackAlreadyAttached                       = errors.New("callback already attached")
	ErrCallbackNotAttached                           = errors.New("callback not attached")
	ErrBreakpointAlreadySet                          = errors.New("breakpoint already set")
	ErrAllBreakpointsInUse                           = errors.New("all breakpoints in use")
	ErrVariableAlreadyInUse                          = errors.New("variable already in use")
	ErrSignalisingMBTransferInError                  = errors.New("signalising MB transfer in error for compatibility reasons")
	ErrNoBufferFound                                 = errors.New("no buffer found")
	ErrRestoringDeviceFailed                         = errors.New("restoring device failed")
	ErrWeakPowerSupplyOnAStation                     = errors.New("weak power supply on a station")
	ErrConfigurationDataNotValid                     = errors.New("configuration data not valid")
	ErrFirmwareNotForThisModule                      = errors.New("firmware not for this module")
	ErrFirmwareDownGradeError                        = errors.New("firmware down grade error")
	ErrModuleIsOldOrMissing                          = errors.New("module is old or missing")
	ErrTelnetInUseNoRedirect                         = errors.New("telnet in use, no redirect")
	ErrShellNotRunningNoRedirect                     = errors.New("shell not running, no redirect")
	ErrSNTPClientNotStarted                          = errors.New("SNTP client not started")
	ErrServiceNotAvailableInBootModeNOAPP            = errors.New("service not available in boot mode NOAPP")
	ErrFieldbusNetworkNumberNotValid                 = errors.New("fieldbus network number not valid")
	ErrFieldbusNetworkNodeIdMacIdNotValid            = errors.New("fieldbus network node id, mac id not valid")
	ErrDataSizeNotValidExceedsMaxSize                = errors.New("data size not valid, exceeds max size")
	ErrCommandTagNotValid                            = errors.New("command tag not valid")
	ErrCommandArgumentsNotValid                      = errors.New("command argument(s) not valid")
	ErrCommandEventQueueOverrun                      = errors.New("command/Event Queue overrun")
	ErrIOBusAccessFailsBadStation                    = errors.New("I/O Bus access fails (bad station)")
	ErrLogbookOfTheSLCReachedCriticalLevel           = errors.New("the logbook of the SLC reached a critical level")
	ErrSERCOSBrokenWire                              = errors.New("SERCOS: Broken wire")
	ErrSERCOSMissingDrive                            = errors.New("SERCOS: Missing drive")
	ErrSERCOSTimeoutWhileWaitingForStartupPhase      = errors.New("SERCOS: Timeout while waiting for startup phase")
	ErrAliveCheckFailedModuleIsDead                  = errors.New("alive check failed. Module is dead")
	ErrFallbackDeviceWasUsedForBooting               = errors.New("fallback device was used for booting")
	ErrMConfigIniNotFoundUsedBakInstead              = errors.New("MConfig.ini not found, used .bak instead")
	ErrErrorCreatingContiguousFile                   = errors.New("error creating contiguous file")
	ErrExceptionSignalOccurred                       = errors.New("exception signal occurred")
	ErrNoConnectionToDevice                          = errors.New("no connection to device")
	ErrInitValueError                                = errors.New("init value error")
	ErrCheckValueError                               = errors.New("check value error")
	ErrDriveSpecificError                            = errors.New("drive specific error")
	ErrHardwareError                                 = errors.New("hardware error")
	ErrCommunicationError                            = errors.New("communication error")
	ErrCyclicSetvalueMonitoringError                 = errors.New("cyclic setvalue monitoring error")
	ErrCycleTimeMonitoringError                      = errors.New("cycle time monitoring error")
	ErrModuleDoesNotSupportColdClimate               = errors.New("a module does not support cold climate")
	ErrAccessDenied                                  = errors.New("access denied")
	ErrWeakUserLevelDeprecated                       = errors.New("weak user level (deprecated: replaced by M_E_ACCDENIED)")
	ErrVHDLSTXERRCodeInformation                     = errors.New("VHD_LST_XERR code information")
	ErrWriteAccessAllowedOnlyOnPrimaryCPU            = errors.New("write access allowed only on primary CPU")
	ErrRIOffline                                     = errors.New("RI offline")
	ErrRINotRedundant                                = errors.New("RI not redundant")
	ErrDataLengthInvalid                             = errors.New("data length invalid")
	ErrApplicationNameMismatch                       = errors.New("application name mismatch")
	ErrNumberOfBlocksDifferent                       = errors.New("number of blocks different")
	ErrMemoryBlocksDifferent                         = errors.New("memory blocks different")
	ErrDeviceNotRedundant                            = errors.New("device not redundant")
	ErrDeviceNetworkNotRedundant                     = errors.New("device network not redundant")
	ErrApplicationError1                             = errors.New("application error1")
	ErrApplicationError2                             = errors.New("application error2")
	ErrApplicationError3                             = errors.New("application error3")
	ErrApplicationFatalError                         = errors.New("application fatal error")
	ErrApplicationDoneError                          = errors.New("application done error")
	ErrApplicationOverload                           = errors.New("application overload")
	ErrCycleOverload                                 = errors.New("cycle overload")
	ErrBCHCommandError                               = errors.New("BCH command error")
	ErrPrimaryAssignmentConflict                     = errors.New("primary assignment conflict")
	ErrSwitchoverCommand                             = errors.New("switchover command")
	ErrApplicationChecksumMismatch                   = errors.New("application checksum mismatch")
	ErrSelfTestSwitchoverCommand                     = errors.New("self test switchover command")
	ErrSelfTestRun                                   = errors.New("self test run")
	ErrSelfTestError                                 = errors.New("self test error")
	ErrProgrammingError                              = errors.New("programming error")
	ErrSelfTestIDNotValid                            = errors.New("self test ID not valid")
	ErrSelfTestRAMError                              = errors.New("self test RAM error")
	ErrSelfTestPLCError                              = errors.New("self test PLC error")
	ErrSelfTestCOMError                              = errors.New("self test COM error")
	ErrCipherMethodNotAvailable                      = errors.New("cipher method not available")
	ErrCouldntCreateCipherHash                       = errors.New("couldn't create cipher hash")
	ErrNoEncryptionLayerFound                        = errors.New("no encryption layer found")
	ErrNoMasterkeyAvailable                          = errors.New("no masterkey available")
	ErrBootPartitionCannotBeEncrypted                = errors.New("boot partition can not be encrypted")
	ErrLoadingCipherLibraryFailed                    = errors.New("loading cipher library failed")
	ErrCipherLibraryVersion                          = errors.New("cipher library version")
	ErrInitCipherLibraryFailed                       = errors.New("init cipher library failed")
	ErrCipherStrengthNotValid                        = errors.New("cipher strength not valid")
	ErrInitCipherAlgorithmFailed                     = errors.New("init cipher algorithm failed")
	ErrErrorEncryptingData                           = errors.New("error encrypting data")
	ErrErrorDecryptingData                           = errors.New("error decrypting data")
	ErrRecoverKeyFileNotFound                        = errors.New("recover key file not found")
	ErrRecoveryKeyFileCRCError                       = errors.New("recovery key file crc error")
	ErrRecoveryPasswordWrong                         = errors.New("recovery password wrong")
	ErrNoProgressAvailable                           = errors.New("no progress available")
	ErrDetailedReasonUnknown                         = errors.New("detailed reason unknown")
	ErrConfigurationOfSNTPServerNotValid             = errors.New("configuration of SNTP server not valid")
	ErrOnlineReconfigNotSupported                    = errors.New("online re-config not supported")
	ErrConfigurationOfCoreCategoryNotValid           = errors.New("configuration of CoreCategory not valid")
	ErrPasswordExpired                               = errors.New("password expired")
	ErrUserIsDisabled                                = errors.New("user is disabled")
	ErrResourceIsCurrentlyBusy                       = errors.New("resource is currently busy")
	ErrErrorCreatingWatchdog                         = errors.New("error creating watchdog")
	ErrErrorCreatingDeletingSemaphore                = errors.New("error creating/deleting semaphore")
	ErrPriorityOutOfRange                            = errors.New("priority out of range (has to be between 1 and 255)")
	ErrOffsetDoesNotMatchTickSyncRate                = errors.New("offset does not match tick/sync rate or is greater than cycle time")
	ErrNextPossibleCycleTimeDiffersMoreThan25Percent = errors.New("next possible cycle time does differ more than 25% from configured cycle time")
	ErrCycleTimeIsTooSmall                           = errors.New("cycle time is too small")
	ErrErrorAttachingAuxSync                         = errors.New("error attaching aux/sync")
	ErrAuxiliaryClockIsTurnedOff                     = errors.New("auxiliary clock is turned off")
	ErrUnknownTaskTriggerSource                      = errors.New("unknown task trigger source")
	ErrUserMainFunctionIsNULL                        = errors.New("user main function is NULL")
	ErrPTPSynchronizationNotActivated                = errors.New("PTP synchronization not activated")
	ErrCoreCategoryNotValid                          = errors.New("core category not valid")
	ErrTaskHasNotBeenStopped                         = errors.New("task has not been stopped")
	ErrSelectedAttributesMustNotBeChanged            = errors.New("selected attributes must not be changed")
	ErrErrorInitRTCChip                              = errors.New("error init RTC chip")
	ErrErrorInitEthernetController                   = errors.New("error init Ethernet controller")
	ErrIPAddressNotValid                             = errors.New("IP-Address not valid")
	ErrTickRateNotValid                              = errors.New("tick rate not valid")
	ErrCPUClockNotValid                              = errors.New("CPU clock not valid")
	ErrNVRamCreateFailed                             = errors.New("NV-Ram create failed")
	ErrNVRamNotPresent                               = errors.New("NV-Ram not present")
	ErrNVRamIsEmpty                                  = errors.New("NV-Ram is empty")
	ErrNVRamNoDOSFileSystem                          = errors.New("NV-Ram no DOS file system")
	ErrRamDiskCreateFailed                           = errors.New("Ram-Disk create failed")
	ErrBatteryVoltageTooLow                          = errors.New("battery voltage too low")
	ErrErrorInitSIODriver                            = errors.New("error init SIO driver")
	ErrDuplicateIPAddress                            = errors.New("duplicate IP address")
	ErrTooMuchEHDEntriesInMCore                      = errors.New("too much EHD entries in MCore")
	ErrPanicSituationOccurredDuringBoot              = errors.New("panic situation occurred during boot")
	ErrInvalidNVRamSignature                         = errors.New("invalid NV-Ram signature")
	ErrNATInitializationFailed                       = errors.New("NAT initialization failed")
	ErrSMARTSelfTestOnCFCFailed                      = errors.New("SMART self-test on CFC failed")
	ErrErrorInDOSFileSystem                          = errors.New("error in DOS file system")
	ErrMConfigIniKeywordTooShort                     = errors.New("MConfig.ini keyword too short")
	ErrMSysIsInvalid                                 = errors.New("MSys is invalid (checksum, not found, ...)")
	ErrSNTPSynchronizationFailure                    = errors.New("SNTP synchronization failure")
	ErrHardwareHandshakeNotPermitted                 = errors.New("hardware handshake is not permitted (probably device is used as console)")
	ErrObjectAlreadyExist                            = errors.New("object already exist")
	ErrObjectNotFound                                = errors.New("object not found")
	ErrSQLErrorOccurred                              = errors.New("SQL error occurred")
	ErrSWModuleAHDIsNotInstalled                     = errors.New("SW-Module AHD is not installed")
	ErrNoSessionAvailable                            = errors.New("no session available")
	ErrUnknown                                       = errors.New("unknown")
)

func description(code uint32) error {
	switch code % 0x0000ffff {
	default:
		return ErrUnknown
	case ErrorOK:
		return ErrDescriptionOK
	case ErrorINPROGRESS:
		return ErrInProgress
	case ErrorCANCELED:
		return ErrCanceled
	case ErrorFAILED:
		return ErrFailed
	case ErrorBADPROG:
		return ErrInvalidProgram
	case ErrorBADPROC:
		return ErrInvalidProcedure
	case ErrorBADARGS:
		return ErrInvalidArgs
	case ErrorBADRPC:
		return ErrInvalidRPC
	case ErrorBADAUTH:
		return ErrInvalidAuth
	case ErrorBADVERS:
		return ErrInvalidVersion
	case ErrorBADPERM:
		return ErrInvalidPermissions
	case ErrorPARM:
		return ErrParameter
	case ErrorNOFILE:
		return ErrFileNotFound
	case ErrorFILEBIG:
		return ErrFileTooBig
	case ErrorFILEEMPTY:
		return ErrFileEmpty
	case ErrorNOSEC:
		return ErrSectionNotFound
	case ErrorNOGRP:
		return ErrGroupNotFound
	case ErrorNOKEY:
		return ErrKeywordNotFound
	case ErrorENDFILE:
		return ErrEndOfFile
	case ErrorNOSEMA:
		return ErrSemTakeProblem
	case ErrorNOPFSET:
		return ErrSetNotFound
	case ErrorNOPFUNIT:
		return ErrUnitNotFound
	case ErrorSMODE:
		return ErrNotAllowedInSysMode
	case ErrorINSTALL:
		return ErrModuleNotInstallable
	case ErrorNOMEM:
		return ErrNotEnoughSystemMemory
	case ErrorNOTSUPP:
		return ErrFunctionNotSupported
	case ErrorNOMODNBR:
		return ErrNoModuleNumberAssigned
	case ErrorTIMEOUT1:
		return ErrTimeoutInReplyQueue
	case ErrorTIMEOUT2:
		return ErrNoAnswerFromSWModule
	case ErrorTIMEOUT3:
		return ErrTimeoutInReplyQueue
	case ErrorNOMOD1:
		return ErrModuleNotFound
	case ErrorNOMOD2:
		return ErrModuleDoesNotRespond
	case ErrorNOMOD3:
		return ErrModuleNotPresentAnymore
	case ErrorNODELTSK:
		return ErrModuleAsTaskNotPresent
	case ErrorNOVXWOBJ:
		return ErrModuleInVxWorksNotPresent
	case ErrorNOOBJ:
		return ErrObjectNotAvailable
	case ErrorBADINDEX:
		return ErrIndexInRequestNotValid
	case ErrorBADADDR:
		return ErrAddressInRequestNotValid
	case ErrorUSERID:
		return ErrUserIdNotValid
	case ErrorLISTID:
		return ErrListIdNotValid
	case ErrorDUPUSER:
		return ErrUserAlreadyPresent
	case ErrorNBELEM:
		return ErrTooManyListElements
	case ErrorCBACK:
		return ErrCallbackAddressNotValid
	case ErrorBADELEM:
		return ErrObjectContainsBadElement
	case ErrorBADNAME:
		return ErrNameNotValid
	case ErrorCARDNB:
		return ErrIOModuleNumberNotValid
	case ErrorDRVID:
		return ErrDriverIdNotValid
	case ErrorDUPRES:
		return ErrResourceAlreadyPresent
	case ErrorFULL:
		return ErrResourceIsFull
	case ErrorNOLIC:
		return ErrNoLicenseForSWModule
	case ErrorNOLICEXP:
		return ErrLicenseExpired
	case ErrorWRONGVERS:
		return ErrBadObjectLibraryVersion
	case ErrorNOREAD:
		return ErrNoReadPermission
	case ErrorNOWRITE:
		return ErrNoWritePermission
	case ErrorBADREAD:
		return ErrReadErrorInFile
	case ErrorBADWRITE:
		return ErrWriteErrorInFile
	case ErrorBADSEEK:
		return ErrSearchErrorInFile
	case ErrorBADCHECK:
		return ErrChecksumNotValid
	case ErrorBADVXWLD:
		return ErrModuleNotLoadableUnderVxWorks
	case ErrorBADMEMLD:
		return ErrModuleNotLoadableIntoMemory
	case ErrorNOLIBREG:
		return ErrLibraryCannotBeRegistered
	case ErrorEMPTY:
		return ErrResourceIsEmpty
	case ErrorBADMODE:
		return ErrModeNotValid
	case ErrorBADOBJ:
		return ErrObjectNotAllowed
	case ErrorLOCKED:
		return ErrAccessToObjectIsLocked
	case ErrorBADIPADDR:
		return ErrObjectInUseByOtherClient
	case ErrorNOENTRY:
		return ErrFunctionInitIsMissing
	case ErrorNOREG:
		return ErrModuleCannotBeRegistered
	case ErrorBADINIT:
		return ErrFunctionInitReturnsError
	case ErrorDEVFULL:
		return ErrBlockDeviceIsFull
	case ErrorBADCOPY:
		return ErrCopyErrorInFile
	case ErrorNOGLOBMEM:
		return ErrNotEnoughGlobalMemory
	case ErrorNOAPPMEM:
		return ErrNotEnoughApplicationMemory
	case ErrorSYSTEM1:
		return ErrInternalSystemError1
	case ErrorBADROUTE:
		return ErrGatewayNotValid
	case ErrorBADLOGIN:
		return ErrLoginNamePasswordNotValid
	case ErrorNOSIODEV:
		return ErrSerialDeviceIsMissing
	case ErrorBADSIODEV:
		return ErrSerialDeviceNotValid
	case ErrorDEVINUSE:
		return ErrDeviceAlreadyInUse
	case ErrorDEVMISS:
		return ErrDeviceNotPresent
	case ErrorNOMODEM:
		return ErrModemNotPresent
	case ErrorNOUNIT:
		return ErrNoMoreUnitAvailable
	case ErrorBADMODEM:
		return ErrModemNotValid
	case ErrorBADPPP1:
		return ErrPPPInitError
	case ErrorBADPPP2:
		return ErrPPPEstablishError
	case ErrorBADSLIP:
		return ErrSLIPInitError
	case ErrorBADPROTO:
		return ErrProtocolNotValid
	case ErrorLNAMEMISS:
		return ErrLocalNameIsMissing
	case ErrorBADHOST:
		return ErrHostNameCannotBeSet
	case ErrorNOHOST:
		return ErrHostDoesNotRespond
	case ErrorBADGATE:
		return ErrGatewayNameCannotBeSet
	case ErrorNOGATE:
		return ErrGatewayDoesNotRespond
	case ErrorBADMOUNT:
		return ErrMountIsBad
	case ErrorBADTYPE:
		return ErrWrongDataType
	case ErrorBADUMOUNT:
		return ErrUnmountIsBad
	case ErrorNODELSYS:
		return ErrSYSModuleCannotBeDeleted
	case ErrorBADPATH:
		return ErrPathNameNotValid
	case ErrorNOSYS1:
		return ErrFileMCoreIsMissing
	case ErrorNOSYS2:
		return ErrFileMSysIsMissing
	case ErrorNOFTP:
		return ErrNoFTPServerOnHost
	case ErrorIPINUSE:
		return ErrIPAddressAlreadyInUse
	case ErrorNOLOGIN:
		return ErrLoginFromServerDenied
	case ErrorNOMCONF:
		return ErrNoMConfigIniOnDevice
	case ErrorUPDATE:
		return ErrFirmwareUpdateNotAllowed
	case ErrorBADFPROG:
		return ErrProgrammingFlashFailed
	case ErrorBADVERIFY:
		return ErrVerifyFlashFailed
	case ErrorBADRANGE:
		return ErrRangeNotValid
	case ErrorBADCPU:
		return ErrNotForThisCPU
	case ErrorOLDADDR:
		return ErrAddressInRequestIsOld
	case ErrorNOBLOCKDEV:
		return ErrBlockDeviceNotPresent
	case ErrorBADFORMAT:
		return ErrErrorInFormatting
	case ErrorBADREOPEN:
		return ErrErrorInReopenAfterFormatting
	case ErrorBADDOSINI:
		return ErrErrorInDOSInitialization
	case ErrorBADBOOTSEC:
		return ErrBootSectorNotSettable
	case ErrorNOFORMAT:
		return ErrFormattingNotAllowed
	case ErrorOLDVERS1:
		return ErrOldSoftwareVersionMCore
	case ErrorOLDVERS2:
		return ErrOldSoftwareVersionMSys
	case ErrorOLDVERS3:
		return ErrOldSoftwareVersionIODriver
	case ErrorBADLOGDEV:
		return ErrLogDeviceNotValid
	case ErrorBADMEMINIT:
		return ErrErrorInMemoryPartitionInit
	case ErrorBADPNCINIT:
		return ErrErrorInPanicHandlerInit
	case ErrorBADWDGINIT:
		return ErrErrorWatchdogHandlerInit
	case ErrorNOERRTOL:
		return ErrSWModuleNotErrorTolerant
	case ErrorBADTCKRATE:
		return ErrTickRateTooLow
	case ErrorBADSPAWN:
		return ErrErrorInTaskSpawn
	case ErrorWDOGON:
		return ErrHWWatchdogMustBeOff
	case ErrorINVSTATE:
		return ErrActionInThisStateNotAllowed
	case ErrorNODE:
		return ErrBadNetworkFieldbusNode
	case ErrorTIMEOUT:
		return ErrTimeoutInFunctionCall
	case ErrorSWREBOOT:
		return ErrSoftwareReboot
	case ErrorWDGREBOOT:
		return ErrWatchdogReboot
	case ErrorNOONLINE:
		return ErrOnlineChangeNotAllowed
	case ErrorDRVMISS:
		return ErrDriverIsMissing
	case ErrorBADDRV:
		return ErrDriverNotLoadable
	case ErrorBADSLOT:
		return ErrStationOrSlotNotValid
	case ErrorBELOWMIN:
		return ErrNumberStationsBelowMinimum
	case ErrorABOVEMAX:
		return ErrNumberStationsAboveMaximum
	case ErrorSTATMISS:
		return ErrMinimumOneStationIsMissing
	case ErrorBADSESS:
		return ErrSessionIdNotValid
	case ErrorBADUSER:
		return ErrLoginUserNameNotValid
	case ErrorNOPWORD:
		return ErrLoginPasswordNotPresent
	case ErrorBADPWORD:
		return ErrLoginPasswordNotValid
	case ErrorBADTIME:
		return ErrLoginTimeFrameNotValid
	case ErrorNOTASK:
		return ErrTaskDoesNotExist
	case ErrorNODEBUG:
		return ErrTaskCannotBeDebugged
	case ErrorNOATTACH:
		return ErrTaskIsNotAttached
	case ErrorNOBKPT:
		return ErrBreakpointNotSet
	case ErrorNOSYM:
		return ErrSymbolNotFound
	case ErrorTSKMISS:
		return ErrTaskNameIsMissing
	case ErrorNOSSTEP:
		return ErrTaskNotOnBreakpoint
	case ErrorHAVECBACK:
		return ErrCallbackAlreadyAttached
	case ErrorNOCBACK:
		return ErrCallbackNotAttached
	case ErrorHAVEBKPT:
		return ErrBreakpointAlreadySet
	case ErrorALLBKPT:
		return ErrAllBreakpointsInUse
	case ErrorVARINUSE:
		return ErrVariableAlreadyInUse
	case ErrorMBTRANS:
		return ErrSignalisingMBTransferInError
	case ErrorNOBUFF:
		return ErrNoBufferFound
	case ErrorBADRESTORE:
		return ErrRestoringDeviceFailed
	case ErrorWEAKPWR:
		return ErrWeakPowerSupplyOnAStation
	case ErrorBADCONFIG:
		return ErrConfigurationDataNotValid
	case ErrorBADFWARE:
		return ErrFirmwareNotForThisModule
	case ErrorDOWNGRADE:
		return ErrFirmwareDownGradeError
	case ErrorMODMISS:
		return ErrModuleIsOldOrMissing
	case ErrorTNETINUSE:
		return ErrTelnetInUseNoRedirect
	case ErrorNOSHELL:
		return ErrShellNotRunningNoRedirect
	case ErrorBADSNTP:
		return ErrSNTPClientNotStarted
	case ErrorNOAPP:
		return ErrServiceNotAvailableInBootModeNOAPP
	case ErrorNETNB:
		return ErrFieldbusNetworkNumberNotValid
	case ErrorNODEID:
		return ErrFieldbusNetworkNodeIdMacIdNotValid
	case ErrorDSIZE:
		return ErrDataSizeNotValidExceedsMaxSize
	case ErrorCMDTAG:
		return ErrCommandTagNotValid
	case ErrorCMDARG:
		return ErrCommandArgumentsNotValid
	case ErrorQUEUEFULL:
		return ErrCommandEventQueueOverrun
	case ErrorIOBUSFAIL:
		return ErrIOBusAccessFailsBadStation
	case ErrorSLCLOGFULL:
		return ErrLogbookOfTheSLCReachedCriticalLevel
	case ErrorBROKENWIRE:
		return ErrSERCOSBrokenWire
	case ErrorDRIVEMISS:
		return ErrSERCOSMissingDrive
	case ErrorSTARTUPTMO:
		return ErrSERCOSTimeoutWhileWaitingForStartupPhase
	case ErrorDEAD:
		return ErrAliveCheckFailedModuleIsDead
	case ErrorFALLBACK:
		return ErrFallbackDeviceWasUsedForBooting
	case ErrorMCONFIGBAK:
		return ErrMConfigIniNotFoundUsedBakInstead
	case ErrorCONTIGFILE:
		return ErrErrorCreatingContiguousFile
	case ErrorEXCEPTION:
		return ErrExceptionSignalOccurred
	case ErrorOFFLINE:
		return ErrNoConnectionToDevice
	case ErrorINITVAL:
		return ErrInitValueError
	case ErrorCHECKVAL:
		return ErrCheckValueError
	case ErrorDRIVE:
		return ErrDriveSpecificError
	case ErrorHARDWARE:
		return ErrHardwareError
	case ErrorCOMMUNIC:
		return ErrCommunicationError
	case ErrorNOSETVAL:
		return ErrCyclicSetvalueMonitoringError
	case ErrorCYCLETIME:
		return ErrCycleTimeMonitoringError
	case ErrorCOLDCLIMATE:
		return ErrModuleDoesNotSupportColdClimate
	case ErrorACCDENIED:
		return ErrAccessDenied
	case ErrorBADUSERLVL:
		return ErrWeakUserLevelDeprecated
	case ErrorXERR:
		return ErrVHDLSTXERRCodeInformation
	case ErrorREDULOCK:
		return ErrWriteAccessAllowedOnlyOnPrimaryCPU
	case ErrorRICONN:
		return ErrRIOffline
	case ErrorRINETREDU:
		return ErrRINotRedundant
	case ErrorLENGTH:
		return ErrDataLengthInvalid
	case ErrorAPPNAME:
		return ErrApplicationNameMismatch
	case ErrorBLOCKNB:
		return ErrNumberOfBlocksDifferent
	case ErrorBLOCKDIFF:
		return ErrMemoryBlocksDifferent
	case ErrorDEVREDU:
		return ErrDeviceNotRedundant
	case ErrorDEVNETREDU:
		return ErrDeviceNetworkNotRedundant
	case ErrorAPPERROR1:
		return ErrApplicationError1
	case ErrorAPPERROR2:
		return ErrApplicationError2
	case ErrorAPPERROR3:
		return ErrApplicationError3
	case ErrorAPPFATAL:
		return ErrApplicationFatalError
	case ErrorAPPDONE:
		return ErrApplicationDoneError
	case ErrorAPPOVLOAD:
		return ErrApplicationOverload
	case ErrorCYCLEOVLOAD:
		return ErrCycleOverload
	case ErrorBCHCMD:
		return ErrBCHCommandError
	case ErrorPRICONFLICT:
		return ErrPrimaryAssignmentConflict
	case ErrorSWITCHOVER:
		return ErrSwitchoverCommand
	case ErrorAPPCS:
		return ErrApplicationChecksumMismatch
	case ErrorSTSWITCHOVER:
		return ErrSelfTestSwitchoverCommand
	case ErrorSTRUN:
		return ErrSelfTestRun
	case ErrorSTERROR:
		return ErrSelfTestError
	case ErrorPROGRAM:
		return ErrProgrammingError
	case ErrorSTID:
		return ErrSelfTestIDNotValid
	case ErrorSTRAM:
		return ErrSelfTestRAMError
	case ErrorSTPLC:
		return ErrSelfTestPLCError
	case ErrorSTCOM:
		return ErrSelfTestCOMError
	case ErrorCRYPT_METADATA_METHOD:
		return ErrCipherMethodNotAvailable
	case ErrorCRYPT_METADATA_HASH:
		return ErrCouldntCreateCipherHash
	case ErrorCRYPT_CRYPT_EXISTS:
		return ErrNoEncryptionLayerFound
	case ErrorCRYPT_CRYPT_MASTERKEY:
		return ErrNoMasterkeyAvailable
	case ErrorCRYPT_CRYPT_BOOTDEV:
		return ErrBootPartitionCannotBeEncrypted
	case ErrorCRYPT_ALGORITHM_LOAD_LIB:
		return ErrLoadingCipherLibraryFailed
	case ErrorCRYPT_ALGORITHM_LIB_VERSION:
		return ErrCipherLibraryVersion
	case ErrorCRYPT_ALGORITHM_INIT_LIB:
		return ErrInitCipherLibraryFailed
	case ErrorCRYPT_ALGORITHM_KEYLENGTH:
		return ErrCipherStrengthNotValid
	case ErrorCRYPT_ALGORITHM_INIT:
		return ErrInitCipherAlgorithmFailed
	case ErrorCRYPT_ALGORITHM_CRYPT:
		return ErrErrorEncryptingData
	case ErrorCRYPT_ALGORITHM_DECRYPT:
		return ErrErrorDecryptingData
	case ErrorCRYPT_KEY_RECOVERYFILE:
		return ErrRecoverKeyFileNotFound
	case ErrorCRYPT_KEY_RECOVERYFILECRC:
		return ErrRecoveryKeyFileCRCError
	case ErrorCRYPT_KEY_RECOVERYPASSWORD:
		return ErrRecoveryPasswordWrong
	case ErrorCRYPT_NOPROGRESS:
		return ErrNoProgressAvailable
	case ErrorCRYPT_UNKNOWN:
		return ErrDetailedReasonUnknown
	case ErrorBADSNTPSERVER:
		return ErrConfigurationOfSNTPServerNotValid
	case ErrorNOONLINECFG:
		return ErrOnlineReconfigNotSupported
	case ErrorBADCORECATEG:
		return ErrConfigurationOfCoreCategoryNotValid
	case ErrorPWDEXPIRED:
		return ErrPasswordExpired
	case ErrorBADLOGINLOCKED:
		return ErrUserIsDisabled
	case ErrorBUSY:
		return ErrResourceIsCurrentlyBusy
	case ErrorTASK_WDOG:
		return ErrErrorCreatingWatchdog
	case ErrorTASK_SEM:
		return ErrErrorCreatingDeletingSemaphore
	case ErrorTASK_PRIO:
		return ErrPriorityOutOfRange
	case ErrorTASK_OFFSET:
		return ErrOffsetDoesNotMatchTickSyncRate
	case ErrorTASK_CYCDIFF:
		return ErrNextPossibleCycleTimeDiffersMoreThan25Percent
	case ErrorTASK_CYCTIME:
		return ErrCycleTimeIsTooSmall
	case ErrorTASK_ATTACH:
		return ErrErrorAttachingAuxSync
	case ErrorTASK_AUXOFF:
		return ErrAuxiliaryClockIsTurnedOff
	case ErrorTASKSource:
		return ErrUnknownTaskTriggerSource
	case ErrorTASK_NOMAINFUNC:
		return ErrUserMainFunctionIsNULL
	case ErrorTASK_NOPTP:
		return ErrPTPSynchronizationNotActivated
	case ErrorTASK_CORECATEG:
		return ErrCoreCategoryNotValid
	case ErrorTASK_RUNNING:
		return ErrTaskHasNotBeenStopped
	case ErrorTASK_CHANGEATTR:
		return ErrSelectedAttributesMustNotBeChanged
	case ErrorRTC:
		return ErrErrorInitRTCChip
	case ErrorETHER:
		return ErrErrorInitEthernetController
	case ErrorIP_ADDR:
		return ErrIPAddressNotValid
	case ErrorTICKRATE:
		return ErrTickRateNotValid
	case ErrorCPUCLOCK:
		return ErrCPUClockNotValid
	case ErrorNVMAKE:
		return ErrNVRamCreateFailed
	case ErrorNVMISS:
		return ErrNVRamNotPresent
	case ErrorNVEMPTY:
		return ErrNVRamIsEmpty
	case ErrorNVNODOS:
		return ErrNVRamNoDOSFileSystem
	case ErrorRDMAKE:
		return ErrRamDiskCreateFailed
	case ErrorBATTLOW:
		return ErrBatteryVoltageTooLow
	case ErrorSIO:
		return ErrErrorInitSIODriver
	case ErrorDUPIP:
		return ErrDuplicateIPAddress
	case ErrorEHDOFLOW:
		return ErrTooMuchEHDEntriesInMCore
	case ErrorBOOTPANIC:
		return ErrPanicSituationOccurredDuringBoot
	case ErrorNVRAMSIG:
		return ErrInvalidNVRamSignature
	case ErrorNATFAIL:
		return ErrNATInitializationFailed
	case ErrorSMART:
		return ErrSMARTSelfTestOnCFCFailed
	case ErrorFILESYSTEM:
		return ErrErrorInDOSFileSystem
	case ErrorSHORTKEY:
		return ErrMConfigIniKeywordTooShort
	case ErrorMSYSINVALID:
		return ErrMSysIsInvalid
	case ErrorSNTPSYNC:
		return ErrSNTPSynchronizationFailure
	case ErrorBADSIOHANDSHAKE:
		return ErrHardwareHandshakeNotPermitted
	case ErrorALLREADYEXIST:
		return ErrObjectAlreadyExist
	case ErrorNOOBJECT:
		return ErrObjectNotFound
	case ErrorSQLERROR:
		return ErrSQLErrorOccurred
	case ErrorNOAHD:
		return ErrSWModuleAHDIsNotInstalled
	case ErrorNOSESSAV:
		return ErrNoSessionAvailable
	}
}

// Error codes.
const (
	ErrorOK                          uint32 = 0          // O.K., no error
	ErrorINPROGRESS                  uint32 = 1          // Still in progress, no error
	ErrorCANCELED                    uint32 = 2          // Canceled by user, no error
	ErrorFAILED                      uint32 = 0x80000100 // Unspecified error
	ErrorBADPROG                     uint32 = 0x80000101 // SMI-Program not valid
	ErrorBADPROC                     uint32 = 0x80000102 // SMI-Procedure not valid
	ErrorBADARGS                     uint32 = 0x80000103 // SMI-Arguments not valid
	ErrorBADRPC                      uint32 = 0x80000104 // SMI-Version not valid
	ErrorBADAUTH                     uint32 = 0x80000105 // SMI-Authentication failed
	ErrorBADVERS                     uint32 = 0x80000106 // SMI-Protocol version bad
	ErrorBADPERM                     uint32 = 0x80000107 // SMI-Permission not valid
	ErrorPARM                        uint32 = 0x80000110 // Parameter not valid
	ErrorNOFILE                      uint32 = 0x80000111 // File not found
	ErrorFILEBIG                     uint32 = 0x80000112 // File too big
	ErrorFILEEMPTY                   uint32 = 0x80000113 // File is empty
	ErrorNOSEC                       uint32 = 0x80000114 // Section not found
	ErrorNOGRP                       uint32 = 0x80000115 // Group not found
	ErrorNOKEY                       uint32 = 0x80000116 // Keyword not found
	ErrorENDFILE                     uint32 = 0x80000117 // End of file reached
	ErrorNOSEMA                      uint32 = 0x80000118 // Problem with semTake
	ErrorNOPFSET                     uint32 = 0x80000119 // Set not found
	ErrorNOPFUNIT                    uint32 = 0x80000120 // Unit not found
	ErrorSMODE                       uint32 = 0x80000121 // Not allowed in this SysMode
	ErrorINSTALL                     uint32 = 0x80000122 // Module not installable
	ErrorNOMEM                       uint32 = 0x80000123 // Not enough system memory
	ErrorNOTSUPP                     uint32 = 0x80000124 // Function not supported
	ErrorNOMODNBR                    uint32 = 0x80000125 // No module number assigned
	ErrorTIMEOUT1                    uint32 = 0x80000126 // Timeout in reply queue
	ErrorTIMEOUT2                    uint32 = 0x80000127 // No answer from SW-Module
	ErrorTIMEOUT3                    uint32 = 0x80000128 // Timeout in reply queue
	ErrorNOMOD1                      uint32 = 0x80000129 // Module not found
	ErrorNOMOD2                      uint32 = 0x8000012A // Module does not respond
	ErrorNOMOD3                      uint32 = 0x8000012B // Module not present any more
	ErrorNODELTSK                    uint32 = 0x8000012C // Module as task not present
	ErrorNOVXWOBJ                    uint32 = 0x8000012D // Module in VxWorks not present
	ErrorNOOBJ                       uint32 = 0x8000012E // Object not available
	ErrorBADINDEX                    uint32 = 0x8000012F // Index in request not valid
	ErrorBADADDR                     uint32 = 0x80000130 // Address in request not valid
	ErrorUSERID                      uint32 = 0x80000131 // User-Id not valid
	ErrorLISTID                      uint32 = 0x80000132 // List-Id not valid
	ErrorDUPUSER                     uint32 = 0x80000133 // User already present
	ErrorNBELEM                      uint32 = 0x80000134 // Too many list elements
	ErrorCBACK                       uint32 = 0x80000135 // Callback-Address not valid
	ErrorBADELEM                     uint32 = 0x80000136 // Object contains bad element
	ErrorBADNAME                     uint32 = 0x80000137 // Name not valid
	ErrorCARDNB                      uint32 = 0x80000138 // IO module number not valid
	ErrorDRVID                       uint32 = 0x80000139 // Driver-ID not valid
	ErrorDUPRES                      uint32 = 0x8000013A // Resource already present
	ErrorFULL                        uint32 = 0x8000013B // Resource is full
	ErrorNOLIC                       uint32 = 0x8000013C // No license for SW-Module
	ErrorNOLICEXP                    uint32 = 0x8000013D // License expired
	ErrorWRONGVERS                   uint32 = 0x8000013E // Bad object/library version
	ErrorNOREAD                      uint32 = 0x80000140 // No read permission
	ErrorNOWRITE                     uint32 = 0x80000141 // No write permission
	ErrorBADREAD                     uint32 = 0x80000142 // Read error in file
	ErrorBADWRITE                    uint32 = 0x80000143 // Write error in file
	ErrorBADSEEK                     uint32 = 0x80000144 // Search error in file
	ErrorBADCHECK                    uint32 = 0x80000145 // Checksum not valid
	ErrorBADVXWLD                    uint32 = 0x80000146 // Module not loadable under VxWorks
	ErrorBADMEMLD                    uint32 = 0x80000147 // Module not loadable into memory
	ErrorNOLIBREG                    uint32 = 0x80000148 // Library can not be registered
	ErrorEMPTY                       uint32 = 0x80000149 // Resource is empty
	ErrorBADMODE                     uint32 = 0x8000014A // Mode not valid
	ErrorBADOBJ                      uint32 = 0x8000014B // Object not allowed
	ErrorLOCKED                      uint32 = 0x8000014C // Access to object is locked
	ErrorBADIPADDR                   uint32 = 0x8000014D // Object in use by other client
	ErrorNOENTRY                     uint32 = 0x8000014E // Function ???_Init is missing
	ErrorNOREG                       uint32 = 0x8000014F // Module can not be registered
	ErrorBADINIT                     uint32 = 0x80000150 // Function ???_Init returns error
	ErrorDEVFULL                     uint32 = 0x80000151 // Block-Device is full
	ErrorBADCOPY                     uint32 = 0x80000152 // Copy error in file
	ErrorNOGLOBMEM                   uint32 = 0x80000153 // Not enough global memory
	ErrorNOAPPMEM                    uint32 = 0x80000154 // Not enough application memory
	ErrorSYSTEM1                     uint32 = 0x80000156 // Internal system error 1
	ErrorBADROUTE                    uint32 = 0x80000157 // Gateway not valid
	ErrorBADLOGIN                    uint32 = 0x80000158 // Login Name/Password not valid
	ErrorNOSIODEV                    uint32 = 0x80000159 // Serial device is missing
	ErrorBADSIODEV                   uint32 = 0x8000015A // Serial device not valid
	ErrorDEVINUSE                    uint32 = 0x8000015B // Device already in use
	ErrorDEVMISS                     uint32 = 0x8000015C // Device not present
	ErrorNOMODEM                     uint32 = 0x8000015D // Modem not present
	ErrorNOUNIT                      uint32 = 0x8000015E // No more unit available
	ErrorBADMODEM                    uint32 = 0x8000015F // Modem not valid
	ErrorBADPPP1                     uint32 = 0x80000160 // PPP init error
	ErrorBADPPP2                     uint32 = 0x80000161 // PPP establish error
	ErrorBADSLIP                     uint32 = 0x80000162 // SLIP init error
	ErrorBADPROTO                    uint32 = 0x80000163 // Protocol not valid
	ErrorLNAMEMISS                   uint32 = 0x80000164 // Local name is missing
	ErrorBADHOST                     uint32 = 0x80000165 // Host name can not be set
	ErrorNOHOST                      uint32 = 0x80000166 // Host does not respond
	ErrorBADGATE                     uint32 = 0x80000167 // Gateway name can not be set
	ErrorNOGATE                      uint32 = 0x80000168 // Gateway does not respond
	ErrorBADMOUNT                    uint32 = 0x80000169 // Mount is bad
	ErrorBADTYPE                     uint32 = 0x8000016A // Wrong data type
	ErrorBADUMOUNT                   uint32 = 0x80000170 // Unmount is bad
	ErrorNODELSYS                    uint32 = 0x80000171 // SYS-Module can not be deleted
	ErrorBADPATH                     uint32 = 0x80000172 // Path/Name not valid
	ErrorNOSYS1                      uint32 = 0x80000173 // File MCore is missing
	ErrorNOSYS2                      uint32 = 0x80000174 // File MSys is missing
	ErrorNOFTP                       uint32 = 0x80000175 // No FTP-Server on host
	ErrorIPINUSE                     uint32 = 0x80000176 // IP-Address already in use
	ErrorNOLOGIN                     uint32 = 0x80000177 // Login from server denied
	ErrorNOMCONF                     uint32 = 0x80000178 // No MConfig.ini on device
	ErrorUPDATE                      uint32 = 0x80000179 // Firmware update not allowed
	ErrorBADFPROG                    uint32 = 0x8000017A // Programming flash failed
	ErrorBADVERIFY                   uint32 = 0x8000017B // Verify flash failed
	ErrorBADRANGE                    uint32 = 0x8000017C // Range not valid
	ErrorBADCPU                      uint32 = 0x8000017D // Not for this CPU
	ErrorOLDADDR                     uint32 = 0x8000017E // Address in request is old
	ErrorNOBLOCKDEV                  uint32 = 0x8000017F // Block Device not present
	ErrorBADFORMAT                   uint32 = 0x80000180 // Error in formatting
	ErrorBADREOPEN                   uint32 = 0x80000181 // Error in reopen after formatting
	ErrorBADDOSINI                   uint32 = 0x80000182 // Error in DOS initialization
	ErrorBADBOOTSEC                  uint32 = 0x80000183 // Boot sector not settable
	ErrorNOFORMAT                    uint32 = 0x80000184 // Formatting not allowed
	ErrorOLDVERS1                    uint32 = 0x80000185 // Old software version MCore
	ErrorOLDVERS2                    uint32 = 0x80000186 // Old software version MSys
	ErrorOLDVERS3                    uint32 = 0x80000187 // Old software version IO-Driver
	ErrorBADLOGDEV                   uint32 = 0x80000188 // Log-Device not valid
	ErrorBADMEMINIT                  uint32 = 0x80000189 // Error in Memory-Partition init
	ErrorBADPNCINIT                  uint32 = 0x8000018A // Error in Panic-Handler init
	ErrorBADWDGINIT                  uint32 = 0x8000018B // Error Watchdog-Handler init
	ErrorNOERRTOL                    uint32 = 0x8000018C // SW-Module not error tolerant
	ErrorBADTCKRATE                  uint32 = 0x8000018D // Tick rate too low
	ErrorBADSPAWN                    uint32 = 0x8000018E // Error in task spawn
	ErrorWDOGON                      uint32 = 0x8000018F // HW-Watchdog must be off
	ErrorINVSTATE                    uint32 = 0x80000190 // Action in this state not allowed
	ErrorNODE                        uint32 = 0x80000191 // Bad network/fieldbus node
	ErrorTIMEOUT                     uint32 = 0x80000192 // Timeout in function call
	ErrorSWREBOOT                    uint32 = 0x80000193 // Software reboot
	ErrorWDGREBOOT                   uint32 = 0x80000194 // Watchdog reboot
	ErrorNOONLINE                    uint32 = 0x80000195 // Online change not allowed
	ErrorDRVMISS                     uint32 = 0x80000196 // Driver is missing
	ErrorBADDRV                      uint32 = 0x80000197 // Driver not loadable
	ErrorBADSLOT                     uint32 = 0x80000198 // Station or slot not valid
	ErrorBELOWMIN                    uint32 = 0x80000199 // Number stations below minimum
	ErrorABOVEMAX                    uint32 = 0x8000019A // Number stations above maximum
	ErrorSTATMISS                    uint32 = 0x8000019B // Minimum one station is missing
	ErrorBADSESS                     uint32 = 0x8000019C // Session ID not valid
	ErrorBADUSER                     uint32 = 0x8000019D // Login user name not valid
	ErrorNOPWORD                     uint32 = 0x8000019E // Login password not present
	ErrorBADPWORD                    uint32 = 0x8000019F // Login password not valid
	ErrorBADTIME                     uint32 = 0x800001A0 // Login time frame not valid
	ErrorNOTASK                      uint32 = 0x800001A1 // Task does not exist
	ErrorNODEBUG                     uint32 = 0x800001A2 // Task can not be debugged
	ErrorNOATTACH                    uint32 = 0x800001A3 // Task is not attached
	ErrorNOBKPT                      uint32 = 0x800001A4 // Breakpoint not set
	ErrorNOSYM                       uint32 = 0x800001A5 // Symbol not found
	ErrorTSKMISS                     uint32 = 0x800001A6 // Task name is missing
	ErrorNOSSTEP                     uint32 = 0x800001A7 // Task not on breakpoint
	ErrorHAVECBACK                   uint32 = 0x800001A8 // Callback already attached
	ErrorNOCBACK                     uint32 = 0x800001A9 // Callback not attached
	ErrorHAVEBKPT                    uint32 = 0x800001AA // Breakpoint already set
	ErrorALLBKPT                     uint32 = 0x800001AB // All breakpoints in use
	ErrorVARINUSE                    uint32 = 0x800001AC // Variable already in use
	ErrorMBTRANS                     uint32 = 0x800001AD // Signalising MB transfer in error for compatibility reasons
	ErrorNOBUFF                      uint32 = 0x800001AE // No buffer found
	ErrorBADRESTORE                  uint32 = 0x800001AF // Restoring device failed
	ErrorWEAKPWR                     uint32 = 0x800001B0 // Weak power supply on a station
	ErrorBADCONFIG                   uint32 = 0x800001B1 // Configuration data not valid
	ErrorBADFWARE                    uint32 = 0x800001B2 // Firmware not for this module
	ErrorDOWNGRADE                   uint32 = 0x800001B3 // Firmware down grade error
	ErrorMODMISS                     uint32 = 0x800001B4 // Module is old or missing
	ErrorTNETINUSE                   uint32 = 0x800001B5 // Telnet in use, no redirect
	ErrorNOSHELL                     uint32 = 0x800001B6 // Shell not running, no redirect
	ErrorBADSNTP                     uint32 = 0x800001B7 // SNTP client not started
	ErrorNOAPP                       uint32 = 0x800001B8 // Service not available in boot mode NOAPP
	ErrorNETNB                       uint32 = 0x800001C0 // Fieldbus network number not valid
	ErrorNODEID                      uint32 = 0x800001C1 // Fieldbus network node id, mac id not valid
	ErrorDSIZE                       uint32 = 0x800001C2 // Data size not valid, exceeds max size
	ErrorCMDTAG                      uint32 = 0x800001C3 // Command tag not valid
	ErrorCMDARG                      uint32 = 0x800001C4 // Command argument(s) not valid
	ErrorQUEUEFULL                   uint32 = 0x800001C5 // Command/Event Queue overrun
	ErrorIOBUSFAIL                   uint32 = 0x800001C6 // I/O Bus access fails (bad station)
	ErrorSLCLOGFULL                  uint32 = 0x800001C7 // The logbook of the SLC reached a critical level
	ErrorBROKENWIRE                  uint32 = 0x800001C8 // SERCOS: Broken wire
	ErrorDRIVEMISS                   uint32 = 0x800001C9 // SERCOS: Missing drive
	ErrorSTARTUPTMO                  uint32 = 0x800001CA // SERCOS: Timeout while waiting for startup phase
	ErrorDEAD                        uint32 = 0x800001CB // Alive check failed. Module is dead
	ErrorFALLBACK                    uint32 = 0x800001D0 // Fallback device was used for booting
	ErrorMCONFIGBAK                  uint32 = 0x800001D1 // MConfig.ini not found, used .bak instead
	ErrorCONTIGFILE                  uint32 = 0x800001D2 // Error creating contiguous file
	ErrorEXCEPTION                   uint32 = 0x800001D3 // Exception signal occurred
	ErrorOFFLINE                     uint32 = 0x800001E0 // No connection to device
	ErrorINITVAL                     uint32 = 0x800001E1 // Init value error
	ErrorCHECKVAL                    uint32 = 0x800001E2 // Check value error
	ErrorDRIVE                       uint32 = 0x800001E3 // Drive specific error
	ErrorHARDWARE                    uint32 = 0x800001E4 // Hardware error
	ErrorCOMMUNIC                    uint32 = 0x800001E5 // Communication error
	ErrorNOSETVAL                    uint32 = 0x800001E6 // Cyclic setvalue monitoring error
	ErrorCYCLETIME                   uint32 = 0x800001E7 // Cycle time monitoring error
	ErrorCOLDCLIMATE                 uint32 = 0x800001E8 // A module does not support cold climate
	ErrorACCDENIED                   uint32 = 0x800001E9 // Access denied
	ErrorBADUSERLVL                  uint32 = 0x800001EA // Weak user level (deprecated: replaced by M_E_ACCDENIED)
	ErrorXERR                        uint32 = 0x800001EB // VHD_LST_XERR code information
	ErrorREDULOCK                    uint32 = 0x800001F0 // Write access allowed only on primary CPU
	ErrorRICONN                      uint32 = 0x800001F1 // RI offline
	ErrorRINETREDU                   uint32 = 0x800001F2 // RI not redundant
	ErrorLENGTH                      uint32 = 0x800001F3 // Data length invalid
	ErrorAPPNAME                     uint32 = 0x800001F4 // Application name mismatch
	ErrorBLOCKNB                     uint32 = 0x800001F5 // Number of blocks different
	ErrorBLOCKDIFF                   uint32 = 0x800001F6 // Memory blocks different
	ErrorDEVREDU                     uint32 = 0x800001F7 // Device not redundant
	ErrorDEVNETREDU                  uint32 = 0x800001F8 // Device network not redundant
	ErrorAPPERROR1                   uint32 = 0x800001F9 // Application error1
	ErrorAPPERROR2                   uint32 = 0x800001FA // Application error2
	ErrorAPPERROR3                   uint32 = 0x800001FB // Application error3
	ErrorAPPFATAL                    uint32 = 0x800001FC // Application fatal error
	ErrorAPPDONE                     uint32 = 0x800001FD // Application done error
	ErrorAPPOVLOAD                   uint32 = 0x800001FE // Application overload
	ErrorCYCLEOVLOAD                 uint32 = 0x800001FF // Cycle overload
	ErrorBCHCMD                      uint32 = 0x80000200 // BCH command error
	ErrorPRICONFLICT                 uint32 = 0x80000201 // Primary assignment conflict
	ErrorSWITCHOVER                  uint32 = 0x80000202 // Switchover command
	ErrorAPPCS                       uint32 = 0x80000203 // Application checksum mismatch
	ErrorSTSWITCHOVER                uint32 = 0x80000204 // Self test switchover command
	ErrorSTRUN                       uint32 = 0x80000205 // Self test run
	ErrorSTERROR                     uint32 = 0x80000206 // Self test error
	ErrorPROGRAM                     uint32 = 0x80000207 // Programming error
	ErrorSTID                        uint32 = 0x80000210 // Self test ID not valid
	ErrorSTRAM                       uint32 = 0x80000211 // Self test RAM error
	ErrorSTPLC                       uint32 = 0x80000212 // Self test PLC error
	ErrorSTCOM                       uint32 = 0x80000213 // Self test COM error
	ErrorCRYPT_METADATA_METHOD       uint32 = 0x80000214 // Cipher method not available
	ErrorCRYPT_METADATA_HASH         uint32 = 0x80000215 // Couldn't create cipher hash
	ErrorCRYPT_CRYPT_EXISTS          uint32 = 0x80000216 // No encryption layer found
	ErrorCRYPT_CRYPT_MASTERKEY       uint32 = 0x80000217 // No masterkey available
	ErrorCRYPT_CRYPT_BOOTDEV         uint32 = 0x80000218 // Boot partition can not be encrypted
	ErrorCRYPT_ALGORITHM_LOAD_LIB    uint32 = 0x80000219 // Loading cipher library failed
	ErrorCRYPT_ALGORITHM_LIB_VERSION uint32 = 0x80000221 // Cipher library version
	ErrorCRYPT_ALGORITHM_INIT_LIB    uint32 = 0x80000222 // Init cipher library failed
	ErrorCRYPT_ALGORITHM_KEYLENGTH   uint32 = 0x80000223 // Cipher strength not valid
	ErrorCRYPT_ALGORITHM_INIT        uint32 = 0x80000224 // Init cipher algorithm failed
	ErrorCRYPT_ALGORITHM_CRYPT       uint32 = 0x80000225 // Error encrypting data
	ErrorCRYPT_ALGORITHM_DECRYPT     uint32 = 0x80000226 // Error decrypting data
	ErrorCRYPT_KEY_RECOVERYFILE      uint32 = 0x80000227 // Recover key file not found
	ErrorCRYPT_KEY_RECOVERYFILECRC   uint32 = 0x80000228 // Recovery key file crc error
	ErrorCRYPT_KEY_RECOVERYPASSWORD  uint32 = 0x80000229 // Recovery password wrong
	ErrorCRYPT_NOPROGRESS            uint32 = 0x80000230 // No Progress available
	ErrorCRYPT_UNKNOWN               uint32 = 0x80000231 // Detailed reason unknown
	ErrorBADSNTPSERVER               uint32 = 0x80000232 // Configuration of SNTP server not valid
	ErrorNOONLINECFG                 uint32 = 0x80000233 // Online re-config not supported
	ErrorBADCORECATEG                uint32 = 0x80000234 // Configuration of CoreCategory not valid
	ErrorPWDEXPIRED                  uint32 = 0x80000235 // Password expired
	ErrorBADLOGINLOCKED              uint32 = 0x80000236 // User is disabled
	ErrorBUSY                        uint32 = 0x80000237 // Resource is currently busy
	ErrorTASK_WDOG                   uint32 = 0x80000240 // Error creating watchdog
	ErrorTASK_SEM                    uint32 = 0x80000241 // Error creating/deleting semaphore
	ErrorTASK_PRIO                   uint32 = 0x80000242 // Priority out of range (has to be between 1 and 255)
	ErrorTASK_OFFSET                 uint32 = 0x80000243 // Offset does not match tick/snyc rate or is greater than cycle time
	ErrorTASK_CYCDIFF                uint32 = 0x80000244 // Next possible cycle time does differ more than 25% from configured cycle time
	ErrorTASK_CYCTIME                uint32 = 0x80000245 // Cycle time is too small
	ErrorTASK_ATTACH                 uint32 = 0x80000246 // Error attaching aux/sync
	ErrorTASK_AUXOFF                 uint32 = 0x80000247 // Auxiliary clock is turned off
	ErrorTASKSource                  uint32 = 0x80000248 // Unknown task trigger source
	ErrorTASK_NOMAINFUNC             uint32 = 0x80000249 // User main function is NULL
	ErrorTASK_NOPTP                  uint32 = 0x8000024A // PTP synchronization not activated
	ErrorTASK_CORECATEG              uint32 = 0x8000024B // Core category not valid
	ErrorTASK_RUNNING                uint32 = 0x8000024C // Task has not been stopped
	ErrorTASK_CHANGEATTR             uint32 = 0x8000024D // Selected attributes must not be changed
	ErrorRTC                         uint32 = 0x80000300 // Error init RTC chip
	ErrorETHER                       uint32 = 0x80000301 // Error init Ethernet controller
	ErrorIP_ADDR                     uint32 = 0x80000302 // IP-Address not valid
	ErrorTICKRATE                    uint32 = 0x80000303 // Tick rate not valid
	ErrorCPUCLOCK                    uint32 = 0x80000304 // CPU clock not valid
	ErrorNVMAKE                      uint32 = 0x80000305 // NV-Ram create failed
	ErrorNVMISS                      uint32 = 0x80000306 // NV-Ram not present
	ErrorNVEMPTY                     uint32 = 0x80000307 // NV-Ram is empty
	ErrorNVNODOS                     uint32 = 0x80000308 // NV-Ram no DOS file system
	ErrorRDMAKE                      uint32 = 0x80000309 // Ram-Disk create failed
	ErrorBATTLOW                     uint32 = 0x8000030A // Battery voltage too low
	ErrorSIO                         uint32 = 0x8000030D // Error init SIO driver
	ErrorDUPIP                       uint32 = 0x8000030E // Duplicate ip address
	ErrorEHDOFLOW                    uint32 = 0x8000030F // Too much EHD entries in MCore
	ErrorBOOTPANIC                   uint32 = 0x80000310 // Panic situation occurred during boot
	ErrorNVRAMSIG                    uint32 = 0x80000311 // Invalid nvram signature
	ErrorNATFAIL                     uint32 = 0x80000312 // NAT initialization failed
	ErrorSMART                       uint32 = 0x80000313 // SMART self-test on CFC failed
	ErrorFILESYSTEM                  uint32 = 0x80000314 // Error in DOS file system
	ErrorSHORTKEY                    uint32 = 0x80000315 // MConfig.ini keyword too short
	ErrorMSYSINVALID                 uint32 = 0x80000316 // MSys is invalid (checksum, not found, ...)
	ErrorSNTPSYNC                    uint32 = 0x80000317 // SNTP synchronization failure
	ErrorBADSIOHANDSHAKE             uint32 = 0x80000318 // Hardware handshake is not permitted (probably device is used as console)
	ErrorALLREADYEXIST               uint32 = 0x80002001 // Object already exist
	ErrorNOOBJECT                    uint32 = 0x80002002 // Object not found
	ErrorSQLERROR                    uint32 = 0x80002003 // SQL error occurred
	ErrorNOAHD                       uint32 = 0x80002004 // SW-Module AHD is not installed
	ErrorNOSESSAV                    uint32 = 0x80002005 // No session available
)
