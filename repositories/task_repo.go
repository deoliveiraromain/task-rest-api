package repositories

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/deoliveiraromain/task-rest-api/models"
)

type TaskRepo struct {
	Coll *mgo.Collection
}

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
	err := r.Coll.Find(bson.M{"name" : query}).One(&result.Data)
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

func (r *TaskRepo) Update(task *models.Task) error {
	err := r.Coll.UpdateId(task.Id, task)
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
