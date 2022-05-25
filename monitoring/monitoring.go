package monitoring

//se prometheus config docs:
// https://prometheus.io/docs/guides/go-application/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/team7mysupermon/mySuperMon_Middleware/storage"
)

var (
	startAutoGenerated storage.StartAutoGenerated
	runAutoGenerated   storage.RunAutoGenerated
	stopAutoGenerated  storage.StopAutoGenerated
)

func Monitor() {
	startRegisterMetrics()
	runRegisterMetrics()
	stopRegisterMetrics()

	fmt.Println("helloooooooooooooooooooooooooooooooooooooo")

	go http.Handle("/metrics", promhttp.Handler())
	fmt.Println("is handling path for metrics")
	
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatalln("Failed to serve metrics on port 9091 --------------------------")
	} else {
		log.Println("Now listening and serving on port 9091 -----------------------")
	}

	fmt.Println("byyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyeeeeeeeeeeeeeee")

	log.Fatal(http.ListenAndServe(":9091", nil))

}


func ParseBody(body []byte, action string) {

	fmt.Println(string(body))

	if action == "start" {

		err := json.Unmarshal(body, &startAutoGenerated)
		if err != nil {
			log.Panicln(err)
		}

		startRecordMetrics()

		fmt.Printf("%+v\n", startAutoGenerated)
	}

	if action == "run" {
		err := json.Unmarshal(body, &runAutoGenerated)
		if err != nil {
			log.Panicln(err)
		}
		runRecordMetrics()
		fmt.Printf("%+v\n", runAutoGenerated)
	}

	if action == "stop" {
		err := json.Unmarshal(body, &stopAutoGenerated)
		if err != nil {
			log.Panicln(err)
		}
		stopRecordMetrics()

		fmt.Printf("%+v\n", stopAutoGenerated)
	}

}