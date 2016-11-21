package api

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/knopt/iot/backend/api/model"
	"github.com/knopt/iot/backend/error"
)

// SetAlarm for given device
func (api *Api) SetAlarm(context *gin.Context) {
	var alarmForm model.AlarmForm

	if err := context.BindJSON(&alarmForm); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	alarmID, err := api.Service.CreateAlarm(&alarmForm)
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
	}

	context.JSON(http.StatusCreated, gin.H{"id": alarmID})
}

// GetAlarm by given id
func (api *Api) GetAlarm(context *gin.Context) {
	id := context.Param("id")

	responseAlarm, err := api.Service.GetAlarm(id)
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
	}

	context.IndentedJSON(http.StatusOK, responseAlarm)
}

// DeleteAlarm by given id
func (api *Api) DeleteAlarm(context *gin.Context) {
	fmt.Printf("DeleteAlarm api call. To pe implemented\n")
	context.JSON(http.StatusOK, gin.H{"status": "request succeded"})
}
