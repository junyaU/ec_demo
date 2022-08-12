package presenter

import (
	"github.com/junyaU/evbus"
	"github.com/labstack/echo/v4"
)

type Popup struct {
	bus evbus.Bus
}

func NewPopup(bus evbus.Bus) *Popup {
	return &Popup{
		bus: bus,
	}
}

func (p Popup) Create(c echo.Context) error {
	return nil
}
