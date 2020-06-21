package server

import (
	"net/http"

	fetchcountries "pepa_pig/internal/fetch_countries"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/fetch_countries", fetchcountries.Handler)

	return r
}
