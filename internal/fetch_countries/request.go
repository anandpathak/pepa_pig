package fetchcountries

type Request struct {
	Source    string    `json:"source" binding:"required"`
	Companies []Company `json:"companies" binding: "required"`
}

type Company struct {
	Name string `json:"name"`
}
