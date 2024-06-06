package bnrdbmigrator

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultSlowThreshold = 10
)

func OpenPostgresConnection(dbUsername, dbPassword, dbHost, dbPort, dbName string, isEnableLog bool, isDisableSSLMode bool) (*gorm.DB, error) {
	db, err := openPostgreSQL(dbUsername, dbPassword, dbHost, dbPort, dbName, isEnableLog, isDisableSSLMode)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openPostgreSQL(dbUsername, dbPassword, dbHost, dbPort, dbName string, isEnableLog, isDisableSSLMode bool) (*gorm.DB, error) {

	cfg := gorm.Config{}

	logLv := logger.Error
	if isEnableLog {
		logLv = logger.Info
	}
	newLogger := NewLogger(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: defaultSlowThreshold,
			LogLevel:      logLv,
			Colorful:      true,
		},
	)
	cfg.Logger = newLogger

	sslMode := "require"
	if isDisableSSLMode {
		sslMode = "disable"
	}

	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + sslMode + " TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &cfg)

	return db, err
}
