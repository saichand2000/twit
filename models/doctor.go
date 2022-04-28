package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Doctor is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Doctor struct {
    ID uuid.UUID `json:"id" db:"id"`
    MigrationFizz string `json:"migration_fizz" db:"migration_fizz"`
    String string `json:"string" db:"string"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d Doctor) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Doctors is not required by pop and may be deleted
type Doctors []Doctor

// String is not required by pop and may be deleted
func (d Doctors) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Doctor) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.MigrationFizz, Name: "MigrationFizz"},
		&validators.StringIsPresent{Field: d.String, Name: "String"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Doctor) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Doctor) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
