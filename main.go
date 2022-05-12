package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/monitoring"
)

/*
Struct to get all information from the bearer token returned by MySuperMon upon login
It must match the json object that we get from MySuperMon
*/
type Token struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

/*
Struct for the calls that the user makes with a Usecase and a ApplicationIdentifier.
The information is gathered from the URL and added to the struct.
Therefore, the end of the struct, the uri:xxx part, must match the definitions in the main method.
*/
type StartAndStopCommand struct {
	Usecase               string `uri:"Usecase" binding:"required"`
	ApplicationIdentifier string `uri:"Appiden" binding:"required"`
}

/*
Struct for the calls that the user makes with a Username and a Password.
The information is gathered from the URL and added to the struct.
Therefore, the end of the struct, the uri:xxx part, must match the definitions in the main method.
*/

type LoginCommand struct {
	Username string `uri:"Username" binding:"required"`
	Password string `uri:"Password" binding:"required"`
}


var (
	// The authentication token needed to be able to get the access token when logging in
	authToken = "Basic cGVyZm9ybWFuY2VEYXNoYm9hcmRDbGllbnRJZDpsamtuc3F5OXRwNjEyMw=="

	/*
	Instantiated when a user calls the login API call.
	Contains the authentication token
	*/
	Tokenresponse Token

	/*
	Closes the goroutine that scrapes the recording.
	The goroutine is started when the user starts the recording
	*/
	quit = make(chan bool)
) 


func main() {
	go monitoring.Monitor()

	router := gin.Default()	

	// The API calls
	router.GET("/Login/:Username/:Password", getAuthToken)
	router.GET("/Start/:Usecase/:Appiden", startRecording)
	router.GET("/Stop/:Usecase/:Appiden", stopRecording)

	// Starts the program
	router.Run("localhost:8999")
}

func startRecording(c *gin.Context) {
	// Creates the command structure by taking information from the URL call
	// TODO: Handle errors
	var command StartAndStopCommand
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var res = Operation(command.Usecase, "start", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Control": "A recording has now started"})

	// Starts the scraping on a seperat thread
	go scrapeWithInterval(command)
}

func stopRecording(c *gin.Context) {
	// Creates the command structure by taking information from the URL call
	// TODO: Handle errors
	var command StartAndStopCommand
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	// Sends true through the quit channel to the goroutine that is scraping the recording
	quit <- true

	var res = Operation(command.Usecase, "stop", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Control": "A recording has now ended"})
}

func getAuthToken(c *gin.Context) {
	var url = "https://app.mysupermon.com/oauth/token"
	method := "POST"

	// Creates the command structure by taking information from the URL call
	// TODO: Handle errors
	var command LoginCommand
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
	json.Unmarshal(body, &Tokenresponse)

	fmt.Println("******************************************** Auth Token ********************************************")
	fmt.Printf("%s : %s\n", Tokenresponse.Type, Tokenresponse.AccessToken)
}

func Operation(usecase string, action string, applicationIdentifier string) *http.Response {
	url := "https://app.mysupermon.com/devaten/data/operation?usecaseIdentifier=" + usecase + "&action=" + action
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	// TODO: Handle errors
	if err != nil {
		fmt.Println(err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: 500,
		}
	}
	req.Header.Add("applicationIdentifier", applicationIdentifier)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Tokenresponse.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: 500,
		}
	}
	defer res.Body.Close()
	

	// TODO: Handle errors
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: 500,
		}
	}

	monitoring.ParseBody(body)

	fmt.Printf("********************************************************** begin %v \n", action)
	fmt.Println(string(body))
	fmt.Printf("********************************************************** end %v \n\n", action)
	return res

}

/*
The function that is called when the user starts the recording
Will every 5 seconds do the run operation, which returns some information about the current recording
*/
func scrapeWithInterval(command StartAndStopCommand) {
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

// Takes a username and a password and generates the string that is needed to login
func generateUserInfo(username string, password string) string {
	var userInfo = "username=" + username + "&password=" + password + "&grant_type=password"

	return userInfo
}

