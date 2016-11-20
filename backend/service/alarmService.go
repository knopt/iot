package service

import (
	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"
	databaseModel "github.com/knopt/iot/backend/database/model"
)

// GetAlarm by id from database and return to API
// TODO: implement
func (service *Service) GetAlarm(alarmID string) *apiModel.AlarmForm {
	dbAlarm := service.db.GetAlarm(bson.ObjectIdHex(alarmID))
	return databaseAlarmToAPIAlarmForm(dbAlarm)
}

// CreateAlarm in database and return its ID to API
func (service *Service) CreateAlarm(alarmForm *apiModel.AlarmForm) string {
	alarm := apiAlarmFormToDatabaseAlarm(alarmForm)
	alarmBsonID := service.db.InsertAlarm(alarm)
	return alarmBsonID.Hex()
}

func apiAlarmFormToDatabaseAlarm(alarmForm *apiModel.AlarmForm) *databaseModel.Alarm {
	return &databaseModel.Alarm{
		AlarmTime:    alarmForm.AlarmTime,
		CreatedAt:    alarmForm.CreatedAt,
		Description:  alarmForm.Description,
		DeviceID:     bson.ObjectIdHex(alarmForm.DeviceID),
		RepeatWeekly: alarmForm.RepeatWeekly,
		Weekday:      alarmForm.Weekday,
	}
}

func databaseAlarmToAPIAlarmForm(alarm *databaseModel.Alarm) *apiModel.AlarmForm {
	return &apiModel.AlarmForm{
		AlarmTime:    alarm.AlarmTime,
		CreatedAt:    alarm.CreatedAt,
		Description:  alarm.Description,
		DeviceID:     alarm.DeviceID.Hex(),
		RepeatWeekly: alarm.RepeatWeekly,
		Weekday:      alarm.Weekday,
	}
}
