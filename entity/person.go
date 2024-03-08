// Package entities contains the entities that are used in the application.
package entity

import "github.com/google/uuid"

type Person struct {
	// ID is the unique identifier of the person.
	ID uuid.UUID
	Name string
	Age int
}