package consts

// 常量
const (
	AppName    = "account"
	AppVersion = "v1.0"

	TraceName = "trace_id" // trace名
)

// apollo 配置
const (
	ApolloAppID     = "account"
	ApolloCluster   = "default"
	ApolloNamespace = "application"
)

// 环境变量名
const (
	AppEnv       = "APP_ENV"
	ApolloUrl    = "APOLLO_URL"
	ApolloSecret = "APOLLO_ACCESS_KEY_SECRET"
	ConfigFile   = "CONFIG_FILE"
)

var EnvMap = map[string]string{
	AppEnv:       "dev",
	ApolloUrl:    "http://apollo.dev.jiangyang.me",
	ApolloSecret: "",
	ConfigFile:   "./configs/config.yaml",
}
