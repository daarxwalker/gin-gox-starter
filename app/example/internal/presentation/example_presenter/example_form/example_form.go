package example_form

import "common/pkg/form"

type ExampleForm struct {
	Value  string      `form:"value" binding:"required"`
	Errors form.Errors `form:"-"`
}
