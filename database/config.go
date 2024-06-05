package database

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	DBName     string
	DBHost     string
	DBUsername string
	DBPassword string
	DBPort     string
	DBDriver   string
}

func (dbConfig *DatabaseConfig) Url() (string, error) {
	if (dbConfig.DBDriver == "" ) || (dbConfig.DBDriver == "mysql") {
		return dbConfig.msqlUrl(), nil
	}

	return "", nil
}

func (dbConfig *DatabaseConfig) msqlUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUsername)
}

func DatabaseSetup() DatabaseConfig {

	name, port, username, password, host, driver := os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_DRIVER")

	return DatabaseConfig{
		DBName:     name,
		DBPort:     port,
		DBHost:     host,
		DBUsername: username,
		DBPassword: password,
		DBDriver:   driver,
	}
}
