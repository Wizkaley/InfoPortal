package model

// Student represents students info as a record in this application
//swagger:model
type Student struct {
	StudentName  string `json:"studentName" bson:"studentName"`
	StudentAge   int    `json:"studentAge,string" bson:"studentAge"`
	StudentMarks int    `json:"studentMarks,string" bson:"studentMarks"`
}
