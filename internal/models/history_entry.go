package models

import (
	"time"
)

type HistoryEntry struct {
	ID        string
	Action    string
	Timestamp time.Time
	Key       string
	Success   bool
}
