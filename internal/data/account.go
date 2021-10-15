package data

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id        uint64 `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

type accountRepo struct {
	db *gorm.DB
}

func (r accountRepo) Get(ctx context.Context,id uint64) (*UserModel, error) {
	user := UserModel{}
	err := r.db.WithContext(ctx).Model(&UserModel{}).Find(&user, "a").Error
	return &user, err
}
