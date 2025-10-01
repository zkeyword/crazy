package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Response 返回基本格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = -1
)

// Result 基础封装：允许指定 HTTP 状态码
func Result(httpStatus int, code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// ========== 成功响应（HTTP 200） ==========

func Ok(c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, message, c)
}

func OkDetailed2(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, 10000, data, message, c)
}

// ========== 失败响应（通用） ==========

func Fail(c *gin.Context) {
	Result(http.StatusOK, ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, ERROR, map[string]interface{}{}, message, c)
}

// FailWithDetailed 允许自定义业务 code，但 HTTP 状态码仍为 200（谨慎使用）
func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, code, data, message, c)
}

// ========== 认证/权限专用失败（语义化 HTTP 状态码） ==========

// 401 Unauthorized：身份认证失败（token 无效、过期、缺失）
func FailWithUnauthorized(message string, c *gin.Context) {
	Result(http.StatusUnauthorized, ERROR, map[string]interface{}{}, message, c)
}

// 403 Forbidden：已认证但无权限
func FailWithForbidden(message string, c *gin.Context) {
	Result(http.StatusForbidden, ERROR, map[string]interface{}{}, message, c)
}

// 400 Bad Request：客户端参数错误
func FailWithBadRequest(message string, c *gin.Context) {
	Result(http.StatusBadRequest, ERROR, map[string]interface{}{}, message, c)
}

// ========== 权限检查 ==========

func CheckPermission(c *gin.Context, permKey string) bool {
	// 1. 检查 permissionKeys
	permKeys, exists := c.Get("permissionKeys")
	if !exists {
		FailWithUnauthorized("未登录或权限信息缺失", c)
		return false
	}

	permKeysStr, ok := permKeys.(string)
	if !ok {
		FailWithUnauthorized("权限信息格式错误", c)
		return false
	}

	// 2. 检查权限
	if !strings.Contains(permKeysStr, "All") && !strings.Contains(permKeysStr, permKey) {
		FailWithForbidden("无此操作权限", c)
		return false
	}

	return true
}
