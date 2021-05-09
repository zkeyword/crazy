package router

import (
	"CRAZY/middleware"
	"CRAZY/router/api"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	// session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	// "github.com/gin-contrib/sessions/redis" 基于redis存储引擎的session
)

// Routers 总路由
func Routers() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 设置 session
	store := cookie.NewStore([]byte("crazy"))
	r.Use(sessions.Sessions("session", store))

	// 设置静态目录
	r.Static("/public", "./public")

	// 设置模板
	r.LoadHTMLGlob("views/**/*")

	// 首页
	r.GET("/", api.GetHTML)

	// 验证码
	r.GET("/captcha", api.GetCaptcha)
	r.POST("/captcha", api.PostCaptcha)

	// 登陆
	r.GET("/login", api.Login)

	// 用户
	r.GET("/user/:id", api.GetUserById)
	r.POST("/user/", api.PostUser)
	r.DELETE("/user/:id", api.DelUserById)
	r.PUT("/user/:id", api.PutUserById)

	// 角色
	r.GET("/role/:id", api.GetRoleById)
	r.POST("/role/", api.PostRole)
	r.DELETE("/role/:id", api.DelRoleById)
	r.PUT("/role/:id", api.PutRoleById)

	// 权限
	r.GET("/permission/:id", api.GetPermissionById)
	r.POST("/permission/", api.PostPermission)
	r.DELETE("/permission/:id", api.DelPermissionById)
	r.PUT("/permission/:id", api.PutPermissionById)

	// api 部分
	apiRouter := r.Group("/api")
	apiRouter.Use(middleware.JWTAuth())
	{
		// 获取配置
		apiRouter.GET("/config", api.GetConfig)
	}

	return r
}
