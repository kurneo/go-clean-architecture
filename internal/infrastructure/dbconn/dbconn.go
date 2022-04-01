package dbconn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kurneo/config"
	"sync"
)

type DBConn struct {
	DB *gorm.DB
}

var (
	dbConnOnce   sync.Once
	dbConnection *DBConn
)

func NewConnection() (*DBConn, error) {
	dbConfig := config.NewDBConfig()
	connection, e := dbConfig.GetConnection()

	if e != nil {
		return nil, e
	}

	var err error

	dbConnOnce.Do(func() {

		dbConnection = &DBConn{}

		switch dbConfig.Driver {
		case config.DBConnectionMysql:
			err = mysqlConnect(connection)
		case config.DBConnectionPostgres:
			err = postgresConnect(connection)
		default:
			err = nil
		}
	})

	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

func mysqlConnect(connection *config.Connection) error {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		connection.Username,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Database,
		connection.Charset,
		connection.ParseTime,
		connection.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	dbConnection.DB = db

	return nil
}

func postgresConnect(connection *config.Connection) error {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		connection.Host,
		connection.Username,
		connection.Password,
		connection.Database,
		connection.Port,
		connection.SSLMode,
		connection.TimeZone,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: false,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	dbConnection.DB = db

	return nil
}
