package router

import (
	"CRAZY/middleware"
	"CRAZY/router/api"

	"github.com/gin-gonic/gin"
)

// Routers 总路由
func Routers() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 设置静态目录
	r.Static("/public", "./public")

	// 设置模板
	r.LoadHTMLGlob("views/**/*")

	// 首页
	r.GET("/", api.GetHTML)

	// 登陆
	r.GET("/login", api.Login)

	r.GET("/p", api.GetNumber)

	// 用户
	r.GET("/user/:id", api.GetUserById)

	// api 部分
	apiRouter := r.Group("/api")
	apiRouter.Use(middleware.JWTAuth())
	{
		// 获取配置
		apiRouter.GET("/config", api.GetConfig)
		apiRouter.GET("/ping", api.GetTags)
	}

	return r
}
