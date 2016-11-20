package service

import "github.com/knopt/iot/backend/database"

// Service connects database and API
type Service struct {
	db *database.Database
}

// NewService creates new service singleton object
func NewService(db *database.Database) *Service {
	return &Service{db}
}
