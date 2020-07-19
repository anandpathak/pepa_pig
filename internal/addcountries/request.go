package addcountries

import "pepa_pig/utils"

type Request struct {
	Source  string `json:"source" binding:"required"`
	Company string `json:"company" binding: "required"`
	Country string `json:"country" binding: "required"`
}

func (r Request) isValid() bool {
	if len(utils.ConvertTolowerCaseWithoutSpace(r.Company)) < 2 ||
		len(utils.ConvertTolowerCaseWithoutSpace(r.Country)) < 2 {
		return false
	}
	return true
}
