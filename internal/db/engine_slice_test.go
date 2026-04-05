package db

import (
	"storageEngine/internal/models"
	"testing"
	"time"
)

type student struct {
	Name string
	Age  int
}

// table testing method
func TestEngineSlice_SaveAndFetch(t *testing.T) {
	// 1. Setup: Initialize the engine
	engine := &EngineSlice{
		store_list: make([]models.Record, 0),
	}

	// 2. Define Test Cases (Table-Driven Design)
	tests := []struct {
		name          string
		key           string
		value         any
		expectedValue any
		shouldFound   bool
	}{
		{
			name:          "Save and Fetch String",
			key:           "user_1",
			value:         "Alice",
			expectedValue: "Alice",
			shouldFound:   true,
		},
		{
			name:          "Key Not Found",
			key:           "non_existent",
			value:         nil,
			expectedValue: nil,
			shouldFound:   false,
		},
		{
			name:          "Save and Fetch Integer",
			key:           "user_2",
			value:         100,
			expectedValue: 100,
			shouldFound:   true,
		},
	}

	// 3. Execution Loop
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.value != nil {
				engine.Save(tc.key, tc.value)
			}

			val, found := engine.Fetch(tc.key)

			// Assertions
			if found != tc.shouldFound {
				t.Errorf("Expected found %v, got %v", tc.shouldFound, found)
			}

			if val != tc.expectedValue {
				t.Errorf("Expected value %v, got %v", tc.expectedValue, val)
			}
		})
	}
}

func TestDeleteRecord(t *testing.T) {
	// initalize the slice engine package and student struct
	record := NewEngineSlice()
	// intialized student structs
	s1 := student{
		Name: "Student1",
		Age:  23,
	}
	s2 := student{
		Name: "Student2",
		Age:  32,
	}
	// added student struct to slice
	record.store_list = append(record.store_list, models.Record{
		Key:       "Student1",
		Value:     s1,
		Timestamp: time.Now(),
	})
	record.store_list = append(record.store_list, models.Record{
		Key:       "Student2",
		Value:     s2,
		Timestamp: time.Now(),
	})
	err := record.Delete("Student2")
	if err != nil {
		t.Fatalf("Deletion is failed %v ", err)
	}

	// will fetch the value for same key which is deleted it has to be false
	_, found := record.Fetch("Student2")
	if found {
		t.Errorf("Record is Fetched so Deletion is not working%v ", err)
	}
}
