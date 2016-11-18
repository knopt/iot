package server

import (
	"os"

	"github.com/knopt/iot/backend/api"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

// StartServer starts server at :8080
func StartServer() {
	router := gin.Default()

	setUpRouter(router)
	router.Run(":8080")
}

func init() {
	log.SetOutput(os.Stderr)
}

func setUpRouter(router *gin.Engine) {
	allRoutes := router.Group("/")
	{
		alarm := allRoutes.Group("alarm")
		alarm.POST("", api.SetAlarm)
	}

	log.Info("router did setup")
}
