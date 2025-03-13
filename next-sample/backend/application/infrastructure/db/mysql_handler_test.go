package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	_, err := NewDBMock()
	assert.NoError(t, err)

}
