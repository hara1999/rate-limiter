// Package database handles db connections
package database

import (
	"time"

	"github.com/hara1999/fluxy/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

// DB struct holds the database connection
type DB struct {
	Database *gorm.DB
}

// DBConnection returns a DB struct with a gorm connection
func DBConnection(dsn string) (*DB, error) {
	logMode := viper.GetBool("DB_LOG_MODE")
	loglevel := gormLog.Silent
	if logMode {
		loglevel = gormLog.Info
	}
	pgConn := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})

	db, err := gorm.Open(pgConn, &gorm.Config{Logger: gormLog.Default.LogMode(loglevel)})

	if err != nil {
		logger.Fatal("database refused %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 1)
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	err = sqlDB.Ping()
	if err != nil {
		logger.Fatal("%v", err)
	}
	logger.Log("database connected")

	return &DB{Database: db}, nil
}
