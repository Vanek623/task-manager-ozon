package server

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
