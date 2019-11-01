package mongodal_test

import (
	mocks "RESTApp/mocks"
	"RESTApp/utils/mongodal"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMongoSession(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sess := mocks.NewMockMgoSessionDAL(mockCtrl)
	var dsa mongodal.MgoDBDAL
	gomock.InOrder(sess.EXPECT().DB("planes").Return(dsa))
	sess.DB("planes")
}

func TestMongoSessDAL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockSess := mocks.NewMockMgoSessionDAL(mockCtrl)
	dbdal := mocks.NewMockMgoDBDAL(mockCtrl)
	mockSess.EXPECT().DB("trial").Return(dbdal).Times(1)
	mockSess.DB("trial")
}

func TestCollectionDAL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := mocks.NewMockMgoDBDAL(mockCtrl)
	coll := mocks.NewMockMgoCollectionDAL(mockCtrl)
	collName := "planes"
	db.EXPECT().C(collName).Return(coll).Times(1)
	db.C(collName)
}

func TestQueryDAL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := mocks.NewMockMgoCollectionDAL(mockCtrl)
	colll := mocks.NewMockMgoQueryDAL(mockCtrl)
	//db.EXPECT().Find(bson.M{"studentName": "ASAP"}).Return(colll).Times(1)
	db.EXPECT().Find(gomock.Any()).Return(colll).Times(1)
	//db.Find(bson.M{"studentName": "ASAP"})
	db.Find(gomock.Any())
}
