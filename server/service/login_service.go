package service

import (
	"gongniu/model"
	"gongniu/serializer"
	"gongniu/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=40"`
}

func (service *LoginService) Login(c *gin.Context) serializer.Response {
	var admin model.Admin
	if err := model.DB.Where("username = ?", service.UserName).First(&admin).Error; err != nil {
		return serializer.Error("账户不存在")
	}
	if !admin.CheckPassword(service.Password) {
		return serializer.Error("密码错误")
	}
	token, err := utils.GenerateJWT(admin.Username, os.Getenv("JWT_SECRET"), uint(admin.Id))
	if err != nil {

		return serializer.Error("账号或密码错误")
	}
	return serializer.Success("登录成功", token)
}
