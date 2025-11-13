package services

import "github.com/boel-go-package/core-domain/cmd/domain/message"

type Services interface {
	List() (message.Message, error)
	Create(entity interface{}) error
	Read(id int) (interface{}, error)
	Update(entity interface{}) error
	Delete(id int) error
}
