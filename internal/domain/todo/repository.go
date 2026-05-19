package todo

import "chi-recap/internal/domain/shared"

// Repository defines the contract for todo persistence
type Repository interface {
	Save(todo *Todo) error
	GetByID(id *shared.ID) (*Todo, error)
	GetAll() ([]*Todo, error)
	Delete(id *shared.ID) error
	Update(todo *Todo) error
}
