package mongo

import (
	"RESTApp/dao"
	"RESTApp/model"
	"RESTApp/utils/mongo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabaseSession(t *testing.T) {

	sess, _ := mongo.GetDataBaseSession("localhost:27017")

	//s := *mgo.
	tst := model.Student{
		StudentName:  "Pretty",
		StudentAge:   56,
		StudentMarks: 99,
	}
	dao.AddStudent(tst, sess, "trial")
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

	oldMgoDial := mongo.MgoDial

	defer func() { mongo.MgoDial = oldMgoDial }()

	assert.Panics(t, func() { mongo.GetDataBaseSession("localhost:2712017") }, "--------")
}

// func testCommon(t *testing.T) {
// 	t.Run("TestGetDatabaseSession", TestGetDatabaseSession)
// 	t.Run("TestGetDatabaseSessionErr", TestGetDatabaseSessionErr)
// }
