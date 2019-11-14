package model

// Plane ...
// swagger:model Plane
type Plane struct {
	// Id for the plane
	// required: true
	//min: 1
	Pid int `json:"id" bson:"id" validate:"required,ne=0"`
	// Name for the Plane
	// required: true
	// min: 3
	Name string `json:"name" bson:"name" validate:"required,min=3"`
	// no. of wheels of the plane
	// required: true
	// min: 1
	NoWheels int `json:"wheels" bson:"wheels" validate:"required,ne=0"`
	// no. of engines of the plane
	// required: true
	// min: 1
	Engines int `json:"engines" bson:"engines" validate:"required,ne=0"`
	// Type of plane
	// required: true
	// min: 3
	PType string `json:"type" bson:"type" validate:"required,min=3"`
}

// GetPlanesAPIResponse ...
//swagger:model GetPlanesAPIResponse
type GetPlanesAPIResponse struct {
	Plane []Plane `json:"Planes"`
}
