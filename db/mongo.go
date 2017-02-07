package db

import (
	"gopkg.in/mgo.v2"
)

//Mongo is a db.DB implementation that talks to MongoDB
type Mongo struct {
	Session        *mgo.Session
	DatabaseName   string
}

// NewRedis initializes and returns a redis DB using the given raw redis.v3 client
func NewMongo(cl *mgo.Session, db string) Mongo {
	return Mongo{Session: cl, DatabaseName:db}
}

/*
func (r *Mongo) FindAll(result interface{}) error {
	return r.session.DB(r.databaseName).C(r.collectionName).Find(bson.M{}).All(result)
}

func (r *Mongo) Find(query interface{}, result interface{}) error {
	res:=r.session.DB(r.databaseName).C(r.collectionName).Find(query).One(result)
	log.Print("RES=>"  + res.Error())
	//return r.session.DB(r.databaseName).C(r.collectionName).Find(query).One(&result)
	return res
}
// Save is the interface implementation
func (r *Mongo) Save(model interface{}) error {
	// Write the model to mongo
	return r.session.DB(r.databaseName).C(r.collectionName).Insert(model)
}

// Delete is the interface implementation
func (r *Mongo) Delete(query interface{}) error {
	return r.session.DB(r.databaseName).C(r.collectionName).Remove(query)
}
*/