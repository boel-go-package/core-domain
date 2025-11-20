package middleware

import (
	"net/http"

	"github.com/boel-go-package/core-domain/cmd/domain/message"
	"github.com/gin-gonic/gin"
)

var (
	active bool = true
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !active {
			res := message.Failed("503_SERVICE_UNAVAILABLE", "Server is not available consume request", nil, nil)

			c.JSON(http.StatusServiceUnavailable, res)
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequestActive() {
	active = true
}

func RequestInactive() {
	active = false
}
