package dao

import "github.com/deadcore/go-blog/model"

type SessionDao interface {
	Get(id string) (model.Session, error)
	Save(session *model.Session)
	FindBySessionToken(token string) (model.Session, error)
	Delete(id string) error
}