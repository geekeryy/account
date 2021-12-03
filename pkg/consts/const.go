package consts

import (
	"github.com/comeonjy/go-kit/pkg/xenv"
)

var EnvMap = map[string]string{
	xenv.AppName:     "account",
	xenv.AppVersion:  "v1.0",
	xenv.TraceName:   "trace_id",
	xenv.ApolloAppID: "account",
	xenv.ApolloUrl:   "http://apollo.dev.jiangyang.me",
	"my_const":       "my_const_value",
}

