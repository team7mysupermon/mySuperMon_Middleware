package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/team7mysupermon/mySuperMon_Middleware/storage"

	"github.com/gin-gonic/gin"
	"github.com/team7mysupermon/mySuperMon_Middleware/monitoring"

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
	router.GET("/Login/:Username/:Password", getAuthToken)
	router.GET("/Start/:Usecase/:Appiden", startRecording)
	router.GET("/Stop/:Usecase/:Appiden", stopRecording)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Starts the program
	err := router.Run(":8999")
	if err != nil {
		return
	}
}

// @BasePath /Start/{Usecase}/{Appiden}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Param Usecase path string true ":Usecase"
// @Param Appiden path string true ":Appiden"
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /Start/{Usecase}/{Appiden} [get]
func startRecording(c *gin.Context) {
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

// @BasePath /Stop/:Usecase/:Appiden"

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /Stop/:Usecase/:Appiden [get]
func stopRecording(c *gin.Context) {
	// Creates the command structure by taking information from the URL call
	// TODO: Handle errors
	var command storage.StartAndStopCommand
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	// Sends true through the quit channel to the goroutine that is scraping the recording
	quit <- true

	var res = Operation(command.Usecase, "stop", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Control": "A recording has now ended"})
}

// @BasePath /Login/:Username/:Password

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /Login/:Username/:Password [get]
func getAuthToken(c *gin.Context) {
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

	monitoring.ParseBody(body, action)

	fmt.Printf("********************************************************** begin %v \n", action)
	fmt.Println(string(body))
	fmt.Printf("********************************************************** end %v \n\n", action)
	return res

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

// Takes a username and a password and generates the string that is needed to login
func generateUserInfo(username string, password string) string {
	var userInfo = "username=" + username + "&password=" + password + "&grant_type=password"

	return userInfo
}
