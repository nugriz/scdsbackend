package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Order is used by pop to map your orders database table to your go code.
type Order struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	BuyerID   uuid.UUID `json:"buyer_id" db:"buyer_id"`
	Quantity  int		`json:"quantity" db:"quantity"`
	Total	  int		`json:"total" db:"total"`
}

// String is not required by pop and may be deleted
func (o Order) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Orders is not required by pop and may be deleted
type Orders []Order

// String is not required by pop and may be deleted
func (o Orders) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *Order) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *Order) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.UUIDIsPresent{Field: o.ProductID, Name: "ProductID"},
		&validators.UUIDIsPresent{Field: o.BuyerID, Name: "BuyerID"},
		&validators.IntIsPresent{Field: o.Quantity, Name: "Quantity"},
		&validators.IntIsPresent{Field: o.Total, Name: "Total"},
	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *Order) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
