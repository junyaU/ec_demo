package popup

import (
	"github.com/demo/layer/domain"
	"github.com/demo/layer/domain/merchandise"
)

const EXHIBITED_EVENT_TYPE = "EXHIBITED"

type ExhibitedEvent struct {
	domain.Event
	merchandiseIds []string
}

func newExhibitedEvent(popupId Id, merchandiseIds []merchandise.Id, version uint64) ExhibitedEvent {
	var ids []string
	for _, id := range merchandiseIds {
		ids = append(ids, id.Identifier())
	}

	return ExhibitedEvent{
		Event:          domain.NewEvent(popupId.Identifier(), version, EXHIBITED_EVENT_TYPE),
		merchandiseIds: ids,
	}
}
