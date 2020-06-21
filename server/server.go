package server

import (
	"fmt"
)

// StartHTTPServer starts the http servier
func StartHTTPServer(port string) {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}
