package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/x/responder"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	"github.com/rentziass/resumaker/models"
)

func DashboardIndex(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")
	u := &models.User{}
	tx := c.Value("tx").(*pop.Connection)
	err := tx.Find(u, uid)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("user", u)
	return c.Render(200, r.HTML("dashboard/index.html"))
}

func DashboardUpdate(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")
	u := &models.User{}
	tx := c.Value("tx").(*pop.Connection)
	err := tx.Find(u, uid)
	if err != nil {
		return errors.WithStack(err)
	}

	// Bind User to the html form elements
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("No transaction found"))
	}

	// Update user
	verrs, err := tx.ValidateAndUpdate(u)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make user available inside the template
		c.Set("user", u)

		// Make the errors available inside the template
		c.Set("errors", verrs)

		// Render again the dashboard/index.html
		res := responder.Wants("html", func(c buffalo.Context) error {
			return c.Render(422, r.HTML("dashboard/index.html"))
		})
		return res.Respond(c)
	}

	c.Flash().Add("success", "Profile successfully updated!")
	return c.Redirect(302, "/dashboard")
}
