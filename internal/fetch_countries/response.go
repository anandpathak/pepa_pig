package fetchcountries

type Response struct {
	Data map[string]string `form:"data"`
}
