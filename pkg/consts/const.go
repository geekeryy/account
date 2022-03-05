package consts

import (
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xerror"
)

var EnvMap = map[string]string{
	xenv.AppName:     "account",
	xenv.AppVersion:  "v1.0",
	xenv.ApolloAppID: "account",
	xenv.ApolloUrl:   "http://apollo.dev.jiangyang.me",
	"my_const":       "my_const_value",
}

const (
	myErr xerror.Code = 11001 // 自定义错误
)
