package api

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// RegisterDevice register device and return id
func (api *Api) RegisterDevice(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "To be implemented.", "mock_id": "123456789012345678901234"})
}
