package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/deadcore/go-blog/dao"
	"gopkg.in/mgo.v2"
)

type mongoUserDao struct {
	context MongoContext
}

func UserDao(context MongoContext) dao.UserDao {
	return &mongoUserDao{
		context: context,
	}
}

func (m *mongoUserDao) Get(id string) (model.User, error) {
	result := model.User{}
	return result, hexifyId(m.collection().FindId(id).One, &result)
}

func (m *mongoUserDao) Save(post model.User) model.User {
	err := m.collection().Insert(&post)

	if err != nil {
		panic(err)
	}

	return post
}

func (m *mongoUserDao) FindByEmailAndPassword(password string, email string) (model.User, error) {
	result := model.User{}
	return result, m.collection().Find(bson.M{"password": password, "email":email}).One(&result)
}

func (m *mongoUserDao) collection() *mgo.Collection {
	return m.context.GetDatabase().C("users")
}