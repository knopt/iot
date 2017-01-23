package model

import "time"

// StatisticForm model for API requests
type StatisticForm struct {
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	DeviceID    string    `json:"device_id"`
	Type        string    `json:"type"`
}
