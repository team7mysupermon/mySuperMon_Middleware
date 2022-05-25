package recording

import (
	"time"

	"github.com/team7mysupermon/mySuperMon_Middleware/storage"
)

var (
	/*
		Closes the goroutine that scrapes the recording.
		The goroutine is started when the user starts the recording
	*/
	quit = make(chan bool)
)

// Takes a username and a password and generates the string that is needed to login
func generateUserInfo(username string, password string) string {
	var userInfo = "username=" + username + "&password=" + password + "&grant_type=password"

	return userInfo
}

/*
The function that is called when the user starts the recording
Will every 5 seconds do the run operation, which returns some information about the current recording
*/
func scrapeWithInterval(command storage.StartAndStopCommand) {
	for {
		select {
		case <-quit:
			return
		default:
			Operation(command.Usecase, "run", command.ApplicationIdentifier)
		}
		time.Sleep(5 * time.Second)

	}
}
