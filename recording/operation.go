package recording

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/team7mysupermon/mySuperMon_Middleware/monitoring"
)



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
	req.Header.Add("Authorization", "Bearer "+ Tokenresponse.AccessToken)
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

