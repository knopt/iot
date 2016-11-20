package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Device database model
type Device struct {
	ID           bson.ObjectId `bson:"_id"`
	Description  string        `bson:"description,omitempty"`
	Name         string        `bson:"name"`
	RegisteredAt time.Time     `bson:"registered_at"`
}
