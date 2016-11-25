package database

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// TODO: Move database config to seperate file, add users to restrict access to DB

// Database struct, keeps reference to running database
type Database struct {
	db *mgo.Database
}

// NewDatabaseConnection returns new connection to database
func NewDatabaseConnection() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	connection := session.DB("iot")

	fmt.Printf("Database connection established correctly")

	return connection
}

// NewDatabase creates global Database object from database connection
func NewDatabase(connection *mgo.Database) *Database {
	connection.C("statistic").EnsureIndexKey("device_id", "type", "date")
	connection.C("alarm").EnsureIndexKey("_id", "device_id", "alarm_time")
	connection.C("device").EnsureIndexKey("_id")
	return &Database{connection}
}
