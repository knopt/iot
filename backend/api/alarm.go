package api

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// SetAlarm sets alarm for given device
// TODO implement
func SetAlarm(context *gin.Context) {
	fmt.Printf("SetAlarm api call. To be implemented\n")
	context.JSON(http.StatusOK, gin.H{"status": "request succeded"})
}
