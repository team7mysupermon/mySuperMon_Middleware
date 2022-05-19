package recording

import (
	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/storage"
)

func StartRecording(c *gin.Context) {
	// Creates the command structure by taking information from the URL call
	// TODO: Handle errors
	var command storage.StartAndStopCommand
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var res = Operation(command.Usecase, "start", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Control": "A recording has now started"})

	// Starts the scraping on a seperat thread
	go scrapeWithInterval(command)
}
