package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Inventory is used by pop to map your inventories database table to your go code.
type Inventory struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	WarehouseID uuid.UUID `json:"warehouse_id" db:"warehouse_id"`
	ProductID   uuid.UUID `json:"product_id" db:"product_id"`
	Stock       int       `json:"stock" db:"stock"`
}

// String is not required by pop and may be deleted
func (i Inventory) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Inventories is not required by pop and may be deleted
type Inventories []Inventory

// String is not required by pop and may be deleted
func (i Inventories) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Inventory) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Inventory) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.UUIDIsPresent{Field: i.WarehouseID, Name: "WarehouseID"},
		&validators.UUIDIsPresent{Field: i.ProductID, Name: "ProductID"},
		&validators.IntIsPresent{Field: i.Stock, Name: "Stock"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Inventory) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
