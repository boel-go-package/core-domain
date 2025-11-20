package repository

type AuthRepository interface {
	Login(dto interface{}) (interface{}, error)
	Logout() error
}
