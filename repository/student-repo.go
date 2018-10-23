package repository

import (
	model "github.com/goplay/model"
)

//MongoStudentIfc ...
type MongoStudentIfc interface {
	FindStudentByID(string) (*model.Student, error)
	CreateStudent(*model.Student) error
	FindAllStudents() (*[]model.Student, error)
}
