package sse

import (
	"CRAZY/utils"
	"CRAZY/utils/db"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var (
	Mutex   sync.Mutex
	Clients map[chan string]bool
)

func init() {
	Clients = make(map[chan string]bool)
}

func SendEvent(c *gin.Context) {
	clientChan := make(chan string)
	Mutex.Lock()
	Clients[clientChan] = true
	Mutex.Unlock()

	defer func() {
		Mutex.Lock()
		delete(Clients, clientChan)
		Mutex.Unlock()
		close(clientChan)
		log.Println("Client disconnected")
	}()

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// fmt.Fprintf(c.Writer, "data: %s\n\n", "Connected successfully")
	// c.Writer.Flush()

	closeNotify := c.Writer.CloseNotify()
	var lastMessage string
	idleTimeout := time.After(30 * time.Second)
	for {
		select {
		case <-closeNotify:
			log.Println("Received close notify")
			return
		case msg := <-clientChan:
			// c.SSEvent("message", msg)
			log.Println("Received message:", msg)
			lastMessage = msg
			fmt.Fprintf(c.Writer, "data: %s\n\n", msg)
			c.Writer.Flush()
		case <-idleTimeout:
			log.Println("Idle timeout")
			if lastMessage != "" {
				fmt.Fprintf(c.Writer, "data: %s\n\n", lastMessage)
				c.Writer.Flush()
			}
			idleTimeout = time.After(30 * time.Second)
		}
	}

	// for {
	// 	fmt.Fprintf(c.Writer, "data: %s\n\n", time.Now().String())
	// 	c.Writer.Flush()
	// 	time.Sleep(time.Second)
	// }
}

type SSEForm struct {
	Msg string `form:"msg" binding:"required"`
}

func PublishHandler(c *gin.Context) {
	var form SSEForm
	err := c.ShouldBind(&form)
	if err == nil {
		conn := db.GetRedis()
		defer conn.Close() // Close the connection when you're done with it.
		res, resErr := conn.Do("PUBLISH", "sseChannel", form.Msg)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

func Subscribe() {
	go func() {
		conn := db.GetRedis()
		defer conn.Close() // Close the connection when you're done with it.
		psc := redis.PubSubConn{Conn: conn}
		psc.Subscribe("sseChannel")

		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				Mutex.Lock()
				for clientMessageChan := range Clients {
					func() {
						defer func() {
							if r := recover(); r != nil {
								log.Println("Recovered in listenForMessages", r)
								delete(Clients, clientMessageChan)
							}
						}()
						clientMessageChan <- string(v.Data)
					}()
				}
				Mutex.Unlock()

			case error:
				return
			}
		}
	}()
}
