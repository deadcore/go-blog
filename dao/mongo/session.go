package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/deadcore/go-blog/dao"
)

type mongoSessionDao struct {
	context MongoContext
}

func SessionDao(context MongoContext) dao.SessionDao {
	return &mongoSessionDao{
		context: context,
	}
}

func (m *mongoSessionDao) Get(id string) (model.Session, error) {
	result := model.Session{}
	return result, m.context.GetDatabase().C("session").Find(bson.M{"_id": id}).One(&result)
}

func (m *mongoSessionDao) Save(post model.Session) model.Session {
	err := m.context.GetDatabase().C("session").Insert(&post)

	if err != nil {
		panic(err)
	}

	return post
}