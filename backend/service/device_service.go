package service

import (
	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"
	databaseModel "github.com/knopt/iot/backend/database/model"
)

// GetDevice by id from database and return to API
func (service *Service) GetDevice(deviceID string) *apiModel.DeviceForm {
	dbDevice := service.db.GetDevice(bson.ObjectIdHex(deviceID))
	return databaseDeviceToAPIDeviceForm(dbDevice)
}

// RegisterDevice in database and return its ID to API
func (service *Service) RegisterDevice(deviceForm *apiModel.DeviceForm) string {
	device := apiDeviceFormToDatabaseDevice(deviceForm)
	deviceBsonID := service.db.InsertDevice(device)
	return deviceBsonID.Hex()
}

func apiDeviceFormToDatabaseDevice(deviceForm *apiModel.DeviceForm) *databaseModel.Device {
	return &databaseModel.Device{
		Name:        deviceForm.Name,
		Description: deviceForm.Description,
	}
}

func databaseDeviceToAPIDeviceForm(device *databaseModel.Device) *apiModel.DeviceForm {
	return &apiModel.DeviceForm{
		ID:           device.ID.Hex(),
		RegisteredAt: device.RegisteredAt,
		Description:  device.Description,
		Name:         device.Name,
	}
}
