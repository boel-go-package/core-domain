package authorization

import "github.com/gin-gonic/gin"

type Authorization interface {
	Validate(c *gin.Context) error
}
