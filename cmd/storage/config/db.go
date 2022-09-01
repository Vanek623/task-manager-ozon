package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// DB структура конфигурации БД
type DB struct {
	Host     string
	Port     string
	UserName string
	Name     string
	Password string

	MaxConnIdleTime time.Duration
	MaxConnLifetime time.Duration
	MinConnections  int32
	MaxConnections  int32
}

var configDB DB

// GetConfigDB чтение конфигурации БД
func GetConfigDB() DB {
	return configDB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// DB configs
	configDB.Host = "localhost"
	configDB.Port = "5432"
	configDB.Name = "task_manager_db"
	configDB.UserName = "user"
	configDB.Password = os.Getenv("DB_PASSWORD")

	// Pool configs
	configDB.MaxConnIdleTime = time.Minute
	configDB.MaxConnLifetime = time.Hour
	configDB.MinConnections = 2
	configDB.MaxConnections = 4
}
