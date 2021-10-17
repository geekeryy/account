package data

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id           uint64 `gorm:"primarykey"`
	UUID         string `gorm:"uniqueIndex;type:varchar(36);not null"`
	WechatOpenid string `gorm:"type:varchar(36)"`
	NickName     string`gorm:"type:varchar(50)"`
	Role         uint64`gorm:"type:uint"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

type accountRepo struct {
	db *gorm.DB
}

func (r accountRepo) Get(ctx context.Context, id uint64) (*UserModel, error) {
	user := UserModel{}
	err := r.db.WithContext(ctx).Model(&UserModel{}).Find(&user, "a").Error
	return &user, err
}
func (r accountRepo) GetByWechatOpenid(ctx context.Context, user *UserModel) error {
	if err:=r.db.WithContext(ctx).Model(&UserModel{}).Where("wechat_openid = ?", user.WechatOpenid).Take(&user).Error;err!=nil && !errors.Is(err,gorm.ErrRecordNotFound){
		return err
	}
	return nil
}
func (r accountRepo) Create(ctx context.Context, user *UserModel) error {
	return r.db.WithContext(ctx).Model(&UserModel{}).Create(&user).Error
}
