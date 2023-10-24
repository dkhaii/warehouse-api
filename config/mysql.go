package config

import (
	"database/sql"
	"fmt"
	"time"
)

func NewMySQLDatabase(cfg Config) (*sql.DB, error) {
	username := cfg.GetString("DB_USERNAME")
	password := cfg.GetString("DB_PASSWORD")
	host := cfg.GetString("DB_HOST")
	port := cfg.GetInt("DB_PORT")
	dbName := cfg.GetString("DB_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, dbName)

	dbConn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(20)
	dbConn.SetConnMaxIdleTime(15 * time.Minute)
	dbConn.SetConnMaxLifetime(90 * time.Minute)

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// migrate -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" -path database/migrations up
