package api

import (
	"CRAZY/utils/db"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func safeDivide(numerator, denominator int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("divide operation failed: %v", r)
		}
	}()

	result = numerator / denominator // 如果 denominator 是 0，这里会发生 panic
	return
}

// GetHTML
func GetHTML(c *gin.Context) {
	// go say("world") // 创建一个新的 Goroutine
	// say("hello")    // 在当前 Goroutine（主 Goroutine）中执行

	// messages := make(chan string)

	// go func() {
	// 	messages <- "hello2"
	// }()

	// msg := <-messages
	// fmt.Println(msg)

	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	c.HTML(200, "home.html", "xxx")
}

func GetTest(c *gin.Context) {
	conn := db.GetRedis()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", "test"))
	if err != nil {
		println(111, err)
		return
	}
	c.String(200, data)
}

// 检查写入是否出错
func checkWriteErr(w gin.ResponseWriter) bool {
	if _, ok := <-w.CloseNotify(); ok {
		return true
	}
	return false
}

func GetSSE(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("start sse")

	// 创建一个通道，用于接收退出信号
	quit := make(chan struct{})
	// 客户端连接关闭时，发送退出信号
	notify := c.Writer.CloseNotify()
	go func() {
		<-notify
		quit <- struct{}{}
	}()

	counter := 0
	for {
		select {
		case <-time.After(time.Second):
			// 每秒发送一次事件
			message := fmt.Sprintf("Event #%d", counter)
			c.SSEvent("message", message)
			counter++
		case <-quit:
			// 收到退出信号，结束循环
			return
		}

		// 检查写入是否出错，如果出错（例如客户端断开连接），则结束循环
		if checkWriteErr(c.Writer) {
			return
		}
	}
}
