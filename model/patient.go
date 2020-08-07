package model
//model to bind data from api response/request
type Patient struct {
	Name  string `json:"name"`
	Tel   string `json:"tel"`
	Email string `json:"email"`
}
