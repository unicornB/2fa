package router

import (
	"gongniu/api"
	"gongniu/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	//静态资源配置
	r.StaticFile("/", "./static/h5/index.html") // 单页应用的入口通常是 index.html
	r.Static("/assets", "./static/h5/assets")   // 根据需要配置其他静态资源路径
	r.Static("/static", "./static/h5/static")   // 根据需要配置其他静态资源路径

	r.StaticFile("/pc", "./static/pc/index.html") // 单页应用的入口通常是 index.html
	r.Static("/pc/css", "./static/pc/css")        // 根据需要配置其他静态资源路径
	r.Static("/pc/js", "./static/pc/js")          // 根据需要配置其他静态资源路径
	r.Static("/pc/fonts", "./static/pc/fonts")    // 根据需要配置其他静态资源路径
	r.Static("/pc/img", "./static/pc/img")        // 根据需要配置其他静态资源路径
	r.NoRoute(func(c *gin.Context) {              // 处理所有未匹配的路由到 index.html，实现前端路由（SPA）
		c.File("./static/h5/index.html") // 提供 index.html 文件，由前端路由处理后续路径
		c.File("./static/pc/index.html") // 提供 index.html 文件，由前端路由处理后续路径
	})
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
