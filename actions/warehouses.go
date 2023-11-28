package actions

import (
	"net/http"
	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// WarehousesShow default implementation.
func WarehousesShow(c buffalo.Context) error {
	id := c.Param("warehouse_id")

	// create a variable to receive the todo
	warehouse := models.Warehouse{}
	// grab the todo from the database
	err := models.DB.Find(&warehouse, id)
	// handle possible error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	//return the data as json
	return c.Render(http.StatusOK, r.JSON(&warehouse))
}

// WarehousesIndex default implementation.
func WarehousesIndex(c buffalo.Context) error {
	warehouses := []models.Warehouse{}
	err := models.DB.Order("id desc").All(&warehouses)
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, warehouses))
}

// WarehousesCreate default implementation.
func WarehousesCreate(c buffalo.Context) error {
	warehouse := &models.Warehouse{}

	// Bind user to the json elements.
	if err := c.Bind(warehouse); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	verrs, err := tx.ValidateAndCreate(warehouse)
	if err != nil {
		return errors.WithStack(err)
	}

	// verrs.HasAny returns true/false depending on whether any errors
	// have been tracked.
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusOK, r.Auto(c, map[string]string{"message": "Warehouse Created"}))
}

// WarehousesDelete default implementation.
func WarehousesDelete(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("warehouses/delete.html"))
}
