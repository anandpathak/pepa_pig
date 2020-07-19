package addcountries

import (
	"fmt"
	"pepa_pig/redis"
	"pepa_pig/utils"

	"github.com/gin-gonic/gin"
)

type service struct {
	context *gin.Context
	req     Request
}

func (s service) Execute() (Response, error) {
	
	redis.Instance.SetNX(s.context,
		fmt.Sprintf("in_review_%s", utils.ConvertTolowerCaseWithoutSpace(s.req.Company)),
		utils.ConvertTolowerCaseWithoutSpace(s.req.Country), 0)
	//redis.Instance.Get(s.context, s.req)
	return Response{
		Message: "Thanks! send for review",
	}, nil
}
