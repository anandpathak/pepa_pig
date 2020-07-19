package addcountries

type Response struct {
	Message string `json:"message" binding:"required"`
}
