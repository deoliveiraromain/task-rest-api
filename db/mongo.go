package db

import "gopkg.in/mgo.v2"

//Mongo is a db.DB implementation that talks to MongoDB
type Mongo struct {
	Session      *mgo.Session
	DatabaseName string
}

// NewRedis initializes and returns a redis DB using the given raw redis.v3 client
func NewMongo(cl *mgo.Session, db string) Mongo {
	return Mongo{Session: cl, DatabaseName: db}
}
