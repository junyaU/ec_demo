package popup

type Store struct {
	ownerId string
}

func NewStore() *Store {
	return &Store{}
}
