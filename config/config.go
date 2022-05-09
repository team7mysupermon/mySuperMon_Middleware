package main

//se prometheus config docs:
// https://prometheus.io/docs/guides/go-application/

import (
	"net/http"
	"time"
	"fmt"
	"log"
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


//declare metric var
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

	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:      "counter",
			Help:      "This is a counter",
		})

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:      "gauge",
			Help:      "This is a gauge",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:      "histogram",
			Help:      "This is a histogram",
		})

	summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name:      "summary",
			Help:      "This is a summary",
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

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		status := http.StatusOK

		defer func() {
			histogram.WithLabelValues(fmt.Sprintf("%d", status)).Observe(time.Since(start).Seconds())
		}()

		if req.Method == http.MethodGet {
			handler.ServeHTTP(w, req)
			return
		}
		status = http.StatusBadRequest

		w.WriteHeader(status)
	})
}



func main() {
	record_metrics()
	cpuTemp.Set(420.69)
	hdFailures.With(prometheus.Labels{"device":"/dev/sda"}).Inc()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

	rand.Seed(time.Now().Unix())

	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "prom_request_time",
		Help: "Time it has taken to retrieve the metrics",
	}, []string{"time"})

	prometheus.Register(histogramVec)

	http.Handle("/metrics", newHandlerWithHistogram(promhttp.Handler(), histogramVec))

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)

			time.Sleep(time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(":2112", nil))
}