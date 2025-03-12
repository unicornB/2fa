package api

import (
	"gongniu/serializer"
	"gongniu/service"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Error("参数错误"))
	}
}
func UserInfo(c *gin.Context) {
	var service service.UserService
	res := service.UserInfo(c)
	c.JSON(200, res)
}
func UserList(c *gin.Context) {
	var service service.UserService
	res := service.UserList(c)
	c.JSON(200, res)
}
func UserEditPass(c *gin.Context) {
	var service service.UserEditService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserEditPass(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Error("参数错误"))
	}
}
func UserAdd(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserAdd(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Error(err.Error()))
	}
}
func UserDel(c *gin.Context) {
	var service service.UserService
	res := service.UserDelete(c)
	c.JSON(200, res)
}
