package curd

import (
	INC "Eduvanz/src/messages"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func DBconnection(DB INC.DBcredentials) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", DB.DBhost, DB.DBport, DB.DBuser, DB.DBpassword, DB.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Println("Successfully connected!")
	INC.UserDB.Connecction = db
}
