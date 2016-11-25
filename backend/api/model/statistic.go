package model

import "time"

// StatisticForm model for API requests
type StatisticForm struct {
	Date        time.Time `json:"date" binding:"required"`
	Value       float64   `json:"value" binding:"required"`
	Description string    `json:"description"`
	DeviceID    string    `json:"device_id" binding:"required"`
	Type        string    `json:"type" binding:"required"`
}
