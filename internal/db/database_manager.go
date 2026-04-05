package db

import (
	"errors"
	"fmt"
	"storageEngine/internal/models"
	"storageEngine/pkg/utils"
	"time"
)

// it takes 2 things interface like which engine
// log enteries
type DBManager struct {
	engine  Storage
	history []models.HistoryEntry
}

func NewDBManager(selectedEngine Storage) *DBManager {
	return &DBManager{
		// Dependency Injection
		engine: selectedEngine,
		// taking 100 as generic we can make it on user to define it
		history: make([]models.HistoryEntry, 0, 100),
	}
}

func (d *DBManager) Set(key string, data any) error {
	trimmedKey, ok := utils.CheckKey(key)
	if !ok {
		return errors.New("key should not be empty")
	}
	// delegating to storage engine
	err := d.engine.Save(trimmedKey, data)
	if err != nil {
		return err
	}
	// saving to history
	d.logs("Save", trimmedKey, true)
	return nil
}

func (d *DBManager) Get(key string) (any, error) {
	trimmedKey, ok := utils.CheckKey(key)
	if !ok {
		return nil, errors.New("key should not be empty")
	}
	// delegating to storage engine
	data, ok := d.engine.Fetch(trimmedKey)
	if !ok {
		return nil, fmt.Errorf("Not able to fetch the record for Key %s", key)
	}
	// saving to history
	d.logs("Fetch", trimmedKey, true)
	return data, nil
}

func (d *DBManager) Remove(key string) error {
	trimmedKey, ok := utils.CheckKey(key)
	if !ok {
		return errors.New("key should not be empty")
	}
	err := d.engine.Delete(trimmedKey)
	if err != nil {
		return err
	}
	d.logs("Delete", trimmedKey, true)
	return nil
}

func (d *DBManager) logs(action string, key string, success bool) {
	d.history = append(d.history, models.HistoryEntry{
		ID:        utils.GenerateID(),
		Action:    action,
		Key:       key,
		Timestamp: time.Now(),
		Success:   success,
	})
}
