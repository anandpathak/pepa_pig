package server

import (
	"net/http"

	"pepa_pig/internal/addcountries"
	fetchcountries "pepa_pig/internal/fetch_countries"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/fetch_countries", fetchcountries.Handler)
	r.POST("/add_country", addcountries.Handler)
	return r
}
