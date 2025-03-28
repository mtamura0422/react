// mysql_handler_test.go
package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
TestNewDB DB接続テスト
*/
func TestNewDB(t *testing.T) {
	_, err := NewTestDB()
	assert.NoError(t, err)

}

/*
TestCloseDB DBクローズテスト
*/
func TestCloseDB(t *testing.T) {
	_, err := NewTestDB()
	assert.NoError(t, err)

	err = CloseDB()
	assert.NoError(t, err)

}
