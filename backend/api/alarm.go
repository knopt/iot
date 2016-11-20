package api

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/knopt/iot/backend/api/model"
)

// SetAlarm for given device
func (api *Api) SetAlarm(context *gin.Context) {
	var alarmForm model.AlarmForm

	if err := context.BindJSON(&alarmForm); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	alarmID := api.Service.CreateAlarm(&alarmForm)

	context.JSON(http.StatusCreated, gin.H{"id": alarmID})
}

// GetAlarm by given id
func (api *Api) GetAlarm(context *gin.Context) {
	id := context.Param("id")

	responseAlarm := api.Service.GetAlarm(id)

	context.IndentedJSON(http.StatusOK, responseAlarm)
}

// DeleteAlarm by given id
func (api *Api) DeleteAlarm(context *gin.Context) {
	fmt.Printf("DeleteAlarm api call. To pe implemented\n")
	context.JSON(http.StatusOK, gin.H{"status": "request succeded"})
}
