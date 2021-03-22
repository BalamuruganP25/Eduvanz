package main

import (
	logservice "Eduvanz/src/Logservice"
	config "Eduvanz/src/config"
	db "Eduvanz/src/curd"
	INC "Eduvanz/src/messages"
	api "Eduvanz/src/server"
	"fmt"
	//"os"
)

func main() {
	fmt.Println("Eduvanz application started")
	config.ReadEnvironmentVariable()
	logservice.CreateLogMainFolder()
	logservice.CreateLogSubFolder()
	logservice.Applicationdetails()
	db.DBconnection(INC.DBconfig)
	api.ApiServer()

}
