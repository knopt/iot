package service

import (
	"errors"

	"github.com/knopt/iot/backend/database"
	"gopkg.in/mgo.v2/bson"
)

// Service connects database and API
type Service struct {
	db *database.Database
}

// NewService creates new service singleton object
func NewService(db *database.Database) *Service {
	return &Service{db}
}

func validateBsonObjectID(id string) error {
	if len(id) != 24 {
		return errors.New("Wrong bson ID length")
	} else if !bson.IsObjectIdHex(id) {
		return errors.New("Wrong format of ID")
	}
	return nil
}
