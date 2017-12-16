package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Access-Token")

		if token == "qwer123" {
			c.Next()
		} else {
			c.AbortWithStatus(401)
		}
	}
}
