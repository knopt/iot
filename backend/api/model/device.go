package model

import "time"

// DeviceForm for API calls
type DeviceForm struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	RegisteredAt time.Time `json:"registered_at"`
}
