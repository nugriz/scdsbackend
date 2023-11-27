package actions

import (
	"net/http"
	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// InventoriesCreate default implementation.
func InventoriesCreate(c buffalo.Context) error {
	inventory := &models.Inventory{}

	// Bind user to the json elements.
	if err := c.Bind(inventory); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	verrs, err := tx.ValidateAndCreate(inventory)
	if err != nil {
		return errors.WithStack(err)
	}

	// verrs.HasAny returns true/false depending on whether any errors
	// have been tracked.
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusOK, r.Auto(c, map[string]string{"message": "Inventory Created"}))
}

// InventoriesUpdate default implementation.
func InventoriesUpdate(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("inventories/update.html"))
}

// InventoriesShow default implementation.
func InventoriesShow(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("inventories/show.html"))
}

// InventoriesIndex default implementation.
func InventoriesIndex(c buffalo.Context) error {
	warehouse_id := c.Param("warehouse_id")
	inventories := []models.Inventory{}
	query := models.DB.Where("supplier_id in (?)", warehouse_id)
	err := query.All(&inventories)

	// Retrieve all Users from the DB. Select all except password.
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, inventories))
}

// InventoriesDelete default implementation.
func InventoriesDelete(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("inventories/delete.html"))
}
