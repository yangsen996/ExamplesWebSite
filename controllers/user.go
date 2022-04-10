package controllers

import (
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"github.com/yangsen996/ExamplesWebSite/comm"
	"github.com/yangsen996/ExamplesWebSite/model"
	"github.com/yangsen996/ExamplesWebSite/server"
	"github.com/yangsen996/ExamplesWebSite/utils"
	"net/http"
)

type userController struct{}

func NewUserController() *userController {
	return &userController{}
}

func (uc *userController) Login(c *gin.Context) {
	var l comm.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusOK, "参数不正确")
		return
	}
	u := model.User{
		UserName: l.Username,
		Password: l.Password,
	}
	if err, user := server.NewUserService().Login(&u); err != nil {
		c.JSON(http.StatusOK, "用户名或密码不正确")
		return
	} else {
		//todo token
		uc.GetToken(c, *user)
	}

}
func (uc *userController) Register(c *gin.Context) {
	var r comm.RegisterReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusOK, "参数不正确")
		return
	}
	//加密
	pwd := utils.MD5([]byte(r.Password))
	//uuid
	uuid := uuid2.NewV4()

	user := model.User{
		UUID:     uuid,
		UserName: r.Username,
		Password: pwd,
	}
	err := server.NewUserService().Register(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, "注册成功")
	}
}
func (uc *userController) GetToken(c *gin.Context, u model.User) {

}
