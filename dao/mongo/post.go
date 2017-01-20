package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/deadcore/go-blog/dao"
	"gopkg.in/mgo.v2"
)

type mongoPostDao struct {
	context MongoContext
}

func PostDao(context MongoContext) dao.PostDao {
	return &mongoPostDao{
		context: context,
	}
}

func (m *mongoPostDao) Get(id string) (model.Post, error) {
	result := model.Post{}
	return result, m.collection().FindId(id).One(&result)
}

func (m *mongoPostDao) Save(post model.Post) model.Post {
	objectId := bson.NewObjectId()
	post.Id = objectId.Hex()

	m.context.GetDatabase().C("users").UpsertId(objectId, &post)
	return post
}

func (m *mongoPostDao) FindAll() ([]model.Post, error) {
	results := make([]model.Post, 1)

	return results, m.collection().Find(bson.M{}).All(&results)
}

func (m *mongoPostDao) collection() *mgo.Collection {
	return m.context.GetDatabase().C("posts")
}
