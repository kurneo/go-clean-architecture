package config

import (
	"errors"
	"os"
	"sync"
)

// Database Config
const (
	DBConnectionMysql    = "mysql"
	DBConnectionPostgres = "postgres"
)

type Connection struct {
	Driver    string
	Host      string
	Port      string
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime string
	Loc       string
	SSLMode   string
	TimeZone  string
}

type DBConfig struct {
	Config
	Driver      string
	Connections []Connection
}

var (
	dbConfigOnce sync.Once
	dbConfig     *DBConfig
)

func (dbConfig *DBConfig) GetConnection() (*Connection, error) {
	var connection *Connection

	for i := 0; i < len(dbConfig.Connections); i++ {
		if dbConfig.Connections[i].Driver == dbConfig.Driver {
			connection = &dbConfig.Connections[i]
			break
		}
	}

	if connection == nil {
		return nil, errors.New("cannot find connection for driver " + dbConfig.Driver)
	}

	return connection, nil
}

func NewDBConfig() *DBConfig {
	dbConfigOnce.Do(func() {
		//Init connections
		mysqlConnection := Connection{
			Driver:    DBConnectionMysql,
			Host:      os.Getenv("DB_HOST"),
			Port:      os.Getenv("DB_PORT"),
			Database:  os.Getenv("DB_DATABASE"),
			Username:  os.Getenv("DB_USERNAME"),
			Password:  os.Getenv("DB_PASSWORD"),
			Charset:   "utf8",
			ParseTime: "True",
			Loc:       "Local",
		}

		postgresConnection := Connection{
			Driver:   DBConnectionPostgres,
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  "disable",
			TimeZone: "Asia/Ho_Chi_Minh",
		}

		dbConfig = &DBConfig{
			Driver:      os.Getenv("DB_CONNECTION"),
			Connections: []Connection{mysqlConnection, postgresConnection},
			Config:      Config{Type: "db"},
		}
	})

	return dbConfig
}
