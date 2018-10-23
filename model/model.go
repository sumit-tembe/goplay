package model

import (
	"sync"

	"github.com/goplay/service/logger"
)

//GenericResponse ...
type GenericResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

//SyncData this must be used in sync
//mux.Lock() and mux.Unlock() will ensure only one operation at time
type SyncData struct {
	Value int
	mux   sync.Mutex
}

//Arithmatic ...
type Arithmatic struct {
	Value     int    `json:"value"`
	Operation string `json:"operation"`
}

//ArithmaticResponse ...
type ArithmaticResponse struct {
	Arithmatics []Arithmatic
	Value       int
}

//SafeOperation ...
func (data *SyncData) SafeOperation(arithmatic Arithmatic) {
	defer data.mux.Unlock()
	data.mux.Lock()
	switch arithmatic.Operation {
	case "+":
		data.Value += arithmatic.Value
	case "-":
		data.Value -= arithmatic.Value
	case "*":
		data.Value *= arithmatic.Value
	case "/":
		data.Value /= arithmatic.Value
	default:
		logger.Errorf("Invalid operation %s", arithmatic.Operation)
	}

}
