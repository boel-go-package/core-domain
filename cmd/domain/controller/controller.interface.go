package controller

import (
	"github.com/boel-go-package/core-domain/cmd/domain/message"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(entity interface{}) error
	List(c *gin.Context)
	Get(id int) (message.Message, error)
	Find(query interface{}) (message.Message, error)
	Update(entity interface{}) error
	Delete(id int) error
}
