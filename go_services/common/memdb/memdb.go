package memdb

import (
	"math/rand"
	"strconv"
)

// This will be an in-memory key value store that holds some data
// This will mostly be for testing purposes, to test APIs before we connect to a remote database service

type Entry struct {
	data any
}

type Db struct {
	Store map[string]Entry
}

func NewKv() *Db {
	newStore := &Db{
		map[string]Entry{},
	}

	populateStore(newStore)

	return newStore
}

func UpsertElement(db *Db, key string, data any) {
	db.Store[key] = Entry{
		data,
	}
}

func populateStore(db *Db) {
	names := []string{
		"Bill", "Bob", "Jack", "Jill",
	}
	eyeColors := []string{
		"Black", "Blue", "Brown", "Green",
	}
	type person struct {
		id       int
		name     string
		eyeColor string
	}
	for i := 0; i < 10; i++ {
		newPerson := person{
			id:       i,
			name:     names[rand.Intn(len(names))],
			eyeColor: eyeColors[rand.Intn(len(eyeColors))],
		}
		UpsertElement(db, strconv.Itoa(i), newPerson)
	}
}
