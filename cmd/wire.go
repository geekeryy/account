//go:build wireinject

package cmd

import (
	"context"

	"github.com/comeonjy/account/configs"
	"github.com/comeonjy/account/internal/infra/persistence"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/google/wire"

	"github.com/comeonjy/account/internal/server"
	"github.com/comeonjy/account/internal/service"
)

func InitApp(ctx context.Context, logger *xlog.Logger) *App {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		newApp,
		configs.ProviderSet,
		persistence.ProviderSet,
	))
}
