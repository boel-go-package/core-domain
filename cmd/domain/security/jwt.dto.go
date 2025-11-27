package security

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTDTO struct {
	DisplayName string `json:"display_name" form:"display_name"`
	Email       string `json:"email" form:"email"`
	Resource    gin.H  `json:"resource" form:"resource"`
	jwt.RegisteredClaims
}

type JWTRefreshDTO struct {
	jwt.RegisteredClaims
}
