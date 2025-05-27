package example_presenter

import (
	"github.com/gin-gonic/gin"
	
	"example/internal/presentation/example_presenter/example_handler"
)

func Register(app gin.IRouter) {
	app.GET("/", example_handler.HandleHome())
	app.GET("/form", example_handler.HandleForm())
	app.POST("/form", example_handler.HandleForm())
}
