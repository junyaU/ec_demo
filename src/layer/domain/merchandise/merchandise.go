package merchandise

type Merchandise struct {
}

func NewMerchandise() (*Merchandise, error) {
	return &Merchandise{}, nil
}
