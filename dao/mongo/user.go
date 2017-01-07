package mongo

import (
	"github.com/deadcore/go-blog/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/deadcore/go-blog/dao"
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
	return result, m.context.GetDatabase().C("posts").Find(bson.M{"_id": id}).One(&result)
}

func (m *mongoUserDao) Save(post model.User) model.User {
	post.Id = bson.NewObjectId().Hex()
	err := m.context.GetDatabase().C("users").Insert(&post)

	if err != nil {
		panic(err)
	}

	return post
}

func (m *mongoUserDao) FindByEmailAndPassword(password [64]byte, email string) (model.User, error) {
	result := model.User{}
	return result, m.context.GetDatabase().C("posts").Find(bson.M{"password": password, "email":email}).One(&result)
}