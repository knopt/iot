package database

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Database struct
type Database struct {
	db *mgo.Database
}

// NewDatabaseConnection returns new connection to database
// TODO: Move database config to seperate file, add users to restrict access to DB
func NewDatabaseConnection() *mgo.Database {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	connection := session.DB("iot")

	fmt.Printf("Database connection established correctly")

	return connection
}

//NewDatabase creates global Database object from database connection
func NewDatabase(connection *mgo.Database) *Database {
	return &Database{connection}
}

//GetAlarm by alarmID
//TODO: implement and move to seperate alarms file
func (database *Database) GetAlarm(alarmID int64) int64 {
	fmt.Printf("Database - GetAlarm. alarmID: %v\n", alarmID)
	return alarmID
}
