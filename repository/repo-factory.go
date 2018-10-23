package repository

import (
	database "github.com/goplay/database"
)

//GetMongotudentRepo ...
func GetMongotudentRepo() MongoStudentIfc {
	repo := new(MongoStudentRepo)
	repo.session = database.MongoSession
	return repo
}
