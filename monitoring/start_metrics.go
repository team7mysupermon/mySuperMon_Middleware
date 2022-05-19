package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	START_STATEMENTS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_STATEMENTS_GAUGE",
			Help: "",
		})

	START_STATEMENT_LATENCY_IN_S_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_STATEMENT_LATENCY_IN_S_GAUGE",
			Help: "",
		})

	START_FILE_IO_LATENCY_IN_S_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_FILE_IO_LATENCY_IN_S_GAUGE",
			Help: "",
		})

	START_CURRENT_CONNECTIONS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_CURRENT_CONNECTIONS_GAUGE",
			Help: "",
		})

	START_DATABASE_SIZE_IN_MB_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_DATABASE_SIZE_IN_MB_GAUGE",
			Help: "",
		})

	START_STATEMENT_AVG_LATENCY_IN_MS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_STATEMENT_AVG_LATENCY_IN_MS_GAUGE",
			Help: "",
		})

	START_APPLICATION_ID_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_APPLICATION_ID_GAUGE",
			Help: "",
		})

	START_FILE_IOS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_FILE_IOS_GAUGE",
			Help: "",
		})

	START_TABLE_SCANS_GAUGE = prometheus.NewGauge( 
		prometheus.GaugeOpts{
			Name: "START_TABLE_SCANS_GAUGE",
			Help: "",
		})

	START_DATA_SOURCE_ID_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_DATA_SOURCE_ID_GAUGE",
			Help: "",
		})

	START_USECASE_IDENTIFIER_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_USECASE_IDENTIFIER_GAUGE",
			Help: "",
		})

	START_UNIQUE_USERS_GAUGE = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "START_UNIQUE_USERS_GAUGE",
			Help: "",
		})
)

func startRegisterMetrics() {
	prometheus.MustRegister(START_STATEMENTS_GAUGE)
	prometheus.MustRegister(START_STATEMENT_LATENCY_IN_S_GAUGE)
	prometheus.MustRegister(START_FILE_IO_LATENCY_IN_S_GAUGE)
	prometheus.MustRegister(START_CURRENT_CONNECTIONS_GAUGE)
	prometheus.MustRegister(START_DATABASE_SIZE_IN_MB_GAUGE)
	prometheus.MustRegister(START_STATEMENT_AVG_LATENCY_IN_MS_GAUGE)
	prometheus.MustRegister(START_APPLICATION_ID_GAUGE)
	prometheus.MustRegister(START_FILE_IOS_GAUGE)
	prometheus.MustRegister(START_TABLE_SCANS_GAUGE)
	prometheus.MustRegister(START_DATA_SOURCE_ID_GAUGE)
	prometheus.MustRegister(START_USECASE_IDENTIFIER_GAUGE)
	prometheus.MustRegister(START_UNIQUE_USERS_GAUGE)
}

func startRecordMetrics() {

	go func() {
		for {
			startdata := startAutoGenerated.StartMetaData.StartDataSourceList[0].StartData

			START_STATEMENTS_GAUGE.Set(startdata.Statements)
			START_STATEMENT_LATENCY_IN_S_GAUGE.Set(startdata.StatementLatencyInS)
			START_FILE_IO_LATENCY_IN_S_GAUGE.Set(startdata.FileIoLatencyInS)
			START_CURRENT_CONNECTIONS_GAUGE.Set(startdata.CurrentConnections)
			START_DATABASE_SIZE_IN_MB_GAUGE.Set(startdata.DatabaseSizeInMb)
			START_STATEMENT_AVG_LATENCY_IN_MS_GAUGE.Set(startdata.StatementAvgLatencyInMs)
			START_APPLICATION_ID_GAUGE.Set(startdata.ApplicationID)
			START_FILE_IOS_GAUGE.Set(startdata.FileIos)
			START_TABLE_SCANS_GAUGE.Set(startdata.TableScans)
			START_DATA_SOURCE_ID_GAUGE.Set(startdata.DataSourceID)
			START_USECASE_IDENTIFIER_GAUGE.Set(startdata.UsecaseIdentifier)
			START_UNIQUE_USERS_GAUGE.Set(startdata.UniqueUsers)
			time.Sleep(5 * time.Second)
		}
	}()
}
