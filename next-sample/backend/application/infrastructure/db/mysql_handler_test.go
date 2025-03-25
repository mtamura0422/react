package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	_, err := NewTestDB()
	assert.NoError(t, err)

}

func TestCloseDB(t *testing.T) {
	_, err := NewTestDB()
	assert.NoError(t, err)

	err = CloseDB()
	assert.NoError(t, err)

}
