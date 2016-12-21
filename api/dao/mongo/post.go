package mongo

import (
	"github.com/deadcore/go-blog/api/model"
	"gopkg.in/mgo.v2/bson"
)

type (
	MongoPostDao struct {
		Context MongoContext
	}
)

func (m *MongoPostDao) Get(id string) (model.Post, error) {
	result := model.Post{}
	return result, m.Context.GetDatabase().C("posts").Find(bson.M{"_id": id}).One(&result)
}

func (m *MongoPostDao) Save(post model.Post) model.Post {
	post.Id = bson.NewObjectId().Hex()
	err := m.Context.GetDatabase().C("posts").Insert(&post)

	if err != nil {
	}

	return post
}
