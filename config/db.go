package config

import (
	"fmt"
	"time"
)

// DatabaseConfig is the structure for the database config
type DatabaseConfig struct {
	databaseName        string
	databaseUser        string
	databasePassword    string
	databaseHost        string
	databasePort        int
	databaseMaxPoolSize int
	readTimeOut         time.Duration
	writeTimeOut        time.Duration
}

func newDatabaseConfig(prefix string) DatabaseConfig {
	return DatabaseConfig{
		databaseHost:        mustHaveString(prefix + "DATABASE_HOST"),
		databaseName:        mustHaveString(prefix + "DATABASE_NAME"),
		databaseUser:        mustHaveString(prefix + "DATABASE_USER"),
		databasePassword:    mustHaveString(prefix + "DATABASE_PASSWORD"),
		databasePort:        mustGetInt(prefix + "DATABASE_PORT"),
		databaseMaxPoolSize: mustGetInt(prefix + "DATABASE_MAX_POOL_SIZE"),
		readTimeOut:         mustGetTimeDurationInMs(prefix + "DATABASE_READ_TIME_OUT_MS"),
		writeTimeOut:        mustGetTimeDurationInMs(prefix + "DATABASE_WRITE_TIME_OUT_MS"),
	}
}

func (dc DatabaseConfig) connectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		dc.databaseName,
		dc.databaseUser,
		dc.databasePassword,
		dc.databaseHost,
		dc.databasePort)
}
