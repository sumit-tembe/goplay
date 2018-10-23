package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/goplay/service/logger"
)

// ...
var (
	DefaultEnvFile = "env.json"
	HTTPPort       string
)

//LoadEnvConfigs load env configs
func LoadEnvConfigs() {
	file, e := ioutil.ReadFile(DefaultEnvFile)
	if e != nil {
		logger.Errorf("File error: %v\n", e)
	}
	envConfig := struct {
		HTTPPort string `json:"HTTP_PORT"`
	}{}
	json.Unmarshal(file, &envConfig)

	//Override values from environment variables if they are set.
	HTTPPort = GetEnv("SERVER_PORT", envConfig.HTTPPort)
}

// GetEnv gets the environment variable if it exists otherwise returns the defaultValue
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
