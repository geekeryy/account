package data

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"

	"account/configs"
	"github.com/comeonjy/go-kit/pkg/xmongo"
)

var ProviderSet = wire.NewSet(NewMongo, NewData, NewWorkRepo)

type Data struct {
	Mongo *mongo.Collection
}

func NewMongo(cfg configs.Interface) *mongo.Collection {
	xmongo.Init(xmongo.Config{
		Username: cfg.Get().MongoUsername,
		Password: cfg.Get().MongoPassword,
		Addr:     cfg.Get().MongoAddr,
		Database: cfg.Get().MongoDatabase,
	})
	return xmongo.Conn("user")
}

func NewData(client *mongo.Collection) *Data {
	return &Data{
		Mongo: client,
	}
}
