package models

import "time"

// Example of typed string constants for enum fields.
// Use this pattern for any field with a fixed set of valid values.
type ExampleStatus string

const (
	ExampleStatusActive   ExampleStatus = "active"
	ExampleStatusInactive ExampleStatus = "inactive"
)

// Example is a placeholder model. Replace with your domain models.
type Example struct {
	ID        string        `json:"id"`
	Status    ExampleStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
}
