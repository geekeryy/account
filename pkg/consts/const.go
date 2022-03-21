package consts

import (
	"github.com/comeonjy/go-kit/pkg/xenv"
)

var EnvMap = map[string]string{
	xenv.AppName:     "account",
	xenv.AppVersion:  "v1.0",
	xenv.ApolloAppID: "account",
}

// Redis Key

const (
	SmsLoginCode = "sms_login_code:%s"
)
