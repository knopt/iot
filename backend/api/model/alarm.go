package model

import "time"

// AlarmForm model
type AlarmForm struct {
	ID           string       `json:"id"`
	AlarmTime    time.Time    `json:"alarm_time" binding:"required"`
	CreatedAt    time.Time    `json:"created_at" binding:"required"`
	Description  string       `json:"description"`
	DeviceID     string       `json:"device_id" binding:"required"`
	RepeatWeekly bool         `json:"repeat_weekly"`
	Weekday      time.Weekday `json:"weekday"`
}
