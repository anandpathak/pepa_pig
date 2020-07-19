package fetchcountries

import (
	"pepa_pig/redis"

	"github.com/gin-gonic/gin"
)

type service struct {
	context *gin.Context
	requst  Request
}

func (s service) Execute() (Response, error) {

	data := make(map[string]string)

	for _, c := range s.requst.Companies {
		data[c.Name] = redis.Instance.Get(s.context, c.Name).Val()
	}
	//redis.Instance.Get(s.context, s.requst)
	return Response{
		Data: data,
	}, nil
}
