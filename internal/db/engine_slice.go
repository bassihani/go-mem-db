package db

import (
	"storageEngine/internal/models"
	"time"
)

type EngineSlice struct {
	store_list []models.Record
}

func NewEngineSlice() *EngineSlice {
	return &EngineSlice{
		store_list: make([]models.Record, 0, 100),
	}
}

func (e *EngineSlice) Save(key string, value any) error {
	var record models.Record
	record = models.Record{
		Key:       key,
		Value:     value,
		Timestamp: time.Now(),
	}
	for i, value := range e.store_list {
		if value.Key == key {
			e.store_list[i] = record
		}
	}
	e.store_list = append(e.store_list, record)
	return nil
}

func (e *EngineSlice) Fetch(key string) (any, bool) {
	for _, value := range e.store_list {
		if value.Key == key {
			return value.Value, true
		}
	}
	return nil, false
}

func (e *EngineSlice) Delete(key string) error {
	for i, value := range e.store_list {
		if value.Key == key {
			e.store_list = append(e.store_list[:i], e.store_list[i+1:]...)
			return nil
		}
	}
	return nil
}
