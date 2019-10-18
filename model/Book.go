package model

// Book ...
type Book struct {
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
	Pages  int    `json:"pages" bson:"pages"`
}
