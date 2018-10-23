package main

import (
	controller "github.com/goplay/controller"
	database "github.com/goplay/database"
	utils "github.com/goplay/utils"
)

func main() {
	// Load Env confings prior running server
	utils.LoadEnvConfigs()
	database.InitMongo()
	router := controller.RegisterRoutes()
	router.Run(":" + utils.HTTPPort)
}
