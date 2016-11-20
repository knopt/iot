package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Alarm model
type Alarm struct {
	ID           bson.ObjectId `bson:"_id"`
	DeviceID     bson.ObjectId `bson:"device_id"`
	AlarmTime    time.Time     `bson:"alarm_time"`
	Description  string        `bson:"description,omitempty"`
	CreatedAt    time.Time     `bson:"created_at"`
	RepeatWeekly bool          `bson:"repeat_weekly"`
	UploadedAt   time.Time     `bson:"uploaded_at"`
	Weekday      time.Weekday  `bson:"weekday"`
}
