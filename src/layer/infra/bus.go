package infra

import (
	"github.com/demo/layer/adapter"
	"github.com/demo/layer/usecase/popup"
	"github.com/junyaU/evbus"
)

var _ adapter.MessagePublisher = (evbus.Bus)(nil)

func InitBus() evbus.Bus {
	bus := evbus.New()
	db := NewDB()
	eventStore := adapter.NewEventStore(db)

	exhibitHandler := popup.NewExhibit(eventStore)
	hostHandler := popup.NewHost(eventStore)

	bus.Subscribe("popup:exhibit", exhibitHandler.Exec)
	bus.Subscribe("popup:host", hostHandler.Exec)

	return bus
}
