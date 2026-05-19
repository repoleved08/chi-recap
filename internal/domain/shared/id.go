package shared

import "github.com/google/uuid"

// ID is a value object representing a unique identifier
type ID struct {
	value string
}

// NewID creates a new ID with a random UUID
func NewID() *ID {
	return &ID{value: uuid.New().String()}
}

// NewIDFromString creates an ID from a string
func NewIDFromString(id string) *ID {
	return &ID{value: id}
}

// String returns the string representation of the ID
func (id *ID) String() string {
	return id.value
}

// Equals checks if two IDs are equal
func (id *ID) Equals(other *ID) bool {
	return id.value == other.value
}
