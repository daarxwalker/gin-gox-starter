package middleware

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"common/pkg/component"
	"common/pkg/env"
	"common/pkg/facade"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if env.Development() {
					fmt.Print(err)
				}
				status := c.Writer.Status()
				if status < 400 {
					status = http.StatusInternalServerError
				}
				facade.Gox(c).Render(status, component.Error(c, status, err.(error)))
				c.Abort()
				return
			}
		}()
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		status := c.Writer.Status()
		facade.Gox(c).Render(status, component.Error(c, status, err))
	}
}
