package router

import (
	"CRAZY/router/api"

	"github.com/gin-gonic/gin"
)

// Routers 总路由
func Routers() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 设置静态目录
	r.Static("/public", "./public")

	// 设置模板
	r.LoadHTMLGlob("views/**/*")


	// api 部分
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/ping", api.GetTags)
		apiRouter.GET("/html", api.GetHTML)
	}
	

	return r
}
