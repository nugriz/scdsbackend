package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Chat is used by pop to map your chats database table to your go code.
type Chat struct {
	ID        	uuid.UUID `json:"id" db:"id"`
	CreatedAt 	time.Time `json:"created_at" db:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at" db:"updated_at"`
	SenderID  	uuid.UUID `json:"sender_id" db:"sender_id"`
	ReceiverID	uuid.UUID `json:"receiver_id" db:"receiver_id"`
	Message     string    `json:"message" db:"message"`
}

// String is not required by pop and may be deleted
func (c Chat) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Chats is not required by pop and may be deleted
type Chats []Chat

// String is not required by pop and may be deleted
func (c Chats) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Chat) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Chat) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Message, Name: "Message"},
		&validators.UUIDIsPresent{Field: c.SenderID, Name: "SenderID"},
		&validators.UUIDIsPresent{Field: c.ReceiverID, Name: "ReceiverID"},

	), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Chat) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
