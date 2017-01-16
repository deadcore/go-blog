package model

type User struct {
	Id       string `bson:"_id,omitempty"`
	Email    string
	Password string `json:"-"`
	Roles    []string
}


func (p User) SetId(id string) {
	p.Id = id
}

func (p User) GetId() string {
	return p.Id
}