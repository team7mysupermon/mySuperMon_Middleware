package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Token struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}
type Command struct {
	Usecase               string `uri:"Usecase" binding:"required"`
	ApplicationIdentifier string `uri:"Appiden" binding:"required"`
}

type Config struct {
	Username              string `json:"MySuperMon_Username"`
	Password              string `json:"MySuperMon_Password"`
	ApplicationIdentifier string `json:"MySuperMon_ApplicationIdentifier"`
	AuthInfo              string `json:"Auth_information"`
}

var Tokenresponse Token

var config Config

var quit = make(chan bool)

func main() {
	go getAuthToken()
	router := gin.Default()
	router.GET("/Start/:Usecase/:Appiden", postStart_Usecase_Appidentifier)
	router.GET("/Stop/:Usecase/:Appiden", postStop_Usecase_Appidentifier)
	router.Run("localhost:8999")

}

func postStart_Usecase_Appidentifier(c *gin.Context) {
	var command Command
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var res = Operation(command.Usecase, "start", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Control": "A recording has now started"})

	go run_scrape_interval(command)

}

func postStop_Usecase_Appidentifier(c *gin.Context) {
	var command Command
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	quit <- true
	var res = Operation(command.Usecase, "stop", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Control": "A recording has now ended"})
}

func Operation(usecase string, action string, applicationIdentifier string) *http.Response {
	url := "https://app.mysupermon.com/devaten/data/operation?usecaseIdentifier=" + usecase + "&action=" + action
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: 500,
		}
	}
	fmt.Printf("********************************************************** begin %v \n", action)
	fmt.Println(string(body))
	fmt.Printf("********************************************************** end %v \n\n", action)
	return res

}

func getAuthToken() {
	readConfig()

	for {
		var url = "https://app.mysupermon.com/oauth/token"
		method := "POST"

		payload := strings.NewReader(generateUserInfo(config.Username, config.Password))

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", config.AuthInfo)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Unmarshal(body, &Tokenresponse)

		fmt.Println("******************************************** Auth Token ********************************************")
		fmt.Printf("%s : %s\n", Tokenresponse.Type, Tokenresponse.AccessToken)

		time.Sleep(time.Second * (time.Duration(Tokenresponse.ExpiresIn) - 100))
	}
}

func readConfig() {
	body, err := os.ReadFile("Devaten.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &config)
}

func generateUserInfo(username string, password string) string {
	var userInfo = "username=" + username + "&password=" + password + "&grant_type=password"

	return userInfo
}

func run_scrape_interval(command Command) {
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
