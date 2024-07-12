package router

import (
	"CRAZY/middleware"
	"CRAZY/router/api/cms"
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

	// 验证码
	r.GET("/captcha", common.GetCaptcha)
	r.POST("/captcha", common.PostCaptcha)

	// r.GET("/test", api.GetTest)
	// r.GET("/sse", sse.SendEvent)
	// r.POST("/sse", sse.PublishHandler)

	// 登陆/注册
	r.POST("/login", system.Login)
	// r.POST("/register", system.Register)

	// api 部分
	systemRouter := r.Group("/system")
	systemRouter.Use(middleware.JWTAuth("system"))
	{
		systemRouter.POST("/upload", system.Upload)
		systemRouter.POST("/logout", system.Logout)
		systemRouter.GET("/permissionKeys", system.GetPermissionKeys)
		systemRouter.GET("/userInfo", system.GetUserPermissionKeys)

		// 用户
		systemRouter.GET("/user", system.GetUser)
		systemRouter.GET("/user/name/:username", system.GetUserByUsername)
		systemRouter.GET("/user/id/:id", system.GetUserById)
		systemRouter.POST("/user", system.PostUser)
		systemRouter.DELETE("/user/:id", system.DelUserById)
		systemRouter.PUT("/user/:id", system.PutUserById)
		systemRouter.PUT("/user/status/:id", system.PutUserStatusById)

		// 角色
		systemRouter.GET("/role", system.GetRole)
		systemRouter.GET("/role/:id", system.GetRoleById)
		systemRouter.POST("/role", system.PostRole)
		systemRouter.DELETE("/role/:id", system.DelRoleById)
		systemRouter.PUT("/role/:id", system.PutRoleById)

		// 权限
		systemRouter.GET("/permission/:id", system.GetPermissionById)
		systemRouter.GET("/permission/:id/tree", system.GetPermissionTreeById)
		systemRouter.POST("/permission", system.PostPermission)
		systemRouter.DELETE("/permission/:id", system.DelPermissionById)
		systemRouter.PUT("/permission/:id", system.PutPermissionById)

		// 角色、用户、权限关联
		systemRouter.GET("/role/:id/permission", system.GetRolePermissionByRoleID)
		systemRouter.POST("/role/:id/permission", system.PostRolePermissionByRoleID)

		systemRouter.GET("/user/:id/role", system.GetRoleUserByUserID)
		systemRouter.GET("/user/:id/permission", system.GetUserRolePermissionByUserId)

		systemRouter.GET("/role/:id/user", system.GetRoleUserByRoleID)
		systemRouter.POST("/role/:id/user/:userId", system.PostRoleUserByRoleIDAndUserID)
		systemRouter.DELETE("/role/:id/user/:userId", system.DeleteRoleUserByRoleIDAndUserID)

		// 用户地址
		systemRouter.GET("/userAddress", shop.GetUserAddress)

		// 其他设置
		systemRouter.GET("/other", system.GetOther)
		systemRouter.GET("/other/:id", system.GetOtherById)
		systemRouter.POST("/other", system.PostOther)
		systemRouter.DELETE("/other/:id", system.DelOtherById)
		systemRouter.PUT("/other/:id", system.PutOtherById)
		systemRouter.GET("/other/export", system.ExportOther)
		systemRouter.POST("/other/import", system.ImportOther)

	}

	// shop 部分
	shopRouter := r.Group("/shop")
	shopRouter.Use(middleware.JWTAuth("shop"))
	{
		// 用户地址
		shopRouter.GET("/userAddress", shop.GetUserAddress)

	}

	// cms 部分
	cmsAdminRouter := r.Group("/cms/admin")
	cmsAdminRouter.Use(middleware.JWTAuth("cms"))
	{
		// 文章
		cmsAdminRouter.GET("/post", cms.GetPost)
		cmsAdminRouter.GET("/post/:id", cms.GetPostById)
		cmsAdminRouter.POST("/post/", cms.PostPost)
		cmsAdminRouter.DELETE("/post/:id", cms.DelPostById)
		cmsAdminRouter.PUT("/post/:id", cms.PutPostById)

		// banner
		cmsAdminRouter.GET("/banner", cms.GetBanner)
		cmsAdminRouter.GET("/banner/:id", cms.GetBannerById)
		cmsAdminRouter.POST("/banner/", cms.PostBanner)
		cmsAdminRouter.DELETE("/banner/:id", cms.DelBannerById)
		cmsAdminRouter.PUT("/banner/:id", cms.PutBannerById)

		// 类别
		cmsAdminRouter.GET("/category", cms.GetCategory)
		cmsAdminRouter.GET("/category/:id", cms.GetCategoryById)
		cmsAdminRouter.POST("/category/", cms.PostCategory)
		cmsAdminRouter.DELETE("/category/:id", cms.DelCategoryById)
		cmsAdminRouter.PUT("/category/:id", cms.PutCategoryById)
	}

	cmsRouter := r.Group("/cms/api")
	{
		cmsRouter.GET("/post", cms.GetPost)
		cmsRouter.GET("/post/news", cms.GetPostNews)
		cmsRouter.GET("/post/:id", cms.GetPostById)
		cmsRouter.GET("/category", cms.GetCategory)
		cmsRouter.GET("/banner", cms.GetBanner)
	}

	return r
}
