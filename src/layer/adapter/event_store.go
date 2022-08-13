package adapter

import (
	"encoding/json"
	"errors"
	"github.com/demo/layer/domain"
)

const TABLE_NAME = "dummy"

type DataHandler interface {
	Push(table string, value interface{}) error
	Get(table, hashKey string) ([]domain.EventModel, error)
}

type EventStore struct {
	DataHandler
}

func NewEventStore(dh DataHandler) *EventStore {
	return &EventStore{
		dh,
	}
}

func (s *EventStore) AppendToStream(e domain.Eventer) error {
	isIllegalEvent := e.Identifier() == "" || e.Type() == ""
	if isIllegalEvent {
		return errors.New("不正なイベントを追加することはできません")
	}

	eventStream, err := s.Get(TABLE_NAME, e.Identifier())
	if err != nil {
		return err
	}

	currentVersion := e.Version()
	isExistEventInStream := len(eventStream) > 0

	isUnExpectedVersion := currentVersion != 0 && isExistEventInStream && eventStream[len(eventStream)-1].GetVersion() != currentVersion
	if isUnExpectedVersion {
		return errors.New("期待していないバージョンです")
	}

	eventJSON, err := json.Marshal(e)
	if err != nil {
		return err
	}

	return s.Push(TABLE_NAME, domain.NewEventModel(e.Identifier(), e.Version()+1, e.Type(), eventJSON))
}

func (s *EventStore) LoadEventStream(id string) ([]domain.EventModel, error) {
	return []domain.EventModel{}, nil
}
