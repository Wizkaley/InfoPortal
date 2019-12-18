package dao

import (
	"RESTApp/model"
	"RESTApp/mongodal"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewMongoDBDAL ...
var NewMongoDBDAL = mongodal.NewMongoDBDAL

// PutPlane inserts a plane to database
func PutPlane(p model.Plane, ds *mgo.Session, dbs string) (err error) {
	session := ds
	clone := session.Clone()
	db := clone.DB(dbs)
	dal := NewMongoDBDAL(db)
	err = dal.C("planes").Insert(p)
	if err != nil {
		log.Print("Could not insert", err)
	}
	defer clone.Close()
	return
}

// GetPlane returns a given plane from the database
func GetPlane(name string, ds *mgo.Session, dbs string) (p model.Plane, err error) {
	session := ds
	clone := session.Clone()

	db := clone.DB(dbs)
	dal := NewMongoDBDAL(db)
	planes := dal.C("planes")
	err = planes.Find(bson.M{"name": name}).One(&p)
	if err != nil {
		log.Println("Could not get a Record with that name!")
	}
	// defer clone.Close()
	return
}

// UpdatePlane updates a plane whose name is given
func UpdatePlane(plane model.Plane, ds *mgo.Session, db string) (model.Plane, error) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	datab := clone.DB(db)
	dal := NewMongoDBDAL(datab)
	err := dal.C("plnes").Update(bson.M{"name": plane.Name}, plane)
	return plane, err
}

// DeletePlane Deletes a given plane
func DeletePlane(name string, ds *mgo.Session, db string) (stat bool) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	datab := clone.DB(db)
	dal := NewMongoDBDAL(datab)
	err := dal.C("planes").Remove(bson.M{"name": name})
	if err != nil {
		return false
	}
	return true
}

// DeletePlaneByID ...
func DeletePlaneByID(id int, ds *mgo.Session, db string) (stat bool) {
	session := ds.Clone()
	defer session.Close()
	datab := session.DB(db)
	dal := NewMongoDBDAL(datab)
	err := dal.C("planes").Remove(bson.M{"id": id})
	if err != nil {
		return false
	}
	return true
}

// GetAllPlanes returns all the planes from the database
func GetAllPlanes(ds *mgo.Session, db string) (planes []model.Plane, err error) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	datab := clone.DB(db)
	dal := NewMongoDBDAL(datab)
	err = dal.C("planes").Find(bson.M{}).All(&planes)
	return
}
