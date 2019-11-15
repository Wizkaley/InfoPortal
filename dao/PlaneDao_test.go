package dao

import (
	"RESTApp/model"
	"RESTApp/utils/mongo"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//var gDB *mgo.Session

var testingdb = "testing"

func TestPutPlane(t *testing.T) {
	testCases := []struct {
		url        string
		db         string
		collection string
		err        error
	}{
		{
			url:        "localhost:27017",
			db:         "testing",
			collection: "planes",
			err:        nil,
		},
	}

	gDB, _ := mongo.GetDataBaseSession(testCases[0].url)
	defer gDB.Close()

	plane := model.Plane{
		Pid:      7,
		Name:     "MIG19",
		NoWheels: 6,
		Engines:  4,
		PType:    "Attack",
	}

	err := PutPlane(plane, gDB, testingdb)
	assert.Equal(t, testCases[0].err, err, "Expected %v but got %v", testCases[0].err, err)

}

func TestGetPlane(t *testing.T) {
	tc := []struct {
		Name string
		err  error
	}{
		{
			Name: "MIG19",
			err:  nil,
		},
	}

	gDB, _ := mongo.GetDataBaseSession("localhost:27017")

	p := GetPlane(tc[0].Name, gDB, testingdb)
	assert.Equal(t, tc[0].Name, p.Name, "Expected %s but got %s", tc[0].Name, p.Name)

	// p1 := GetPlane(tc[1].Name, gDB)
	// assert.Equal
	defer gDB.Close()
}

func TestUpdatePlane(t *testing.T) {
	pl := model.Plane{
		Pid:      8,
		Name:     "Boeing 777",
		NoWheels: 24,
		Engines:  8,
		PType:    "Cargo",
	}

	gDB, _ := mongo.GetDataBaseSession("localhost:27017")

	defer gDB.Close()

	gP := GetPlane(pl.Name, gDB, testingdb)
	//fmt.Println("...............................getPlane", gP)
	gP.Pid = pl.Pid
	gP.NoWheels = pl.NoWheels
	gP.Engines = pl.Engines
	gP.PType = pl.PType
	p, _ := UpdatePlane(gP, gDB)
	//fmt.Println("...............................updatePlane", p)
	assert.Equal(t, p.Pid, pl.Pid, "Exepected %s but got %v", p.Pid, pl.Pid)
}

func TestRemovePlane(t *testing.T) {
	gDB, _ := mongo.GetDataBaseSession("localhost:27017")
	defer gDB.Close()
	err := DeletePlane("MIG19", gDB, testingdb)
	assert.Equalf(t, true, err, "Expected %s but got %s", true, err)
}

func TestRemovePlaneErr(t *testing.T) {
	gDB, _ := mongo.GetDataBaseSession("localhost:27017")
	defer gDB.Close()

	err := DeletePlane("name", gDB, testingdb)
	assert.Error(t, errors.New("not found"), err, "..")
}

func TestGetAllPlanes(t *testing.T) {
	gDB, _ := mongo.GetDataBaseSession("localhost:27017")
	defer gDB.Close()

	var p model.Plane
	p.Pid = 2
	p.Name = "Charter"
	p.NoWheels = 8
	p.Engines = 6
	p.PType = "Commercial"
	PutPlane(p, gDB, testingdb)
	planes, err := GetAllPlanes(gDB, testingdb)

	if planes == nil {
		t.Errorf("Failed to Get All Planes: %v", err)
	}
}

func TestDeleteByID(t *testing.T) {
	gDB, _ := mongo.GetDataBaseSession("localhost:27017")

	// err := DeletePlaneByID(7, gDB, testingdb)
	// assert.Equal(t, true, err)

	err := DeletePlaneByID(4, gDB, testingdb)
	assert.Equal(t, false, err)
}

func TestDeleteByIDErr(t *testing.T) {
	gDB, _ := mongo.GetDataBaseSession("localhost:27017")

	err := DeletePlaneByID(1221321343, gDB, testingdb)
	assert.Error(t, errors.New("not found"), err, "...")
}
