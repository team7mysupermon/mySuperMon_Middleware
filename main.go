package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type token struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}

var Tokenresponse token

// this code must be removed before being publish to git store it in a separate file
//************************************************************************************************

//************************************************************************************************

func main() {

	getAuthToken()

	start_recording()
	go run_scrape_interval()
	time.Sleep(time.Second * 30)

	stop_recording()

}

func getAuthToken() {

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
}

// here are start and stop recording but we dont have a access_token to give them so they return 401( unauthorized )

func start_recording() {
	fmt.Println("******************************************** Start Recording Interval ******************************")

	url := "https://app.mysupermon.com/devaten/data/startRecording?usecaseIdentifier=this-is-api-test"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("applicationIdentifier", ApplicationIdentifier)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Tokenresponse.Access_token)

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
	fmt.Println(string(body))
}

func stop_recording() {

	url := "https://app.mysupermon.com/devaten/data/stopRecording?usecaseIdentifier=this-is-api-test&inputSource"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("applicationIdentifier", ApplicationIdentifier)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Tokenresponse.Access_token)

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
	fmt.Println("******************************************** Stop Recording Interval *********************************")
	fmt.Println(string(body))
}

func run_scrape_interval() {
	for {

		time.Sleep(5 * time.Second)
		url := "https://app.mysupermon.com/devaten/data/getRunSituation"
		method := "GET"

		payload := strings.NewReader("")

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("applicationIdentifier", ApplicationIdentifier)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+Tokenresponse.Access_token)

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
		fmt.Println("******************************************** Run Interval ********************************************")
		fmt.Println(string(body))
		fmt.Println("******************************************** Stop Interval *******************************************\n")

	}
}
