package service

import (
	"errors"

	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"
	databaseModel "github.com/knopt/iot/backend/database/model"
)

// GetAlarm by id from database and return to API
func (service *Service) GetAlarm(alarmID string) (*apiModel.AlarmForm, error) {
	dbAlarm, err := service.db.GetAlarm(bson.ObjectIdHex(alarmID))
	if err != nil {
		return nil, err
	}
	return databaseAlarmToAPIAlarmForm(dbAlarm), nil
}

// CreateAlarm in database and return its ID to API
func (service *Service) CreateAlarm(alarmForm *apiModel.AlarmForm) (*string, error) {
	alarm, err := apiAlarmFormToDatabaseAlarm(alarmForm)
	if err != nil {
		return nil, err
	}
	alarmBsonID, err := service.db.InsertAlarm(alarm)
	if err != nil {
		return nil, err
	}
	alarmHexBsonID := alarmBsonID.Hex()
	return &alarmHexBsonID, nil
}

func apiAlarmFormToDatabaseAlarm(alarmForm *apiModel.AlarmForm) (*databaseModel.Alarm, error) {
	if len(alarmForm.DeviceID) != 24 {
		return nil, errors.New("Wrong deviceID length")
	} else if !bson.IsObjectIdHex(alarmForm.DeviceID) {
		return nil, errors.New("Wrong format of deviceID")
	}
	return &databaseModel.Alarm{
		AlarmTime:    alarmForm.AlarmTime,
		CreatedAt:    alarmForm.CreatedAt,
		Description:  alarmForm.Description,
		DeviceID:     bson.ObjectIdHex(alarmForm.DeviceID),
		RepeatWeekly: alarmForm.RepeatWeekly,
		Weekday:      alarmForm.Weekday,
	}, nil
}

func databaseAlarmToAPIAlarmForm(alarm *databaseModel.Alarm) *apiModel.AlarmForm {
	return &apiModel.AlarmForm{
		ID:           alarm.ID.Hex(),
		AlarmTime:    alarm.AlarmTime,
		CreatedAt:    alarm.CreatedAt,
		Description:  alarm.Description,
		DeviceID:     alarm.DeviceID.Hex(),
		RepeatWeekly: alarm.RepeatWeekly,
		Weekday:      alarm.Weekday,
	}
}
