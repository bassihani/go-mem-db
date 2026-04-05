package db

import (
	"storageEngine/internal/models"
	"time"
)

type EngineMap struct {
	engineMap map[string]models.Record
}

// Constructor function to ensure the map is initialized
func NewEngineMap() *EngineMap {
	return &EngineMap{
		engineMap: make(map[string]models.Record),
	}
}

func (e *EngineMap) Save(key string, value any) error {
	e.engineMap[key] = models.Record{
		Key:       key,
		Value:     value,
		Timestamp: time.Now(),
	}
	return nil
}

func (e *EngineMap) Fetch(key string) (any, bool) {
	record, ok := e.engineMap[key]
	if !ok {
		return nil, false
	}
	return record.Value, true
}

func (e *EngineMap) Delete(key string) error {
	delete(e.engineMap, key)
	return nil
}
