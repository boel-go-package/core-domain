package util

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GetCookie(c *gin.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie
}

func SetCookie(c *gin.Context, name string, value string, expires time.Duration) {
	c.SetCookie(name, value, int(expires.Seconds()), "/", "", false, true)
}

func DeleteCookie(c *gin.Context, name string) {
	c.SetCookie(name, "", -1, "/", "", false, true)
}
