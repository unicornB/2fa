package api

import (
	"2fa.com/model"
	"2fa.com/serializer"
	"2fa.com/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBindJSON(&userService); err == nil {
		res := userService.Login(c)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(200, serializer.Error("参数错误"))
	}
}
func UserSendEmailCode(c *gin.Context) {
	var userSendMailService service.UserSendMailService
	if err := c.ShouldBindJSON(&userSendMailService); err == nil {
		res := userSendMailService.SendEmail(c)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(200, serializer.Error("参数错误"))
	}
}
func UserGetMe(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	c.JSON(200, user)
}
