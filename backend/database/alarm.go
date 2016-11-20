package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/knopt/iot/backend/database/model"
)

// GetAlarm by alarmID
func (database *Database) GetAlarm(alarmID bson.ObjectId) *model.Alarm {
	var alarm model.Alarm
	err := database.db.C("alarm").FindId(alarmID).One(&alarm)
	if err != nil {
		panic(err)
	}
	return &alarm
}

// InsertAlarm into database
func (database *Database) InsertAlarm(alarm *model.Alarm) bson.ObjectId {
	alarm.ID = bson.NewObjectId()
	alarm.UploadedAt = time.Now()
	if err := database.db.C("alarm").Insert(alarm); err != nil {
		panic(err)
	}

	return alarm.ID
}
