package popup

import (
	"errors"
	"github.com/demo/layer/domain/merchandise"
	"github.com/demo/layer/domain/owner"
	"github.com/demo/layer/domain/popup"
	"github.com/demo/layer/usecase"
)

type exhibit struct {
	store usecase.EventRepository
}

func NewExhibit(r usecase.EventRepository) *exhibit {
	return &exhibit{
		store: r,
	}
}

func (e *exhibit) Exec(command ExhibitCommand) error {
	ownerEvents, err := e.store.LoadEventStream(command.ownerId)
	if err != nil {
		return err
	}
	o, err := owner.Get(ownerEvents)
	if err != nil {
		return err
	}

	popupEvents, err := e.store.LoadEventStream(command.popupId)
	if err != nil {
		return err
	}
	p, err := popup.Get(popupEvents)
	if err != nil {
		return err
	}

	if !p.IsCorrectOwner(o.Identify()) {
		return errors.New("not the correct owner of the pop-up store")
	}

	var ids []merchandise.Id
	for i, id := range command.merchandiseIds {
		merchandiseEvents, err := e.store.LoadEventStream(id)
		if err != nil {
			return err
		}
		m, err := merchandise.Get(merchandiseEvents)
		if err != nil {
			return err
		}

		if !o.CanExhibitMerchandise(m.OwnerIdentify()) {
			return errors.New("contains merchandises that cannot be exhibited")
		}

		ids[i] = m.Identify()
	}

	exhibitedEvent, err := p.ExhibitMerchandise(ids)
	if err != nil {
		return err
	}

	return e.store.AppendToStream(exhibitedEvent)
}

type ExhibitCommand struct {
	popupId        string
	ownerId        string
	merchandiseIds []string
}

func NewExhibitedCommand(popupId, ownerId string, merchandiseId []string) ExhibitCommand {
	return ExhibitCommand{
		popupId:        popupId,
		ownerId:        ownerId,
		merchandiseIds: merchandiseId,
	}

}
