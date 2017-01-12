package api

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/gin-gonic/gin.v1"
)

// GetCurrentTime returns current time in unix second
func (api *Api) GetCurrentTime(context *gin.Context) {
	resultString := fmt.Sprintf("*%v*", time.Now().Unix())

	context.JSON(http.StatusOK, resultString)
}
