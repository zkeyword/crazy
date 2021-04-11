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

	// startTime := time.Now()
	// d, _ := time.ParseDuration(strconv.Itoa(7*24) + "h")
	// endTime := startTime.Add(d)
	// fmt.Println(xor.Enc(endTime.Format(config.SysTimeform)))
	// x := xor.Dec("80398965be5d736a81399b67a157747688398a")
	// fmt.Println(x)

	// flag.Parse()

	// 启动mysql
	defer db.CloseMysql()
	fmt.Print("Start Mysql...\r")
	checkErr("Start Mysql", db.StartMysql(config.DbConfig.Dsn, config.DbConfig.MaxIdle, config.DbConfig.MaxOpen))
	fmt.Print("Start Mysql Success!!!\n")

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
