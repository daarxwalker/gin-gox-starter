package home_view

import (
	"context"
	"fmt"
	
	. "github.com/daarxwalker/gox"
	
	"common/pkg/component"
)

func Home(c context.Context, ok bool) Node {
	return component.Layout(
		c,
		Div(
			Class("bg-slate-100 text-slate-900 w-screen h-screen grid place-items-center text-center font-semibold"),
			Text(fmt.Sprintf("Go Gin Gox Starter: %t", ok)),
		),
	)
}
