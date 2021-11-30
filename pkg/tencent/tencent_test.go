package tencent_test

import (
	"context"
	"testing"

	"github.com/comeonjy/account/configs"
	"github.com/comeonjy/account/pkg/tencent"
)

func TestNewTenSms(t *testing.T) {
	c:=configs.NewConfig(context.Background())
	if err:=tencent.NewTenSms(c).SendLoginCode("15881315861","1234");err!=nil{
		t.Error(err)
	}
}