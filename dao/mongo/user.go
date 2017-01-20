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
	user := model.User{}
	err := m.collection().FindId(id).One(&user)

	return user, err
}

func (m *mongoUserDao) Save(user model.User) model.User {
	objectId := bson.NewObjectId()
	user.Id = objectId.Hex()

	m.context.GetDatabase().C("users").UpsertId(objectId, &user)
	return user
}

func (m *mongoUserDao) FindByEmailAndPassword(password string, email string) (model.User, error) {
	result := model.User{}
	return result, m.collection().Find(bson.M{"password": password, "email":email}).One(&result)
}

func (m mongoUserDao) collection() *mgo.Collection {
	return m.context.GetDatabase().C("users")
}