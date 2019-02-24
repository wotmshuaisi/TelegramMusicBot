package music

// Item music item struct
type Item struct {
	ID        string // uuid4
	URL       string
	Title     string
	Performer string
}

// API interface of music API
type API interface {
	List(text string) (*[]*Item, error)
	Get(id string) (*Item, error)
	// ...
}
