package recording

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/storage"
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

func GetAuthToken(c *gin.Context) {
	var url = "https://app.mysupermon.com/oauth/token"
	method := "POST"

	// Creates the command structure by taking information from the URL call
	// TODO: Handle errors
	var command storage.LoginCommand
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	// Generates the user info string
	payload := strings.NewReader(generateUserInfo(command.Username, command.Password))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	// TODO: Handle all errors
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authToken)

	// TODO: Handle all errors
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// TODO: Handle all errors
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &Tokenresponse)
	if err != nil {
		return
	}

	fmt.Println("******************************************** Auth Token ********************************************")
	fmt.Printf("%s : %s\n", Tokenresponse.Type, Tokenresponse.AccessToken)
}

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

