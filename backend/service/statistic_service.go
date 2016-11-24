package service

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	apiModel "github.com/knopt/iot/backend/api/model"

	databaseModel "github.com/knopt/iot/backend/database/model"
)

// InsertStatistic from api to db
func (service *Service) InsertStatistic(statisticForm *apiModel.StatisticForm) error {
	statistic, err := apiStatisticFormToDatabaseStatistic(statisticForm)
	fmt.Printf("service %v\n", statistic)
	if err != nil {
		return err
	}
	err = service.db.InsertStatistic(statistic)
	if err != nil {
		return err
	}
	return nil
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

	bsonDeviceID, err := stringToBsonObjectID(deviceID)
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

func stringToBsonObjectID(id string) (bson.ObjectId, error) {
	if len(id) != 24 {
		return bson.NewObjectId(), errors.New("Wrong bson ID length")
	} else if !bson.IsObjectIdHex(id) {
		return bson.NewObjectId(), errors.New("Wrong format of deviceID")
	}
	return bson.ObjectIdHex(id), nil
}

func apiStatisticFormToDatabaseStatistic(statisticForm *apiModel.StatisticForm) (*databaseModel.Statistic, error) {
	if len(statisticForm.DeviceID) != 24 {
		return nil, errors.New("Wrong deviceID length")
	} else if !bson.IsObjectIdHex(statisticForm.DeviceID) {
		return nil, errors.New("Wrong format of deviceID")
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
