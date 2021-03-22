package config

import (
	INC "Eduvanz/src/messages"
	"fmt"
	"os"
	"strconv"
)

func ReadEnvironmentVariable() {
	fmt.Println("Reading Environment Variable")
	_, DBnameStatus := os.LookupEnv("Eduvanz_DATABASENAME")
	if !DBnameStatus {
		fmt.Println("DBnameStatus Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.DBconfig.DBname = os.Getenv("Eduvanz_DATABASENAME")
		fmt.Println("database name:", INC.DBconfig.DBname)
	}
	_, DBpasswordStatus := os.LookupEnv("Eduvanz_DATABASEPASSWORD")
	if !DBpasswordStatus {
		fmt.Println("database  password Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.DBconfig.DBpassword = os.Getenv("Eduvanz_DATABASEPASSWORD")
		fmt.Println("database password:", INC.DBconfig.DBpassword)
	}
	_, DBuserstatus := os.LookupEnv("Eduvanz_DATABASEUSER")
	if !DBuserstatus {
		fmt.Println("data base user   Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.DBconfig.DBuser = os.Getenv("Eduvanz_DATABASEUSER")
		fmt.Println("database user name:", INC.DBconfig.DBuser)
	}
	_, DBPortStatus := os.LookupEnv("Eduvanz_DATABASEPORT")
	if !DBPortStatus {
		fmt.Println("database port  Not found in Environment variable file")
		os.Exit(3)
	} else {
		getDBport, err := strconv.Atoi(os.Getenv("Eduvanz_DATABASEPORT"))
		if err != nil {
			fmt.Println("Port read error")
			os.Exit(3)
		} else {
			INC.DBconfig.DBport = getDBport
			fmt.Println("database port:", INC.DBconfig.DBport)
		}
	}
	_, HostStatus := os.LookupEnv("Eduvanz_DATABASEHOST")
	if !HostStatus {
		fmt.Println("Host  Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.DBconfig.DBhost = os.Getenv("Eduvanz_DATABASEHOST")
		fmt.Println("database host:", INC.DBconfig.DBhost)
	}
	_, ApiIpStatus := os.LookupEnv("Eduvanz_API_IP")
	if !ApiIpStatus {
		fmt.Println("api ip  Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Apidetails.ApiIp = os.Getenv("Eduvanz_API_IP")
		fmt.Println("api ip :", INC.Apidetails.ApiIp)
	}
	_, PortStatus := os.LookupEnv("Eduvanz_API_PORT")
	if !PortStatus {
		fmt.Println("api port  Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Apidetails.ApiPort = os.Getenv("Eduvanz_API_PORT")
		fmt.Println("api port :", INC.Apidetails.ApiPort)
	}
	_, logfoldernamestatus := os.LookupEnv("Eduvanz_LOG_FOLDER_NAME")
	if !logfoldernamestatus {
		fmt.Println("logfoldername Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Logserviceconfig.Logfoldername = os.Getenv("Eduvanz_LOG_FOLDER_NAME")
		fmt.Println("logfoldername :", INC.Logserviceconfig.Logfoldername)
	}
	_, logfilepathstatus := os.LookupEnv("Eduvanz_LOG_FILEPATH")
	if !logfilepathstatus {
		fmt.Println("logfilepath Not found in Environment variable file")
		os.Exit(3)
	} else {
		INC.Logserviceconfig.Logfilepath = os.Getenv("Eduvanz_LOG_FILEPATH")
		fmt.Println("logfilepath :", INC.Logserviceconfig.Logfilepath)
	}
}
