package services

import (
	"github.com/knopt/iot/backend/database"
)

//Service connects database and API
type Service struct {
	db *database.Database
}

//GetAlarm get alarm by id from database
//TODO: implement
func (service *Service) GetAlarm(alarmID int64) int64 {
	return service.db.GetAlarm(alarmID)
}
