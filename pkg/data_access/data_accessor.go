package data_access

import "github.com/google/uuid"

// Represents a generic blueprint (contract) that a data accessor should follow.
// O is the expected output type.
// I is the expected input type. In otherwords a DTO type of object.
type DataAccessor[O any, I any] interface {
	GetAll() ([]O, error)
	GetById(uuid.UUID) (O, error)
	Add(data *I) (string, error)
	Update(data O) error
	Delete(id uuid.UUID) error
}
