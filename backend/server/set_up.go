package server

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-contrib/cors.v1"
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

	CORSConfig := setupCORS()

	router := gin.Default()
	setUpRouter(router, api, CORSConfig)
	router.Run(ip)
}

func init() {
	log.SetOutput(os.Stderr)
}

func setUpRouter(router *gin.Engine, api *apis.Api, corsConfig *cors.Config) {
	CORSHandler := cors.New(*corsConfig)

	router.Use(CORSHandler)
	allRoutes := router.Group("/")
	{
		alarm := allRoutes.Group("alarm")
		{
			alarm.GET("get/id/:id", api.GetAlarm)
			alarm.GET("get/nearest/id/:id", api.GetNearestAlarmByDevice)
			alarm.GET("get/nearest/string/id/:id", api.GetNearestAlarmByDeviceString)
			alarm.GET("get/device/:device/all", api.GetAlarmsByDevice)
			alarm.POST("set", api.SetAlarm)
			alarm.DELETE(":device/:id", api.DeleteAlarm)
		}
		device := allRoutes.Group("device")
		{
			device.POST("register", api.RegisterDevice)
			device.GET("get/id/:id", api.GetDevice)
			device.GET("get/first/id/:id", api.GetNearestAlarmByDevice)
			device.GET("get/all", api.GetDevices)
		}
		statistics := allRoutes.Group("statistics")
		{
			statistics.POST("/", api.InsertStatistic)
			statistics.GET("/get/device/:id/date/from/:from/to/:to/type/:type", api.GetStatisticsByDeviceDataType)
			statistics.GET("/get/device/:id/date/to/:to/type/:type", api.GetStatisticsByDeviceDataType)
			statistics.GET("/get/device/:id/date/from/:from/type/:type", api.GetStatisticsByDeviceDataType)
			statistics.GET("/get/types/device/:id", api.GetStatisticsTypes)
		}
	}

	log.Info("router did setup")
}

func setupCORS() *cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	return &config
}
