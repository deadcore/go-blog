package model

type User struct {
	Id       string
	Email    string
	Password string `json:"-"`
	Roles    []string
}

func (p *User) SetId(id string) {
	p.Id = id
}