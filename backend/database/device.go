package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/knopt/iot/backend/database/model"
)

// GetDevice by deviceID
func (database *Database) GetDevice(deviceID bson.ObjectId) (*model.Device, error) {
	var device model.Device
	err := database.db.C("device").FindId(deviceID).One(&device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

// GetDevices get all devices from db
func (database *Database) GetDevices() ([]*model.Device, error) {
	var devices []*model.Device

	err := database.db.C("device").Find(nil).All(&devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// InsertDevice into database
func (database *Database) InsertDevice(device *model.Device) (*bson.ObjectId, error) {
	device.ID = bson.NewObjectId()
	device.RegisteredAt = time.Now()
	if err := database.db.C("device").Insert(device); err != nil {
		return nil, err
	}

	return &device.ID, nil
}
