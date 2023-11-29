package actions

import (
	"net/http"
	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// OrdersCreate default implementation.
func OrdersCreate(c buffalo.Context) error {
	order := &models.Order{}

	// Bind user to the json elements.
	if err := c.Bind(order); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	verrs, err := tx.ValidateAndCreate(order)
	if err != nil {
		return errors.WithStack(err)
	}

	// verrs.HasAny returns true/false depending on whether any errors
	// have been tracked.
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusOK, r.Auto(c, map[string]string{"message": "Order Created"}))
}

// OrdersIndex default implementation.
func OrdersList(c buffalo.Context) error {
	o := []models.Order{}
	err := models.DB.Order("id desc").All(&o)
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, o))
}

func OrdersIndex(c buffalo.Context) error {
	buyer_id := c.Param("buyer_id")
	orders := []models.Order{}
	query := models.DB.Where("buyer_id in (?)", buyer_id)
	err := query.All(&orders)

	// Retrieve all Users from the DB. Select all except password.
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, orders))
}

// OrdersShow default implementation.
func OrdersShow(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("orders/show.html"))
}
