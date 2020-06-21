package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

// StartHTTPServer starts the http servier
func StartHTTPServer(port string) {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}
