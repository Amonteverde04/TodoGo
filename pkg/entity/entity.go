package entity

import (
	"time"
)

// Reusable and expected fields for ALL database entities
type Entity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
