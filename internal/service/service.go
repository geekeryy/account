package service

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc/metadata"

	v1 "account/api/v1"
	"account/configs"
	"account/internal/data"
	"github.com/comeonjy/go-kit/pkg/xlog"
)

var ProviderSet = wire.NewSet(NewAccountService)

type AccountService struct {
	v1.UnimplementedAccountServer
	conf     configs.Interface
	logger   *xlog.Logger
	workRepo data.WorkRepo
}

func NewAccountService(conf configs.Interface, logger *xlog.Logger, workRepo data.WorkRepo) *AccountService {
	return &AccountService{
		conf:     conf,
		workRepo: workRepo,
		logger:   logger,
	}
}

func (svc *AccountService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *AccountService) Ping(ctx context.Context, in *v1.Empty) (*v1.Result, error) {
	return &v1.Result{
		Code:    200,
		Message: "pong",
	}, nil
}
