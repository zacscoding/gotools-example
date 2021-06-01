package database

import (
	"context"
	"github.com/zacscoding/gotools-example/person/model"
)

// PersonDB is a interface of person to operate with database
// The Save() method is used to persist person data to storage.
//go:generate mockery --name PersonDB --output ./mocks --filename person_mock.go
type PersonDB interface {
	// Save saves a given person to database.
	Save(ctx context.Context, person *model.Person) error

	// FindByEmail returns a person if find by given email, otherwise returns an error.
	FindByEmail(ctx context.Context, email string) (*model.Person, error)
}
