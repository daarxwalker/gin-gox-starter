package example_handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"example/internal/presentation/example_presenter/example_view"
	
	"common/pkg/facade"
	"example/internal/infrastructure/repository/example_repository"
)

func HandleHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := example_repository.GetExample(c, facade.DB(c))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		facade.Page(c).SetTitle("Home")
		facade.Gox(c).Render(http.StatusOK, example_view.HomeView(c, result == 1))
	}
}
