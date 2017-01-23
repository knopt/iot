package service

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"
	databaseModel "github.com/knopt/iot/backend/database/model"
)

// GetAlarm by id from database and return to API
func (service *Service) GetAlarm(alarmID string) (*apiModel.AlarmForm, error) {
	if err := validateBsonObjectID(alarmID); err != nil {
		return nil, err
	}
	dbAlarm, err := service.db.GetAlarm(bson.ObjectIdHex(alarmID))
	if err != nil {
		return nil, err
	}
	return databaseAlarmToAPIAlarmForm(dbAlarm), nil
}

// GetAlarmsByDevice get alarms by device
func (service *Service) GetAlarmsByDevice(deviceID string) ([]*apiModel.AlarmForm, error) {
	if err := validateBsonObjectID(deviceID); err != nil {
		return nil, err
	}
	dbAlarms, err := service.db.GetAlarmsByDevice(bson.ObjectIdHex(deviceID))
	if err != nil {
		return nil, err
	}
	return arrayOfDatabaseAlarmsToArrayOfAPIAlarms(dbAlarms), nil
}

// GetNearestAlarmByDevice get first alarm to ring
func (service *Service) GetNearestAlarmByDevice(deviceID string) (*apiModel.AlarmForm, error) {
	if err := validateBsonObjectID(deviceID); err != nil {
		return nil, err
	}
	dbAlarm, err := service.db.GetNearestAlarm(bson.ObjectIdHex(deviceID), time.Now())
	if err != nil {
		return nil, err
	}
	apiAlarm := databaseAlarmToAPIAlarmForm(dbAlarm)

	return apiAlarm, nil
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

// DeleteAlarm by deviceID
func (service *Service) DeleteAlarm(alarmID string) error {
	if err := validateBsonObjectID(alarmID); err != nil {
		return err
	}
	err := service.db.DeleteAlarm(bson.ObjectIdHex(alarmID))

	return err
}

// DeleteNearestAlarm by deviceID
func (service *Service) DeleteNearestAlarm(deviceID string) error {
	if err := validateBsonObjectID(deviceID); err != nil {
		return err
	}
	err := service.db.DeleteNearestAlarm(bson.ObjectIdHex(deviceID))

	return err
}

func arrayOfDatabaseAlarmsToArrayOfAPIAlarms(dbAlarms []*databaseModel.Alarm) []*apiModel.AlarmForm {
	apiAlarms := make([]*apiModel.AlarmForm, 0)

	for _, dbAlarm := range dbAlarms {
		apiAlarm := databaseAlarmToAPIAlarmForm(dbAlarm)
		apiAlarms = append(apiAlarms, apiAlarm)
	}

	return apiAlarms
}

func apiAlarmFormToDatabaseAlarm(alarmForm *apiModel.AlarmForm) (*databaseModel.Alarm, error) {
	if err := validateBsonObjectID(alarmForm.DeviceID); err != nil {
		return nil, err
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
