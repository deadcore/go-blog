package context

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/dao/mongo"
)

type DaoContext interface {
	PostDao() dao.PostDao
	SessionDao() dao.SessionDao
	UserDao() dao.UserDao
}

type daoContext struct {
	postDao    dao.PostDao
	sessionDao dao.SessionDao
	userDao    dao.UserDao
}

func (d *daoContext) PostDao() dao.PostDao {
	return d.postDao
}

func (d *daoContext) SessionDao() dao.SessionDao {
	return d.sessionDao
}

func (d *daoContext) UserDao() dao.UserDao {
	return d.userDao
}


func NewDaoContext(configuration Configuration) DaoContext {
	var config = configuration.MongoConfiguration

	mongoContext := getMongoContext(config)

	return &daoContext{
		postDao: mongo.PostDao(mongoContext),
		userDao: mongo.UserDao(mongoContext),
		sessionDao: mongo.SessionDao(mongoContext),
	}
}

func getMongoContext(mongoConfiguration *MongoConfiguration) mongo.MongoContext {
	return mongo.NewMongoContext(mongoConfiguration.Host, mongoConfiguration.Database)
}