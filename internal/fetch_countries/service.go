package fetchcountries

import "github.com/gin-gonic/gin"

type service struct {
	context *gin.Context
	requst  Request
}

func (s service) Execute() (Response, error) {

	return Response{
		Data: map[string]string{
			"amazon":   "us",
			"flipkart": "india",
		},
	}, nil
}
