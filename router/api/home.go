package api

import (
	"github.com/gin-gonic/gin"
)

// GetHTML
func GetHTML(c *gin.Context) {
	c.HTML(200, "home.html", "xxx")
}
