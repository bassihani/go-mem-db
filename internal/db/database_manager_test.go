package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//created mock engine

type MockEngine struct{}

func (m *MockEngine) Save(k string, v any) error { return nil }
func (m *MockEngine) Fetch(k string) (any, bool) { return "MockData", true }
func (m *MockEngine) Delete(k string) error      { return nil }

func Test_HistoryLogs(t *testing.T) {
	MockEngine := MockEngine{}
	engine := NewDBManager(&MockEngine)
	engine.Set("Student1", "Richards")
	assert.Equal(t, 1, len(engine.history))
	value := engine.history[0]
	assert.Equal(t, value.Action, "Save")
	assert.Equal(t, value.Success, true)
	assert.NotEmpty(t, value.ID)
}
