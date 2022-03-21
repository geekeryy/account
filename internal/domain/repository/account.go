// Package repository @Description  TODO
// @Author  	 jiangyang
// @Created  	 2022/3/20 3:38 PM
package repository

import (
	"context"

	"github.com/comeonjy/account/internal/domain/entry"
)

type AccountRepo interface {
	Get(ctx context.Context, id uint64) (*entry.UserModel, error)
	GetByWechatOpenid(ctx context.Context, user *entry.UserModel) error
	Create(ctx context.Context, user *entry.UserModel) error
	Updates(ctx context.Context, user *entry.UserModel) error
	GetByMobile(ctx context.Context, user *entry.UserModel) error
	GetByEmail(ctx context.Context, user *entry.UserModel) error
	GetByAccount(ctx context.Context, account string) (*entry.UserModel, error)
}
