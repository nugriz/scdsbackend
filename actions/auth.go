package actions

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthLogin default implementation.
func AuthLogin(c buffalo.Context) error {
	// User Model
	user := &models.User{}
	// Get the JWT Key Secret from .env file.
	secret := os.Getenv("JWT_SECRET")

	// Use Bind function to User model.
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	// Save var of Request JSON Post.
	username := user.Username
	password := user.Password

	// We check if the username or password are not empty.
	if username == "" || password == "" {
		return c.Error(http.StatusBadRequest, errors.New("Username and password cannot be empty"))
	}

	// Get the DB connection from the context.
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Find user with the username.
	q := tx.Select("id, username, password").Where("username= ?", username)
	err := q.First(user)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email.
			return c.Error(http.StatusUnauthorized, errors.New("Invalid username or password"))
		}
		return errors.WithStack(err)
	}

	// Get hashed password from db.
	PasswordHash := user.Password

	// Confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(PasswordHash), []byte(password))
	if err != nil {
		return c.Error(http.StatusUnauthorized, errors.New("Invalid username or password"))
	}

	// Generate token with 2 hours expiration time.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(http.StatusAccepted, r.Auto(c, map[string]string{"token": tokenString}))
}
