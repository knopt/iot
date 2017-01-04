package service

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"
	databaseModel "github.com/knopt/iot/backend/database/model"
)

// GetDevice by id from database and return to API
func (service *Service) GetDevice(deviceID string) (*apiModel.DeviceForm, error) {
	if err := validateBsonObjectID(deviceID); err != nil {
		return nil, err
	}
	dbDevice, err := service.db.GetDevice(bson.ObjectIdHex(deviceID))
	if err != nil {
		return nil, err
	}
	return databaseDeviceToAPIDeviceForm(dbDevice), nil
}

// GetDevices get every device in db
func (service *Service) GetDevices() ([]*apiModel.DeviceForm, error) {
	dbDevices, err := service.db.GetDevices()
	if err != nil {
		fmt.Printf("error in getdevices")
		return nil, err
	}
	return arrayOfDatabaseDevicesToArrayOfAPIDevices(dbDevices), nil
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

func arrayOfDatabaseDevicesToArrayOfAPIDevices(dbDevices []*databaseModel.Device) []*apiModel.DeviceForm {
	apiDevices := make([]*apiModel.DeviceForm, 0)

	for _, dbDevice := range dbDevices {
		apiDevice := databaseDeviceToAPIDeviceForm(dbDevice)
		apiDevices = append(apiDevices, apiDevice)
	}

	return apiDevices
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
