package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Response  返回基本格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const SUCCESS = 0
const ERROR = -1

// Result 基本封装
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

// 根据permissionKeys处理权限
func CheckPermission(c *gin.Context, permKey string) bool {
	PermissionKeys, _ := c.Get("permissionKeys")
	PermissionKeysStr, _ := PermissionKeys.(string)
	found := false
	for _, v := range strings.Split(PermissionKeysStr, ",") {
		if v == "All" || v == permKey {
			found = true
			break
		}
	}
	if !found {
		err := errors.New("not permission")
		FailWithMessage(err.Error(), c)
	}
	return found
}
