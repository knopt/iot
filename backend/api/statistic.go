package api

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/knopt/iot/backend/api/model"
	"github.com/knopt/iot/backend/error"
)

// GetStatisticsByDeviceDataType by given requests parameters
func (api *Api) GetStatisticsByDeviceDataType(context *gin.Context) {
	deviceID := context.Param("id")
	dateFrom := context.Param("from")
	dateTo := context.Param("to")
	dataType := context.Param("type")

	responseStatistics, err := api.Service.GetStatistics(deviceID, dateFrom, dateTo, dataType)
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
		return
	}

	context.IndentedJSON(http.StatusOK, responseStatistics)
}

//InsertStatistic from api statistic form
func (api *Api) InsertStatistic(context *gin.Context) {
	var statisticForm model.StatisticForm

	if err := context.BindJSON(&statisticForm); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := api.Service.InsertStatistic(&statisticForm)
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
		return
	}

	context.String(http.StatusOK, "Success")
}

//GetStatisticsTypes by deviceID
func (api *Api) GetStatisticsTypes(context *gin.Context) {
	deviceID := context.Param("id")

	responseTypes, err := api.Service.GetStatisticsTypes(deviceID)
	if err != nil {
		error.Handler(&error.Error{Code: http.StatusBadRequest, Err: err}, context)
		return
	}

	context.IndentedJSON(http.StatusOK, responseTypes)
}
