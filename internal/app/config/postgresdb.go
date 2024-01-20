package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormLogger "gorm.io/gorm/logger"
)

type DBPostgreOption struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	MaxPoolSize     int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// NewPostgreDatabase return gorm dbmap object with postgre options param
func NewPostgreDatabase(option DBPostgreOption) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		option.Host,
		option.Username,
		option.Password,
		option.DBName,
		option.Port,
	)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(getLoggerLevel(AppConfig().GetString("GORM_LOGGER"))),
	})
	if err != nil {
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = pgsqlDB.Ping()
	if err != nil {
		return nil, err
	}

	pgsqlDB.SetMaxOpenConns(option.MaxPoolSize)
	pgsqlDB.SetConnMaxLifetime(option.ConnMaxLifetime)
	pgsqlDB.SetMaxIdleConns(option.MaxIdleConns)

	return db, nil
}

// getLoggerLevel return gorm log level setup based from env/config of the app
func getLoggerLevel(v string) gormLogger.LogLevel {
	switch v {
	case "error":
		return gormLogger.Error
	case "warn":
		return gormLogger.Warn
	case "info":
		return gormLogger.Info
	default:
		return gormLogger.Silent
	}
}
