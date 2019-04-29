package db

import (
	"fmt"
	"os"

	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
)

var conn *dbr.Connection

// db initialization
func init() {
	path, err := homedir.Dir()
	if err != nil {
		fmt.Println("dberr:", err)
		os.Exit(0)
	}
	dbPath := path + "/page.db"
	_, err = os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	connection, err := dbr.Open("sqlite3", dbPath, nil)
	if err != nil {
		fmt.Println("error happened in connection:", err)
		os.Exit(0)
	}
	conn = connection
}

// GetConnection returns db connection
func GetConnection() *dbr.Connection {
	return conn
}
