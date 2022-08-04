package server

import "time"

// GRPC and HTTP config
const (
	// IP адрес сервера
	IP = "localhost"

	// Port порт сервера
	Port = ":8081"

	// HTTPPort порт сервера
	HTTPPort = ":8080"

	// FullAddress Полный адрес сервера
	FullAddress = IP + Port

	// FullHTTPAddress Полный адрес сервера
	FullHTTPAddress = IP + HTTPPort

	// ConnectionType тип соединения
	ConnectionType = "tcp"
)

const (
	hostDB   = "localhost"
	portDB   = 5432
	userName = "task_manager_db"
	nameDB   = "task_manager_db"

	maxConnIdleTime = time.Minute
	maxConnLifetime = time.Hour
	minConnections  = 2
	maxConnections  = 4
)
