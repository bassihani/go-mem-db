package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CheckKey(t *testing.T) {
	givenValue := "   Student1 "
	key, ok := CheckKey(givenValue)
	assert.Equal(t, ok, true)
	require.NotEqual(t, givenValue, key)
}

func Test_Empty_CheckKey(t *testing.T) {
	givenValue := ""
	key, ok := CheckKey(givenValue)
	assert.Equal(t, ok, false)
	require.Equal(t, givenValue, key)
}

func TestGenerateUUID(t *testing.T) {
	ID := GenerateID()
	require.NotEmpty(t, ID)
	require.IsType(t, "", ID)
}
