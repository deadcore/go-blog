package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
)

type MongoSessionDao struct {
	Context MongoContext
}

func (m *MongoSessionDao) Get(id string) (model.Session, error) {
	result := model.Session{}
	return result, m.Context.GetDatabase().C("session").Find(bson.M{"_id": id}).One(&result)
}

func (m *MongoSessionDao) Save(post model.Session) model.Session {
	err := m.Context.GetDatabase().C("session").Insert(&post)

	if err != nil {
		panic(err)
	}

	return post
}