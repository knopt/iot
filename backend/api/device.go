package api

import (
	"net/http"

	"github.com/knopt/iot/backend/api/model"

	"gopkg.in/gin-gonic/gin.v1"
)

// GetDevice by given id
func (api *Api) GetDevice(context *gin.Context) {
	id := context.Param("id")

	responseDevice := api.Service.GetDevice(id)

	context.IndentedJSON(http.StatusOK, responseDevice)
}

// RegisterDevice in database and return its ID
func (api *Api) RegisterDevice(context *gin.Context) {
	var deviceForm model.DeviceForm

	if err := context.BindJSON(&deviceForm); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	deviceID := api.Service.RegisterDevice(&deviceForm)

	context.JSON(http.StatusCreated, gin.H{"id": deviceID})
}
