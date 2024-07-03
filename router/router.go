package router

import (
	"CRAZY/middleware"
	"CRAZY/router/api/common"
	"CRAZY/router/api/shop"
	"CRAZY/router/api/system"

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
	// r.Static("/public", "./public")

	// 设置模板
	// r.LoadHTMLGlob("views/**/*")

	// 首页
	// r.GET("/", common.GetHTML)

	// 登陆/注册
	r.POST("/login", system.Login)
	// r.POST("/register", system.Register)

	// 验证码
	r.GET("/captcha", common.GetCaptcha)
	r.POST("/captcha", common.PostCaptcha)

	// r.GET("/test", api.GetTest)
	// r.GET("/sse", sse.SendEvent)
	// r.POST("/sse", sse.PublishHandler)

	// api 部分
	apiRouter := r.Group("/api")
	apiRouter.Use(middleware.JWTAuth())
	{
		apiRouter.POST("/common/upload", common.Upload)
		apiRouter.POST("/common/logout", system.Logout)
		// TOTO: 修改路由
		apiRouter.GET("/common/permissionKeys", system.GetPermissionKeys)
		apiRouter.GET("/common/userInfo", system.GetUserPermissionKeys)

		// 用户
		apiRouter.GET("/user", system.GetUser)
		apiRouter.GET("/user/name/:username", system.GetUserByUsername)
		apiRouter.GET("/user/id/:id", system.GetUserById)
		apiRouter.POST("/user", system.PostUser)
		apiRouter.DELETE("/user/:id", system.DelUserById)
		apiRouter.PUT("/user/:id", system.PutUserById)
		apiRouter.PUT("/user/status/:id", system.PutUserStatusById)

		// 角色
		apiRouter.GET("/role", system.GetRole)
		apiRouter.GET("/role/:id", system.GetRoleById)
		apiRouter.POST("/role", system.PostRole)
		apiRouter.DELETE("/role/:id", system.DelRoleById)
		apiRouter.PUT("/role/:id", system.PutRoleById)

		// 权限
		apiRouter.GET("/permission/:id", system.GetPermissionById)
		apiRouter.GET("/permission/:id/tree", system.GetPermissionTreeById)
		apiRouter.POST("/permission", system.PostPermission)
		apiRouter.DELETE("/permission/:id", system.DelPermissionById)
		apiRouter.PUT("/permission/:id", system.PutPermissionById)

		// 角色、用户、权限关联
		apiRouter.GET("/role/:id/permission", system.GetRolePermissionByRoleID)
		apiRouter.POST("/role/:id/permission", system.PostRolePermissionByRoleID)

		apiRouter.GET("/user/:id/role", system.GetRoleUserByUserID)
		apiRouter.GET("/user/:id/permission", system.GetUserRolePermissionByUserId)

		apiRouter.GET("/role/:id/user", system.GetRoleUserByRoleID)
		apiRouter.POST("/role/:id/user/:userId", system.PostRoleUserByRoleIDAndUserID)
		apiRouter.DELETE("/role/:id/user/:userId", system.DeleteRoleUserByRoleIDAndUserID)

		// 用户地址
		apiRouter.GET("/userAddress", shop.GetUserAddress)

		// 其他设置
		apiRouter.GET("/other", system.GetOther)
		apiRouter.GET("/other/:id", system.GetOtherById)
		apiRouter.POST("/other", system.PostOther)
		apiRouter.DELETE("/other/:id", system.DelOtherById)
		apiRouter.PUT("/other/:id", system.PutOtherById)
		apiRouter.GET("/other/export", system.ExportOther)
		apiRouter.POST("/other/import", system.ImportOther)

	}

	return r
}
