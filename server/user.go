package server

import (
	"errors"
	"github.com/yangsen996/ExamplesWebSite/global"
	"github.com/yangsen996/ExamplesWebSite/model"
	"github.com/yangsen996/ExamplesWebSite/utils"
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

// Register 注册
func (userService *userService) Register(u model.User) (err error) {
	var user model.User
	if res := global.G_DB.Where("user_name=?", u.UserName).Find(&user).RowsAffected; res != 0 {
		return errors.New("用户已注册")
	}
	err = global.G_DB.Create(&u).Error
	return err
}

//Login 登录
func (userService *userService) Login(u *model.User) (err error, userInfo *model.User) {
	var user model.User

	pwd := utils.MD5([]byte(u.Password))
	err = global.G_DB.Where("username=? and password=?", u.UserName, pwd).First(&user).Error
	return err, &user
}
