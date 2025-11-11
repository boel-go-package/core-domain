package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammad-hasby-golang-package/gin-domain/cmd/domain/message"
)

type Controller interface {
	Create(entity interface{}) error
	List(c *gin.Context)
	Get(id int) (message.Message, error)
	Find(query interface{}) (message.Message, error)
	Update(entity interface{}) error
	Delete(id int) error
}
