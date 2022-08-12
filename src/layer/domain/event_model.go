package domain

import (
	"encoding/json"
	"time"
)

type EventModel struct {
	ID        string
	Version   uint64
	CreatedAt time.Time
	EventType string
	EventData json.RawMessage
}

func NewEventModel(id string, version uint64, eventType string, eventData json.RawMessage) EventModel {
	return EventModel{
		ID:        id,
		Version:   version,
		EventType: eventType,
		EventData: eventData,
		CreatedAt: time.Now(),
	}
}

func (m EventModel) GetVersion() uint64 {
	return m.Version
}
