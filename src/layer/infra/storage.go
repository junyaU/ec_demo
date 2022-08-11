package infra

type Storage struct {
	conn interface{}
}

func NewStorage() *Storage {
	return &Storage{}
}
