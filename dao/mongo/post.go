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
	return result, m.collection().FindId(bson.ObjectIdHex(id)).One(&result)
}

func (m *mongoPostDao) Save(post model.Post) model.Post {
	var objectId bson.ObjectId
	if (bson.IsObjectIdHex(post.Id)) {
		objectId = bson.ObjectIdHex(post.Id)
	} else {
		objectId = bson.NewObjectId()
		post.Id = objectId.Hex()
	}

	m.collection().UpsertId(objectId, &post)
	return post
}

func (m *mongoPostDao) FindAll() ([]model.Post, error) {
	results := make([]model.Post, 1)

	return results, m.collection().Find(bson.M{}).All(&results)
}

func (m *mongoPostDao) Delete(id string) error {
	return m.collection().RemoveId(bson.ObjectIdHex(id))
}

func (m *mongoPostDao) collection() *mgo.Collection {
	return m.context.GetDatabase().C("posts")
}
