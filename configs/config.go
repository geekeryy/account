package configs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"

	"account/pkg/consts"
	"github.com/comeonjy/go-kit/pkg/xconfig"
	"github.com/comeonjy/go-kit/pkg/xconfig/apollo"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewConfig)

var _cfg atomic.Value

type Interface interface {
	Get() Config
}

func (Config) Get() Config {
	return _cfg.Load().(Config)
}

func NewConfig(ctx context.Context) Interface {
	c := xconfig.New(
		xconfig.WithContext(ctx),
		xconfig.WithSource(apollo.NewSource(xenv.GetEnv(consts.ApolloUrl), consts.ApolloAppID, consts.ApolloCluster, consts.ApolloNamespace, xenv.GetEnv(consts.ApolloSecret))),
		xconfig.WithDecoder(json.Unmarshal),
	)
	var tempConf Config
	if err := c.Scan(&tempConf); err != nil {
		panic(fmt.Sprintf("config scan: %v", err))
	}
	_cfg.Store(tempConf)

	if err := c.Watch(func(config *xconfig.Config) {
		if err := config.Scan(&tempConf); err != nil {
			log.Println("config watch exit:", err)
			return
		}
		if err := tempConf.Validate(); err != nil {
			log.Println("invalid config value:", err)
			return
		}
		_cfg.Store(tempConf)
	}); err != nil {
		panic(fmt.Sprintf("config watch: %v", err))
	}

	return _cfg.Load().(Config)
}
