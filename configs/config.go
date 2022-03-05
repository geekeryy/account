package configs

import (
	"context"
	"encoding/json"
	"log"

	"github.com/comeonjy/go-kit/pkg/xconfig"
	"github.com/comeonjy/go-kit/pkg/xconfig/apollo"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewConfig)

type Config struct {
	Mode                 string `json:"mode"`
	GrpcAddr             string `json:"grpc_addr"`
	HttpAddr             string `json:"http_addr"`
	PprofAddr            string `json:"pprof_addr"`
	ApmUrl               string `json:"apm_url"`
	MysqlConf            string `json:"mysql_conf"`
	TenSecretId          string `json:"ten_secret_id"`
	TenSecretKey         string `json:"ten_secret_key"`
	TenSmsConf           string `json:"ten_sms_conf"`
	WechatMiniAppid      string `json:"wechat_mini_appid"`
	WechatMiniSecret     string `json:"wechat_mini_secret"`
	JwtKey               string `json:"jwt_key"`
	YunpianApiKey        string `json:"yunpian_api_key"`
	RedisOption          string `json:"redis_option"`
	Email                string `json:"email"`

}

func (c *Config) Validate() error {
	return nil
}

// Interface 对外暴露接口（用于功能扩展）
type Interface interface {
	Get() Config
	xconfig.ReloadConfigInterface
}

// 内部配置类载体
type config struct {
	xconfig.IConfig
}

// Get 获取配置
func (c *config) Get() Config {
	return c.LoadValue().(Config)
}

func (c *config) ReloadConfig() error {
	return c.ReLoad()
}

func storeHandler(data []byte) interface{} {
	conf := Config{}
	if err := json.Unmarshal(data, &conf); err != nil {
		log.Println(err)
	}
	if err := conf.Validate(); err != nil {
		log.Println(err)
	}
	return conf
}

// NewConfig 获取配置
func NewConfig(ctx context.Context) Interface {
	cfg := &config{
		xconfig.New(ctx,
			apollo.NewSource(xenv.GetEnv(xenv.ApolloUrl), xenv.GetEnv(xenv.ApolloAppID), xenv.GetApolloCluster("default"), xenv.GetApolloSecret(), xenv.GetApolloNamespace("grpc"), xenv.GetApolloNamespace("common")),
			storeHandler),
	}
	return cfg
}
