package api

import (
	"CRAZY/utils/db"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

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
