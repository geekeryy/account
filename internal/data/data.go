package data

import (
	"account/configs"
	"github.com/comeonjy/go-kit/pkg/xmysql"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

type Data struct {
	Account *gorm.DB
}

func newAccountMysql(cfg configs.Interface) *gorm.DB {
	return xmysql.New(cfg.Get().MysqlConf)
}

func NewData(cfg configs.Interface) *Data {
	return &Data{
		Account: newAccountMysql(cfg),
	}
}
