package model

import "time"

// AlarmForm model
type AlarmForm struct {
	DeviceID     int64        `json:"device_id" binding:"required"`
	Time         time.Time    `json:"time" binding:"required"`
	Day          time.Weekday `json:"weekday" binding:"required"`
	RepeatWeekly bool         `json:"repeat_weekly" binding:"required"`
	CreatedAt    time.Time    `json:"creation_date_time" binding:"required"`
}
