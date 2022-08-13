package usecase

import "github.com/demo/layer/domain"

type EventRepository interface {
	AppendToStream(e domain.Eventer) error
	LoadEventStream(id string) ([]domain.EventModel, error)
}
