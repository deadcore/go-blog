package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/deadcore/go-blog/dao"
	"gopkg.in/mgo.v2"
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
	return result, hexifyId(m.collection().FindId(id).One, &result)
}

func (m *mongoSessionDao) FindBySessionToken(sessionToken string) (model.Session, error) {
	result := model.Session{}
	return result, hexifyId(m.collection().Find(bson.M{"token": sessionToken}).One, &result)
}

func (m *mongoSessionDao) Save(session *model.Session) {
	if err := m.collection().Insert(&session); err != nil {
		panic(err)
	}
}

func (m *mongoSessionDao) Delete(id string) error {
	return m.collection().RemoveId(id)
}

func (m *mongoSessionDao) collection() *mgo.Collection {
	return m.context.GetDatabase().C("session")
}
