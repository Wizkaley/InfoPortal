package model


type Plane struct{
	Pid int `json:"id"`
	NoWheels int `json:"wheels"`
	Engines int `json:"engines"`
	PType string `json:"PlaneType"`
}