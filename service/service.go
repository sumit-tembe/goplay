package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/goplay/model"
	repository "github.com/goplay/repository"
	"github.com/goplay/service/logger"
)

var getMongoStudentRepo = repository.GetMongotudentRepo

//SyncData ...
var SyncData = &model.SyncData{Value: 1}

//Root ...
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, "root")
}

//SyncArithmatic ...
func SyncArithmatic(c *gin.Context) {
	var arithmatics []model.Arithmatic
	err := json.NewDecoder(c.Request.Body).Decode(&arithmatics)
	if err != nil {
		logger.Errorf("Error reading request body in SyncArithmatic; Reason: %v", err)
		return
	}
	// Sumulating multiple operation on SyncData concurrenlty
	go func() {
		for _, arithmatic := range arithmatics {
			SyncData.SafeOperation(arithmatic)
		}
	}()

	rsp := model.ArithmaticResponse{Arithmatics: arithmatics, Value: SyncData.Value}
	c.JSON(http.StatusOK, rsp)
}
