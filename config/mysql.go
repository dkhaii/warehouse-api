package config

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

func NewMySQLDatabase(cfg Config) (*sql.DB, error) {
	username := cfg.GetString("DB_USERNAME")
	password := cfg.GetString("DB_PASSWORD")
	host := cfg.GetString("DB_HOST")
	port := cfg.GetString("DB_PORT")
	dbName := cfg.GetString("DB_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)

	dbPool, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	dbPool.SetMaxOpenConns(10)
	dbPool.SetMaxIdleConns(20)
	dbPool.SetConnMaxIdleTime(15 * time.Minute)
	dbPool.SetConnMaxLifetime(90 * time.Minute)

	err = dbPool.Ping()
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

func NewMySQLCloudDatabase(cfg Config) (*sql.DB, error) {
	dbUser := cfg.GetString("DB_USERNAME")
	dbPassword := cfg.GetString("DB_PASSWORD")
	dbName := cfg.GetString("DB_DATABASE")
	cloudSQLConn := cfg.GetString("CLOUDSQL_CONNECTION_NAME")
	usePrivate := cfg.GetString("PRIVATE_IP")

	dialer, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}

	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return dialer.Dial(ctx, cloudSQLConn, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	dbPool.SetMaxOpenConns(10)
	dbPool.SetMaxIdleConns(20)
	dbPool.SetConnMaxIdleTime(15 * time.Minute)
	dbPool.SetConnMaxLifetime(90 * time.Minute)

	err = dbPool.Ping()
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

// migrate -database "mysql://development:development@tcp(localhost:3306)/cozy_warehouse" -path database/migrations up
