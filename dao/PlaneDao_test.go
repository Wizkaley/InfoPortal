package dao

import (
	mocks "RESTApp/mocks"
	"RESTApp/model"
	"RESTApp/mongodal"
	"RESTApp/utils"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//var gDB *mgo.Session

var mockMongo *mocks.MockMgoDBDAL
var testingdb = "testing"

func getMockMongoDAL(database *mgo.Database) mongodal.MgoDBDAL {
	return mockMongo
}

func TestPutPlane(t *testing.T) {
	gDB, _ := utils.GetDataBaseSession("localhost:27017")

	NewMongoDBDAL = getMockMongoDAL

	mockCtrl := gomock.NewController(t)
	mockMongo = mocks.NewMockMgoDBDAL(mockCtrl)

	// Error Condition Check
	mockPColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("planes").Return(mockPColl).Times(2)
	err := errors.New("insert error")

	mockPColl.EXPECT().Insert(gomock.Any()).Return(err).Times(1)

	plane := model.Plane{}

	_ = PutPlane(plane, gDB, "testing")

	// Success Condition Check
	mockPColl.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	_ = PutPlane(plane, gDB, "testing")

	// Cleanup
	gDB.Close()
	mockCtrl.Finish()
	NewMongoDBDAL = mongodal.NewMongoDBDAL
}

func TestGetPlane(t *testing.T) {
	gDB, _ := utils.GetDataBaseSession("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)
	mockPColl := mocks.NewMockMgoCollectionDAL(mockCtrl)

	mockFOQry := mocks.NewMockMgoQueryDAL(mockCtrl)

	// Error Condition Check
	err := errors.New("Find Error")
	p := model.Plane{}
	mockMongo.EXPECT().C("planes").Return(mockPColl).AnyTimes()
	mockPColl.EXPECT().Find(gomock.Any()).Return(mockFOQry).AnyTimes()
	mockFOQry.EXPECT().One(&p).Return(err).AnyTimes()

	_, _ = GetPlane("MIG-21", gDB, "testing")

	// Success Condition Check
	mockFOQry.EXPECT().One(&p).Return(nil).AnyTimes()
	// mockCtrl.Finish()

	_, _ = GetPlane("MIG-21", gDB, "testing")
	mockCtrl.Finish()
	gDB.Close()

}

func TestUpdatePlane(t *testing.T) {

	gDB, _ := utils.GetDataBaseSession("localhost:27017")
	defer gDB.Close()

	//var p model.Plane
	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)
	mockPColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("planes").Return(mockPColl).AnyTimes()

	err := errors.New("update error")
	pl := model.Plane{
		Pid:      8,
		Name:     "Boeing 777",
		NoWheels: 24,
		Engines:  8,
		PType:    "Cargo",
	}

	mockPColl.EXPECT().Update(bson.M{"name": pl.Name}, pl).Return(err).AnyTimes()
	_, _ = UpdatePlane(pl, gDB, "testing")

	mockPColl.EXPECT().Update(bson.M{"name": pl.Name}, pl).Return(nil).AnyTimes()
	_, _ = UpdatePlane(pl, gDB, "testing")
	mockCtrl.Finish()
}

func TestRemovePlane(t *testing.T) {
	gDB, _ := utils.GetDataBaseSession("localhost:27017")
	defer gDB.Close()

	mockCtrl := gomock.NewController(t)
	mockMongo = mocks.NewMockMgoDBDAL(mockCtrl)

	mockPColl := mocks.NewMockMgoCollectionDAL(mockCtrl)

	err := errors.New("remove error")
	mockMongo.EXPECT().C("planes").Return(mockPColl).AnyTimes()
	mockPColl.EXPECT().Remove(gomock.Any()).Return(err).AnyTimes()

	DeletePlane("MIG19", gDB, testingdb)

	mockPColl.EXPECT().Remove(gomock.Any()).Return(nil).AnyTimes()
	DeletePlane("MIG19", gDB, testingdb)
	mockCtrl.Finish()
}

func TestGetAllPlanes(t *testing.T) {
	gDB, _ := utils.GetDataBaseSession("localhost:27017")
	defer gDB.Close()

	mockCtrl := gomock.NewController(t)
	mockPColl := mocks.NewMockMgoCollectionDAL(mockCtrl)

	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)
	mockAllQry := mocks.NewMockMgoQueryDAL(mockCtrl)
	err := errors.New("Get All error")
	mockMongo.EXPECT().C(gomock.Any).Return(mockPColl).Times(1)
	mockPColl.EXPECT().Find(gomock.Any()).Return(mockAllQry).Times(2)
	mockAllQry.EXPECT().All(gomock.Any()).Return(err).Times(1)

	// var pl model.Plane
	_, err = GetAllPlanes(gDB, "testing")
	if err != nil {
		t.Errorf("Test Failed :: %v", err)
	}
	mockAllQry.EXPECT().All(gomock.Any()).Return(nil).Times(1)
	_, err = GetAllPlanes(gDB, "testing")
	if err != nil {
		t.Errorf("Test Failed :: %v", err)
	}
}
func TestDeleteByID(t *testing.T) {
	gDB, _ := utils.GetDataBaseSession("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)
	mockPColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("planes").Return(mockPColl).AnyTimes()
	// err := errors.New("Remove Error")
	// mockPColl.EXPECT().Remove(gomock.Any()).Return(err).AnyTimes()

	// _ = DeletePlaneByID(4, gDB, testingdb)

	mockPColl.EXPECT().Remove(gomock.Any()).Return(nil).AnyTimes()
	_ = DeletePlaneByID(4, gDB, testingdb)
	mockCtrl.Finish()

}
