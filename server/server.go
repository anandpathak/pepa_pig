package server

import (
	"fmt"
	"pepa_pig/redis"
)

// StartHTTPServer starts the http servier
func StartHTTPServer(port string) {

	redis.Connect()
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}
