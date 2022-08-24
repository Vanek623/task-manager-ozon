package storage

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// ConfigDB структура конфигурации БД
type ConfigDB struct {
	Host     string
	Port     string
	UserName string
	Name     string
	Password string
}

var configDB ConfigDB

// GetConfig чтение конфигурации БД
func GetConfig() ConfigDB {
	return configDB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	configDB.Host = "localhost"
	configDB.Port = "5432"
	configDB.Name = "task_manager_db"
	configDB.UserName = "user"
	configDB.Password = os.Getenv("DB_PASSWORD")
}

const (
	serviceAddress = "localhost:8082"
	connectionType = "tcp"

	maxConnIdleTime = time.Minute
	maxConnLifetime = time.Hour
	minConnections  = 2
	maxConnections  = 4
)
