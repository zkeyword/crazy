package main

import (
	"CRAZY/config"
	"CRAZY/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
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
