package storage

import "time"

/*
	Struct to get all information from the bearer token returned by MySuperMon upon login
	It must match the json object that we get from MySuperMon
*/
type Token struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

/*
	Struct for the calls that the user makes with a Usecase and a ApplicationIdentifier.
	The information is gathered from the URL and added to the struct.
	Therefore, the end of the struct, the uri:xxx part, must match the definitions in the main method.
*/
type StartAndStopCommand struct {
	Usecase               string `uri:"Usecase" binding:"required"`
	ApplicationIdentifier string `uri:"Appiden" binding:"required"`
}

/*
	Struct for the calls that the user makes with a Username and a Password.
	The information is gathered from the URL and added to the struct.
	Therefore, the end of the struct, the uri:xxx part, must match the definitions in the main method.
*/
type LoginCommand struct {
	Username string `uri:"Username" binding:"required"`
	Password string `uri:"Password" binding:"required"`
}

type RunIdentificationData struct {
	IdNumber              int    `json:"idNum"`
	UsecaseIdentifier     string `json:"usecaseIdentifier"`
	ApplicationID         int    `json:"applicationId"`
	ApplicationName       string `json:"applicationName"`
	ApplicationIdentifier string `json:"applicationIdentifier"`
	//RunSituationResult    []RunClientData `json:"runSituationResult"`
}

type RunClientData struct {
	DataSourceId int              `json:"dataSourceId"`
	DatabaseType string           `json:"databaseType"`
	DatabaseName string           `json:"databaseName"`
	SchemaName   string           `json:"schemaName"`
	HostURL      string           `json:"hostUrl"`
	Data         RunRecordingData `json:"data"`
}

type RunRecordingData struct {
	SumRowsAffected        int        `json:"SUM_ROWS_AFFECTED"`
	SumSelectRange         int        `json:"SUM_SELECT_RANGE"`
	SumLockTime            float32    `json:"SUM_LOCK_TIME"`
	SumSortRows            int        `json:"SUM_SORT_ROWS"`
	SumErrors              int        `json:"SUM_ERRORS"`
	SumRowsSent            *int       `json:"SUM_ROWS_SENT"`
	SumSelectScan          int        `json:"SUM_SELECT_SCAN"`
	SumNoGoodIndexUsed     int        `json:"SUM_NO_GOOD_INDEX_USED"`
	ExecTimeMax            *time.Time `json:"EXEC_TIME_MAX"`
	SumSortScan            int        `json:"SUM_SORT_SCAN"`
	SumSelectRangeCheck    int        `json:"SUM_SELECT_RANGE_CHECK"`
	SumTimerWait           float64    `json:"SUM_TIMER_WAIT"`
	UsecaseIdentifier      string     `json:"USECASE_IDENTIFIER"`
	StartTimeStamp         time.Time  `json:"STARTTIMESTMAP"`
	SumRowsExamined        *int       `json:"SUM_ROWS_EXAMINED"`
	SumSelectFullJoin      int        `json:"SUM_SELECT_FULL_JOIN"`
	SumNoIndexUsed         int        `json:"SUM_NO_INDEX_USED"`
	CountStar              int        `json:"COUNT_STAR"`
	SumSelectFullRangeJoin int        `json:"SUM_SELECT_FULL_RANGE_JOIN"`
	SumSortMergePasses     int        `json:"SUM_SORT_MERGE_PASSES"`
	SumSortRange           int        `json:"SUM_SORT_RANGE"`
}
