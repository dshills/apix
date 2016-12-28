// Package store provides access to data storage
package store

import (
	"fmt"
	"log"

	// MySQL driver
	"github.com/dshills/apix/config"
	"github.com/jmoiron/sqlx"
)

var localDB *sqlx.DB

// Config will accept a config.Database and Connect to the db
func Config(con *config.Database) error {
	// Database
	if con.Host == "" || con.Name == "" || con.User == "" {
		return fmt.Errorf("%v %v %v are required", con.Prefix+"_DB_HOST", con.Prefix+"_DB_NAME", con.Prefix+"_DB_USER")
	}
	return ConnectDSN(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v%v", con.User, con.Password, con.Host, con.Port, con.Name, con.Params))
}

// Connect to a MySQL instance
func Connect(user, pass, name, host string, port int, params string) error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v%v", user, pass, host, port, name, params)
	return ConnectDSN(dsn)
}

// ConnectDSN to a MySQL instance using a DSN
func ConnectDSN(dsn string) error {
	if localDB != nil {
		return fmt.Errorf("Database already connected")
	}
	var err error
	localDB, err = sqlx.Connect("mysql", dsn)
	return err
}

// DB returns the sqlx.DB conection
// if connection is nil panics
func DB() *sqlx.DB {
	if localDB == nil {
		log.Fatal("Database is not connected")
	}
	return localDB
}

// IsConnected checks for database connection
// if public for checking during tests.
func IsConnected() bool {
	if localDB == nil {
		return false
	}
	return true
}

// CloseDB will close the current connection
func CloseDB() {
	localDB.Close()
	localDB = nil
}
