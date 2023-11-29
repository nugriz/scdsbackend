package actions

import (
	"net/http"
	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// ChatsCreate default implementation.
func ChatsCreate(c buffalo.Context) error {
	chat := &models.Chat{}

	// Bind user to the json elements.
	if err := c.Bind(chat); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	verrs, err := tx.ValidateAndCreate(chat)
	if err != nil {
		return errors.WithStack(err)
	}

	// verrs.HasAny returns true/false depending on whether any errors
	// have been tracked.
	if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusOK, r.Auto(c, map[string]string{"message": "Chat Created"}))
}

// ChatsIndex default implementation.
func ChatsIndex(c buffalo.Context) error {
	chats := []models.Chat{}
	err := models.DB.All(&chats)
	if err != nil {
		return errors.WithStack(err)
	}
	// Add the paginator to the context so it can be used in the template.
	//c.Set("pagination", q.Paginator)

	return c.Render(http.StatusOK, r.Auto(c, chats))
}

