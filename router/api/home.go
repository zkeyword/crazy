package api

import (
	"CRAZY/utils/db"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// GetHTML
func GetHTML(c *gin.Context) {
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
