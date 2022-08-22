package storage

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
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

func GetConfig() ConfigDB {
	return configDB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	configDB.Host = "localhost"
	configDB.Port = "5432"
	configDB.UserName = "task_manager_db"
	configDB.Name = "task_manager_db"
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
