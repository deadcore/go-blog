package model

type Post struct {
	Id      string `bson:"_id,omitempty"`
	Content string
	Title   string
}

func (p Post) SetId(id string) {
	p.Id = id
}

func (p Post) GetId() string {
	return p.Id
}
