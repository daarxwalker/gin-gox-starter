package component

import (
	"context"
	
	. "github.com/daarxwalker/gox"
	
	"common/pkg/facade"
)

func Layout(c context.Context, children ...Node) Node {
	return Html(
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
			Fragment(children...),
		),
	)
}
