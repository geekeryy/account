// Package entry @Description  TODO
// @Author  	 jiangyang
// @Created  	 2022/3/20 3:39 PM
package entry

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id           uint64 `gorm:"primarykey"`
	UUID         string `gorm:"uniqueIndex;type:varchar(36);not null"`
	Mobile       string `gorm:"uniqueIndex;type:varchar(11)"`
	Email        string `gorm:"uniqueIndex;type:varchar(200)"`
	Password     string `gorm:"type:varchar(200)"`
	WechatOpenid string `gorm:"type:varchar(36)"`
	NickName     string `gorm:"type:varchar(50)"`
	AvatarUrl    string `gorm:"type:varchar(200)"`
	Role         uint64 `gorm:"type:uint"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
