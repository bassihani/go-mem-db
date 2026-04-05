package models

import "time"

type Record struct {
	Key       string
	Value     any
	Timestamp time.Time
}
