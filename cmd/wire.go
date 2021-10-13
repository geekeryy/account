//go:build wireinject
// +build wireinject

package cmd

import (
	"context"

	"account/configs"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/google/wire"

	"account/internal/data"
	"account/internal/server"
	"account/internal/service"
)

func InitApp(ctx context.Context,logger *xlog.Logger) *App {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		newApp,
		configs.ProviderSet,
		data.ProviderSet,
	))
}
