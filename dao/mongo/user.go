package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
)

type MongoUserDao struct {
	Context MongoContext
}

func (m *MongoUserDao) Get(id string) (model.User, error) {
	result := model.User{}
	return result, m.Context.GetDatabase().C("posts").Find(bson.M{"_id": id}).One(&result)
}

func (m *MongoUserDao) Save(post model.User) model.User {
	post.Id = bson.NewObjectId().Hex()
	err := m.Context.GetDatabase().C("users").Insert(&post)

	if err != nil {
		panic(err)
	}

	return post
}