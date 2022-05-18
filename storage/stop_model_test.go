package storage_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/team7mysupermon/mySuperMon_Middleware/storage"
)

var (
	STOP_TEST = []byte(`{
		"status":"SUCCESS",
		"responseCode":200,
		"data":{
			"applicationId":203,
			"applicationName":"Konakart_Application",
			"applicationIdentifier":"d9cdf882-6f1e-45d3-b8ca-2d1b19d9712e",
			"idNum":1682,
			"usecaseIdentifier":"MIDDELWAREAPITEST",
			"startTimestamp":"2022-05-09T12:37:06.000+0000",
			"dataSourceList":[
				{
					"dataSourceId":232,
					"databaseType":"MySQL",
					"databaseName":"konakart",
					"schemaName":null,
					"hostUrl":"konakart",
					"data":{
						"STATEMENTS":31.0,
						"STATEMENT_LATENCY_IN_S":0.02,
						"FILE_IO_LATENCY_IN_S":0.0,
						"CURRENT_CONNECTIONS":0.0,
						"DATABASE_SIZE_IN_MB":4.6,
						"STATEMENT_AVG_LATENCY_IN_MS":0.53,
						"APPLICATION_ID":203.0,
						"FILE_IOS":12.0,
						"TABLE_SCANS":7.0,
						"DATA_SOURCE_ID":232.0,
						"USECASE_IDENTIFIER":0.0,
						"UNIQUE_USERS":1.0
					},
					"valueObjectList":[
						{
							"fieldName":"spl_per_sec",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"timespend",
							"newValue":99.0,
							"oldValue":27.0,
							"deviation":-266.67
						},
						{
							"fieldName":"sumWarnings",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumTimerWait",
							"newValue":0.0049,
							"oldValue":0.0021,
							"deviation":-133.33
						},
						{
							"fieldName":"sumSortScan",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSortRow",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSortRange",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSortMergePasses",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSelectScan",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSelectRangeCheck",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSelectFullRangeJoin",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSelectFullJoin",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumRowsSent",
							"newValue":23.0,
							"oldValue":8.0,
							"deviation":-187.5
						},
						{
							"fieldName":"sumRowsExamined",
							"newValue":11.0,
							"oldValue":4.0,
							"deviation":-175.0
						},
						{
							"fieldName":"sumRowsAffected",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumNoIndexUsed",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumNoGoodIndexUsed",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumLockTime",
							"newValue":1.0E-4,
							"oldValue":2.0E-4,
							"deviation":50.0
						},
						{
							"fieldName":"sumErrors",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumCreatedTmpTables",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumCreatedTmpDiskTables",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"rowsSentAvg",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"minTimerWait",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"maxTimerWait",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"execTimeTotal",
							"newValue":0.0049,
							"oldValue":0.0021,
							"deviation":-133.33
						},
						{
							"fieldName":"execTimeMax",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"execTimeAvgMS",
							"newValue":-0.0292,
							"oldValue":-0.0292,
							"deviation":0.0
						},
						{
							"fieldName":"execCount",
							"newValue":23.0,
							"oldValue":8.0,
							"deviation":-187.5
						},
						{
							"fieldName":"avgTimerWait",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						},
						{
							"fieldName":"sumSelectRange",
							"newValue":0.0,
							"oldValue":0.0,
							"deviation":0.0
						}
					]
				}
			]
		},
		"errorMessage":null,
		"errorCode":null,
		"reportLink":"http://app.mysupermon.com/#/report/view/MIDDELWAREAPITEST/1682"
	
	}`)
)

func setUpStopData() storage.StopAutoGenerated{
	var stopAutoGenerated storage.StopAutoGenerated

	err := json.Unmarshal(STOP_TEST, &stopAutoGenerated)
		if err != nil {
			log.Panicln(err)
		}

	return stopAutoGenerated
}

func TestStopAutoGenerated(t *testing.T) {
	data := setUpStopData()

	assert.Equal(t, data.Status, "SUCCESS")
	assert.Equal(t, data.ResponseCode, 200)
	assert.Equal(t, data.ErrorMessage, nil)
	assert.Equal(t, data.ErrorCode, nil)
	assert.Equal(t, data.ReportLink, "http://app.mysupermon.com/#/report/view/MIDDELWAREAPITEST/1682")
}

func TestStopMetaData(t *testing.T) {
	data := setUpStopData().StopMetaData

	assert.Equal(t, data.ApplicationID, 203)
	assert.Equal(t, data.ApplicationName, "Konakart_Application")
	assert.Equal(t, data.ApplicationIdentifier, "d9cdf882-6f1e-45d3-b8ca-2d1b19d9712e")
	assert.Equal(t, data.IDNum, 1682)
	assert.Equal(t, data.UsecaseIdentifier, "MIDDELWAREAPITEST")
	assert.Equal(t, data.StartTimestamp, "2022-05-09T12:37:06.000+0000")
}

/* func TestStopSituationResult(t *testing.T) {
	data := setUpStartData().StartMetaData.StartDataSourceList[0]

	assert.Equal(t, data.DataSourceID, 232)
	assert.Equal(t, data.DatabaseType, "MySQL")
	assert.Equal(t, data.DatabaseName, "konakart")
	assert.Equal(t, data.SchemaName, "konakart")
	assert.Equal(t, data.HostURL, "34.88.216.230")
}

func TestStopData(t *testing.T){
	data := setUpStartData().StartMetaData.StartDataSourceList[0].StartData

	assert.Equal(t, data.Statements, 31.0)
	assert.Equal(t, data.StatementLatencyInS, 0.02)
	assert.Equal(t, data.FileIoLatencyInS, 0.0)
	assert.Equal(t, data.CurrentConnections, 0.0)
	assert.Equal(t, data.DatabaseSizeInMb, 4.6)

	assert.Equal(t, data.StatementAvgLatencyInMs, 0.53)
	assert.Equal(t, data.ApplicationID, 203.0)
	assert.Equal(t, data.FileIos, 12.0)
	assert.Equal(t, data.TableScans, 7.0)
	assert.Equal(t, data.DataSourceID, 232.0)

	assert.Equal(t, data.UsecaseIdentifier, 0.0)
	assert.Equal(t, data.UniqueUsers, 1.0)
} */