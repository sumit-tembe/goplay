package database

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/mgo.v2"

	"github.com/goplay/service/logger"
	utils "github.com/goplay/utils"
)

// ...
var (
	DefaultCredsFile = "mongo-creds.json"
	dbConfig         *mgo.DialInfo
	MongoSession     *mgo.Session
)

//InitMongo ...
func InitMongo() {
	var err error
	MongoSession, err = ConnectToMongoDB()
	if err != nil {
		logger.Errorf("Error connecting to DB: [%v]", err)
		panic(err)
	}
}

//GetDBConfigs ...
func GetDBConfigs() *mgo.DialInfo {
	if dbConfig != nil {
		return dbConfig
	}
	file, e := ioutil.ReadFile(DefaultCredsFile)
	if e != nil {
		logger.Errorf("File error: %v\n", e)
		log.Panic(e)
	}
	json.Unmarshal(file, &dbConfig)

	//Override values from environment variables if they are set.
	dbConfig.Database = utils.GetEnv("Database", dbConfig.Database)
	dbConfig.Username = utils.GetEnv("Username", dbConfig.Username)
	dbConfig.Password = utils.GetEnv("Password", dbConfig.Password)
	return dbConfig
}

//ConnectToMongoDB ...
func ConnectToMongoDB() (*mgo.Session, error) {
	GetDBConfigs()
	session, err := mgo.DialWithInfo(dbConfig)
	if err != nil {
		logger.Errorf("Error connecting to DB: [%v]", err)
		panic(err)
	}
	return session, err
}
