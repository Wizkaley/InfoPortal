package dao

import (
	"RESTApp/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var abc = false

// PutPlane inserts a plane to database
func PutPlane(p model.Plane, ds *mgo.Session) (err error) {
	session := ds
	clone := session.Clone()
	err = clone.DB("trial").C("planes").Insert(p)
	// defer func() {
	// 	r := recover()
	// 	if r != nil {
	// 		log.Print("Error While Inserting Entry")
	// 	}
	// }()
	defer clone.Close()
	return
}

// GetPlane returns a given plane from the database
func GetPlane(name string, ds *mgo.Session) (p model.Plane) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	clone.DB("trial").C("planes").Find(bson.M{"name": name}).One(&p)
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
func DeletePlane(name string, ds *mgo.Session) (stat bool) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	err := clone.DB("trial").C("planes").Remove(bson.M{"name": name})
	if err != nil {
		return false
	}
	return true
}

// DeletePlaneByID ...
func DeletePlaneByID(id int, ds *mgo.Session) (stat bool) {
	session := ds.Clone()
	defer session.Clone()
	err := session.DB("trial").C("planes").Remove(bson.M{"id": id})
	if err != nil || abc {
		return false
	}
	return true
}

// GetAllPlanes returns all the planes from the database
func GetAllPlanes(ds *mgo.Session) (planes []model.Plane, err error) {
	session := ds
	clone := session.Clone()
	defer clone.Close()
	err = clone.DB("trial").C("planes").Find(bson.M{}).All(&planes)
	return
}
