package config

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func NewMySQLDatabase(configuration Config) (*sql.DB, error) {
	username := configuration.Get("DB_USERNAME")
	password := configuration.Get("DB_PASSWORD")
	host := configuration.Get("DB_HOST")
	portStr := configuration.Get("DB_PORT")
	dbName := configuration.Get("DB_DATABASE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)

	dbConn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(5)
	dbConn.SetConnMaxLifetime(1 * time.Hour)

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
