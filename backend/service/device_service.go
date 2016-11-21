package service

import (
	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"
	databaseModel "github.com/knopt/iot/backend/database/model"
)

// GetDevice by id from database and return to API
func (service *Service) GetDevice(deviceID string) (*apiModel.DeviceForm, error) {
	dbDevice, err := service.db.GetDevice(bson.ObjectIdHex(deviceID))
	if err != nil {
		return nil, err
	}
	return databaseDeviceToAPIDeviceForm(dbDevice), nil
}

// RegisterDevice in database and return its ID to API
func (service *Service) RegisterDevice(deviceForm *apiModel.DeviceForm) (*string, error) {
	device := apiDeviceFormToDatabaseDevice(deviceForm)
	deviceBsonID, err := service.db.InsertDevice(device)
	if err != nil {
		return nil, err
	}

	deviceHexBsonID := deviceBsonID.Hex()
	return &deviceHexBsonID, nil
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
