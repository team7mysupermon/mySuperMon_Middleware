package main

import (
	"github.com/team7mysupermon/mySuperMon_Middleware/recording"

	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/monitoring"
)

func main() {
	go monitoring.Monitor()

	router := gin.Default()

	// The API calls
	router.GET("/Login/:Username/:Password", recording.GetAuthToken)
	router.GET("/Start/:Usecase/:Appiden", recording.StartRecording)
	router.GET("/Stop/:Usecase/:Appiden", recording.StopRecording)

	// Starts the program
	err := router.Run(":8999")
	if err != nil {
		return
	}
}











