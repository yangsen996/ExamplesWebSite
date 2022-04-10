package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Model
	UUID           uuid.UUID `json:"uuid" gorm:"comment:用户uuid"`
	UserName       string    `json:"user_name" gorm:"not null;comment:用户登录名"`
	Password       string    `json:"-"  gorm:"not null;comment:用户登录密码"` // 用户登录密码
	Phone          string    `json:"phone" gorm:"not null;default:'';uniqueIndex:idx_phone;comment:用户手机号"`
	Email          string    `json:"email" gorm:"not null;default:'';comment:用户邮箱"`
	EmailValidated int       `json:"email_validated" gorm:"type:tinyint(1);comment:邮箱是否验证"`
	PhoneValidated int       `json:"phone_validated" gorm:"type:tinyint(1);comment:手机是否验证"`
	LasActive      time.Time `json:"las_active" gorm:"comment:最后活跃时间"`
	Profile        string    `json:"profile" gorm:"type:text;comment:用户属性"`
	Status         int       `json:"status" gorm:"type:int;not null;default:0;index:idx_status;comment:账户状态(启用/禁用/锁定/标记删除)"`
}

func (User) TableName() string {
	return "user"
}
