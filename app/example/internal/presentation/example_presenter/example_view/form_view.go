package example_view

import (
	"context"
	"net/http"
	
	. "github.com/daarxwalker/gox"
	
	"common/pkg/component"
	"example/internal/presentation/example_presenter/example_form"
)

func FormView(c context.Context, exampleForm example_form.ExampleForm, success bool) Node {
	return component.Layout(
		c,
		Div(
			Class("bg-slate-100 w-screen h-screen grid place-items-center"),
			Div(
				Class("space-y-8"),
				Div(
					Class("font-semibold"),
					Text("Form"),
				),
				If(
					success,
					Div(
						Class("alert alert-success"),
						Text("Success"),
					),
				),
				Form(
					Method(http.MethodPost),
					Action("/form"),
					Class("space-y-4"),
					Fieldset(
						Class("fieldset"),
						Label(For("value"), Text("Value")),
						Input(
							CLSX{
								"input":       true,
								"input-error": len(exampleForm.Errors.GetErrors("value")) > 0,
							},
							Id("value"),
							Type("text"),
							Name("value"),
							Value(exampleForm.Value),
						),
						If(
							len(exampleForm.Errors.GetErrors("value")) > 0,
							Div(
								Class("text-red-500 text-sm"),
								Text(exampleForm.Errors.GetErrors("value")),
							),
						),
					),
					Button(Type("submit"), Class("btn btn-primary w-full"), Text("Submit")),
				),
				Div(
					Class("text-center"),
					A(Class("text-blue-500 underline hover:no-underline"), Href("/"), Text("Back to homepage")),
				),
			),
		),
	)
}
