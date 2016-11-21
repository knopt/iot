package error

import "gopkg.in/gin-gonic/gin.v1"

// Error struct to store occured errors
type Error struct {
	Code int
	Err  error
}

// Handler handle error
func Handler(err *Error, context *gin.Context) {
	context.AbortWithError(err.Code, err.Err)
}
