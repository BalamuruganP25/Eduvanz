package curd

import (
	INC "Eduvanz/src/messages"
	"fmt"
	"strconv"
)

func InsertRegistrationDetails(input INC.Register) int {

	RegistrationInsert := "insert into meetup(name,age,date,profession,locality,address) values ($1,$2,$3,$4,$5,$6)"
	_, err := INC.UserDB.Connecction.Exec(RegistrationInsert, input.Name, input.Age, input.Date, input.Profession, input.Locality, input.Address)
	if err != nil {
		INC.Logger.Println("create role details err -->", err)
		return 1
	}

	return 0

}

func GetParticipantsList(pagenumber int) (int, []INC.MeetUpList) {

	var Registration INC.MeetUpList
	var RegistrationList []INC.MeetUpList
	GetRegistration := "select id,name,age,date,profession,locality,address from meetup LIMIT 10 OFFSET " + strconv.Itoa(pagenumber)

	rows, err := INC.UserDB.Connecction.Query(GetRegistration)
	for rows.Next() {
		err = rows.Scan(&Registration.Id, &Registration.Name, &Registration.Age, &Registration.Date, &Registration.Profession, &Registration.Locality, &Registration.Address)
		if err != nil {
			fmt.Println("get user details scan error -->", err)
			return 1, RegistrationList
		}

		RegistrationList = append(RegistrationList, Registration)

	}
	return 0, RegistrationList

}

func UpdateParticipantsDetails(update INC.MeetUpList) int {

	UpdateQuery := "update  meetup set name=$1,age=$2,date=$3,profession=$4,locality=$5,address=$6 where id =$7"

	_, err := INC.UserDB.Connecction.Exec(UpdateQuery, update.Name, update.Age, update.Date, update.Profession, update.Locality, update.Address, update.Id)
	if err != nil {
		INC.Logger.Println("UpdateParticipants details err -->", err)
		return 1
	}
	return 0

}
