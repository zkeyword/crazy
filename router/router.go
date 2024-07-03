package router

import (
	"CRAZY/middleware"
	"CRAZY/router/api"

	// "CRAZY/utils/db"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
	// session存储引擎
	// "github.com/gin-contrib/sessions/cookie"
	// "github.com/gin-contrib/sessions/redis" //基于redis存储引擎的session
)

// Routers 总路由
func Routers() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 设置 session
	// store := cookie.NewStore([]byte("crazy"))
	// store, _ := redis.NewStoreWithPool(db.GetRedisPool(), []byte("crazy"))
	// r.Use(sessions.Sessions("session", store))

	// 设置静态目录
	r.Static("/public", "./public")

	// 设置模板
	r.LoadHTMLGlob("views/**/*")

	// 首页
	r.GET("/", api.GetHTML)

	// 登陆/注册
	r.POST("/login", api.Login)
	r.POST("/register", api.Register)

	// 验证码
	r.GET("/captcha", api.GetCaptcha)
	r.POST("/captcha", api.PostCaptcha)

	// r.GET("/test", api.GetTest)
	// r.GET("/sse", sse.SendEvent)
	// r.POST("/sse", sse.PublishHandler)

	// api 部分
	apiRouter := r.Group("/api")
	apiRouter.Use(middleware.JWTAuth())
	{
		apiRouter.POST("/common/upload", api.Upload)
		apiRouter.POST("/common/logout", api.Logout)
		apiRouter.GET("/common/permissionKeys", api.GetPermissionKeys)
		apiRouter.GET("/common/userInfo", api.GetUserPermissionKeys)

		// 用户
		apiRouter.GET("/user", api.GetUser)
		apiRouter.GET("/user/name/:username", api.GetUserByUsername)
		apiRouter.GET("/user/id/:id", api.GetUserById)
		apiRouter.POST("/user", api.PostUser)
		apiRouter.DELETE("/user/:id", api.DelUserById)
		apiRouter.PUT("/user/:id", api.PutUserById)
		apiRouter.PUT("/user/status/:id", api.PutUserStatusById)

		// 角色
		apiRouter.GET("/role", api.GetRole)
		apiRouter.GET("/role/:id", api.GetRoleById)
		apiRouter.POST("/role", api.PostRole)
		apiRouter.DELETE("/role/:id", api.DelRoleById)
		apiRouter.PUT("/role/:id", api.PutRoleById)

		// 权限
		apiRouter.GET("/permission/:id", api.GetPermissionById)
		apiRouter.GET("/permission/:id/tree", api.GetPermissionTreeById)
		apiRouter.POST("/permission", api.PostPermission)
		apiRouter.DELETE("/permission/:id", api.DelPermissionById)
		apiRouter.PUT("/permission/:id", api.PutPermissionById)

		// 角色、用户、权限关联
		apiRouter.GET("/role/:id/permission", api.GetRolePermissionByRoleID)
		apiRouter.POST("/role/:id/permission", api.PostRolePermissionByRoleID)

		apiRouter.GET("/user/:id/role", api.GetRoleUserByUserID)
		apiRouter.GET("/user/:id/permission", api.GetUserRolePermissionByUserId)

		apiRouter.GET("/role/:id/user", api.GetRoleUserByRoleID)
		apiRouter.POST("/role/:id/user/:userId", api.PostRoleUserByRoleIDAndUserID)
		apiRouter.DELETE("/role/:id/user/:userId", api.DeleteRoleUserByRoleIDAndUserID)

		// 用户地址
		apiRouter.GET("/userAddress", api.GetUser)

		// 其他设置
		apiRouter.GET("/other", api.GetOther)
		apiRouter.GET("/other/:id", api.GetOtherById)
		apiRouter.POST("/other", api.PostOther)
		apiRouter.DELETE("/other/:id", api.DelOtherById)
		apiRouter.PUT("/other/:id", api.PutOtherById)
		apiRouter.GET("/other/export", api.ExportOther)
		apiRouter.POST("/other/import", api.ImportOther)

	}

	return r
}
