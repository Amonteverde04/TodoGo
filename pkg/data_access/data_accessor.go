package data_access

// Represents a generic blueprint (contract) that a data accessor should follow.
// O is the expected output type.
// I is the expected input type. In otherwords a DTO type of object.
type DataAccessor[O any, I any] interface {
	GetAll() ([]*O, error)
	GetById(id string) (*O, error)
	Add(data *I) (string, error)
	Update(id int, data *I) error
	Delete(id int) error
}
