package utils

import (
	"strings"

	"github.com/google/uuid"
)

func CheckKey(key string) (string, bool) {
	trimmedKey := strings.TrimSpace(key)
	if trimmedKey == "" {
		return key, false
	}
	return trimmedKey, true
}

func GenerateID() string {
	return uuid.New().String()
}
