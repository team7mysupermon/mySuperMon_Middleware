package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Token struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}
type Command struct {
	Usecase               string `uri:"Usecase" binding:"required"`
	ApplicationIdentifier string `uri:"Appiden" binding:"required"`
}

var Tokenresponse Token
var quit = make(chan bool)

// this code must be removed before being publish to git store it in a separate file
//************************************************************************************************

//************************************************************************************************

func main() {
	go getAuthToken()
	router := gin.Default()
	router.GET("/Start/:Usecase/:Appiden", postStart_Usscase_Appidentifier)
	router.GET("/Stop/:Usecase/:Appiden", postStop_Usscase_Appidentifier)
	router.Run("localhost:8999")

}

func postStart_Usscase_Appidentifier(c *gin.Context) {
	var command Command
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var res = Operation(command.Usecase, "start", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Controle": "A recording has now startede"})

	go run_scrape_interval(command)

}

func postStop_Usscase_Appidentifier(c *gin.Context) {
	var command Command
	if err := c.ShouldBindUri(&command); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	quit <- true
	var res = Operation(command.Usecase, "stop", command.ApplicationIdentifier)
	c.JSON(res.StatusCode, gin.H{"Controle": "A recording has now Ended"})
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
	req.Header.Add("Authorization", "Bearer "+Tokenresponse.Access_token)
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
	for {
		var url = "https://app.mysupermon.com/oauth/token"
		method := "POST"

		payload := strings.NewReader(userinfo)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", Authinfo)

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
		fmt.Printf("%s : %s\n", Tokenresponse.Token_type, Tokenresponse.Access_token)

		time.Sleep(time.Second * (time.Duration(Tokenresponse.Expires_in) - 100))
	}
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
