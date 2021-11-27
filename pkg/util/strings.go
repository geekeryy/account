// Package util @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/11/27 10:18 下午
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
