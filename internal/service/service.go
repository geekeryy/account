package service

import (
	"context"
	"encoding/json"
	"errors"

	v1 "account/api/v1"
	"account/configs"
	"account/internal/data"
	"account/pkg/redis"
	"account/pkg/wechat"
	"account/pkg/yunpian"
	"github.com/comeonjy/go-kit/pkg/xjwt"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/google/wire"
	"google.golang.org/grpc/metadata"
)

var ProviderSet = wire.NewSet(NewAccountService)

type AccountService struct {
	v1.UnimplementedAccountServer
	conf        configs.Interface
	logger      *xlog.Logger
	accountRepo data.AccountRepo
	mini        *wechat.Mini
	sms         *yunpian.Client
	redis       *redis.Client
}

func NewAccountService(conf configs.Interface, accountRepo data.AccountRepo, logger *xlog.Logger) *AccountService {
	xjwt.Init(conf.Get().JwtKey)
	return &AccountService{
		conf:        conf,
		accountRepo: accountRepo,
		logger:      logger,
		mini:        wechat.NewMini(conf.Get().WechatMiniAppid, conf.Get().WechatMiniSecret),
		sms:         yunpian.NewClient(conf.Get().YunpianApiKey),
		redis:       redis.NewClient(conf.Get().MysqlConf),
	}
}

func (svc *AccountService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *AccountService) Ping(ctx context.Context, in *v1.Empty) (*v1.Result, error) {
	var msg string
	bus, err := svc.getCurrentUser(ctx)
	marshal, err := json.Marshal(bus)
	if err != nil {
		msg = err.Error()
	} else {
		msg = string(marshal)
	}
	return &v1.Result{
		Code:    200,
		Message: msg,
	}, nil
}

func (svc *AccountService) getCurrentUser(ctx context.Context) (*xjwt.Business, error) {
	bus := xjwt.Business{}
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		if tokens := mdIn.Get("Token"); len(tokens) > 0 {
			err := xjwt.ParseToken(tokens[0], &bus)
			return &bus, err
		}
	}
	return nil, errors.New("Token not found ")
}
