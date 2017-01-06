package model

import "time"

type Session struct {
	UserId       string
	Token        string
	Roles        []string
	Expiry       time.Time
}