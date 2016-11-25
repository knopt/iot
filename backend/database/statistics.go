package database

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/knopt/iot/backend/database/model"
)

// GetStatisticsDeviceFromToType by deviceId, dataFrom, dateTo and dataType
func (database *Database) GetStatisticsDeviceFromToType(deviceID bson.ObjectId, dateFrom time.Time, dateTo time.Time, dataType string) ([]*model.Statistic, error) {
	var allStatistics []*model.Statistic
	fmt.Printf("GET STATS: %v %v %v %v\n", deviceID, dateFrom, dateTo, dataType)

	err := database.db.C("statistic").Find(
		bson.M{
			"device_id": deviceID,
			"type":      dataType,
			"date": bson.M{
				"$gt": dateFrom,
				"$lt": dateTo,
			},
		},
	).All(&allStatistics)
	if err != nil {
		return nil, err
	}
	return allStatistics, nil
}

// InsertStatistic into database
func (database *Database) InsertStatistic(statistic *model.Statistic) error {
	fmt.Printf("%v\n", statistic)
	if err := database.db.C("statistic").Insert(statistic); err != nil {
		return err
	}

	return nil
}
