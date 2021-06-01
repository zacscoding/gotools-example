package model

import (
	"encoding/json"
	"time"
)

// Person is a model of person
type Person struct {
	// ID is a unique identifier for a person. i.e sequence of database
	ID uint `json:"id"`

	// Name is name of person.
	Name string `json:"name"`

	// Email is emaol of a person.
	Email string `json:"email"`

	// CreatedAt is the saved time.
	CreatedAt time.Time `json:"createdAt"`

	// UpdateAt is the last updated time and equals to CreatedAt until update.
	UpdatedAt time.Time `json:"updatedAt"`
}

// String returns json string value
func (p Person) String() string {
	data, _ := json.Marshal(p)
	return string(data)
}
