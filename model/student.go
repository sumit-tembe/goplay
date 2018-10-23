package model

//Student ...
type Student struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	MiddleName string `json:"middlename"`
}
