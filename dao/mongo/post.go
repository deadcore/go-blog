package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/deadcore/go-blog/dao"
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
	return result, m.context.GetDatabase().C("posts").Find(bson.M{"_id": id}).One(&result)
}

func (m *mongoPostDao) Save(post model.Post) model.Post {
	post.Id = bson.NewObjectId().Hex()
	err := m.context.GetDatabase().C("posts").Insert(&post)

	if err != nil {
		panic(err)
	}

	return post
}

func (m *mongoPostDao) FindAll() ([]model.Post, error) {
	results := make([]model.Post, 1)

	return results, m.context.GetDatabase().C("posts").Find(bson.M{}).All(&results)
}
