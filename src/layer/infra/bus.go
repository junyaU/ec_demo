package infra

import (
	"github.com/junyaU/evbus"
)

func InitBus() evbus.Bus {
	bus := evbus.New()

	return bus
}
