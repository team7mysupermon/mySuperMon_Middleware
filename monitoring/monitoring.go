package monitoring

//se prometheus config docs:
// https://prometheus.io/docs/guides/go-application/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/team7mysupermon/mySuperMon_Middleware/storage"
)

//declare metric var
var (
	SUM_ROWS_AFFECTED_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "SUM_ROWS_AFFECTED",
			Help: "Value of the no. of times a row is affected during monitoring",
		})

	SUM_SELECTED_RANGE_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "SUM_SELECTED_RANGE",
			Help: "Value of the no. of selected ranges during monitoring",
		})

	runRecordingData storage.RunRecordingData
)

// registers metrics to expose
func registerMetrics() {
	prometheus.MustRegister(SUM_ROWS_AFFECTED_GAUGE)
	prometheus.MustRegister(SUM_SELECTED_RANGE_GAUGE)
}

func recordMetrics() {

	go func() {
		for {
			//SUM_ROWS_AFFECTED_GAUGE.Set(runRecordingData.SumRowsAffected)
			//SUM_SELECTED_RANGE_GAUGE.Set(runRecordingData.SumSelectRange)
			time.Sleep(5 * time.Second)
		}
	}()

}

func Monitor() {
	registerMetrics()
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}

	log.Fatal(http.ListenAndServe(":9090", nil))
	//rand.Seed(time.Now().Unix())
	//http.Handle("/metrics", newHandlerWithHistogram(promhttp.Handler(), histogramVec))
}

func ParseBody(body []byte) {
	var startRecordingValue storage.RunRecordingData

	err := json.Unmarshal(body, &startRecordingValue)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", startRecordingValue)
}

/* func newHandlerWithHistogram(handler http.Handler, histogram *prometheus.HistogramVec) http.Handler {
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
} */
