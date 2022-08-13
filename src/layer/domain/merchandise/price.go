package merchandise

import (
	"errors"
	"math"
)

type price struct {
	cost   int
	margin int
}

func newPrice(itemType ItemType, toribun int) (price, error) {
	if toribun < 0 || toribun > 5000 {
		return price{}, errors.New("margin cannot exceed the range of 0~5000 yen")
	}

	var itemCost int
	switch itemType {
	case T_SHIRT, SWEAT_SHIRT, HOODIE, JACKET:
		itemCost = 2200

	case CAP, BUCKET_HAT, SANDALS:
		itemCost = 1600

	case PHONE_CASE, TOTE_BAG:
		itemCost = 1400
	}

	return price{
		cost:   itemCost,
		margin: toribun,
	}, nil
}

func (p price) displayTotalAmount() int {
	total := p.cost + p.margin
	tax := 1.1
	return int(math.Floor(float64(total) * tax))
}
