package model

type Post struct {
	Id      string `bson:"_id,omitempty"`
	Content string
	Title   string
}
