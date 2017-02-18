package repositories

import (
	"github.com/deoliveiraromain/task-rest-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskRepo struct {
	Coll *mgo.Collection
}

//TaskRepoColl : Name of task collection in mongo
var TaskRepoColl = "todo"

func (r *TaskRepo) All() (models.TaskCollection, error) {
	result := models.TaskCollection{[]models.Task{}}
	err := r.Coll.Find(nil).All(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TaskRepo) FindByName(query string) (models.TaskResource, error) {
	result := models.TaskResource{}
	err := r.Coll.Find(bson.M{"name": query}).One(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TaskRepo) Create(task *models.Task) error {
	id := bson.NewObjectId()
	_, err := r.Coll.UpsertId(id, task)
	if err != nil {
		return err
	}
	task.Id = id

	return nil
}

func (r *TaskRepo) Update(taskDb *models.Task, task *models.Task) error {
	err := r.Coll.UpdateId(taskDb.Id, task)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepo) Delete(id string) error {
	err := r.Coll.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
