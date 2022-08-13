package merchandise

import "github.com/demo/layer/domain/owner"

type Merchandise struct {
	id         Id
	ownerId    owner.Id
	name       name
	itemType   ItemType
	printImage printImage
	price      price
	printArea  PrintArea
}

func NewMerchandise() (*Merchandise, error) {
	return &Merchandise{}, nil
}
