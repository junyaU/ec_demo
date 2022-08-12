package merchandise

type Merchandise struct {
	id        Id
	name      string
	logo      string
	itemType  ItemType
	printArea PrintArea
}

func NewMerchandise() (*Merchandise, error) {
	return &Merchandise{}, nil
}
