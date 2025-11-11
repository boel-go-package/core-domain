package services

import "github.com/muhammad-hasby-golang-package/gin-domain/cmd/domain/message"

type Services interface {
	List() (message.Message, error)
	Create(entity interface{}) error
	Read(id int) (interface{}, error)
	Update(entity interface{}) error
	Delete(id int) error
}
