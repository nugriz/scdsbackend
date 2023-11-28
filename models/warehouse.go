package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Warehouse is used by pop to map your warehouses database table to your go code.
type Warehouse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Location  string    `json:"location" db:"location"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	SellerID  uuid.UUID `json:"seller_id" db:"seller_id"`
}

// String is not required by pop and may be deleted
func (w Warehouse) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Warehouses is not required by pop and may be deleted
type Warehouses []Warehouse

// String is not required by pop and may be deleted
func (w Warehouses) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Warehouse) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *Warehouse) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: w.Location, Name: "Location"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Warehouse) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
