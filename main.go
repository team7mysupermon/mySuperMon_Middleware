package main

import (
	"github.com/team7mysupermon/mySuperMon_Middleware/recording"

	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/monitoring"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/team7mysupermon/mySuperMon_Middleware/docs"
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
