package storage

import "time"

const (
	serviceAddress = "localhost:8082"
	connectionType = "tcp"

	hostDB   = "localhost"
	portDB   = 5432
	userName = "task_manager_db"
	nameDB   = "task_manager_db"

	maxConnIdleTime = time.Minute
	maxConnLifetime = time.Hour
	minConnections  = 2
	maxConnections  = 4
)
