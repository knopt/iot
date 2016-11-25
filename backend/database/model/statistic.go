package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Statistic database model
type Statistic struct {
	DeviceID    bson.ObjectId `bson:"device_id"`
	Date        time.Time     `bson:"date"`
	Description string        `bson:"description,omitempty"`
	Value       float64       `bson:"value"`
	Type        string        `bson:"type"`
}
