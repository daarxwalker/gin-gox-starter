package example_view

import (
	"context"
	"fmt"
	
	. "github.com/daarxwalker/gox"
	
	"common/pkg/component"
)

func HomeView(c context.Context, ok bool) Node {
	return component.Layout(
		c,
		Div(
			Class("bg-slate-100 w-screen h-screen grid place-items-center text-center"),
			Div(
				Class("space-y-8"),
				Div(
					Class("font-semibold"),
					Text(fmt.Sprintf("Go Gin Gox Starter: %t", ok)),
				),
				Div(
					A(Class("text-blue-500 underline hover:no-underline"), Href("/form"), Text("Form")),
				),
			),
		),
	)
}
