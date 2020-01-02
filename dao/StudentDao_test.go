package dao

import (
	mocks "RESTApp/mocks"
	"RESTApp/model"
	"RESTApp/utils"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddStudent(t *testing.T) {
	// utils.InitConfig()
	gDB, _ := utils.GetDataBaseSessionWithURI("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)

	mockSColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("Students").Return(mockSColl).AnyTimes()
	err := errors.New("Insert Error")
	mockSColl.EXPECT().Insert(model.Student{}).Return(err).AnyTimes()

	s := model.Student{}
	_ = AddStudent(s, gDB, "testing")

	mockSColl.EXPECT().Insert(gomock.Any()).Return(nil).AnyTimes()
	_ = AddStudent(s, gDB, "testing")

	mockCtrl.Finish()
	gDB.Close()

}

func TestRemoveStudent(t *testing.T) {
	// utils.InitConfig()
	gDB, _ := utils.GetDataBaseSessionWithURI("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)

	mockSColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	err := errors.New("Remove Error")
	mockMongo.EXPECT().C("Students").Return(mockSColl).AnyTimes()
	mockSColl.EXPECT().Remove(gomock.Any()).Return(err).AnyTimes()

	_ = RemoveByName("test", gDB, "testing")

	mockSColl.EXPECT().Remove(gomock.Any()).Return(nil).AnyTimes()

	_ = RemoveByName("test", gDB, "testing")

	mockCtrl.Finish()
}

func TestGetByName(t *testing.T) {
	// utils.InitConfig()
	gDB, _ := utils.GetDataBaseSessionWithURI("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)

	mockSColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("Students").Return(mockSColl).AnyTimes()

	mockFOQry := mocks.NewMockMgoQueryDAL(mockCtrl)
	mockSColl.EXPECT().Find(gomock.Any()).Return(mockFOQry).AnyTimes()

	err := errors.New("Find One Mock Error")
	mockFOQry.EXPECT().One(gomock.Any()).Return(err).AnyTimes()

	_, _ = GetByName("ASAP", gDB, testingdb)

	mockFOQry.EXPECT().One(gomock.Any()).Return(nil).AnyTimes()

	_, _ = GetByName("ASAP", gDB, testingdb)

	mockCtrl.Finish()
}

func TestGetAll(t *testing.T) {
	// utils.InitConfig()
	gDB, _ := utils.GetDataBaseSessionWithURI("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)

	mockSColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("Students").Return(mockSColl).AnyTimes()

	mockFOQry := mocks.NewMockMgoQueryDAL(mockCtrl)
	mockSColl.EXPECT().Find(gomock.Any()).Return(mockFOQry).AnyTimes()

	err := errors.New("Find One Mock Error")
	mockFOQry.EXPECT().All(gomock.Any()).Return(err).AnyTimes()

	_, _ = GetAll(gDB, "testing")

	mockFOQry.EXPECT().All(gomock.Any()).Return(nil).AnyTimes()

	_, _ = GetAll(gDB, "testing")
	mockCtrl.Finish()
}

func TestUpdateStudent(t *testing.T) {
	// utils.InitConfig()
	gDB, _ := utils.GetDataBaseSessionWithURI("localhost:27017")

	mockCtrl := gomock.NewController(t)
	mockMongo := mocks.NewMockMgoDBDAL(mockCtrl)

	mockSColl := mocks.NewMockMgoCollectionDAL(mockCtrl)
	mockMongo.EXPECT().C("Students").Return(mockSColl).AnyTimes()

	err := errors.New("Update Mock Error")
	mockSColl.EXPECT().Update(gomock.Any(), gomock.Any()).Return(err).AnyTimes()

	var s model.Student
	_ = UpdateStudent(s, gDB, "testing")

	mockSColl.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	_ = UpdateStudent(s, gDB, "testing")

	mockCtrl.Finish()
}
