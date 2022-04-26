package main

//se prometheus config docs:
// https://prometheus.io/docs/guides/go-application/

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_events",
		Help: "Total no. of processed events",
	})
)

func record_metrics() {
	
	go func () {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

}


func main() {
	record_metrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}