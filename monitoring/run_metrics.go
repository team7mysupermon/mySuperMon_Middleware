package monitoring

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (

	//RUN METRICS :
	RUN_SUM_ROWS_AFFECTED_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_ROWS_AFFECTED",
			Help: "Value of the no. of times a row is affected during monitoring",
		})

	RUN_SUM_SELECTED_RANGE_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SELECTED_RANGE",
			Help: "",
		})

	RUN_SUM_LOCK_TIME_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_LOCK_TIME",
			Help: "",
		})

	RUN_SUM_SORT_ROWS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SORT_ROWS",
			Help: "",
		})

	RUN_SUM_ERRORS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_ERRORS",
			Help: "",
		})

	RUN_SUM_ROWS_SENT_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_ROWS_SENT",
			Help: "",
		})

	RUN_SUM_SELECT_SCAN_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SELECT_SCAN",
			Help: "",
		})

	RUN_SUM_NO_GOOD_INDEX_USED_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_NO_GOOD_INDEX_USED",
			Help: "",
		})

	RUN_EXEC_TIME_MAX_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_EXEC_TIME_MAX",
			Help: "",
		})

	RUN_SUM_SORT_SCAN_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SUM_SORT_SCAN",
			Help: "",
		})

	RUN_SUM_SELECT_RANGE_CHECK_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SELECT_RANGE_CHECK",
			Help: "",
		})

	RUN_SUM_TIMER_WAIT_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_TIMER_WAIT",
			Help: "",
		})

	RUN_USECASE_IDENTIFIER_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_USECASE_IDENTIFYER",
			Help: "",
		})

	RUN_START_TIMESTAMP_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_START_TIMESTAMP",
			Help: "",
		})

	RUN_SUM_ROWS_EXAMINED_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_ROWS_EXAMINED",
			Help: "",
		})

	RUN_SUM_SELECT_FULL_JOIN_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SELECT_FULL_JOIN",
			Help: "",
		})

	RUN_SUM_NO_INDEX_USED_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_NO_INDEX_USED",
			Help: "",
		})

	RUN_COUNT_STAR_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_COUNT_STAR",
			Help: "",
		})

	RUN_SUM_SELECT_FULL_RANGE_JOIN_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SELECT_FULL_RANGE_JOIN",
			Help: "",
		})

	RUN_SUM_SORT_MERGE_PASSES_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SORT_MERGE_PASSES",
			Help: "",
		})

	RUN_SUM_SORT_RANGE_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "RUN_SUM_SORT_RANGE",
			Help: "",
		})
)

// registers metrics to expose
func runRegisterMetrics() {
	prometheus.MustRegister(RUN_SUM_ROWS_AFFECTED_GAUGE)
	prometheus.MustRegister(RUN_SUM_SELECTED_RANGE_GAUGE)
	prometheus.MustRegister(RUN_SUM_LOCK_TIME_GAUGE)
	prometheus.MustRegister(RUN_SUM_SORT_ROWS_GAUGE)
	prometheus.MustRegister(RUN_SUM_ERRORS_GAUGE)
	prometheus.MustRegister(RUN_SUM_ROWS_SENT_GAUGE)
	prometheus.MustRegister(RUN_SUM_SELECT_SCAN_GAUGE)
	prometheus.MustRegister(RUN_SUM_NO_GOOD_INDEX_USED_GAUGE)
	prometheus.MustRegister(RUN_EXEC_TIME_MAX_GAUGE)
	prometheus.MustRegister(RUN_SUM_SORT_SCAN_GAUGE)
	prometheus.MustRegister(RUN_SUM_SELECT_RANGE_CHECK_GAUGE)
	prometheus.MustRegister(RUN_SUM_TIMER_WAIT_GAUGE)
	//prometheus.MustRegister(RUN_SUM_LOCK_TIME_GAUGE)
	prometheus.MustRegister(RUN_USECASE_IDENTIFIER_GAUGE)
	prometheus.MustRegister(RUN_START_TIMESTAMP_GAUGE)
	prometheus.MustRegister(RUN_SUM_ROWS_EXAMINED_GAUGE)
	prometheus.MustRegister(RUN_SUM_SELECT_FULL_JOIN_GAUGE)
	prometheus.MustRegister(RUN_SUM_NO_INDEX_USED_GAUGE)
	prometheus.MustRegister(RUN_COUNT_STAR_GAUGE)
	prometheus.MustRegister(RUN_SUM_SELECT_FULL_RANGE_JOIN_GAUGE)
	prometheus.MustRegister(RUN_SUM_SORT_MERGE_PASSES_GAUGE)
	prometheus.MustRegister(RUN_SUM_SORT_RANGE_GAUGE)
}

func runRecordMetrics() {

	go func() {
		for {
			runData := runAutoGenerated.RunMetaData.RunSituationResult[0].RunData

			RUN_SUM_ROWS_AFFECTED_GAUGE.Set(float64(runData.SumRowsAffected))
			RUN_SUM_SELECTED_RANGE_GAUGE.Set(float64(runData.SumSelectRange))
			RUN_SUM_LOCK_TIME_GAUGE.Set(runData.SumLockTime)
			RUN_SUM_SORT_ROWS_GAUGE.Set(float64(runData.SumSortRows))
			RUN_SUM_ERRORS_GAUGE.Set(float64(runData.SumErrors))
			RUN_SUM_ROWS_SENT_GAUGE.Set(float64(runData.SumSelectScan))
			RUN_SUM_SELECT_SCAN_GAUGE.Set(float64(runData.SumSelectScan))
			RUN_SUM_NO_GOOD_INDEX_USED_GAUGE.Set(float64(runData.SumNoGoodIndexUsed))
			//RUN_EXEC_TIME_MAX_GAUGE.Set(runData.ExecTimeMax)
			RUN_SUM_SORT_SCAN_GAUGE.Set(float64(runData.SumSortScan))
			RUN_SUM_SELECT_RANGE_CHECK_GAUGE.Set(float64(runData.SumSelectRangeCheck))
			RUN_SUM_TIMER_WAIT_GAUGE.Set(runData.SumTimerWait)
			//RUN_USECASE_IDENTIFIER_GAUGE.Set(runData.UsecaseIdentifier)
			//RUN_START_TIMESTAMP_GAUGE.Set(runData.Starttimestmap)
			RUN_SUM_ROWS_EXAMINED_GAUGE.Set(float64(runData.SumRowsAffected))
			RUN_SUM_SELECT_FULL_JOIN_GAUGE.Set(float64(runData.SumSelectFullJoin))
			RUN_SUM_NO_INDEX_USED_GAUGE.Set(float64(runData.SumNoIndexUsed))
			RUN_COUNT_STAR_GAUGE.Set(float64(runData.CountStar))
			RUN_SUM_SELECT_FULL_RANGE_JOIN_GAUGE.Set(float64(runData.SumSelectFullRangeJoin))
			RUN_SUM_SORT_MERGE_PASSES_GAUGE.Set(float64(runData.SumSortMergePasses))
			RUN_SUM_SORT_RANGE_GAUGE.Set(float64(runData.SumSortRange))

			time.Sleep(5 * time.Second)
		}
	}()

}
