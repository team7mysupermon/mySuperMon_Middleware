package main

import (

	"github.com/team7mysupermon/mySuperMon_Middleware/storage"

	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/monitoring"
	"github.com/team7mysupermon/mySuperMon_Middleware/recording"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/team7mysupermon/mySuperMon_Middleware/docs"
)

var (
	// The authentication token needed to be able to get the access token when logging in
	authToken = "Basic cGVyZm9ybWFuY2VEYXNoYm9hcmRDbGllbnRJZDpsamtuc3F5OXRwNjEyMw=="

	/*
		Instantiated when a user calls the login API call.
		Contains the authentication token
	*/
	Tokenresponse storage.Token

	/*
		Closes the goroutine that scrapes the recording.
		The goroutine is started when the user starts the recording
	*/
	quit = make(chan bool)
)

func main() {
	go monitoring.Monitor()
	docs.SwaggerInfo.BasePath = ""
	router := gin.Default()

	// The API calls
	router.GET("/Login/:Username/:Password", recording.GetAuthToken)
	router.GET("/Start/:Usecase/:Appiden", recording.StartRecording)
	router.GET("/Stop/:Usecase/:Appiden", recording.StopRecording)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Starts the program
	err := router.Run(":8999")
	if err != nil {
		return
	}
}
