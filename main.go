package main

import (
	"CRAZY/config"
	"CRAZY/router"
	"CRAZY/utils/db"

	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func main() {

	// fmt.Println("请输入账号:")
	// reader := bufio.NewReader(os.Stdin)
	// user, _ := reader.ReadString('\n')
	// user = strings.TrimSuffix(user, "\n")

	// secret := googleAuthenticator.NewGoogleAuth().GetSecret()
	// code, err := googleAuthenticator.NewGoogleAuth().GetCode(secret)

	// qrCodeURL := googleAuthenticator.NewGoogleAuth().GetQrcodeUrl(user, secret)
	// desktop.Open(qrCodeURL)

	// fmt.Println("请输入CODE:")
	// readerCode := bufio.NewReader(os.Stdin)
	// userCode, _ := readerCode.ReadString('\n')
	// userCode = strings.TrimSuffix(userCode, "\n")

	// fmt.Println(code, userCode, secret)
	// if code != userCode {
	// 	os.Exit(1)
	// }

	// utils.WriteFile()
	// utils.ReadFile()

	ENV := os.Getenv("MY_ENV_VAR")
	fmt.Println("The value of MY_ENV_VAR is:", ENV)

	// 启动mysql
	if ENV == "prod" {
		db.StartMysql(config.ProdDbConfig.Dsn, config.ProdDbConfig.MaxIdle, config.ProdDbConfig.MaxOpen, config.ProdDbConfig.LogMode)
	} else {
		db.StartMysql(config.DbConfig.Dsn, config.DbConfig.MaxIdle, config.DbConfig.MaxOpen, config.DbConfig.LogMode)
	}

	// 启动redis
	// defer db.CloseRedis()
	// if ENV == "prod" {
	// 	db.StartRedis(config.ProdRedisDbConfig.Addr, config.ProdRedisDbConfig.Password, config.ProdRedisDbConfig.DB, config.ProdRedisDbConfig.MaxIdle, config.ProdRedisDbConfig.MaxOpen)
	// } else {
	// 	db.StartRedis(config.RedisDbConfig.Addr, config.RedisDbConfig.Password, config.RedisDbConfig.DB, config.RedisDbConfig.MaxIdle, config.RedisDbConfig.MaxOpen)
	// }

	// 启动sse Subscribe
	// sse.Subscribe()

	// 创建文件日志，按天分割，日志文件仅保留一周
	w, err := rotatelogs.New(config.LogPath)
	checkErr("CreateRotateLog", err)

	// 设置日志
	logrus.SetOutput(w)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)

	// 设置gin
	gin.SetMode(config.ServerMode)
	r := router.Routers()
	r.MaxMultipartMemory = config.MaxMultipartMemory

	server := &http.Server{
		Addr:           config.ServerPort,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening http://0.0.0.0%s", config.ServerPort)

	server.ListenAndServe()
}

func checkErr(errMsg string, err error) {
	if err != nil {
		fmt.Printf("%s Error: %v\n", errMsg, err)
		os.Exit(1)
	}
}
