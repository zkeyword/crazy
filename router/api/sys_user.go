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
	Role     string `form:"role"`
}

// PostUser 新增用户
func PostUser(c *gin.Context) {
	var form UserForm
	err := c.ShouldBind(&form)
	if form.Role == "" {
		// TODO: https://www.cnblogs.com/xinliangcoder/p/11234017.html 自定义验证器
		utils.FailWithMessage("Key: 'UserForm.Role' Error:Field validation for 'Role' failed on the 'lt' 0", c)
		return
	}
	if err == nil {
		Model := &model.User{
			Username: form.Username,
			Password: xor.Enc(form.Password),
			Status:   form.Status,
		}
		res, resErr := services.NewUserService.Create(Model, form.Role)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// DelUserById 删除用户
func DelUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	resErr := services.NewUserService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

type PutUserForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Status   int    `form:"status"`
	Role     string `form:"role"`
}

// PutUserById 修改用户
func PutUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form PutUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.User{
			Username: form.Username,
			Password: xor.Enc(form.Password),
		}
		res, resErr := services.NewUserService.PutUserById(id, Model, form.Role)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetUserById 获取用户
func GetUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := services.NewUserService.Get(id)
	utils.OkDetailed(res, "success", c)
}
