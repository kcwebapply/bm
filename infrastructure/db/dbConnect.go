package db

import (
	"os"

	"github.com/gocraft/dbr"
	"github.com/kcwebapply/bm/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
)

var conn *dbr.Connection

// db initialization
func init() {
	path, err := homedir.Dir()
	if err != nil {
		util.LoggingError(err.Error())
		os.Exit(0)
	}
	dbPath := path + "/page.db"
	_, err = os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		util.LoggingError(err.Error())
	}

	connection, err := dbr.Open("sqlite3", dbPath, nil)
	if err != nil {
		util.LoggingError(err.Error())

	}
	conn = connection
}

// GetConnection returns db connection
func GetConnection() *dbr.Connection {
	return conn
}
