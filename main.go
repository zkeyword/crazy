package main

import (
	"CRAZY/config"
	"CRAZY/router"
	"CRAZY/utils/xor"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func main() {
	// utils.WriteFile()
	// utils.ReadFile()

	startTime := time.Now()
	d, _ := time.ParseDuration(strconv.Itoa(7*24) + "h")
	endTime := startTime.Add(d)
	fmt.Println(xor.Enc(endTime.Format(config.SysTimeform)))
	x := xor.Dec("80398965be5d736a81399b67a157747688398a")
	fmt.Println(x)

	flag.Parse()

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
