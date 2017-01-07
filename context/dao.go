package context

import (
	"github.com/deadcore/go-blog/dao"
	"github.com/deadcore/go-blog/dao/mongo"
	"github.com/deadcore/go-blog/dao/memory"
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


func NewMongoDaoContext(host string, database string) DaoContext {
	mongoContext := getMongoContext(host, database)

	return &daoContext{
		postDao: mongo.PostDao(mongoContext),
		userDao: mongo.UserDao(mongoContext),
		sessionDao: mongo.SessionDao(mongoContext),
	}
}

func NewInMemoryDaoContext() DaoContext {
	return &daoContext{
		postDao: memory.PostDao(),
		//userDao: memory.UserDao(),
		//sessionDao: memory.SessionDao(),
	}
}

func getMongoContext(host string, database string) mongo.MongoContext {
	return mongo.NewMongoContext(host, database)
}