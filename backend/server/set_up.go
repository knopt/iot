package server

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"

	apis "github.com/knopt/iot/backend/api"
	db "github.com/knopt/iot/backend/database"
	services "github.com/knopt/iot/backend/service"
)

// StartServer starts server at given port
func StartServer(ip string) {
	dbConnection := db.NewDatabaseConnection()
	database := db.NewDatabase(dbConnection)
	service := services.NewService(database)
	api := apis.NewApi(service)
	router := gin.Default()
	setUpRouter(router, api)
	router.Run(ip)
}

func init() {
	log.SetOutput(os.Stderr)
}

func setUpRouter(router *gin.Engine, api *apis.Api) {
	allRoutes := router.Group("/")
	{
		alarm := allRoutes.Group("alarm")
		{
			alarm.GET("get/:id", api.GetAlarm)
			alarm.POST("set", api.SetAlarm)
			alarm.DELETE(":device/:id", api.DeleteAlarm)
		}
		device := allRoutes.Group("device")
		{
			device.POST("register", api.RegisterDevice)
			device.GET("get/:id", api.GetDevice)
		}
		statistics := allRoutes.Group("statistics")
		{
			statistics.POST("/", api.InsertStatistic)
			statistics.GET("/device/:id/date/from/:from/to/:to/type/:type", api.GetStatisticsByDeviceDataType)
			statistics.GET("/device/:id/date/to/:to/type/:type", api.GetStatisticsByDeviceDataType)
			statistics.GET("/device/:id/date/from/:from/type/:type", api.GetStatisticsByDeviceDataType)
		}
	}

	log.Info("router did setup")
}
