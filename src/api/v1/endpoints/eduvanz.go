package endpoints

import (
	curd "Eduvanz/src/curd"
	INC "Eduvanz/src/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Register(process *gin.Context) {
	var registration INC.Register
	if err := process.BindJSON(&registration); err != nil {
		INC.Logger.Println("Bind json error in Register request -->", err)
		process.JSON(http.StatusBadRequest, "BadRequest")
		return
	}
	if registration.GuestCount < 0 || registration.GuestCount > 2 {
		process.JSON(http.StatusInternalServerError, "Only allowed GuestCount 0-2")
		return
	}
	s := strings.ToLower(registration.Profession)
	if s != "student" && s != "employed" {
		fmt.Println("Profession -->", strings.ToLower(registration.Profession))
		process.JSON(http.StatusInternalServerError, "Not allowed Profession")
		return
	}
	result := curd.InsertRegistrationDetails(registration)
	if result == 1 {
		process.JSON(http.StatusInternalServerError, "faliure")
		return
	}

	process.JSON(http.StatusOK, "success")
	return
}
func GetParticipants(process *gin.Context) {
	query_string := process.Request.URL.Query()
	page, _ := strconv.Atoi(query_string["page"][0])
	result, record := curd.GetParticipantsList(page)
	if result == 1 {
		process.JSON(http.StatusInternalServerError, "InternalServerError")
		return
	}
	process.JSON(http.StatusOK, record)
	return
}
func UpdateParticipants(process *gin.Context) {
	var UpdateRegistration INC.MeetUpList
	if err := process.BindJSON(&UpdateRegistration); err != nil {
		INC.Logger.Println("Bind json error in UpdateRegistration request -->", err)
		process.JSON(http.StatusBadRequest, "BadRequest")
		return
	}
	result := curd.UpdateParticipantsDetails(UpdateRegistration)
	if result == 1 {
		process.JSON(http.StatusInternalServerError, "InternalServerError")
		return
	}
	process.JSON(http.StatusOK, "success")
	return
}
