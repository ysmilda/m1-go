package res

import (
	"net"
	"time"

	"github.com/ysmilda/m1-go/modules/msys"
)

type (
	UserAuth struct {
		Toolname       string `binary:"length:12"`
		Username       string `binary:"length:64"`
		LoginTimestamp msys.Timestamp
		LastAlive      msys.Timestamp
		IPAddress      net.IP `binary:"length:4"`
		Active         bool
		Local          bool
		ServerPort     uint16
		LastActivity   msys.Timestamp `binary:"skip:8"`
		LoginSessionID uint32
	}

	// UserAccess contains information about the user access.
	UserAccess struct {
		Group                      uint8
		Level                      uint8
		Priority                   uint8 `binary:"skip:1"`
		SystemPermissions          Permissions
		ApplicationPermissions     Permissions
		AppData                    int32
		PasswordValidityTime       time.Duration `binary:"length:4,unit:days"`
		DaysTillPasswordExpiration time.Duration `binary:"length:4,unit:days,skip:4"`
	}

	UserData struct {
		Access UserAccess
		Auth   UserAuth
	}

	UserAttributes struct {
		Group                  uint8
		Level                  uint8
		Priority               uint8
		Disabled               bool
		SystemPermissions      Permissions
		ApplicationPermissions Permissions
		ValidFrom              time.Time     `binary:"length:4,unit:seconds"`
		ValidUntil             time.Time     `binary:"length:4,unit:seconds"`
		PasswordValidityTime   time.Duration `binary:"length:4,unit:days,skip:16"`
	}

	UserInfo struct {
		IsLocked bool `binary:"skip:15"`
	}

	User struct {
		Name       string `binary:"length:64"`
		Attributes UserAttributes
		Info       UserInfo
	}
)
