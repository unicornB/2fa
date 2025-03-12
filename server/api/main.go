package api

import (
	"gongniu/serializer"
	"gongniu/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var service service.LoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Error("参数错误"))
	}

}

func GetBizData(c *gin.Context) {
	var service service.BizDataService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetBizData(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Error("参数错误"))
	}
}
func UpdateBizData(c *gin.Context) {
	var service service.BizDataService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateBizData(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.Error("参数错误"))
	}
}
