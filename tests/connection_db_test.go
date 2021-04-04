package tests

import (
	"library-system/db"
	"testing"
)

func TestGetConnection(t *testing.T) {
	testConnection := db.GetConnection()
	defer testConnection.Close()
}
