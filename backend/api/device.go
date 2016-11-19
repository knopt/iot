package api

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// RegisterDevice register device and return id
func RegisterDevice(context *gin.Context) {
	fmt.Printf("RegisterDevice api call. To be implemented\n")
	context.JSON(http.StatusOK, gin.H{"id": "1"})
}
