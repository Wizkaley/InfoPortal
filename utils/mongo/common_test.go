package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabaseSession(t *testing.T) {

	sess, _ := GetDataBaseSession("localhost:27017")

	//s := *mgo.
	dbName, _ := sess.DatabaseNames()

	var name string
	for _, val := range dbName {
		if val == "trial" {
			name = val
		}
		continue
	}

	assert.Equalf(t, "trial", name, "Expected %s but got %s", "trial", name)
}

func TestGetDatabaseSessionErr(t *testing.T) {

	oldMgoDial := mgoDial

	defer func() { mgoDial = oldMgoDial }()

	assert.Panics(t, func() { GetDataBaseSession("localhost:2712017") }, "--------")
}

func testCommon(t *testing.T) {
	t.Run("TestGetDatabaseSession", TestGetDatabaseSession)
	t.Run("TestGetDatabaseSessionErr", TestGetDatabaseSessionErr)
}
