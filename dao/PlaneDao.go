package dao

import (
	"RESTApp/model"
	"RESTApp/utils/mongo"
	"RESTApp/utils/mongodal"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var abc = false

// PutPlane inserts a plane to database
func PutPlane(p model.Plane, ds *mgo.Session, db string) (err error) {
	session := ds
	clone := session.Clone()
	//err = clone.DB("trial").C("planes").Insert(p)
	ds, err = mongo.GetDataBaseSession("localhost:27017")
	err = mongodal.NewMongoDBDAL(ds.DB(db)).C("planes").Insert(p)
	if err != nil {
		log.Print("Could not insert", err)
	}
	defer clone.Close()
	return
}

// GetPlane returns a given plane from the database
func GetPlane(name string, ds *mgo.Session, db string) (p model.Plane) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	clone.DB(db).C("planes").Find(bson.M{"name": name}).One(&p)
	return
}

// UpdatePlane updates a plane whose name is given
func UpdatePlane(plane model.Plane, ds *mgo.Session) (model.Plane, error) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	err := clone.DB("trial").C("plnes").Update(bson.M{"name": plane.Name}, plane)
	return plane, err
}

// DeletePlane Deletes a given plane
func DeletePlane(name string, ds *mgo.Session, db string) (stat bool) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	err := clone.DB(db).C("planes").Remove(bson.M{"name": name})
	if err != nil {
		return false
	}
	return true
}

// DeletePlaneByID ...
func DeletePlaneByID(id int, ds *mgo.Session, db string) (stat bool) {
	session := ds.Clone()
	defer session.Clone()
	err := session.DB(db).C("planes").Remove(bson.M{"id": id})
	if err != nil || abc {
		return false
	}
	return true
}

// GetAllPlanes returns all the planes from the database
func GetAllPlanes(ds *mgo.Session, db string) (planes []model.Plane, err error) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	err = clone.DB(db).C("planes").Find(bson.M{}).All(&planes)
	return
}
