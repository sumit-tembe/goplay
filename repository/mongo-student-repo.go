package repository

import (
	model "github.com/goplay/model"
	"github.com/goplay/service/logger"
	"gopkg.in/mgo.v2/bson"
)

//MongoStudentRepo ...
type MongoStudentRepo struct {
	BaseRepo
}

//...
const (
	Database   = "admin"
	Collection = "student"
)

//CreateStudent ...
func (repo *MongoStudentRepo) CreateStudent(stud *model.Student) error {
	c := repo.session.DB(Database).C(Collection)
	err := c.Insert(stud)
	if err != nil {
		logger.Errorf("Error Inserting Student details [%v]; Reason: [%v]", stud, err)
	}
	return err
}

//FindStudentByID ...
func (repo *MongoStudentRepo) FindStudentByID(ID string) (*model.Student, error) {
	stud := model.Student{}
	c := repo.session.DB(Database).C(Collection)
	err := c.Find(bson.M{"id": ID}).One(&stud)
	if err != nil {
		logger.Errorf("Error fetching Student details by id [%s]; Reason: [%v]", ID, err)
	}
	return &stud, err
}

//FindAllStudents ...
func (repo *MongoStudentRepo) FindAllStudents() (*[]model.Student, error) {
	stud := []model.Student{}
	c := repo.session.DB(Database).C(Collection)
	err := c.Find(nil).All(&stud)
	if err != nil {
		logger.Errorf("Error fetching All Student; Reason: [%v]", err)
	}
	return &stud, err
}
