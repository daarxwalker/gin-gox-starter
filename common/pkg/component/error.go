package component

import (
	"context"
	
	. "github.com/daarxwalker/gox"
)

func Error(c context.Context, status int, err error) Node {
	return Html(
		Head(
			Title(Text(status)),
			Meta(Name("viewport"), Content("width=device-width,initial-scale=1.0")),
			Link(
				Rel("icon"),
				Href("/static/favicon.ico"),
			),
			Link(
				Rel("stylesheet"),
				Type("text/css"),
				Href(GetAssetsURL(c, "css")),
			),
			Script(
				Defer(),
				Src(GetAssetsURL(c, "js")),
			),
		),
		Body(
			Div(
				Class("bg-slate-100 w-screen h-screen grid place-items-center text-center text-red-600"),
				Div(
					Class("space-y-4"),
					Div(Class("font-semibold"), Text(status)),
					Div(Class("text-sm"), Text(err)),
				),
			),
		),
	)
}
