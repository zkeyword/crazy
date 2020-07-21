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

	r.GET("/ping", api.GetTags)

	return r
}
