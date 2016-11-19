package server

import (
	"os"

	"github.com/knopt/iot/backend/api"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

// StartServer starts server at given port
func StartServer(port string) {
	router := gin.Default()

	setUpRouter(router)
	router.Run(port)
}

func init() {
	log.SetOutput(os.Stderr)
}

func setUpRouter(router *gin.Engine) {
	allRoutes := router.Group("/")
	{
		alarm := allRoutes.Group("alarm")
		{
			alarm.POST("set", api.SetAlarm)
			alarm.DELETE(":device/:id", api.DeleteAlarm)
		}
		device := allRoutes.Group("device")
		{
			device.GET("/register", api.RegisterDevice)
		}
	}

	log.Info("router did setup")
}
