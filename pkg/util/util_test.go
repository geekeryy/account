// Package util @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/11/28 8:04 下午
package util_test

import (
	"fmt"
	"testing"

	"github.com/comeonjy/account/pkg/util"
)

func TestMd5(t *testing.T) {
	fmt.Println(util.Md5("123456"))
}
