package models

import "gopkg.in/mgo.v2/bson"

type Task struct {
	Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	//Due       time.Time `json:"due"`
}

type TaskCollection struct {
	Data []Task `json:"data"`
}

type TaskResource struct {
	Data Task `json:"data"`
}



