package utils

import (
	"testing"
)

func TestGetDatabaseSession(t *testing.T) {

	//InitConfig()
	sess, err := GetDataBaseSessionWithURI("localhost:27017")
	if err != nil {
		t.Fatal("Expected Session but got Empty")
	}
	sess.Close()
}
