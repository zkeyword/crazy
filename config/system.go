package config

import "time"

// ServerPort web服务端口
var ServerPort = ":9000"

// ServerMode gin开发模式
// var ServerMode = "debug"
var ServerMode = "release"

// LogPath 日志路径
var LogPath = "runtime/logs/%Y%m%d.log"

// SysTimeform 时间格式化字符串
const SysTimeform string = "2006-01-02 15:04:05"

// SysTimeformShort 日期时间格式化字符串
const SysTimeformShort string = "2006-01-02"

// SysTimeLocation 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// 设置文件上传大小
var MaxMultipartMemory int64 = 100 << 20 // 设置最大上传大小为100M

// // Sys web系统配置结构体
// type Sys struct {
// 	AppAdmin      string
// 	UserIDKey     string
// 	UserStructKey string
// }

// // SysConfig web系统配置
// var SysConfig = &Sys{
// 	AppAdmin:      "admin",
// 	UserIDKey:     "UserID",
// 	UserStructKey: "User",
// }

var JWTExpiresAt = time.Now().Add(24 * 30 * time.Hour).Unix()
var JWTSignKey = "CRAZY"
