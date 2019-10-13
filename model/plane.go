package model

// Plane ...
// swagger: model
type Plane struct {
	Pid      int    `json:"id,string" bson:"id"`
	Name     string `json:"name" bson:"name"`
	NoWheels int    `json:"wheels,string" bson:"wheels"`
	Engines  int    `json:"engines,string" bson:"engines"`
	PType    string `json:"type" bson:"type"`
}
