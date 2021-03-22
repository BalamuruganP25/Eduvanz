package server

import (
	logservice "Eduvanz/src/Logservice"
	endpoints "Eduvanz/src/api/v1/endpoints"
	config "Eduvanz/src/config"
	db "Eduvanz/src/curd"
	INC "Eduvanz/src/messages"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransaction(t *testing.T) {
	config.ReadEnvironmentVariable()
	logservice.CreateLogMainFolder()
	logservice.CreateLogSubFolder()
	logservice.Applicationdetails()
	db.DBconnection(INC.DBconfig)
	router := gin.New()
	router.POST("/participants", endpoints.Register)
	Input := INC.Register{
		Name:       "bala",
		Address:    "bf",
		Age:        23,
		Profession: "student",
		GuestCount: 2,
		Locality:   "sure",
		Date:       "Mon Mar 22 2021 11:48:25 GMT+0530 (India Standard Time)",
	}
	RequestJson, err := json.Marshal(Input)
	if err != nil {
		fmt.Println(err)
	}
	req, _ := http.NewRequest("POST", "/participants", bytes.NewBuffer(RequestJson))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body.String())
	router.PUT("/participants", endpoints.UpdateParticipants)
	Update := INC.MeetUpList{
		Name:       "bala",
		Address:    "bf",
		Age:        23,
		Profession: "student",
		GuestCount: 2,
		Locality:   "sure",
		Date:       "Mon Mar 22 2021 11:48:25 GMT+0530 (India Standard Time)",
		Id:         1,
	}
	UpdateJson, err := json.Marshal(Update)
	if err != nil {
		fmt.Println(err)
	}
	updatereq, _ := http.NewRequest("PUT", "/participants", bytes.NewBuffer(UpdateJson))
	updateresp := httptest.NewRecorder()
	router.ServeHTTP(updateresp, updatereq)
	fmt.Println(updateresp.Body.String())
	router.GET("/participants", endpoints.GetParticipants)
	getreq, _ := http.NewRequest("GET", "/participants?page=0", nil)
	getresp := httptest.NewRecorder()
	router.ServeHTTP(getresp, getreq)
	fmt.Println(getresp.Body.String())
}
