package api

import (
	"CRAZY/model"
	"CRAZY/services"
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Status   int    `form:"status" binding:"required"`
}

// PostUser 新增用户
func PostUser(c *gin.Context) {
	var form UserForm
	err := c.ShouldBind(&form)
	if err == nil {
		User := &model.User{
			Username: form.Username,
			Password: xor.Enc(form.Password),
			Status:   form.Status,
		}
		user, createErr := services.NewUserService.Create(User)
		if createErr == nil {
			utils.OkDetailed(gin.H{
				"user": user,
			}, "success", c)
		} else {
			utils.FailWithMessage(createErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// DelUserById 删除用户
func DelUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, _ := services.NewUserService.DeleteById(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

type PutUserForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Status   int    `form:"status"`
}

// PutUserById 修改用户
func PutUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form PutUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		User := &model.User{
			Username: form.Username,
			Password: xor.Enc(form.Password),
		}
		user, _ := services.NewUserService.PutUserById(id, User)
		utils.OkDetailed(gin.H{
			"user": user,
		}, "success", c)
	}

}

// GetUserById 获取用户
func GetUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewUserService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}
