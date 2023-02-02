package dto

type InputUser struct {
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

type InputLogin struct {
	UserName string
	Password string
}
