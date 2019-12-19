package utils

import (
	"testing"
)

func TestGetDatabaseSession(t *testing.T) {

	InitConfig()
	sess, err := GetDataBaseSession()
	if err != nil {
		t.Fatal("Expected Session but got Empty")
	}
	sess.Close()
}
