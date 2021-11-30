// Package yunpian @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/11/27 9:29 下午
package yunpian_test

import (
	"context"
	"testing"

	"github.com/comeonjy/account/configs"
	"github.com/comeonjy/account/pkg/yunpian"
)

func TestSendCode(t *testing.T) {
	cfg := configs.NewConfig(context.Background())
	if err := yunpian.NewClient(cfg.Get().YunpianApiKey).SendCode("15881314861",1234);err != nil {
		t.Error(err)
	}
}
