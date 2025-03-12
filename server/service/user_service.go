package service

import (
	"gongniu/model"
	"gongniu/serializer"
	"gongniu/utils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	ID       uint   `form:"id" json:"id"`
	UserName string `form:"username" json:"username" binding:"min=3,max=30"`
	Password string `form:"password" json:"password" binding:"min=5,max=40"`
}

func (service *UserService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("username = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.Error("账户不存在")
	}
	if !user.CheckPassword(service.Password) {
		return serializer.Error("密码错误")
	}
	token, err := utils.GenerateJWT(user.Username, os.Getenv("JWT_SECRET_USER"), uint(user.ID))
	if err != nil {

		return serializer.Error("账号或密码错误")
	}
	return serializer.Success("登录成功", token)
}

func (service *UserService) UserInfo(c *gin.Context) serializer.Response {
	user, _ := c.Get("user")
	newUser := user.(*model.User)
	newUser.Password = ""
	return serializer.Success("获取用户信息成功", newUser)
}
func (service *UserService) UserList(c *gin.Context) serializer.Response {
	var users []model.User
	//分页查询
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * 10
	if err := model.DB.Limit(10).Offset(offset).Find(&users).Error; err != nil {
		return serializer.Error("查询用户列表失败")
	}
	return serializer.Success("查询用户列表成功", users)
}

type UserEditService struct {
	ID       uint   `form:"id" json:"id"`
	Password string `form:"password" json:"password" binding:"min=5,max=40"`
}

// 修改密码
func (service *UserEditService) UserEditPass(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("id =?", service.ID).First(&user).Error; err != nil {
		return serializer.Error("账户不存在")
	}
	user.SetPassword(service.Password)
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.Error("修改密码失败")
	}
	return serializer.Success("修改密码成功", nil)
}

// 添加用户
func (service *UserService) UserAdd(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("username =?", service.UserName).First(&user).Error; err == nil {
		return serializer.Error("账户已存在")
	}
	user.Username = service.UserName
	user.SetPassword(service.Password)
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Error("添加用户失败")
	}
	return serializer.Success("添加用户成功", nil)
}
func (service *UserService) UserDelete(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("id =?", c.Param("id")).First(&user).Error; err != nil {
		return serializer.Error("账户不存在")
	}
	if err := model.DB.Delete(&user).Error; err != nil {
		return serializer.Error("删除用户失败")
	}
	return serializer.Success("删除用户成功", nil)
}
