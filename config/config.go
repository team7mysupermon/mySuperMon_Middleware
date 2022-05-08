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
		Name: "processed_events_counter",
		Help: "Total no. of processed events",
	})

	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)
)

func record_metrics() {
	
	go func () {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

}

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main() {
	record_metrics()
	cpuTemp.Set(420.69)
	hdFailures.With(prometheus.Labels{"device":"/dev/sda"}).Inc()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}