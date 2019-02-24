package music

import (
	"github.com/google/uuid"
)

// Item music item struct
type Item struct {
	ID        uuid.UUID // uuid4
	URL       string
	Title     string
	Performer string
}

// API interface of music API
type API interface {
	List(text string) (*[]Item, error)
	// Get(text string) (PerMusic, error)
	// ...
}
