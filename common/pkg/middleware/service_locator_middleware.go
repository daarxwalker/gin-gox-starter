package middleware

import (
	"github.com/gin-gonic/gin"
)

func ServiceLocator(services func(c *gin.Context) map[string]any) gin.HandlerFunc {
	return func(c *gin.Context) {
		for serviceToken, service := range services(c) {
			c.Set(serviceToken, service)
		}
		c.Next()
	}
}
