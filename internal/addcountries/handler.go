package addcountries

import (
	"github.com/gin-gonic/gin"
)

// Handler handle the request for fetch country
func Handler(c *gin.Context) {
	var req Request
	c.BindJSON(&req)
	s := service{
		context: c,
		req:     req,
	}
	if !req.isValid() {
		c.JSON(400, gin.H{"error": "Not a valid company/country"})
		return
	}
	res, err := s.Execute()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": res.Message})
}
