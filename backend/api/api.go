package api

import (
	"github.com/knopt/iot/backend/service"
)

// Api single
type Api struct {
	Service *service.Service
}

//NewApi create new Api model
func NewApi(newService *service.Service) *Api {
	return &Api{newService}
}
