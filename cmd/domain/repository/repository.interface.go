package repository

import "github.com/muhammad-hasby-golang-package/gin-domain/cmd/domain/message"

type Repository interface {
	Create(entity interface{}) error
	Read(id int) (message.Message, error)
	Update(entity interface{}) error
	Delete(id int) error
}
