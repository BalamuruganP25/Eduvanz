package server

import (
	endpoints "Eduvanz/src/api/v1/endpoints"
	INC "Eduvanz/src/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "*")
	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			INC.ErrorDetails.ErrorCode = 401
			INC.ErrorDetails.ErrorMessage = "Unauthorized"
			c.AbortWithStatusJSON(401, INC.ErrorDetails)
		}
		c.Next()
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(200)
	}
}
func ApiServer() {
	fmt.Println("Api server start")
	Apiconfig := INC.Apidetails.ApiIp + ":" + INC.Apidetails.ApiPort
	fmt.Println("Apiconfig -->", Apiconfig)
	router := gin.Default()
	router.Use(CORS)
	gin.SetMode(gin.ReleaseMode)
	router.POST("/participants", endpoints.Register)
	router.GET("/participants", endpoints.GetParticipants)
	router.PUT("/participants", endpoints.UpdateParticipants)
	router.Run(Apiconfig)
}
