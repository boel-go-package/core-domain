package middleware

import (
	"net/http"

	"github.com/boel-go-package/core-domain/cmd/domain/message"
	"github.com/gin-gonic/gin"
)

type CorrelationIdMiddleware struct {
}

func (c CorrelationIdMiddleware) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.Request.Header.Get("X-Correlation-ID")
		if correlationId == "" {

			res := message.Failed("400_CORRELATION_ID_REQUIRED", "X-Correlation-ID is required", nil, nil)

			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}

		c.Request.Header.Set("X-Correlation-ID", correlationId)
		c.Writer.Header().Set("X-Correlation-ID", correlationId)

		c.Next()
	}
}
