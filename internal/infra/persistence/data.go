package persistence

import (
	"log"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmysql"

	"github.com/comeonjy/account/configs"
	"github.com/comeonjy/account/internal/domain/entry"
)

var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

type Data struct {
	Account *gorm.DB
}

func NewData(cfg configs.Interface, logger *xlog.Logger) *Data {
	return &Data{
		Account: newAccountMysql(cfg, logger),
	}
}

func newAccountMysql(cfg configs.Interface, logger *xlog.Logger) *gorm.DB {
	db := xmysql.New(cfg.Get().MysqlConf, logger)
	if err := db.AutoMigrate(&entry.UserModel{}); err != nil {
		log.Fatalln("AutoMigrate AccountModel err:", err)
	}
	return db
}
