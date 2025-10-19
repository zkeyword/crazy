package router

import (
	"CRAZY/middleware"
	"CRAZY/router/api/cms"
	"CRAZY/router/api/common"
	"CRAZY/router/api/shop"
	"CRAZY/router/api/sys"
	"net/http"

	// "CRAZY/utils/db"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
	// session存储引擎
	// "github.com/gin-contrib/sessions/cookie"
	// "github.com/gin-contrib/sessions/redis" //基于redis存储引擎的session
)

// Routers 总路由
func Routers(viewsFS http.FileSystem) *gin.Engine {
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
	// r.LoadHTMLGlob("views/**/*")
	r.LoadHTMLFS(viewsFS, "views/**/*")

	// 首页
	r.GET("/", common.GetHTML)

	// 验证码
	r.GET("/captcha", common.GetCaptcha)
	r.POST("/captcha", common.PostCaptcha)

	// r.GET("/test", api.GetTest)
	// r.GET("/sse", sse.SendEvent)
	// r.POST("/sse", sse.PublishHandler)

	// 登陆/注册
	r.POST("/login", sys.Login)
	// r.POST("/register", sys.Register)

	// api 部分
	SysRouter := r.Group("/sys")
	SysRouter.Use(middleware.JWTAuth())
	{
		SysRouter.POST("/refresh", sys.Refresh)
		SysRouter.POST("/logout", sys.Logout)
		SysRouter.POST("/logoutAll", sys.LogoutAll)
		SysRouter.POST("/upload", sys.Upload)
		SysRouter.GET("/permissionKeys", sys.GetPermissionKeys)
		SysRouter.GET("/userInfo", sys.GetUserPermissionKeys)

		// 用户
		SysRouter.GET("/user", sys.GetUser)
		SysRouter.GET("/user/name/:username", sys.GetUserByUsername)
		SysRouter.GET("/user/id/:id", sys.GetUserById)
		SysRouter.POST("/user", sys.PostUser)
		SysRouter.DELETE("/user/:id", sys.DelUserById)
		SysRouter.PUT("/user/:id", sys.PutUserById)
		SysRouter.PUT("/user/status/:id", sys.PutUserStatusById)

		// 角色
		SysRouter.GET("/role", sys.GetRole)
		SysRouter.GET("/role/:id", sys.GetRoleById)
		SysRouter.POST("/role", sys.PostRole)
		SysRouter.DELETE("/role/:id", sys.DelRoleById)
		SysRouter.PUT("/role/:id", sys.PutRoleById)

		// 权限
		SysRouter.GET("/permission/:id", sys.GetPermissionById)
		SysRouter.GET("/permission/:id/tree", sys.GetPermissionTreeById)
		SysRouter.POST("/permission", sys.PostPermission)
		SysRouter.DELETE("/permission/:id", sys.DelPermissionById)
		SysRouter.PUT("/permission/:id", sys.PutPermissionById)

		// 角色、用户、权限关联
		SysRouter.GET("/role/:id/permission", sys.GetRolePermissionByRoleID)
		SysRouter.POST("/role/:id/permission", sys.PostRolePermissionByRoleID)

		SysRouter.GET("/user/:id/role", sys.GetRoleUserByUserID)
		SysRouter.GET("/user/:id/permission", sys.GetUserRolePermissionByUserId)

		SysRouter.GET("/role/:id/user", sys.GetRoleUserByRoleID)
		SysRouter.POST("/role/:id/user/:userId", sys.PostRoleUserByRoleIDAndUserID)
		SysRouter.DELETE("/role/:id/user/:userId", sys.DeleteRoleUserByRoleIDAndUserID)

		// 其他设置
		SysRouter.GET("/other", sys.GetOther)
		SysRouter.GET("/other/:id", sys.GetOtherById)
		SysRouter.POST("/other", sys.PostOther)
		SysRouter.DELETE("/other/:id", sys.DelOtherById)
		SysRouter.PUT("/other/:id", sys.PutOtherById)
		SysRouter.GET("/other/export", sys.ExportOther)
		SysRouter.POST("/other/import", sys.ImportOther)
	}

	// shop 部分
	shopRouter := r.Group("/shop")
	shopRouter.Use(middleware.JWTAuth())
	{
		// 用户地址
		shopRouter.GET("/userAddress", shop.GetUserAddress)

	}

	// cms 部分
	cmsAdminRouter := r.Group("/cms/admin")
	cmsAdminRouter.Use(middleware.JWTAuth())
	{
		// 文章
		cmsAdminRouter.GET("/post", cms.GetPost)
		cmsAdminRouter.GET("/post/:id", cms.GetPostById)
		cmsAdminRouter.POST("/post", cms.PostPost)
		cmsAdminRouter.DELETE("/post/:id", cms.DelPostById)
		cmsAdminRouter.PUT("/post/:id", cms.PutPostById)

		// banner
		cmsAdminRouter.GET("/banner", cms.GetBanner)
		cmsAdminRouter.GET("/banner/:id", cms.GetBannerById)
		cmsAdminRouter.POST("/banner", cms.PostBanner)
		cmsAdminRouter.DELETE("/banner/:id", cms.DelBannerById)
		cmsAdminRouter.PUT("/banner/:id", cms.PutBannerById)

		// 类别
		cmsAdminRouter.GET("/category", cms.GetCategory)
		cmsAdminRouter.GET("/category/:id", cms.GetCategoryById)
		cmsAdminRouter.POST("/category", cms.PostCategory)
		cmsAdminRouter.DELETE("/category/:id", cms.DelCategoryById)
		cmsAdminRouter.PUT("/category/:id", cms.PutCategoryById)
	}

	cmsRouter := r.Group("/cms/api")
	{
		cmsRouter.GET("/post", cms.GetPost)
		cmsRouter.GET("/post/:id", cms.GetPostById)
		cmsRouter.GET("/category", cms.GetCategory)
		cmsRouter.GET("/banner", cms.GetBanner)
	}

	return r
}
