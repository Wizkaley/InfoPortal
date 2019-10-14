package model

// Plane ...
// swagger:model Plane
type Plane struct {
	// Id for the plane
	// required: true
	//min: 1
	Pid int `json:"id,string" bson:"id"`
	// Name for the Plane
	// required: true
	// min: 3
	Name string `json:"name" bson:"name"`
	// no. of wheels of the plane
	// required: true
	// min: 1
	NoWheels int `json:"wheels,string" bson:"wheels"`
	// no. of engines of the plane
	// required: true
	// min: 1
	Engines int `json:"engines,string" bson:"engines"`
	// Type of plane
	// required: true
	// min: 3
	PType string `json:"type" bson:"type"`
}

// GetPlanesAPIResponse ...
//swagger:model GetPlanesAPIResponse
type GetPlanesAPIResponse struct {
	Plane []Plane `json:"Planes"`
}
