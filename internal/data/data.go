package data

import (
	"context"
	"log"

	"account/configs"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmysql"
	"github.com/google/wire"
	"gorm.io/gorm"
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
	if err := db.AutoMigrate(&UserModel{}); err != nil {
		log.Fatalln("AutoMigrate AccountModel err:", err)
	}
	return db
}

type AccountRepo interface {
	Get(ctx context.Context,id uint64) (*UserModel, error)
	GetByWechatOpenid(ctx context.Context,user *UserModel) error
	Create(ctx context.Context,user *UserModel) error
	Updates(ctx context.Context, user *UserModel) error
}

func NewAccountRepo(data *Data) AccountRepo {
	return &accountRepo{db: data.Account}
}
