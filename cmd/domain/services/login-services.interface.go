package services

type LoginServices interface {
	Login(dto interface{}) (interface{}, error)
	Logout() error
}
