package persistence

import (
	"context"
	"errors"

	"github.com/comeonjy/account/internal/domain/entry"
	"github.com/comeonjy/account/internal/domain/repository"
	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(data *Data) repository.AccountRepo {
	return &accountRepo{db: data.Account}
}

func (r accountRepo) Get(ctx context.Context, id uint64) (*entry.UserModel, error) {
	user := entry.UserModel{}
	err := r.db.WithContext(ctx).Model(&entry.UserModel{}).Find(&user, "a").Error
	return &user, err
}
func (r accountRepo) GetByWechatOpenid(ctx context.Context, user *entry.UserModel) error {
	if err := r.db.WithContext(ctx).Model(&entry.UserModel{}).Where("wechat_openid = ?", user.WechatOpenid).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
func (r accountRepo) GetByMobile(ctx context.Context, user *entry.UserModel) error {
	if err := r.db.WithContext(ctx).Model(&entry.UserModel{}).Where("mobile = ?", user.Mobile).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

func (r accountRepo) GetByEmail(ctx context.Context, user *entry.UserModel) error {
	if err := r.db.WithContext(ctx).Model(&entry.UserModel{}).Where("email = ?", user.Email).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

func (r accountRepo) GetByAccount(ctx context.Context, account string) (*entry.UserModel, error) {
	user := entry.UserModel{}
	if err := r.db.WithContext(ctx).Model(&entry.UserModel{}).Where("mobile=? or email=? ", account, account).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func (r accountRepo) Create(ctx context.Context, user *entry.UserModel) error {
	return r.db.WithContext(ctx).Model(&entry.UserModel{}).Create(&user).Error
}
func (r accountRepo) Updates(ctx context.Context, user *entry.UserModel) error {
	return r.db.WithContext(ctx).Model(&entry.UserModel{}).Where("uuid = ?", user.UUID).Updates(&user).Error
}
