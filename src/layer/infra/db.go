package infra

import (
	"github.com/demo/layer/adapter"
	"github.com/demo/layer/domain"
)

type db struct {
	conn interface{}
}

func NewDB() adapter.DataHandler {
	return &db{}
}

func (d *db) Push(table string, value interface{}) error {
	return nil
}

func (d *db) Get(table, hashKey string) ([]domain.EventModel, error) {
	return []domain.EventModel{}, nil
}
