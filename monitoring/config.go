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
)

var (
	startRecordingValue StartRecordingValues
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
	/*
		SumSelectRange         int       `json:"SUM_SELECT_RANGE"`
		SumLockTime            time.Time `json:"SUM_LOCK_TIME"`
		SumSortRows            int       `json:"SUM_SORT_ROWS"`
		SumErrors              int       `json:"SUM_ERRORS"`
		SumRowsSent            int       `json:"SUM_ROWS_SENT"`
		SumSelectScan          int       `json:"SUM_SELECT_SCAN"`
		SumNoGoodIndexUsed     int       `json:"SUM_NO_GOOD_INDEX_USED"`
		ExecTimeMax            time.Time `json:"EXEC_TIME_MAX"`
		SumSortScan            int       `json:"SUM_SORT_SCAN"`
		SumSelectRangeCheck    int       `json:"SUM_SELECT_RANGE_CHECK"`
		SumTimerWait           time.Time `json:"SUM_TIMER_WAIT"`
		UsecaseIdentifier      string    `json:"USECASE_IDENTIFIER"`
		StartTimeStamp          time.Time `json:"STARTTIMESTMAP"`
		SumRowsExamined        int       `json:"SUM_ROWS_EXAMINED"`
		SumSelectFullJoin      int       `json:"SUM_SELECT_FULL_JOIN"`
		SumNoIndexUsed         int       `json:"SUM_NO_INDEX_USED"`
		CountStar              int       `json:"COUNT_STAR"`
		SumSelectFullRangeJoin int       `json:"SUM_SELECT_FULL_RANGE_JOIN"`
		SumSortMergePasses     int       `json:"SUM_SORT_MERGE_PASSES"`
		SumSortRange           int       `json:"SUM_SORT_RANGE"` */

)

func recordMetrics() {

	go func() {
		for {
			SUM_ROWS_AFFECTED_GAUGE.Set(startRecordingValue.SumRowsAffected)
			SUM_SELECTED_RANGE_GAUGE.Set(startRecordingValue.SumSelectRange)
			time.Sleep(5 * time.Second)
		}
	}()

}

func registerMetricsToPrometheus() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(SUM_ROWS_AFFECTED_GAUGE)
	prometheus.MustRegister(SUM_SELECTED_RANGE_GAUGE)
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

type StartRecordingValues struct {
	SumRowsAffected        float64   `json:"SUM_ROWS_AFFECTED"`
	SumSelectRange         float64   `json:"SUM_SELECT_RANGE"`
	SumLockTime            time.Time `json:"SUM_LOCK_TIME"`
	SumSortRows            float64   `json:"SUM_SORT_ROWS"`
	SumErrors              int       `json:"SUM_ERRORS"`
	SumRowsSent            int       `json:"SUM_ROWS_SENT"`
	SumSelectScan          int       `json:"SUM_SELECT_SCAN"`
	SumNoGoodIndexUsed     int       `json:"SUM_NO_GOOD_INDEX_USED"`
	ExecTimeMax            time.Time `json:"EXEC_TIME_MAX"`
	SumSortScan            int       `json:"SUM_SORT_SCAN"`
	SumSelectRangeCheck    int       `json:"SUM_SELECT_RANGE_CHECK"`
	SumTimerWait           time.Time `json:"SUM_TIMER_WAIT"`
	UsecaseIdentifier      string    `json:"USECASE_IDENTIFIER"`
	StartTimeStamp         time.Time `json:"STARTTIMESTMAP"`
	SumRowsExamined        int       `json:"SUM_ROWS_EXAMINED"`
	SumSelectFullJoin      int       `json:"SUM_SELECT_FULL_JOIN"`
	SumNoIndexUsed         int       `json:"SUM_NO_INDEX_USED"`
	CountStar              int       `json:"COUNT_STAR"`
	SumSelectFullRangeJoin int       `json:"SUM_SELECT_FULL_RANGE_JOIN"`
	SumSortMergePasses     int       `json:"SUM_SORT_MERGE_PASSES"`
	SumSortRange           int       `json:"SUM_SORT_RANGE"`
}

func ParseBody(body []byte) {
	var startRecordingValue StartRecordingValues

	err := json.Unmarshal(body, &startRecordingValue)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", startRecordingValue)

	setStartRecordingValues(startRecordingValue)
}

func setStartRecordingValues(_startRecordingValue StartRecordingValues) {
	startRecordingValue = _startRecordingValue
}

func Monitor() {
	registerMetricsToPrometheus()
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		return
	}

	//rand.Seed(time.Now().Unix())
	//http.Handle("/metrics", newHandlerWithHistogram(promhttp.Handler(), histogramVec))

	log.Fatal(http.ListenAndServe(":2112", nil))
}
