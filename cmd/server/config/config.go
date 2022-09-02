package config

// GRPC and HTTP config
const (
	address = "localhost"

	portGRPC    = ":8081"
	addressGRPC = address + portGRPC

	portHTTP    = ":8083"
	addressHTTP = address + portHTTP

	connectionType = "tcp"

	storageAddress = "localhost:8082"

	swaggerDir = "/home/ivan/GolandProjects/TaskBot/swagger"
)

var (
	brokers = []string{"localhost:9095"}
)
