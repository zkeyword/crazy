package utils

import (
	"CRAZY/utils/db"
	"fmt"
	"net/http"
	"strconv"
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
	PermissionKeys, exists := c.Get("permissionKeys")
	if !exists {
		FailWithMessage("Permission keys not found", c)
		return false
	}

	PermissionKeysStr, ok := PermissionKeys.(string)
	if !ok {
		FailWithMessage("Permission keys are not a string", c)
		return false
	}

	if !strings.Contains(PermissionKeysStr, "All") && !strings.Contains(PermissionKeysStr, permKey) {
		FailWithMessage("No permission", c)
		return false
	}

	userID, exists := c.Get("userID")
	if !exists {
		FailWithMessage("User ID not found", c)
		return false
	}
	userIDUint, _ := userID.(uint)
	value, err := db.GetKey("UserLoginStatus" + strconv.FormatUint(uint64(userIDUint), 10))
	fmt.Println(value)
	if err != nil || value == "-1" {
		FailWithMessage("Permission expired", c)
		return false
	}

	return true
}
