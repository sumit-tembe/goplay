package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/goplay/controller/auth"
	model "github.com/goplay/model"
	service "github.com/goplay/service"
	"github.com/goplay/service/logger"
)

//RegisterRoutes register controller routes
func RegisterRoutes() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/", service.Root)

	//Authorized APIs
	v1Authorized := router.Group("/v1")
	v1Authorized.POST("/auth", auth.Authenticate)
	v1Authorized.POST("/re-auth", auth.RefreshJWTToken)
	v1Authorized.GET("/", auth.AuthorizationMiddleware(), service.Root)
	v1Authorized.POST("/arithmatic", auth.AuthorizationMiddleware(), service.SyncArithmatic)

	v1Authorized.POST("/students", auth.RunAsRootUser(), HandleCreateStudents)
	v1Authorized.GET("/students", auth.RunAsRootUser(), HandleGetAllStudents)
	v1Authorized.GET("/students/:id", auth.RunAsRootUser(), HandleFindStudentsByID)
	return router
}

//HandleCreateStudents ...
func HandleCreateStudents(c *gin.Context) {
	var stud model.Student
	err := json.NewDecoder(c.Request.Body).Decode(&stud)
	if err != nil {
		logger.Errorf("Error reading request body in HandleCreateStudents; Reason: %v", err)
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	err = service.CreateStudent(&stud)
	if err != nil {
		c.JSON(http.StatusOK, "error")
		return
	}
	c.JSON(http.StatusOK, stud)
}

//HandleFindStudentsByID ...
func HandleFindStudentsByID(c *gin.Context) {
	ID := c.Params.ByName("id")
	stud, err := service.FindStudentByID(ID)
	if err != nil {
		c.JSON(http.StatusOK, "error")
		return
	}

	c.JSON(http.StatusOK, stud)
}

//HandleGetAllStudents ...
func HandleGetAllStudents(c *gin.Context) {
	stud, err := service.FindAllStudents()
	if err != nil {
		c.JSON(http.StatusOK, "error")
		return
	}

	c.JSON(http.StatusOK, stud)
}
