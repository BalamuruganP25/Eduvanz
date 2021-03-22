package messages

import (
	"database/sql"
	"log"
)

type DBcredentials struct {
	DBhost     string
	DBport     int
	DBuser     string
	DBpassword string
	DBname     string
}

var DBconfig DBcredentials

type UserApicredentials struct {
	ApiIp   string
	ApiPort string
}

var Apidetails UserApicredentials

type UserDBconnection struct {
	Connecction *sql.DB
}

var UserDB UserDBconnection

type logservicedetails struct {
	Logfoldername string
	Logfilepath   string
}

var Logserviceconfig logservicedetails

var Logger *log.Logger

type ErrorResponse struct {
	ErrorCode    int    `json:"errorcode"`
	ErrorMessage string `json:"errormessage"`
}

var ErrorDetails ErrorResponse
