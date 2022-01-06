package web

type PersonCreateRequest struct {
	FirstName string `validate:"required,max=200,min=1" json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `validate:"required,gte=0,lte=100" json:"age"`
}
