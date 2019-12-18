package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDatabaseSession(t *testing.T) {

	sess, _ := GetDataBaseSession("localhost:27017")

	//s := *mgo.
	if sess == nil {
		t.Fatal("Expected Session but got Empty")
	}
}

func TestGetDatabaseSessionErr(t *testing.T) {

	oldMgoDial := MgoDial

	defer func() { MgoDial = oldMgoDial }()

	assert.Panics(t, func() { GetDataBaseSession("localhost:2712017") }, "--------")
}
