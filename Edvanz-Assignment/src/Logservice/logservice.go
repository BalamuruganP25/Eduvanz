package Logservice

import (
	INC "Eduvanz/src/messages"
	"fmt"
	"log"
	"os"
	"time"
)

// create  Main log folder
func CreateLogMainFolder() {
	maindir := INC.Logserviceconfig.Logfilepath + INC.Logserviceconfig.Logfoldername + "/"
	if _, err := os.Stat(maindir); os.IsNotExist(err) {
		err = os.Mkdir(maindir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	}
}

// create sub log folder
func CreateLogSubFolder() {
	currentTime := time.Now().Local()
	subdirname := currentTime.Format("2006-01-02")
	fmt.Println("subdirname", subdirname)
	subdir := INC.Logserviceconfig.Logfilepath + INC.Logserviceconfig.Logfoldername + "/" + subdirname + "/"
	if _, err := os.Stat(subdir); os.IsNotExist(err) {
		err = os.Mkdir(subdir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	}
	createFile(subdir)
}

//create log file
func createFile(dir_path string) {
	// detect if file exists
	currentTime := time.Now().Local()
	currentdate := currentTime.Format("2006_01_02_15_04_05")
	filename := dir_path + "logdetails_" + currentdate + ".log"
	fmt.Println("file name  ==>", filename)
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var _, err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		//defer file.Close()
	}
	writelog(filename)
}

// create write log object
func writelog(dir_name string) {
	f, err := os.OpenFile(dir_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	INC.Logger = log.New(f, "", log.LstdFlags)
	//fmt.Println("INC.Logger -->",INC.Logger)
	//INC.Loggerinfo = log.New(f, "Info", log.LstdFlags)
}
func Applicationdetails() {
	currentTime := time.Now().Local()
	currentdate := currentTime.Format("2006-01-02 15:04:05")
	INC.Logger.Println("\t\t\tApplication Name-Jikkosoft")
	INC.Logger.Println("\t\t\tversion-1.0")
	INC.Logger.Println("\t\t\tApplication start date and time-", currentdate)
	INC.Logger.Println("\t\t\tDeveloper-P.Bala Murugan")
}
