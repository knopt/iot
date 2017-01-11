package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/knopt/iot/backend/database/model"
)

// GetAlarm by alarmID
func (database *Database) GetAlarm(alarmID bson.ObjectId) (*model.Alarm, error) {
	var alarm model.Alarm
	err := database.db.C("alarm").FindId(alarmID).One(&alarm)
	if err != nil {
		return nil, err
	}
	return &alarm, nil
}

// GetNearestAlarm get first alarm to ring
func (database *Database) GetNearestAlarm(deviceID bson.ObjectId, date time.Time) (*model.Alarm, error) {
	var alarms []*model.Alarm
	err := database.db.C("alarm").Find(
		bson.M{
			"device_id": deviceID,
			"alarm_time": bson.M{
				"$gte": date,
			},
		},
	).Sort("alarm_time").All(&alarms)
	if err != nil {
		return nil, err
	}
	return alarms[0], nil
}

// GetAlarmsByDevice returns array of alarms
func (database *Database) GetAlarmsByDevice(deviceID bson.ObjectId) ([]*model.Alarm, error) {
	var alarms []*model.Alarm

	err := database.db.C("alarm").Find(
		bson.M{
			"device_id": deviceID,
		},
	).All(&alarms)

	if err != nil {
		return nil, err
	}

	return alarms, nil
}

// InsertAlarm into database
func (database *Database) InsertAlarm(alarm *model.Alarm) (*bson.ObjectId, error) {
	alarm.ID = bson.NewObjectId()
	alarm.UploadedAt = time.Now()
	if err := database.db.C("alarm").Insert(alarm); err != nil {
		return nil, err
	}

	return &alarm.ID, nil
}
