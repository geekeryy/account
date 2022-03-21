package test

import (
	"context"
	"testing"

	"github.com/comeonjy/go-kit/grpc/reloadconfig"
	"google.golang.org/grpc"
)

func TestService_Ping(t *testing.T) {
	dial, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	_, err = reloadconfig.NewReloadConfigClient(dial).ReloadConfig(context.Background(), &reloadconfig.Empty{})
	if err != nil {
		t.Error(err)
		return
	}
}
