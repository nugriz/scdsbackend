package actions

import (
	"net/http"

	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"

	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	// User Model
	user := &models.User{}

	// Bind user to the json elements.
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}
	// Validate and create the data.
	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return errors.WithStack(err)
	}

	// verrs.HasAny returns true/false depending on whether any errors
	// have been tracked.
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusCreated, r.Auto(c, map[string]string{"message": "User Created"}))
}

// UsersRead default implementation.
// UsersRead default implementation.
func UsersRead(c buffalo.Context) error {
	users := &models.Users{}

	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	// Add Order for date.
	q := tx.PaginateFromParams(c.Params()).Order("created_at asc")

	// Retrieve all Users from the DB. Select all except password.
	if err := q.Select(
		"id",
		"created_at",
		"updated_at",
		"username",
		"email",
	).All(users); err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, users))
}
