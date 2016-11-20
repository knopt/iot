package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/knopt/iot/backend/database/model"
)

// GetDevice by deviceID
func (database *Database) GetDevice(deviceID bson.ObjectId) *model.Device {
	var device model.Device
	err := database.db.C("device").FindId(deviceID).One(&device)
	if err != nil {
		panic(err)
	}
	return &device
}

// InsertDevice into database
func (database *Database) InsertDevice(device *model.Device) bson.ObjectId {
	device.ID = bson.NewObjectId()
	device.RegisteredAt = time.Now()
	if err := database.db.C("device").Insert(device); err != nil {
		panic(err)
	}

	return device.ID
}
