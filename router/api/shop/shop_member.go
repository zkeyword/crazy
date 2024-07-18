package shop

import (
	sysUserService "CRAZY/services/system/sys_user"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMemberList 获取会员列表
func GetMemberList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := html.EscapeString(c.Query("username"))
	res, _ := sysUserService.Get(page, pageSize, username)

	utils.OkDetailed(res, "success", c)

}
