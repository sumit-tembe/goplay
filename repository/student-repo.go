package repository

import (
	model "github.com/goplay/model"
)

//MongoStudentIfc ...
//mockgen -destination=mocks/mock_studentrepo.go -mock_names MongoStudentIfc=MongoStudentRepo -package=mocks github.com/goplay/repository MongoStudentIfc
type MongoStudentIfc interface {
	FindStudentByID(string) (*model.Student, error)
	CreateStudent(*model.Student) error
	FindAllStudents() (*[]model.Student, error)
}
