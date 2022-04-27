package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	test()

	start_recording()

	stop_recording()

	// this is made by anton and dont know what it does thinks some docker stuff
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	//     fmt.Fprintf(w, "Hi")
	// })

	// log.Fatal(http.ListenAndServe(":8081", nil))
}

func test() {

	// This is a control it calls a my supermon page allowed to all and gets the status code. it schould allways return 200 ok

	//the url can be change to any http get request
	var url = "https://app.mysupermon.com/#/authentication/login"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v : get status = %v \n\n", url, resp.Status)

	//*****************************************************************
	// this is a attemt to get a access_token from mysupermon. (to get a token it must be a post)

	url = "https://app.mysupermon.com/oauth/token"
	method := "POST"

	// INSERT NODEpade code here

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// insert the req header from the note file

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

// here are start and stop recording but we dont have a access_token to give them so they return 401( unauthorized )

func start_recording() {
	resp, err := http.Get("https://app.mysupermon.com/devaten/data/startRecording")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("we got " + resp.Status)
}

func stop_recording() {
	resp, err := http.Get("https://app.mysupermon.com/devaten/data/stopRecording?usecaseIdentifier=PIZZA&inputSource=PIZZA&frocefullyStop=false")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("we got " + resp.Status)
}
