package api

import (
	sysUserService "CRAZY/services/sys_user"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAddressForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Status   int    `form:"status" binding:"required"`
	RoleIDs  string `form:"roleIds"`
}

// GetUser 获取用户列表
func GetUserAddress(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := html.EscapeString(c.Query("username"))
	res, _ := sysUserService.Get(page, pageSize, username)
	utils.OkDetailed(res, "success", c)
}
