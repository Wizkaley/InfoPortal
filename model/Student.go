package model


//Model for Student
type Student struct{
	StudentName string `json:"studentName" bson:"studentName"`
	StudentAge int `json:"studentAge,string" bson:"studentAge"`
	StudentMarks int `json:"studentMarks,string" bson:"studentMarks"`
}



