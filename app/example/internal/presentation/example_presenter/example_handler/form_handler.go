package example_handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"common/pkg/form"
	"example/internal/presentation/example_presenter/example_form"
	"example/internal/presentation/example_presenter/example_view"
	
	"common/pkg/facade"
)

func HandleForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		var exampleForm example_form.ExampleForm
		if c.Request.Method == http.MethodPost {
			if err := c.ShouldBind(&exampleForm); err != nil {
				exampleForm.Errors = form.CreateErrors(err)
			}
			if len(exampleForm.Errors) == 0 {
				exampleForm = example_form.ExampleForm{}
			}
		}
		facade.Page(c).SetTitle("Form example")
		facade.Gox(c).Render(
			http.StatusOK,
			example_view.FormView(c, exampleForm, c.Request.Method == http.MethodPost && len(exampleForm.Errors) == 0),
		)
	}
}
