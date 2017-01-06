package model

type User struct {
	Id    string `bson:"_id,omitempty"`
	Email string
}