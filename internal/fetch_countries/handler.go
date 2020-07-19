package fetchcountries

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Handler handle the request for fetch country
func Handler(c *gin.Context) {
	var req Request
	c.BindJSON(&req)
	s := service{
		context: c,
		requst:  req,
	}
	res, err := s.Execute()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": res.Data})
}
