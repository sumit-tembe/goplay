package service

import (
	model "github.com/goplay/model"
)

//CreateStudent ...
func CreateStudent(stud *model.Student) error {
	err := getMongoStudentRepo().CreateStudent(stud)
	return err
}

//FindStudentByID ...
func FindStudentByID(ID string) (*model.Student, error) {
	stud, err := getMongoStudentRepo().FindStudentByID(ID)
	return stud, err
}

//FindAllStudents ...
func FindAllStudents() (*[]model.Student, error) {
	stud, err := getMongoStudentRepo().FindAllStudents()
	return stud, err
}
