package database

import (
	"context"
	"fmt"
	"github.com/zacscoding/gotools-example/person/model"
	"sync"
	"time"
)

type InmemoryPersonDB struct {
	idSeq   uint
	mutex   sync.Mutex
	persons map[string]*model.Person
}

// NewInmemoryPersonDB returns a in-memory based PersonDB.
func NewInmemoryPersonDB() PersonDB {
	db := &InmemoryPersonDB{
		persons: make(map[string]*model.Person),
	}

	for i := 1; i < 5; i++ {
		_ = db.Save(context.TODO(), &model.Person{
			Name:  fmt.Sprintf("user-%d", i),
			Email: fmt.Sprintf("user%d@gmail.com", i),
		})
	}
	return db
}

func (i *InmemoryPersonDB) Save(_ context.Context, person *model.Person) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if _, ok := i.persons[person.Email]; ok {
		return ErrKeyConflict
	}

	data := clonePerson(person)
	data.ID = i.idSeq
	i.idSeq++
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	i.persons[data.Email] = data
	return nil
}

func (i *InmemoryPersonDB) FindByEmail(_ context.Context, email string) (*model.Person, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	p, ok := i.persons[email]
	if !ok {
		return nil, ErrNotFound
	}
	return clonePerson(p), nil
}

func clonePerson(p *model.Person) *model.Person {
	return &model.Person{
		ID:        p.ID,
		Name:      p.Name,
		Email:     p.Email,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
