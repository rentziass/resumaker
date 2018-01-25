package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	if UserLoggedIn(c) {
		return c.Redirect(302, "/dashboard")
	}
	return c.Render(200, r.HTML("index.html"))
}
