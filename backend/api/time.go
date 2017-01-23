package api

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/gin-gonic/gin.v1"
)

// GetCurrentTime returns current time in unix second
func (api *Api) GetCurrentTime(context *gin.Context) {
	// fake - time zone - controller
	oneHour := time.Duration(time.Hour)
	resultString := fmt.Sprintf("*%v*", time.Now().Add(oneHour).Unix())

	fmt.Printf("sending time: %v\n", resultString)

	context.JSON(http.StatusOK, resultString)
}
