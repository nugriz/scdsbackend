package actions

import (
	"net/http"
	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// ProductsShow default implementation.
func ProductsShow(c buffalo.Context) error {
	id := c.Param("product_id")

	// create a variable to receive the todo
	product := models.Product{}
	// grab the todo from the database
	err := models.DB.Find(&product, id)
	// handle possible error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	//return the data as json
	return c.Render(http.StatusOK, r.JSON(&product))
}

// ProductsCreate default implementation.
func ProductsCreate(c buffalo.Context) error {
	product := &models.Product{}

	// Bind user to the json elements.
	if err := c.Bind(product); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	verrs, err := tx.ValidateAndCreate(product)
	if err != nil {
		return errors.WithStack(err)
	}

	// verrs.HasAny returns true/false depending on whether any errors
	// have been tracked.
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusOK, r.Auto(c, map[string]string{"message": "Product Created"}))
}

func ProductsIndexBySupplier(c buffalo.Context) error {
	supplier_id := c.Param("supplier_id")
	products := []models.Product{}
	query := models.DB.Where("supplier_id in (?)", supplier_id)
	err := query.All(&products)

	// Retrieve all Users from the DB. Select all except password.
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, products))
}

// ProductsIndex default implementation.
func ProductsIndex(c buffalo.Context) error {
	products := []models.Product{}
	err := models.DB.All(&products)
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, products))
}

// ProductsDelete default implementation.
func ProductsDelete(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("products/delete.html"))
}
