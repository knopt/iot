package service

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"

	databaseModel "github.com/knopt/iot/backend/database/model"
)

// InsertStatistic from api to db
func (service *Service) InsertStatistic(statisticForm *apiModel.StatisticForm) error {
	statistic, err := apiStatisticFormToDatabaseStatistic(statisticForm)
	if err != nil {
		return err
	}
	err = service.db.InsertStatistic(statistic)
	if err != nil {
		return err
	}
	return nil
}

// GetStatisticsTypes by deviceID
func (service *Service) GetStatisticsTypes(deviceID string) ([]string, error) {
	if err := validateBsonObjectID(deviceID); err != nil {
		return nil, err
	}

	bsonDeviceID := bson.ObjectIdHex(deviceID)

	return service.db.GetStatisticsTypes(bsonDeviceID)
}

// GetStatistics by id from database and return to API
func (service *Service) GetStatistics(deviceID string, dateFrom string, dateTo string, dataType string) ([]*apiModel.StatisticForm, error) {
	var err error
	var dbDateTo, dbDateFrom time.Time

	if len(dateTo) > 0 {
		dbDateTo, err = time.Parse(time.RFC3339, dateTo)
		if err != nil {
			return nil, err
		}
	} else {
		dbDateTo = time.Now()
	}
	if len(dateFrom) > 0 {
		dbDateFrom, err = time.Parse(time.RFC3339, dateFrom)
		if err != nil {
			return nil, err
		}
	} else {
		dbDateFrom = time.Time{}
	}

	if err = validateBsonObjectID(deviceID); err != nil {
		return nil, err
	}

	bsonDeviceID := bson.ObjectIdHex(deviceID)
	if err != nil {
		return nil, err
	}

	statistics, err := service.db.GetStatisticsDeviceFromToType(bsonDeviceID, dbDateFrom, dbDateTo, dataType)
	if err != nil {
		return nil, err
	}

	apiStatistics := arrayOfDatabaseStatisticsToArrayOfApiStatistics(statistics)

	return apiStatistics, nil
}

func apiStatisticFormToDatabaseStatistic(statisticForm *apiModel.StatisticForm) (*databaseModel.Statistic, error) {
	if err := validateBsonObjectID(statisticForm.DeviceID); err != nil {
		return nil, err
	}

	return &databaseModel.Statistic{
		DeviceID:    bson.ObjectIdHex(statisticForm.DeviceID),
		Value:       statisticForm.Value,
		Description: statisticForm.Description,
		Type:        statisticForm.Type,
		Date:        statisticForm.Date,
	}, nil
}

func databaseStatisticToAPIStatisticForm(statistics *databaseModel.Statistic) *apiModel.StatisticForm {
	return &apiModel.StatisticForm{
		DeviceID:    statistics.DeviceID.Hex(),
		Value:       statistics.Value,
		Description: statistics.Description,
		Type:        statistics.Type,
		Date:        statistics.Date,
	}
}

// arrayOfDatabaseStatisticsToArrayOfApiStatistics converts db Statistics array to api Statistics array
func arrayOfDatabaseStatisticsToArrayOfApiStatistics(dbStatistics []*databaseModel.Statistic) []*apiModel.StatisticForm {
	apiStatistics := make([]*apiModel.StatisticForm, 0)

	for _, dbStatistic := range dbStatistics {
		apiStatistic := databaseStatisticToAPIStatisticForm(dbStatistic)
		apiStatistics = append(apiStatistics, apiStatistic)
	}

	return apiStatistics
}
