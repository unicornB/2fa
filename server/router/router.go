package router

import (
	"gongniu/api"
	"gongniu/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors())
	//日志请求
	r.Use(middleware.ReqLogMiddleware())
	//签名验证
	r.Use(middleware.SignatureMiddleware(os.Getenv("SIGN_SECRET_KEY")))
	r.POST("/login", api.Login)
	r.GET("/bizData/:bizKey", api.GetBizData)
	r.PUT("/bizData/UpdateBizData", api.UpdateBizData)
	r.GET("/admin/userlist", api.UserList)
	r.PUT("/admin/userpass", api.UserEditPass)
	r.POST("/admin/useradd", api.UserAdd)
	r.DELETE("/admin/userdel/:id", api.UserDel)
	// 路由分组
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", api.UserLogin)
		v1.Use(middleware.CurrentUser())
		auth := v1.Group("")
		auth.Use(middleware.AuthUserRequired())
		{
			auth.GET("/user", api.UserInfo)
		}
	}
	return r
}
