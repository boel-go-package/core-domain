package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Find(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
