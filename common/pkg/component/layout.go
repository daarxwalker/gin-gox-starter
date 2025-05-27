package component

import (
	"context"
	
	. "github.com/daarxwalker/gox"
	
	"common/pkg/facade"
)

func Layout(c context.Context, children ...Node) Node {
	return Html(
		Data(Name("theme"), Value("light")),
		Head(
			Title(Text(facade.Page(c).GetTitle())),
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
			Class("text-slate-900"),
			Fragment(children...),
		),
	)
}
