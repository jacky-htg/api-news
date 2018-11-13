package repositories

import (
	"database/sql"

	"github.com/jacky-htg/api-news/config"
	"github.com/jacky-htg/api-news/libraries"
)

var db *sql.DB
var err error

func init() {
	// Create an sql.DB and check for errors
	db, err = sql.Open(config.GetString("database.mysql.driverName"), config.GetString("database.mysql.dataSourceName"))
	libraries.CheckError(err)

	// Test the connection to the database
	err = db.Ping()
	libraries.CheckError(err)
}
