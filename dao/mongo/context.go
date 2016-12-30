package mongo

import (
	"gopkg.in/mgo.v2"
)

// RequestScope contains the application-specific information that are carried around in a request.
type MongoContext interface {
	GetDatabase() *mgo.Database
	Close()
}

type mongoContext struct {
	sessions *mgo.Session
	database string
}

func (rs *mongoContext) GetDatabase() *mgo.Database {
	return rs.sessions.DB(rs.database)
}

func (rs *mongoContext) Close() {
	rs.sessions.Close()
}

func getSession(host string) *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(host)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func NewMongoContext(host string, database string) MongoContext {
	session := getSession(host)

	return &mongoContext{
		sessions: session,
		database: database,
	}
}