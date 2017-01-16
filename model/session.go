package model

import "time"

type Session struct {
	Id     string `bson:"_id,omitempty"`
	UserId string
	Token  string
	Expiry time.Time
}

func (p Session) SetId(id string) {
	p.Id = id
}

func (p Session) GetId() string {
	return p.Id
}