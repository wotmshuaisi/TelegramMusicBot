package music

// Item music item struct
type Item struct {
	ID        string // uuid4
	URL       string
	Title     string
	Performer string
	Duration  int
}

// API interface of music API
type API interface {
	ListItem(text string) (*[]*Item, error)
	GetURL(id string) (string, error)
	RemoveItem(l []*Item, index int) []*Item
	// ...
}
