package repository

import mgo "gopkg.in/mgo.v2"

//BaseRepo ...
type BaseRepo struct {
	session *mgo.Session
}
