package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rentziass/resumaker/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
