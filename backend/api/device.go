package api

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/knopt/iot/backend/api/model"
	"github.com/knopt/iot/backend/error"
)

// GetDevice by given id
func (api *Api) GetDevice(context *gin.Context) {
	id := context.Param("id")

	responseDevice, err := api.Service.GetDevice(id)
	if err != nil {
		fmt.Printf("error in api")
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
	}

	context.IndentedJSON(http.StatusOK, responseDevice)
}

// GetDevices from db and convert them to api forms
func (api *Api) GetDevices(context *gin.Context) {
	responseDevices, err := api.Service.GetDevices()
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
	}

	context.IndentedJSON(http.StatusOK, responseDevices)
}

// RegisterDevice in database and return its ID
func (api *Api) RegisterDevice(context *gin.Context) {
	var deviceForm model.DeviceForm

	if err := context.BindJSON(&deviceForm); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	deviceID, err := api.Service.RegisterDevice(&deviceForm)
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
	}

	context.JSON(http.StatusCreated, gin.H{"id": deviceID})
}
