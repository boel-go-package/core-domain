package authorization

import "github.com/gin-gonic/gin"

type Authorization interface {
	Validate() gin.HandlerFunc
}
